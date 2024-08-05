package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oklog/ulid/v2"
	"github.com/rs/cors"
	"github.com/urfave/cli/v2"

	"github.com/wombatDaiquiri/wombat/wombat/blog"
	database2 "github.com/wombatDaiquiri/wombat/wombat/database"
	"github.com/wombatDaiquiri/wombat/ww"
)

type tracer struct {
}

func (t tracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	fmt.Printf("\n~~~> QUERY\n---------------\n%v\n---------------\n%+#v\n~~~", data.SQL, data.Args)
	return ctx
}

func (t tracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	// pass
}

func main() {
	fmt.Println("Hello, World!")

	// NEXT TODO: MAKE IT LOOK ARTISTIC XD
	// NEXT TODO: SUPPORT SAVING DRAFTS
	// NEXT TODO: SUPPORT SCHEDULING POSTS
	// NEXT TODO: WRITE 5 POSTS?

	app := initializeCLIApp()
	err := app.Run(os.Args)
	if err != nil {
		panic(fmt.Errorf("run application: %w", err))
	}
}

func initializeCLIApp() *cli.App {
	app := &cli.App{
		Name:  "wombat",
		Usage: "setup & run the http server hosting daiquiri.fyi",
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "run the http server",
				Action: func(ctx *cli.Context) error {
					dbURL := os.Getenv("DATABASE_URL")
					fmt.Printf("\n\n\n\nDBURL: %v\n\n\n", dbURL)
					poolConfig, err := pgxpool.ParseConfig(dbURL)
					if err != nil {
						panic(fmt.Errorf("parse connection config: %w", err))
					}
					poolConfig.ConnConfig.Tracer = tracer{}

					db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
					if err != nil {
						panic(fmt.Errorf("could not connect to database: %w", err))
					}
					defer db.Close()

					router := chi.NewRouter()
					router.Use(middleware.Logger)

					router.Get("/", func(w http.ResponseWriter, r *http.Request) {
						indexFile, err := os.Open("index.html")
						if err != nil {
							http.Error(w, "could not read index.html", http.StatusInternalServerError)
							panic(err)
							return
						}
						defer indexFile.Close()

						_, err = io.Copy(w, indexFile)
						if err != nil {
							http.Error(w, "could not write index.html", http.StatusInternalServerError)
							panic(err)
							return
						}
					})
					router.Post("/post", func(w http.ResponseWriter, r *http.Request) {
						body, err := io.ReadAll(r.Body)
						if err != nil {
							http.Error(w, "could not read request body", http.StatusInternalServerError)
							panic(err)
							return
						}

						var blogpost database2.Blogpost
						err = json.Unmarshal(body, &blogpost)
						if err != nil {
							http.Error(w, "could not unmarshal request body", http.StatusInternalServerError)
							panic(err)
							return
						}

						blogpostULID := ulid.Make().String()
						_, err = db.Exec(context.Background(), `
							INSERT INTO blogposts
							    (ulid, title, slug, lead, content_markdown, publication_date)
							VALUES ($1, $2, $3, $4, $5, $6)`,
							blogpostULID, blogpost.Title, blogpost.Slug, blogpost.Lead, blogpost.ContentMarkdown, ww.ValueOrNil(blogpost.PublicationDate))
						if err != nil {
							panic(err)
							http.Error(w, "could not insert blogpost", http.StatusInternalServerError)
							panic(err)
							return
						}

						var insertedBlogpost database2.Blogpost
						err = pgxscan.Get(context.Background(), db, &insertedBlogpost, `
							SELECT *	
							FROM blogposts	
							WHERE ulid = $1`, blogpostULID)
						if err != nil {
							panic(err)
							http.Error(w, "could not fetch inserted blogpost", http.StatusInternalServerError)
							panic(err)
							return
						}

						returnedBlogpost := blog.NewPost(insertedBlogpost)
						insertedBlogpostJSON, err := json.MarshalIndent(returnedBlogpost, "", "\t")
						if err != nil {
							http.Error(w, "could not marshal inserted blogpost", http.StatusInternalServerError)
							panic(err)
							return
						}

						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(http.StatusCreated)
						_, err = w.Write(insertedBlogpostJSON)
						if err != nil {
							http.Error(w, "could not write inserted blogpost", http.StatusInternalServerError)
							panic(err)
							return
						}
					})

					router.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
						var dbBlogposts []database2.Blogpost
						err := pgxscan.Select(context.Background(), db, &dbBlogposts, `
							SELECT *
							FROM blogposts	
							WHERE publication_date IS NOT NULL AND publication_date <= NOW()
							ORDER BY publication_date DESC`)
						if err != nil {
							http.Error(w, "could not fetch blogposts", http.StatusInternalServerError)
							panic(err)
							return
						}

						blogpostsJSON, err := json.MarshalIndent(blog.NewPosts(dbBlogposts), "", "\t")
						if err != nil {
							http.Error(w, "could not marshal blogposts", http.StatusInternalServerError)
							panic(err)
							return
						}

						w.Header().Set("Content-Type", "application/json")
						_, err = w.Write(blogpostsJSON)
						if err != nil {
							http.Error(w, "could not write blogposts", http.StatusInternalServerError)
							panic(err)
							return
						}
					})

					return http.ListenAndServe(":2137", cors.Default().Handler(router))
				},
			},
			{
				Name:  "migrate",
				Usage: "run migrations on the database, preparing it for running the HTTP server",
				Action: func(ctx *cli.Context) error {
					dbURL := os.Getenv("DATABASE_URL")
					fmt.Printf("\n\n\n\nDBURL: %v\n\n\n", dbURL)
					poolConfig, err := pgxpool.ParseConfig(dbURL)
					if err != nil {
						panic(fmt.Errorf("parse connection config: %w", err))
					}
					poolConfig.ConnConfig.Tracer = tracer{}

					db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
					if err != nil {
						panic(fmt.Errorf("could not connect to database: %w", err))
					}
					defer db.Close()

					database2.Migrate(ctx.Context, db)
					return nil
				},
			},
		},
	}

	return app
}
