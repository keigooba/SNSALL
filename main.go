package main

import (
	"log"
	"snsall/app/controllers"
)

func main() {
	CmdFlag()
	log.Println(controllers.StartWebServer())
}
