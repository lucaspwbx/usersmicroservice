package main

import "github.com/lucasweiblen/usersmicroservice/service"

func main() {
	service := service.UserService{}
	service.Run()
}
