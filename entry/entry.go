package entry

import (
	"bytes"
	"diary/utils"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func firstLine(s []byte) []byte {
	if index := bytes.IndexRune(s, '\n'); index != -1 {
		return s[:index]
	}
	return s
}

func Entry(args []string) {
	if !utils.IsInDiary() {
		panic("Not in diary directory")
	}
	flagSet := flag.NewFlagSet("entry", flag.PanicOnError)
	name := flagSet.String("name", "", "The name of this diary entry")
	f, err := ioutil.TempFile(".", "diary-entry")
	check(err)

	cmd := exec.Command("vim", f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	check(cmd.Run())

	entry, err := ioutil.ReadFile(f.Name())
	check(err)

	if len(*name) == 0 {
		*name = string(firstLine(entry))
	}
	fmt.Println(*name)
	fmt.Println(string(entry))

	check(os.Remove(f.Name()))
}
