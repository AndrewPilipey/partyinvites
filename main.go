package main

import "fmt"

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
} // создание структуры - определение пользовательских типов данных

var responses = make([]*Rsvp, 0, 10) //определение среза для ответов гостей (без фикс.длины, т.к. неизвестно сколько будет гостей)

func main() {
	fmt.Println("TODO: add some features")
}
