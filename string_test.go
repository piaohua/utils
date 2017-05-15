/**********************************************************
 * Author        :
 * Email         :
 * Last modified : 2016-09-15 16:22:18
 * Filename      : string.go
 * Description   : 常用的字符串工具方法
 * *******************************************************/
package utils

import "testing"

func Test_Int2Str(t *testing.T) {
	s := Int2Str(10)
	t.Log(s)
	if s != "10" {
		t.FailNow()
	}
}

func Test_Base62encode(t *testing.T) {
	var n uint64 = 10
	s := Base62encode(n)
	i := Base62decode(s)
	t.Log(s, i)
	if i != n {
		t.FailNow()
	}
}

func Test_Split(t *testing.T) {
	str := "test.tar.gz"
	s := Split(str, "tar.gz")
	t.Log(str)
	t.Logf("%#v", s)
	path := "assets/files/3123.png"
	s = Split(path, "assets/files/")
	t.Log(path)
	t.Logf("%#v", s)
}
