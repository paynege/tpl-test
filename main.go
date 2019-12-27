package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type test struct {
	Name  string
	Value string
}

func main() {
	t := test{"Hello", "World!"}
	s := `192.168.1.100,
192.168.1.105,
192.168.1.102,
192.168.1.103`
	t.Value = s
	tpl, err := template.New("test.tpl").ParseFiles("test.tpl")
	if err != nil {
		log.Fatalln(err)
	}

	tpFile, err := os.Create("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(tpFile, t)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tpl)
}
