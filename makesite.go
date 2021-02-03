package main

import (
	"io/ioutil"
	"os"
	"html/template"
	"flag"
	"fmt"
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


func writeTemplate(fileContent string, templateFile string, newFile string) {
	tmpl := template.Must(template.New(templateFile).ParseFiles(templateFile))
	content := Data{Content: string(fileContent)}
	
	f, err := os.Create(newFile)
	check(err)
	defer f.Close()

	err = tmpl.Execute(os.Stdout, content)
	check(err)

	f.Close()
}


func main() {
	writeTemplate(readFile("first-post.txt"), "template.tmpl", "first-post.html")
	
	filePtr := flag.String("file", "defaultValue", "File to HTML.")
	flag.Parse()
	save(filePtr)
	fmt.Println("file:", *filePtr)
}