package main

import (
	"html/template"
	"io/ioutil"
	"os"
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

func renderTemplate(fileContent string, templateFile string, newFile string) {
	tmpl := template.Must(template.New(templateFile).ParseFiles(templateFile))
	content := Data{Content: string(fileContent)}

	f, err := os.Create(newFile)
	check(err)
	defer f.Close()

	err = tmpl.Execute(f, content)
	check(err)

	f.Close()
}

func main() {
	renderTemplate(readFile("first-post.txt"), "template.tmpl", "first-post.html")

	filePtr := flag.StringVar(...., "file", "defaultValue", "Text file name to turn to HTML page.")
	flag.Parse()
	save(filePtr)
	
	
	fmt.Println("file:", *filePtr)
}
