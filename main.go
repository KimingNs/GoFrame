package main

import (
	_ "GoGF/boot"
	_ "GoGF/router"
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
)

func main() {
	//testJson()
	s := g.Server()
	s.Run()

}

func testJson() {
	j := gjson.New(nil)
	j.Set("code", 200)
	j.Set("message", "success")
	j.Set("data.name", "ztm")
	j.Set("data.age", 20)
	fmt.Println(j.MustToJsonString())

	if e, err := gjson.DecodeToJson(j.MustToJsonString()); err != nil {
		fmt.Println(err)
	} else {
		e.Set("code", 0)
		e.Set("data.mobile", "15107152048")
		fmt.Println(e.MustToJsonString())
	}

}
