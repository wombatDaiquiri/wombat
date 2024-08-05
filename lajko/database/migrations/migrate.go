package migrations

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Migration struct {
	ID             string
	MigrationQuery string
	RollbackQuery  string
}

type AppliedMigration struct {
	ID            string
	RollbackQuery string

	SchemaBefore string
	SchemaAfter  string
}

var migrations = []Migration{
	{
		ID: "20240522124900",
		MigrationQuery: `
		CREATE TABLE actors (
			uri TEXT PRIMARY KEY,
			displayed_name TEXT,
			avatar_uri TEXT
		);

		CREATE TABLE content (
			uri TEXT PRIMARY KEY,
			author_uri TEXT REFERENCES actors(uri),
			in_response_to_uri TEXT REFERENCES content(uri),
			parser_version TEXT,
			lajko_html TEXT,
			attachments JSONB
		);

		CREATE TABLE likes (
			who_uri TEXT REFERENCES actors(uri),
			what_uri TEXT REFERENCES content(uri),
			PRIMARY KEY (who_uri, what_uri)
		);`,
		RollbackQuery: `
			DROP TABLE likes;
			DROP TABLE content;
			DROP TABLE actors;
		`,
	},
}

// Migrate executes all migrations
func Migrate(ctx context.Context, db *pgxpool.Pool) {
	var migrationTableExists bool
	err := pgxscan.Get(ctx, db, &migrationTableExists, `
		SELECT EXISTS (
		   SELECT FROM pg_tables
		   WHERE  schemaname = 'public'
		   AND    tablename  = 'migrations'
	    );`)
	if err != nil {
		panic(fmt.Errorf("could not check if migrations table exists: %w", err))
	}

	if !migrationTableExists {
		_, err := db.Exec(ctx, `CREATE TABLE migrations(
				id VARCHAR(14) PRIMARY KEY,
				rollback_query TEXT,
		
				schema_before TEXT,
			   	schema_after TEXT
			);`)
		if err != nil {
			panic(fmt.Errorf("could not create migrations table: %w", err))
		}
	}

	var appliedMigrations []AppliedMigration
	err = pgxscan.Select(ctx, db, &appliedMigrations, `SELECT * FROM migrations ORDER BY id ASC;`)
	if err != nil {
		panic(fmt.Errorf("could not fetch applied migrations: %w", err))
	}

	// TODO: this should be some suffix but it's for later
	//  it also needs to give options to resolve conflicts, i.e. a migration that's not in code is in the database (abort / rollback + apply current table)
	//  or a migration that's in database goes further than what you've got in code (do nothing/rollback to current state)
	//  FOR NOW - JUST PANIC IF CODE IS NOT A SUPERSET OF DATABASE STATE

	for i, migration := range appliedMigrations {
		if migration.ID != migrations[i].ID {
			panic(fmt.Errorf("migration %s is in the database but not in the code, please resolve manually", migration.ID))
		}
	}
	migrationsToApply := migrations[len(appliedMigrations):]

	fmt.Printf("\n\n\n\n\n\n\n\nscan error: %v\nAPPLIED MIGRATIONS;%v\n\n\n\n", err, appliedMigrations)

	// verify if schema changes

	for _, migration := range migrationsToApply {
		tx, err := db.BeginTx(ctx, pgx.TxOptions{})
		if err != nil {
			panic(fmt.Errorf("could not start transaction: %w", err))
		}
		defer tx.Rollback(ctx)
		currentSchemaChecksum := md5Checksum(getSchemaJSON(ctx, tx))
		fmt.Println("schema checksum before migration:", currentSchemaChecksum)

		_, err = tx.Exec(ctx, migration.MigrationQuery)
		if err != nil {
			panic(fmt.Errorf("could not apply migration: %w", err))
		}
		newSchemaChecksum := md5Checksum(getSchemaJSON(ctx, tx))
		fmt.Println("schema checksum after migration:", newSchemaChecksum)

		if currentSchemaChecksum == newSchemaChecksum {
			panic(fmt.Errorf("migration %s does not change the schema, please resolve manually", migration.ID))
		}

		_, err = tx.Exec(ctx, migration.RollbackQuery)
		if err != nil {
			panic(fmt.Errorf("could not rollback migration: %w", err))
		}
		rollbackedSchemaChecksum := md5Checksum(getSchemaJSON(ctx, tx))
		fmt.Println("schema checksum after rollback:", rollbackedSchemaChecksum)

		fmt.Printf("checksums: %s %s %s\n", currentSchemaChecksum, newSchemaChecksum, rollbackedSchemaChecksum)

		if rollbackedSchemaChecksum != currentSchemaChecksum {
			panic(fmt.Errorf("migration %s does not rollback correctly, please resolve manually", migration.ID))
		}

		// all ok - apply migration
		_, err = tx.Exec(ctx, migration.MigrationQuery)
		if err != nil {
			panic(fmt.Errorf("could not apply migration: %w", err))
		}

		// insert migration to migrations table
		_, err = tx.Exec(ctx, `INSERT INTO migrations (id, rollback_query, schema_before, schema_after) VALUES ($1, $2, $3, $4);`,
			migration.ID, migration.RollbackQuery, currentSchemaChecksum, newSchemaChecksum)
		if err != nil {
			panic(fmt.Errorf("could not insert applied migration: %w", err))
		}

		err = tx.Commit(ctx)
		if err != nil {
			panic(fmt.Errorf("could not commit transaction: %w", err))
		}
	}
}

