package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main(){
	config.LoadDotEnv()
	fmt.Println(config.StringConnection)
	fmt.Printf("Api Running %d", config.Port)
	router := router.Build()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}