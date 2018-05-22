package main

import (
	"io/ioutil"
	"log"
	"net/http"

	protocheck "github.com/vishrayne/go-proto-sample"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Incoming request ---")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln("Unable to read body: ", err)
		}

		book := protocheck.FromProto(data)
		log.Println("book: ", book)
	})

	log.Println("Mock server listening at port 3000. ^C to quit")
	http.ListenAndServe(":3000", nil)
}
