package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func IsInDiary() bool {
	for {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		if dir == "/" {
			return false
		}
		fmt.Println(filepath.Join(dir, ".diary"))
		f, err := os.Stat(filepath.Join(dir, ".diary"))
		if f != nil && f.IsDir() {
			return true
		}
		if err = os.Chdir(".."); err != nil {
			log.Fatal(err)
		}
	}
}
