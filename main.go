package main

import (
	"tugas-acp/configs"	
	"tugas-acp/routes"
)


func main() {
	configs.InitDB()

	e := routes.New()
	e.Start(":8000")
}