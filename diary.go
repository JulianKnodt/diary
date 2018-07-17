package main

import (
	"diary/entry"
	create "diary/init"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := os.Args[1:]
	switch flag.Arg(0) {
	case "init":
		create.Init(args)
	case "entry":
		entry.Entry(args)
	case "edit":
		// edit()
	case "search":
		// search()
	case "help":
		fallthrough
	default:
		fmt.Printf(`
Error: Unknown command %s
      
Known Commands:
  entry: Open an editor to create an entry
  edit: Edit an entry
  search: Search through all entries
`,
			flag.Arg(1),
		)
	}

}