func md5Checksum(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

type tableField struct {
	OrdinalPosition        int     `json:"ordinal_position"`
	ColumnName             string  `json:"column_name"`
	DataType               string  `json:"data_type"`
	ColumnDefault          *string `json:"column_default"`
	IsNullable             *string `json:"is_nullable"`
	CharacterMaximumLength *int    `json:"character_maximum_length"`
	NumericPrecision       *int    `json:"numeric_precision"`
}

// pgx.Tx is an interface
func getSchemaJSON(ctx context.Context, db pgx.Tx) string {
	schema := map[string]interface{}{
		"version": "0.0.1",
	}

	var tables []string
	err := pgxscan.Select(ctx, db, &tables, listTables)
	if err != nil {
		panic(fmt.Errorf("could not list tables: %w", err))
	}

	tablesSchema := map[string]interface{}{}
	for _, table := range tables {
		var fields []tableField
		err := pgxscan.Select(ctx, db, &fields, listTableFields, table)
		if err != nil {
			panic(fmt.Errorf("could not list fields: %w", err))
		}
		tablesSchema[table] = fields
	}
	schema["tables"] = tablesSchema

	var views []string
	err = pgxscan.Select(ctx, db, &views, listViews)
	if err != nil {
		panic(fmt.Errorf("could not list views: %w", err))
	}

	schema["views"] = views
	b, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		panic(fmt.Errorf("could not marshal schema: %w", err))
	}

	return string(b)
}

const (
	// source: https://www.alberton.info/postgresql_meta_info.html
	listTables = `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_type = 'BASE TABLE'
			AND table_schema NOT IN ('pg_catalog', 'information_schema');`
	listTableFields = `
		  SELECT ordinal_position,
				 column_name,
				 data_type,
				 column_default,
				 is_nullable,
				 character_maximum_length,
				 numeric_precision
			FROM information_schema.columns
			WHERE table_name = $1
			ORDER BY ordinal_position;`
	// TODO: maybe we want to list view something to see if it changes?
	listViews = `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_type = 'VIEW'
			AND table_schema NOT IN ('pg_catalog', 'information_schema')
   			AND table_name !~ '^pg_';`

	listTableIndexes = `
		SELECT
			c.relname AS index_name
		FROM
			pg_class AS a
			JOIN pg_index AS b ON (a.oid = b.indrelid)
			JOIN pg_class AS c ON (c.oid = b.indexrelid)
		WHERE
			a.relname = 'test2';`

	listTableConstraints = `
		SELECT constraint_name, constraint_type
		FROM information_schema.table_constraints
		WHERE table_name = 'test2';`
	detailedConstraintInfo = `
		SELECT tc.constraint_name,
			tc.constraint_type,
			tc.table_name,
			kcu.column_name,
			tc.is_deferrable,
			tc.initially_deferred,
			rc.match_option AS match_type,
			rc.update_rule AS on_update,
			rc.delete_rule AS on_delete,
			ccu.table_name AS references_table,
			ccu.column_name AS references_field
    	FROM information_schema.table_constraints tc
		LEFT JOIN information_schema.key_column_usage kcu
       		ON tc.constraint_catalog = kcu.constraint_catalog
      		AND tc.constraint_schema = kcu.constraint_schema
      		AND tc.constraint_name = kcu.constraint_name
		LEFT JOIN information_schema.referential_constraints rc
		   	ON tc.constraint_catalog = rc.constraint_catalog
		  	AND tc.constraint_schema = rc.constraint_schema
		  	AND tc.constraint_name = rc.constraint_name
		LEFT JOIN information_schema.constraint_column_usage ccu
		   	ON rc.unique_constraint_catalog = ccu.constraint_catalog
		  	AND rc.unique_constraint_schema = ccu.constraint_schema
		  	AND rc.unique_constraint_name = ccu.constraint_name
			WHERE tc.table_name = 'testconstraints2'
		  	AND tc.constraint_name = 'testconstraints_id_fk';`
)
