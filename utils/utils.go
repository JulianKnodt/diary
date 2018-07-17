package utils

import (
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
		f, err := os.Stat(filepath.Join(dir, ".diary"))
		if f != nil && f.IsDir() {
			return true
		}
		if err = os.Chdir(".."); err != nil {
			log.Fatal(err)
		}
	}
}

func DiaryPath() string {
	for {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		if dir == "/" {
			return ""
		}
		f, err := os.Stat(filepath.Join(dir, ".diary"))
		if f != nil && f.IsDir() {
			return filepath.Join(dir, ".diary")
		}
		if err = os.Chdir(".."); err != nil {
			log.Fatal(err)
		}
	}

}
