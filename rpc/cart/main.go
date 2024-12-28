package main

import (
	cart "eshop/kitex_gen/cart/cartservice"
	"log"
)

func main() {
	svr := cart.NewServer(new(CartServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
