package main

import "contacts_service/startup"

func main() {
	server := startup.NewServer()
	server.Start()
}
