/**
 * Created by Michael on 2015/8/4.
 */
package utils

import (
	"encoding/json"
	"testing"
)

func Test_copy(t *testing.T) {
	a := &AA{A: 999}
	b := &AA{}
	err := Clone(b, a)
	t.Log(a, b, err)
}

func TestPWD(t *testing.T) {
	t.Log(AalidataPwd("dolo0425"))

}

func TestPhone(t *testing.T) {
	t.Log(PhoneRegexp("8601593533372"))
}

type AA struct {
	CC
	A int `json:"a"`
}

type BB interface {
	Decode(b *[]byte) error
	Encode() (*[]byte, error)
}

type CC struct{}

func (this *CC) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CC) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

/*
{"code":0,"data":{
	"country":"\u4e2d\u56fd",//国家
	"country_id":"CN",
	"area":"\u534e\u5357",//区域
	"area_id":"800000",
	"region":"\u5e7f\u4e1c\u7701",//所在省
	"region_id":"440000",
	"city":"\u6df1\u5733\u5e02",//所在市
	"city_id":"440300",
	"county":"\u5b9d\u5b89\u533a",//所在县
	"county_id":"440306",
	"isp":"\u7535\u4fe1",
	"isp_id":"100017",
	"ip":"119.137.54.77"}}
*/

func TestIPAddr(t *testing.T) {
	ip, err := GetInternalIP()
	t.Log("ip -> ", ip, err)
	ip, err = GetInternalIP2()
	t.Log("ip -> ", ip, err)
	ip, err = GetExternalIP()
	t.Log("ip -> ", ip, err)
	res, err := GetIPAddrByTaoBao(ip)
	t.Logf("res -> %+v, err -> %v", res, err)
	t.Logf("res -> %#v, err -> %v", res, err)
}
