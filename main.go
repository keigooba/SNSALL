package main

import (
	"log"
	"snsall/app/controllers"
)

func main() {
	controllers.CmdFlag()
	log.Println(controllers.StartWebServer())
}
