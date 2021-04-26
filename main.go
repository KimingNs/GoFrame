package main

import (
	_ "GoGF/boot"
	_ "GoGF/router"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/crypto/gaes"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
)

func main() {
	//testJson()
	//testMysql()
	s := g.Server()
	s.SetPort(8199)
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

func testMysql() {
	m := gmap.New()
	m.Sets(g.MapAnyAny{
		"partner_token": "e1e0e7e2dceaea871f5a436b3770fa1f",
		"params_data":   "rBBjCeXwWPp9gm1Qk/BPUSmkmK0fX+BI0qjoFAFI9Nodg2yWxyzp1NDE0RWw7QjXocLormzB8E+7mKtnqe5fkpCizguFketcOMlN+/K3qew5LA9STBvlgf8BtRFtiC81",
	})
	request_json, _ := json.Marshal(m)
	j, _ := gjson.DecodeToJson(request_json)
	//查询默认配置
	//partner_info, err := g.DB().Model("t_partner_info").Where("partner_token = ?", j.GetString("partner_token")).One()
	//查询cogito_tool配置
	partner_info, err := g.DB("cogito_tool").Model("t_partner_info").Where("partner_token = ?", j.GetString("partner_token")).One()

	if err != nil {
		return
	}
	aes_key := partner_info.GMap().GetVar("aes_key")
	aes_iv := partner_info.GMap().GetVar("aes_iv")
	//partner_id := partner_info.GMap().GetVar("partner_id")

	//插入数据保存
	//_, err = g.DB().Model("user").Data(model.User{Password: aes_key.String(), Nickname: aes_iv.String()}).Insert()
	//if err != nil {
	//	return
	//}

	//加密解密
	//str := gjson.New(nil)
	//str.Set("user_name", "zhangtianming")
	//str.Set("user_mobile", "15107152048")
	//str.Set("user_idcard", "420105199810050410")
	//encrypt, err := gaes.EncryptCBC([]byte(str.MustToJsonString()), aes_key.Bytes(), aes_iv.Bytes())
	//if err != nil {
	//	return
	//}
	//hex.EncodeToString(encrypt)
	decrypt, err := gaes.DecryptCBC([]byte(j.GetString("params_data")), aes_key.Bytes(), aes_iv.Bytes())
	fmt.Println(string(decrypt))

}
