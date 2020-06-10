package main

import (
	"flag"
	"io"
	"os"
	"os/exec"
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
	flag.Parse()
	args := flag.Args()
	xmlFile := args[0]
	template := args[1]
	outputFile := args[2]
	zip := exec.Command("7za", "a", template, xmlFile)
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
