/**********************************************************
 * Author        :
 * Email         :
 * Last modified : 2016-09-15 17:29:51
 * Filename      : random.go
 * Description   : 常用的随机数工具方法
 * *******************************************************/
package utils

import (
	"reflect"
	"testing"
)

func TestRand_01(t *testing.T) {
	bs0 := RandomCreateBytes(16)
	bs1 := RandomCreateBytes(16)

	t.Log(string(bs0), string(bs1))
	if string(bs0) == string(bs1) {
		t.FailNow()
	}

	bs0 = RandomCreateBytes(4, []byte(`a`)...)

	if string(bs0) != "aaaa" {
		t.FailNow()
	}
}

func TestRand_02(t *testing.T) {
	bs0 := RandInt()
	bs1 := RandIntN(16)

	t.Log(bs0, bs1)
	if reflect.ValueOf(bs0).Kind() != reflect.Int {
		t.FailNow()
	}
	if reflect.ValueOf(bs1).Kind() != reflect.Int {
		t.FailNow()
	}
}

func TestRand_03(t *testing.T) {
	bs0 := RandFloat32()
	bs1 := RandFloat64()

	t.Log(bs0, bs1)
	if reflect.ValueOf(bs0).Kind() != reflect.Float32 {
		t.FailNow()
	}
	if reflect.ValueOf(bs1).Kind() != reflect.Float64 {
		t.FailNow()
	}
}
