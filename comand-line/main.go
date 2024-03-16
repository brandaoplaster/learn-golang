package main

import (
	"comand-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("start app...")
	application := app.Gerar()
	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
