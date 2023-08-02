package main

import "example/web-service-gin/api"

func main() {
	s := api.NewServer()

	s.Start()
}
