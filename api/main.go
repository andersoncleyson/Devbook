package main

import (
	"fmt"
	"api/src/router"
	"api/src/config"
	"log"
	"net/http"
)

func main(){
	config.Carregar()

	r := router.Gerar()

	fmt.Printf("Server is running in port %d", config.Porta)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}