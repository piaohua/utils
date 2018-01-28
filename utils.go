package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"unicode/utf8"
	"unsafe"
)

// USE AT YOUR OWN RISK

/**
 * Get Auth
 * @return []rune
 */
func GetAuth() []rune {
	var list []rune
	for i := 0; i < 6; i++ {
		ran := RandIntN(122-97+1) + 97
		list = append(list, rune(ran))
	}
	return list
}

/**
 * Generated Salt
 * @return string
 */
func GetSalt() string {
	return RandString(6)
}

/**
 * 验证是否邮箱
 * @param mail string
 * @return bool
 */
func EmailRegexp(mail string) bool {
	if mail != "" {
		reg := regexp.MustCompile(`^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(\.[a-zA-Z0-9_-]+)$`)
		return reg.FindString(mail) != ""
	}
	return false
}

/**
 * 验证是否手机
 * @param phone string
 * @return bool
 */
func PhoneRegexp(phone string) bool {
	if phone != "" {
		reg := regexp.MustCompile(`^(86)*0*1\d{10}$`)
		return reg.FindString(phone) != ""
	}
	return false
}

/**
 * 验证是否手机
 * @param phone string
 * @return bool
 */
func PhoneValidate(phone string) bool {
	if phone == "" {
		return false
	}
	reg := regexp.MustCompile(`^((1[3,4,5,7,8][0-9])|(14[5,7])|(17[0,6,7,8])|(19[7]))\d{8}$`)
	return reg.MatchString(phone)
}

/**
 * 验证账号是否合法
 * @param account string
 * @return bool
 */
func AccountRegexp(account string) bool {
	if account != "" {
		reg := regexp.MustCompile(`^[a-zA-Z0-9]{6,8}$`)
		return reg.FindString(account) != ""
	}
	return false
}

/**
 * 验证只能由数字字母下划线组成的5-17位密码字符串
 * @param name string
 * @return bool
 */
func AalidataPwd(name string) (b bool) {
	if name != "" {
		//reg := regexp.MustCompile(`^[a-zA-Z0-9_]*$`)
		reg := regexp.MustCompile(`^[a-zA-Z_]\w{5,17}$`)
		b = reg.FindString(name) != ""
	}
	return
}

var IllegalNameRune [13]rune

func InitIllegalNameRune() {
	IllegalNameRune[0] = 0x00  //\0
	IllegalNameRune[1] = 0x09  //\t
	IllegalNameRune[2] = 0x5f  //_
	IllegalNameRune[3] = 0x20  //space
	IllegalNameRune[4] = 0x22  //"
	IllegalNameRune[5] = 0x60  //`
	IllegalNameRune[6] = 0x1a  //ctrl+z
	IllegalNameRune[7] = 0x0a  //\n
	IllegalNameRune[8] = 0x0d  //\r
	IllegalNameRune[9] = 0x27  //'
	IllegalNameRune[10] = 0x25 //%
	IllegalNameRune[11] = 0x5c //\
	IllegalNameRune[12] = 0x2c //,
}

var hasIllegalNameRune = func(c rune) bool {
	for _, v := range IllegalNameRune {
		if v == c {
			return true
		}
	}
	return false
}

/**
 * 验证名字
 * @param name string
 * @param maxcount int
 * @return bool
 */
func LegalName(name string, maxcount int) bool {
	if !utf8.ValidString(name) {
		return false
	}

	num := len([]rune(name)) + len([]byte(name))
	result := float64(num) / 4.0
	sum := int(result + 0.99)

	if sum > maxcount*2 {
		return false
	}
	return strings.IndexFunc(name, hasIllegalNameRune) == -1
}

/**
 * 把时间戳转换成头像存储目录
 * @param t int64
 * @param userid int
 * @param headname int64
 * @return (string, string)
 */
func TimeToHeadphpoto(t int64, userid int, headname int64) (string, string) {
	var str, name string
	str = Unix2Fstr(t)

	str = "./headpic/" + str + "/" + Itoa(userid)
	if headname == 0 {
		name = "/130_" + Itoa(userid) + ".jpg"
	} else {
		name = "/" + Itoa(int(headname)) + ".jpg"
	}
	return str, name
}

/**
 * 把时间戳转换成头像存储目录
 * @param t int64
 * @param userid int
 * @return string
 */
func TimeToPhpotoPath(t int64, userid int) string {
	var str string
	str = Unix2Fstr(t)
	return "./photo/" + str + "/" + Itoa(userid)
}

/**
 * 利用用户id生成数值类型唯一激活码
 * @param userid string
 * @return uint32
 */
// TODO:bug
func UseridCovToInvate(userid string) uint32 {
	useridbyte := []byte(userid)
	useridbyte = useridbyte[len(useridbyte)-4:]
	timestr := []byte(Itoa(int(Timestamp())))
	timestr = timestr[len(timestr)-5:]
	useridbyte = append(useridbyte, timestr...)
	code, _ := Atoi(string(useridbyte))
	return uint32(code)
}

/**
 * Bytes2String force casts a []byte to a string.
 * @param b []byte
 * @return (s string)
 */
