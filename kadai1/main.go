package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var directory string

	flags := flag.NewFlagSet("imgconv", flag.ContinueOnError)

	flags.StringVar(&directory, "directory", "", "")
	flags.StringVar(&directory, "d", "", "")

	if err := flags.Parse(os.Args[1:]); err != nil {

	}

	fmt.Println(directory)
}
