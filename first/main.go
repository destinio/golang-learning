package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Person struct {
	name string
	id int
	email string
}

	
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func newPerson(name, email string) {
	p := Person{name: name, email: email, id: 123}

	fmt.Println(p.name)
}

func getData() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
	
}

func main() {
	newPerson("destin", "dleeinc@gmail.com")
	getData()
}