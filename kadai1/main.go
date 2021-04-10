package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	var directory string

	flags := flag.NewFlagSet("imgconv", flag.ContinueOnError)

	flags.StringVar(&directory, "directory", "", "")
	flags.StringVar(&directory, "d", "", "")

	if err := flags.Parse(os.Args[1:]); err != nil {

	}

	if f, err := os.Stat(directory); os.IsNotExist(err) || !f.IsDir() {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if !file.IsDir() && path.Ext(file.Name()) == ".jpg" {
			paths = append(paths, file.Name())
		}
	}

	for _, filepath := range paths {
		fmt.Println(filepath)
	}
}
