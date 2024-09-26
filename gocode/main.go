package main

import (
	"fmt"

	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

//This md2html executable must run in the repo folder where "templates" reside, not inside gocode.
//Ex.: ./gocode/md2html

func main() {

	inputNames := []string{"."}

	// Create empty lists
	inputFiles := make([]string, len(inputNames))
	templateFiles := make([]string, len(inputNames))
	outputFiles := make([]string, len(inputNames))

	// Fill them
	for i, name := range inputNames {
		
		inputFiles[i] = name + "/" + "index.md"
		templateFiles[i] = name + "/" + "index_template.html"
		outputFiles[i] = name + "/" + "index.html"
	}

	// Convert each Markdown file to HTML and save to the corresponding output file
	for i, inputFile := range inputFiles {

		fmt.Printf("Processing file %s...", inputFile)

		// Read in the Markdown file
		md, err := os.ReadFile(inputFile)
		check(err)

		// Convert the Markdown to HTML
		//md = []byte(md)

		// Create a parser
		extensions := parser.CommonExtensions | parser.Footnotes
		parser := parser.NewWithExtensions(extensions)

		html := markdown.ToHTML(md, parser, nil)

		// Read in the template HTML file
		template, err := os.ReadFile(templateFiles[i])
		check(err)

		// Replace the {src} placeholder with the HTML
		output := strings.Replace(string(template), "{src0}", string(html), -1)

		// Write the output to the specified file
		err = os.WriteFile(outputFiles[i], []byte(output), 0644)
		check(err)
		
		fmt.Printf("Done.\n")
	}
}
