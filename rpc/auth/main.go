package main

import (
	auth "eshop/kitex_gen/auth/authservice"
	"log"
)

func main() {
	svr := auth.NewServer(new(AuthServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}