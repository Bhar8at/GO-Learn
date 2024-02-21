package main

import (
	"fmt"
	"os"
)

// Learning Web-dev here

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {

	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600) // Here 0600 indicates read and write permission for current user only

}

func LoadPage(title string) (*Page, error) {

	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func file1() {

	p1 := &Page{
		Title: "newpage",
		Body:  []byte("this is a sample page"),
	}
	p1.save()

	p2, _ := LoadPage("newpage")
	fmt.Println(string(p2.Body))

	// How to build the exec file

	/*
		$ go build wiki.go
		$ ./wiki
		This is a sample Page.
	*/

}
