/**********************************************************
 * Author        :
 * Email         :
 * Last modified : 2016-09-15 16:22:18
 * Filename      : string.go
 * Description   : 常用的字符串工具方法
 * *******************************************************/
package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

/**
 * 随机长度数值字符串
 * @param i int
 * @return s string
 */
func RandStr(i int) (s string) {
	for i > 0 {
		s += strconv.FormatInt(RandInt64N(10), 10)
		i--
	}
	return
}

/**
 * 整形转换成字符串
 * @param i int
 * @return s string
 */
func Int2Str(i int) string {
	return strconv.Itoa(i)
}

func Itoa(i int) string {
	return Int2Str(i)
}

/**
 * 字符串转换成整形
 * @param s string
 * @return (int, error)
 */
func Str2Int(s string) (int, error) {
	return strconv.Atoi(s)
}

func Atoi(s string) (int, error) {
	return Str2Int(s)
}

/**
 * 分离字符串为一个slices
 * @param s string
 * @param sep string
 * @return []string
 */
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

/**
 * 截取字符串
 * @param string str
 * @param begin int
 * @param length int
 * @return int 长度
 */
func SubStr(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}

/**
 * 字符串加法
 * @param numStr string
 * @return string
 */
func StringAdd(numStr string) string {
	runeArr := []rune(numStr)
	length := len(runeArr)
	add := true
	for i := length - 1; i >= 0; i-- {
		if runeArr[i] < 57 {
			runeArr[i]++
			add = false
			break
		} else {
			runeArr[i] = 48
		}
	}
	if add {
		runeArr = append([]rune{49}, runeArr...)
	}
	return string(runeArr)
}

/**
 * 整形转换成字节
 * @param n int
 * @return []byte
 */
func IntToBytes(n int) []byte {
	var x int32 = int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

/**
 * 字节转换成整形
 * @param b []byte
 * @return int
 */
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

/**
 * int64整形转换成字节
 * @param n int64
 * @return []byte
 */
func Int64ToBytes(n int64) []byte {
	x := int64(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

/**
 * 字节转换成int64整形
 * @param b []byte
 * @return int64
 */
func BytesToInt64(b []byte) int64 {
	bytesBuffer := bytes.NewBuffer(b)
	var x int64
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int64(x)
}

/**
 * 切片中字符串第一个位置
 * @param arr []string
 * @param str string
 * @return int
 */
func SliceIndexOf(arr []string, str string) int {
	var index int = -1
	arrlen := len(arr)
	for i := 0; i < arrlen; i++ {
		if arr[i] == str {
			index = i
			break
		}
	}
	return index
}

/**
 * 切片中字符串最后一个位置
 * @param arr []string
 * @param str string
 * @return int
 */
func SliceLastIndexOf(arr []string, str string) int {
	var index int = -1
	for arrlen := len(arr) - 1; arrlen > -1; arrlen-- {
		if arr[arrlen] == str {
			index = arrlen
			break
		}
	}
	return index
}

/**
 * 从字符串切片中移除指定切片字符串
 * @param oriArr []string
 * @param removeArr []string
 * @return []string
 */
func SliceRemoveFormSlice(oriArr []string, removeArr []string) []string {
	endArr := oriArr[:]
	for _, value := range removeArr {
		index := SliceIndexOf(endArr, value)
		if index != -1 {
			endArr = append(endArr[:index], endArr[index+1:]...)
		}
	}
	return endArr
}

/**
 * 生成32位md5字符串
 * @param s string
 * @return string
 */
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

/**
 * 生成Guid字符串
 * @return string
 */
func GetGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""

	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

var base = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var flipbase = flip(base)
var baselen = len(base)

/**
 * 翻转切片字符串
 * @param s string
 * @return map[string]uint64
 */
func flip(s []string) map[string]uint64 {
	f := make(map[string]uint64)
	for index, value := range s {
		f[value] = uint64(index)
	}
	return f
}

/**
 * 翻转切片字符串
 * @param num uint64
 * @return string
 */
func Base62encode(num uint64) string {
	baseStr := ""
	for {
		if num <= 0 {
			break
		}

		i := num % uint64(baselen)
		baseStr += base[i]
		num = (num - i) / uint64(baselen)
	}
	return baseStr
}

/**
 * 翻转切片字符串
 * @param base62 string
 * @return uint64
 */
func Base62decode(base62 string) uint64 {
	var rs uint64 = 0
	len := uint64(len(base62))
	var i uint64
	for i = 0; i < len; i++ {
		rs += flipbase[string(base62[i])] * uint64(math.Pow(float64(baselen), float64(i)))
	}
	return rs
}

/**
 * Md5加密
 * @param text string
 * @return string
 */
func Md5Copy(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

/**
 * Md5加密
 * @param text string
 * @return string
 */
func Md5(text string) string {
	h := md5.New()
	h.Write([]byte(text))                 // 需要加密的字符串
	return hex.EncodeToString(h.Sum(nil)) // 输出加密结果
}

func Md5Buf(buf []byte) string {
	hashMd5 := md5.New()
	hashMd5.Write(buf)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func Md5File(reader io.Reader) string {
	var buf = make([]byte, 4096)
	hashMd5 := md5.New()
	for {
		n, err := reader.Read(buf)
		if err == io.EOF && n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			break
		}
		hashMd5.Write(buf[:n])
	}
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func Base64Encode(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

func Base64Decode(data string) string {
	b, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return ""
	}
	return string(b)
}

// func Join(args []string, sep string) string {
// 	return strings.Join(args, sep)
// }
//
// func TrimSpace(s string) string {
// 	return strings.TrimSpace(s)
// }

// 驼峰式写法转为下划线写法
func UnderscoreName(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// 下划线写法转为驼峰写法
func CamelName(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// 搜索字符串出现的位置
func SearchString(slice []string, s string) int {
	for i, v := range slice {
		if s == v {
			return i
		}
	}
	return -1
}

// GenSign 根据输入参数进行签名
func GenSign(args url.Values, secretKey string) string {
	keys := make([]string, 0, len(args))
	for k := range args {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys))

	buffer := NewBuffer()
	for _, k := range keys {
		buffer.Append(k).Append("=").Append(args.Get(k))
	}

	buffer.Append(secretKey)

	return Md5(buffer.String())
}
