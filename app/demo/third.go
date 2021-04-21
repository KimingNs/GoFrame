package demo

import "github.com/gogf/gf/frame/g"

type thirdApi struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    g.Map  `json:"data"`
}
