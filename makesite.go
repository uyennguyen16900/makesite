package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
)

// Data is
type Data struct {
	Content string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(filename string) string {
	fileContent, err := ioutil.ReadFile(filename)
	check(err)
	return string(fileContent)
}

func renderTemplate(fileContent string, templateFile string, file string) {
	tmpl := template.Must(template.New(templateFile).ParseFiles(templateFile))
	content := Data{Content: string(fileContent)}

	filename := strings.Split(file, ".txt")[0]
	f, err := os.Create(filename + ".html")
	check(err)
	defer f.Close()

	err = tmpl.Execute(f, content)
	check(err)
}

func getFiles(directory string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".txt") {
			fmt.Println(file.Name())
		}
	}
}

func convertMdtoHTML(file string) {
	content, err := ioutil.ReadFile(file)
	check(err)

	html := markdown.ToHTML(content, nil, nil)
	// fmt.Printf(string(html))

	fileOut := strings.Split(file, ".md")[0] + ".html"
	writeError := ioutil.WriteFile(fileOut, html, 0644)
	if writeError != nil {
		log.Fatalf("Could not write to %s", fileOut)
	}
}

func main() {
	var fileName string
	flag.StringVar(&fileName, "file", "defaultValue", "Text file name to turn to HTML page.")
	flag.Parse()
	// save(filePtr)

	if fileName == "" {
		panic("Missing file name!")
	}
	renderTemplate(readFile(fileName), "template.tmpl", fileName)

	var directory string
	flag.StringVar(&directory, "dir", "", "Find all .txt files in the given directory.")
	flag.Parse()

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".txt") {
			fmt.Println(file.Name())
			renderTemplate(readFile(file.Name()), "template.tmpl", file.Name())
		}
	}

	convertMdtoHTML("README.md")

}
