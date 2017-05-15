package utils

import (
	"fmt"
	"testing"
)

func Test_AES(t *testing.T) {
	aesEnc := AesEncrypt{}
	aesEnc.SetKey([]byte("aalk;lkasjd;lkfj;alk"))
	doc := []byte("abcde。")
	arrEncrypt, err := aesEnc.Encrypt(doc)
	fmt.Println(string(arrEncrypt))
	if err != nil {
		fmt.Println(string(arrEncrypt))
		return
	}
	strMsg, err := aesEnc.Decrypt(arrEncrypt)
	if err != nil {
		fmt.Println(string(arrEncrypt))
		return
	}
	fmt.Println(string(strMsg))
}
