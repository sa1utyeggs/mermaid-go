package main

import (
	"github.com/sa1utyeggs/mermaid-go/mermaid"
	"github.com/spf13/afero"
	"log"
	"testing"
)

func TestY(t *testing.T) {
	output := "output.svg"
	filename := "demo/flowchart.mmd"
	fs := afero.NewOsFs()
	file, err := afero.ReadFile(fs, filename)
	fs = afero.NewMemMapFs()
	if err != nil {
		log.Fatal("Error: reading input file")
	}
	result := mermaid.Execute(string(file))

	outfile, err := fs.Create(output)
	if err != nil {
		log.Fatal("Error: creating output file")
	}
	outfile.Write([]byte(result))
}
