package prototest

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
)

func Hello() {
	fmt.Println("Ping!")
}

func TestProto(fname string) {
	// serialize
	book := &AddressBook{}
	p := CreateMockPerson("John Doe", "jdoe@example.com")
	book.People = append(book.People, p)
	fmt.Println("book: ", book)

	out := ToProto(book)

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write addressbook: ", err)
	}

	// deserialize
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file: ", err)
	}

	book2 := FromProto(in)
	fmt.Println("book2: ", book2)
}

func CreateMockPerson(name string, email string) *Person {
	return &Person{
		Id:    1234,
		Name:  name,
		Email: email,
		Phones: []*Person_PhoneNumber{
			{Number: "555-4321", Type: Person_HOME},
		},
	}
}

func ToProto(ab *AddressBook) []byte {
	out, err := proto.Marshal(ab)
	if err != nil {
		log.Fatalln("failed to encode address book: ", err)
		return nil
	}

	return out
}

func FromProto(buf []byte) *AddressBook {
	book := &AddressBook{}
	if err := proto.Unmarshal(buf, book); err != nil {
		log.Fatalln("Failed to parse address book: ", err)
		return nil
	}

	return book
}
