package ww

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"strings"
)

func prettyPrintPrefix() {
	strings.Repeat("\n", 5)
	strings.Repeat("-", 80)
	fmt.Printf("\n%s\n", debug.Stack())
	strings.Repeat("~", 40)
	fmt.Print("\n")
}

func prettyPrintSuffix() {
	fmt.Print("\n")
	strings.Repeat("-", 80)
	strings.Repeat("\n", 5)
}

func PJSON(v any) {
	prettyPrintPrefix()
	defer prettyPrintSuffix()

	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Printf("unexpected marshalling error (type=%T value=%+#v): %v\n", v, v, err)
		return
	}
	fmt.Print(string(b))
}

func P(format string, args ...interface{}) {
	prettyPrintPrefix()
	defer prettyPrintSuffix()

	fmt.Printf(format, args...)
}
