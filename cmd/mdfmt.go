package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	write = flag.Bool("w", false, "write the result to the file")
	diff  = flag.Bool("w", false, "print the changes")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: mdfmt [flags] [path ...]\n")
	flag.PrintDefaults()
}

func isMarkdownFile(f os.FileInfo) bool {
	name := f.Name()
	return !f.IsDir() && !strings.HasPrefix(name, ".") && (strings.HasSuffix(name, ".md") || strings.HasSuffix(name, ".markdown"))
}

func visitFile(path string, f os.FileInfo, err error) error {
	if err == nil && isMarkdownFile(f) {
		err = processFile(path, nil, os.Stdout, false)
	}
	return err
}

func walkDir(path string) error {
	return filepath.Walk(path, visitFile)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		if err := processFile("<standard input>", os.Stdin, os.Stdout, true); err != nil {
			log.Fatalln(err)
		}
		return
	}
	for i := 0; i < flag.NArg(); i++ {
		path := flag.Arg(i)
		switch dir, err := os.Stat(path); {
		case err != nil:
			log.Fatalln(err)
		case dir.IsDir():
			if err := walkDir(path); err != nil {
				log.Fatalln(err)
			}
		default:
			if err := processFile(path, nil, os.Stdout, false); err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func processFile(fileName string, in io.Reader, out io.Writer, stdin bool) error {
	return nil
}
