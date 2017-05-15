package utils

import "testing"

func Test_XXTEA(t *testing.T) {
	str := "Hello World! 你好，中国！"
	key := "1234567890"
	encrypt_data := Encrypt([]byte(str), []byte(key))
	//fmt.Println(base64.StdEncoding.EncodeToString(encrypt_data))
	decrypt_data := Decrypt(encrypt_data, []byte(key))
	t.Log(len(encrypt_data), len(decrypt_data))
	t.Log(string(encrypt_data))
}
