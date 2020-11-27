package main

import (
	"github.com/zacfeng/dcard-take-home-test/routers"
)

func main() {
	router := routers.SetupRouter()
	router.Run()

}
