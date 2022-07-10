package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main(){
	fmt.Printf("Api Running")
	router := router.Build()
	port := "3000"
	log.Fatal(http.ListenAndServe(":"+port, router))
}