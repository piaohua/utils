#  utils

**Install**
```
cd $GOPATH
go get
```

### Documentation
* aes.go : aes
* all.go : all applys a function
* any.go : any applys a function
* apply.go : apply applys a function
* caller.go : GetFuncName get function name
* compose.go : 
* compress.go : gzip / ungzip
* debug.go : 
* err.go : 
* file.go : 
* filter.go : Filter apply a function
* goerr.go : 错误生成
* heap.go : heap队列
* helper.go : helper
* ini.go : 读取ini配置文件
* mail.go : mail
* option.go : 
* os.go : 执行系统命令
* partial.go : 
* queue.go : 
* random.go : 常用的随机数工具方法
* randutil.go : 产生随机字符串工具方法
* reduce.go : Reduce applys a function
* safemap.go : 
* slice.go : 
* sorting.go : 常用排序算法
* string.go : 常用的字符串工具方法
* structandmap.go : 
* timer_queue.go : 
* to.go  : go数据类型到字符数组的转换
* utils.go : 常用的工具方法，如：验证是否为邮箱等
* xxtea.go : 字符串加密，解密
* times.go : 常用的时间工具方法

---

### example
>参照php的date()函数和strtotime()函数
>>t := utils.StrToTime("2012-11-12 23:32:01")

>字符串转换为time.Time类型
>>t := utils.StrToLocalTime("2012-11-12 23:32:01")
>>t := utils.StrToLocalTime("2012-11-12")

>原生Go实现字符串转换为time.Time类型
>>t := time.Date(2012, 11, 12, 23, 32, 01, 0, time.Local)
>>t := time.Date(2012, 11, 12, 0, 0, 0, 0, time.Local)

>time.Time类型格式化为字符串
>>now := time.Now()
>>strTime := utils.Format("Y-m-d H:i:s", now)

>原生Go实现time.Time类型格式化为字符串
>>strTime := time.Now().Format("2006-01-02 15:04:05")

---

**TODO**

**License**
* LICENSE

**Reference**
* github.com/choleraehyq/gofunctools/functools
* github.com/polaris1119/times
* github.com/polaris1119/goutils
* github.com/jmcvetta/randutil
* github.com/astaxie/beego
* test result
