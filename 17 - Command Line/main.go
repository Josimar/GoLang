package main

import (
	"fmt"
	"app/commandline"
	"log"
	"os"
)

func main(){
	fmt.Println("Start point")

	aplicacao := app.Gerar()

	erro := aplicacao.Run(os.Args)
	if erro != nil{
		log.Fatal(erro)
	}
}