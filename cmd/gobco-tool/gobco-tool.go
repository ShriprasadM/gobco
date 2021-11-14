package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/junhwi/gobco/instrument"
)

func getFd(out string) (*os.File, error) {
	if out == "" {
		return os.Stdout, nil
	} else {
		return os.Create(out)
	}
}

func runGobco() {

	cmd := flag.NewFlagSet("gobco", flag.ExitOnError)
	// Register all flags same as go tool cover
	outPtr := cmd.String("o", "", "file for output; default: stdout")
	version := cmd.String("V", "", "print version and exit")
	cmd.String("mode", "", "coverage mode: set, count, atomic")
	coverVar := cmd.String("var", "Cov", "name of coverage variable to generate (default \"Cov\")")
	cmd.Parse(os.Args[2:])
	files := cmd.Args()
	// files := []string{"./shri_test/module1/module1.go"}
	// files := []string{"./..."} // iterate over all packages
	if *version != "" {
		fmt.Println("cover version go1.13.1")
	} else {
		filepath.Walk(files[0], func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			fd, err := getFd(*outPtr)
			err = instrument.Instrument(path, fd, *coverVar)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
			return err
		})
		os.Exit(0)
	}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gobco: ")

	// tool := ""
	tool := os.Args[1]
	args := os.Args[2:]
	// tool := "/Users/shriprasad/go/bin/gobco"
	// args := []string{"./shri_test/module1/module1.go"}

	toolName := path.Base(tool)
	// toolName := "cover"
	if toolName == "cover" {
		runGobco()
	} else {
		cmd := exec.Command(tool, args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
	os.Exit(0)
}
