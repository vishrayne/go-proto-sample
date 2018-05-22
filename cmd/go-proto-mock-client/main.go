package main

import (
	"bytes"
	"log"
	"net/http"

	protocheck "github.com/vishrayne/go-proto-sample"
)

func main() {
	book := &protocheck.AddressBook{}
	book.People = append(book.People, protocheck.CreateMockPerson("John Doe", "jdoe@example.com"))
	book.People = append(book.People, protocheck.CreateMockPerson("Jane Doe", "jdoe@example.com"))

	buf := protocheck.ToProto(book)
	_, err := http.Post("http://localhost:3000", "", bytes.NewBuffer(buf))

	if err != nil {
		log.Fatalln(err)
	}
}
