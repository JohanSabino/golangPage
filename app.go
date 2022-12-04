package main

import (
	"fmt"
	"log"
	"net/http"
)

type Users struct{
	Name string
	Skills string
	Age int
}


func main() {
	
	fs:= http.FileServer(http.Dir("../GoPage"))
	http.Handle("/", fs)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}

	
	//Creating server
	fmt.Println("El servidor está corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000/")
	http.ListenAndServe("localhost:3000", nil)

}


