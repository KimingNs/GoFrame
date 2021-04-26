package demo

import (
	"encoding/hex"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/crypto/gaes"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"time"
)

var loc, _ = time.LoadLocation("Asia/Jakarta")

type ThirdApiRequest struct {
	PartnerToken string `json:"partner_token"`
	ParamsData   string `json:"params_data"`
}

type ThirdApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    g.Map  `json:"data"`
}

func (c *ThirdApiRequest) SelfCheck(r *ghttp.Request) {
	request_json, _ := gjson.DecodeToJson(r.GetBodyString())
	//partner_info, err := g.DB("cogito_tool").Model("t_partner_info").Fields("aes_key,aes_iv,partner_id").Where("partner_token=", request_json.GetString("partner_token")).One()
	partner_info, err := g.DB("cogito_tool").Model("t_partner_info").Fields("aes_key,aes_iv,partner_id").FindOne("partner_token=", request_json.GetString("partner_token"))
	if err != nil {
		return
	}
	aes_key := partner_info.GMap().GetVar("aes_key")
	aes_iv := partner_info.GMap().GetVar("aes_iv")
	//partner_id := partner_info.GMap().GetVar("partner_id")

	decodeString, err := hex.DecodeString(request_json.GetString("params_data"))
	if err != nil {
		return
	}
	params_data, _ := gaes.DecryptCBC(decodeString, aes_key.Bytes(), aes_iv.Bytes())
	json, err := gjson.DecodeToJson(params_data)
	if err != nil {
		return
	}
	//r.Response.WriteJson(json)
	m := gmap.New()
	m.Sets(g.MapAnyAny{
		"user_name":   json.GetString("user_name"),
		"user_mobile": json.GetString("user_mobile"),
		"user_idcard": json.GetString("user_idcard"),
		"cdate":       time.Now().In(loc).Format("2006-01-02"),
		"ctime":       time.Now().In(loc).Format("2006-01-02 15:04:05"),
	})
	_, err = g.DB("cogito_tool").Model("t_local_cache_blacklist").Data(m).Save()
	if err != nil {
		return
	}
	r.Response.WriteJson(ThirdApiResponse{Code: 0, Message: "success", Data: nil})

}
