package main

import (
	_ "GoGF/boot"
	_ "GoGF/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	s := g.Server()
	s.Run()

}
