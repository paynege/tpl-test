package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type test struct {
	Name string
	Value string
}

func main() {
	t := test{"Hello", "World!"}
	tpl, err := template.New("test.tpl").ParseFiles("test.tpl")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, t)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tpl)
}