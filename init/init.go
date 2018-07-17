package init

import (
	"flag"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/*
Initializes a diary in the given folder
*/
func Init(args []string) {
	flagSet := flag.NewFlagSet("init", flag.PanicOnError)
	directory := flagSet.String("path", ".", "The path to initialize the directory in")
	check(flagSet.Parse(args))

	check(os.Chdir(*directory))
	check(os.Mkdir(".diary", os.ModeDir|os.ModePerm))
}