func Bytes2String(b []byte) (s string) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

/**
 * String2Bytes force casts a string to a []byte
 * @param s string
 * @return (b []byte)
 */
func String2Bytes(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

/**
 * 用gob进行数据编码
 * @param data interface{}
 * @return ([]byte, error)
 */
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

/**
 * 用gob进行数据解码
 * @param data []byte
 * @param to interface{}
 * @return error
 */
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

/**
 * 对象深度拷贝
 * @param dst interface{}
 * @param src interface{}
 * @return error
 */
func Clone(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

/**
 * 数值类型转成点分结构的IP地址
 * @param ipnr string
 * @return uint32
 * @eg: t.Log((InetTontoa(3232235966).String()))
 */
func InetTontoa(ipnr uint32) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)
	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

/**
 * 字节类型的IP地址转成数值类型
 * @param ipnr string
 * @return uint32
 * @eg: t.Log((InetTobton(net.IPv4(192,168,1,190))))
 */
func InetTobton(ipnr net.IP) uint32 {
	bits := Split(ipnr.String(), ".")

	b0, _ := Atoi(bits[0])
	b1, _ := Atoi(bits[1])
	b2, _ := Atoi(bits[2])
	b3, _ := Atoi(bits[3])

	var sum uint32

	sum += uint32(b0) << 24
	sum += uint32(b1) << 16
	sum += uint32(b2) << 8
	sum += uint32(b3)

	return sum
}

/**
 * 点分结构的IP地址转成数值类型
 * @param ipnr string
 * @return uint32
 * @eg t.Log((InetToaton("192.168.1.190")))
 */
func InetToaton(ipnr string) uint32 {
	bits := Split(ipnr, ".")

	b0, _ := Atoi(bits[0])
	b1, _ := Atoi(bits[1])
	b2, _ := Atoi(bits[2])
	b3, _ := Atoi(bits[3])

	var sum uint32

	sum += uint32(b0) << 24
	sum += uint32(b1) << 16
	sum += uint32(b2) << 8
	sum += uint32(b3)

	return sum
}

/**
 * 将一个IPV4的字符串互联网协议转换成数字格式
 * @param ip_address string
 * @return float64
 */
func Ip2Long(ip_address string) float64 {
	ip := Split(ip_address, ".")
	if len(ip) != 4 {
		return 0
	}
	ip0, _ := Atoi(ip[0])
	ip1, _ := Atoi(ip[1])
	ip2, _ := Atoi(ip[2])
	ip3, _ := Atoi(ip[3])
	return float64(ip3 + 256*ip2 + 256*256*ip1 + 256*256*256*ip0)
}

/**
 * 将一个数字格式转换成IPV4的字符串
 * @param long float64
 * @return ip string
 */
func Long2Ip(long float64) string {
	var ip string = ""
	if long < 0 || long > 4294967295 {
		return ip
	}
	var i float64
	for i = 3; i >= 0; i-- {
		ip += Itoa(int(long / math.Pow(256, i)))
		if i > 0 {
			ip += "."
		}
		long -= float64(int(long/math.Pow(256, i))) * math.Pow(256, i)
	}
	return ip
}

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

// RemoteIp 返回远程客户端的 IP，如 192.168.1.1
func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

// Ip2long 将 IPv4 字符串形式转为 uint32
func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

func GetInternalIP() (inip string, err error) {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		return
	}
	defer conn.Close()
	inip = strings.Split(conn.LocalAddr().String(), ":")[0]
	return
}

func GetInternalIP2() (inip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				inip = ipnet.IP.String()
				break
			}
		}
	}
	return
}

func GetExternalIP() (exip string, err error) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	exip = string(bytes.TrimSpace(b))
	return
}

//通过淘宝API获取ip物理地址
type IPTaoBao struct {
	Code int          `json:"code"`
	Data IPTaoBaoData `json:"data"`
}

type IPTaoBaoData struct {
	Country    string `json:country`
	Country_id string `json:country_id`
	Area       string `json:area`
	Area_id    string `json:area_id`
	Region     string `json:region`
	Region_id  string `json:region_id`
	City       string `json:city`
	City_id    string `json:city_id`
	County     string `json:county`
	County_id  string `json:county_id`
	Isp        string `json:isp`
	Isp_id     string `json:isp_id`
	Ip         string `json:ip`
}

func GetIPAddrByTaoBao(ip string) (res *IPTaoBao, err error) {
	res = new(IPTaoBao)
	url := "http://ip.taobao.com/service/getIpInfo.php?ip=" + ip
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes.TrimSpace(b), &res)
	return
}

//
func GetPhoto(str string) string {
	if str == "" {
		return str
	}
	//return strings.Join(Split(str, "wx.qlogo.cn/mmopen/"), "nn.18bn.cn/mmopen/?photo=")
	//return strings.Join(Split(str, "wx.qlogo.cn/mmopen/"), "nnyl.iy00.cn/mmopen/?photo=")
	return strings.Join(Split(str, "wx.qlogo.cn/mmopen/"), "qpnn.qd92.cn/mmopen/?photo=")
}
