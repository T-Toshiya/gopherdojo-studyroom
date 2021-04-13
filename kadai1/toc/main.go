package main

import (
	"flag"
	"fmt"
	"github.com/T-Toshiya/gopherdojo-studyroom/kadai1/converter"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	var directory string
	var from string
	var to string

	flags := flag.NewFlagSet("imgconv", flag.ContinueOnError)

	flags.StringVar(&directory, "directory", "", "please specify directory")
	flags.StringVar(&directory, "d", "", "please specify directory")

	flags.StringVar(&from, "from", "jpg", "please specify before conversion format")
	flags.StringVar(&from, "f", "jpg", "please specify before conversion format")

	flags.StringVar(&to, "to", "png", "please specify after conversion format")
	flags.StringVar(&to, "t", "png", "please specify after conversion format")

	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	if directory == "" {
		log.Fatal("You need to specify directory")
	}

	if from == to {
		log.Fatal("You cannot specify same format")
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
		if !file.IsDir() && path.Ext(file.Name()) == "."+from {
			paths = append(paths, file.Name())
		}
	}

	c := &converter.Converter{
		BeforeFmt: from,
		AfterFmt:  to,
		Directory: directory,
	}

	for _, filepath := range paths {
		c.FilePath = filepath
		err := c.Convert()
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Convert success")
}
