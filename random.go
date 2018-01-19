/**********************************************************
 * Author        :
 * Email         :
 * Last modified : 2016-09-15 17:29:51
 * Filename      : random.go
 * Description   : 常用的随机数工具方法
 * *******************************************************/
package utils

import (
	r "crypto/rand"
	"math/rand"
)

func o() *rand.Rand {
	return rand.New(rand.NewSource(TimestampNano()))
}

/**
 * 获取int类型随机数
 * @return r int
 */
func RandInt() (r int) {
	return (o().Int())
}

/**
 * 获取[0,n)int类型随机数
 * @param n int
 * @return r int
 */
func RandIntN(n int) (r int) {
	return (o().Intn(n))
}

/**
 * 获取int32类型随机数
 * @return r int32
 */
func RandInt32() (r int32) {
	return (o().Int31())
}

/**
 * 获取uint32类型随机数
 * @return r uint32
 */
func RandUint32() (r uint32) {
	return (o().Uint32())
}

/**
 * 获取int32类型随机数
 * @param n int32
 * @return r int32
 */
func RandInt32N(n int32) (r int32) {
	return (o().Int31n(n))
}

/**
 * 获取int64类型随机数
 * @return r int64
 */
func RandInt64() (r int64) {
	return (o().Int63())
}

/**
 * 获取int64类型随机数
 * @param n int64
 * @return r int64
 */
func RandInt64N(n int64) (r int64) {
	return (o().Int63n(n))
}

/**
 * 获取float32类型随机数
 * @return r float32
 */
func RandFloat32() (r float32) {
	return (o().Float32())
}

/**
 * 获取float64类型随机数
 * @return r float64
 */
func RandFloat64() (r float64) {
	return (o().Float64())
}

/**
 * Chan中存放随机数
 */
var RandInt32Chan chan int32
var RandUint32Chan chan uint32

func Init() {
	RandInt32Chan = make(chan int32, 10)
	RandUint32Chan = make(chan uint32, 10)
	go func() {
		for {
			select {
			case RandInt32Chan <- RandInt32():
			case RandUint32Chan <- RandUint32():
			}
		}
	}()
}

/**
 * 从RandInt32Chan中获取int32类型随机数
 * @return int32
 */
func GetRandInt32() int32 {
	return <-RandInt32Chan
}

/**
 * 从RandUint32Chan中获取uint32类型随机数
 * @return uint32
 */
func GetRandUint32() uint32 {
	return <-RandUint32Chan
}

/**
 * Rand Send
 * @return
 */
func Seed() {
	rand.Seed(TimestampNano())
}

var alphaNum = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)

/**
 * RandomCreateBytes generate random []byte by specify chars.
 * @param n int
 * @param alphabets []byte
 * @return []byte
 */
func RandomCreateBytes(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum
	}
	var bytes = make([]byte, n)
	var randBy bool
	if num, err := r.Read(bytes); num != n || err != nil {
		Seed()
		randBy = true
	}
	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[rand.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return bytes
}
