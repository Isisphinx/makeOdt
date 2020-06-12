package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func copy(src, dst string) error {

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	if _, err := io.Copy(destination, source); err != nil {
		return err
	}
	return nil
}

func main() {
	open := flag.Bool("open", false, "Open the file after creation")

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	flag.Parse()
	args := flag.Args()
	xmlFile := args[0]
	template := args[1]
	outputFile := args[2]
	zip := exec.Command(dir+"\\7za", "a", template, xmlFile)
	if err := zip.Run(); err != nil {
		panic(err)
	}
	if err := copy(template, outputFile); err != nil {
		panic(err)
	}
	if *open {
		cmd := exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", outputFile)
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}
}
