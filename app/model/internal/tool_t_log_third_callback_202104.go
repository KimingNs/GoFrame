// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// ToolTLogThirdCallback202104 is the golang structure for table t_log_third_callback_202104.
type ToolTLogThirdCallback202104 struct {
	Id             int         `orm:"id,primary"      json:"id"`             //
	AppPackage     string      `orm:"app_package"     json:"appPackage"`     //
	PartnerId      string      `orm:"partner_id"      json:"partnerId"`      //
	PartnerIndex   string      `orm:"partner_index"   json:"partnerIndex"`   //
	ServiceId      string      `orm:"service_id"      json:"serviceId"`      //
	ProviderId     string      `orm:"provider_id"     json:"providerId"`     //
	RequestEncode  string      `orm:"request_encode"  json:"requestEncode"`  //
	RequestDecode  string      `orm:"request_decode"  json:"requestDecode"`  //
	ResponseEncode string      `orm:"response_encode" json:"responseEncode"` //
	ResponseDecode string      `orm:"response_decode" json:"responseDecode"` //
	Cdate          *gtime.Time `orm:"cdate"           json:"cdate"`          //
	Ctime          *gtime.Time `orm:"ctime"           json:"ctime"`          //
	Time           string      `orm:"time"            json:"time"`           //
	ResponseCode   int         `orm:"response_code"   json:"responseCode"`   //
}