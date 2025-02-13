package main

import (
	"html/template"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
} // создание структуры - определение пользовательских типов данных

var responses = make([]*Rsvp, 0, 10) //определение среза для ответов гостей (без фикс.длины, т.к. неизвестно сколько будет гостей)
var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	// TODO - load templates here
}

func main() {
	loadTemplates()
}
