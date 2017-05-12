# pfop


[![Go Report Card](https://goreportcard.com/badge/github.com/qiniu/pfop)](https://goreportcard.com/report/github.com/qiniu/pfop)

## 简介

pfop 是七牛持久化数据处理的请求接口，具体文档可以参考[持久化数据处理文档](http://developer.qiniu.com/code/v6/api/dora-api/pfop/pfop.html)。
本工具是命令行工具，方便开发者在调试服务的时候使用，具体的代码开源，利用到了[七牛的存储SDK](https://github.com/qiniu/api.v6)。

## 下载
可以下载已经编译好的各个平台的可执行文件。

[点击下载](http://devtools.qiniu.com/pfop-v1.2.zip)


如果需要自己编译，可以设置好 $GOPATH ，然后使用如下命令下载依赖库：

```
go get github.com/qiniu/api.v6
go get github.com/qiniu/rpc
```

然后使用`go build pfop.go`进行编译。

## 使用

```
Usage of pfop:
  -ak="": access key
  -sk="": secret key
  -bucket="": bucket name
  -key="": file key
  -fops="": joined fops
  -pipe="": pipeline to use
  -url="": notify url
  -zone="nb": api zone [nb, bc, hn, na0]
  -force: force to redo
```

|参数|描述|
|---|-----|
|ak|七牛API的 AccessKey|
|sk|七牛API的 SecretKey|
|bucket|待处理文件所在空间|
|key|待处理文件的名称|
|fops|持久化数据处理的指令|
|pipe|数据处理的私有队列|
|url|数据处理完成的结果通知地址|
|zone|空间所在的机房，这个参数可以不填|
|force|是否强制执行指令，避免旧指令缓存|

## 多机房

|机房|zone值|
|----|----|
|华东|nb|
|华北|bc|
|华南|hn|
|北美|na0|

## 示例

我们对华北空间 `if-bc` 里面的文件 `qiniu.mp4` 做 m3u8 切片，同时截取一个封面图片。

```
$ ./pfop_darwin_amd64 -ak 'u-vclPNwctgp033dF6RSlKxpi1q-UL2yPfw8FtkM' -sk 'xxxfJETfd4PwumJdPgawsEoSeWvAAO5KlJDk5S22' -bucket 'if-bc' -key 'qiniu.mp4' -fops 'avthumb/m3u8|saveas/aWYtYmM6cWluaXUubTN1OA==;vframe/jpg/offset/10|saveas/aWYtYmM6cWluaXUuanBn' -pipe 'jemy' -zone 'bc' -force

See http://api-z1.qiniu.com/status/get/prefop?id=z1.57cfe243f51b822f9501fe47
```
上面的命令适用于Linux和Mac环境，如果是Windows的环境，请给参数加上双引号，而不是单引号。

可以通过访问输出的链接地址，查看数据处理的结果，也可以在上面指令发送之前，指定参数`url`来主动接受七牛的处理完成通知。

处理结果如下，具体含义参考文档[数据处理结果查询](http://developer.qiniu.com/code/v6/api/dora-api/pfop/prefop.html)。

```
{
  "code": 0,
  "desc": "The fop was completed successfully",
  "id": "z1.57cfe243f51b822f9501fe47",
  "inputBucket": "if-bc",
  "inputKey": "qiniu.mp4",
  "items": [
    {
      "cmd": "avthumb/m3u8|saveas/aWYtYmM6cWluaXUubTN1OA==",
      "code": 0,
      "desc": "The fop was completed successfully",
      "hash": "Fh-6iZGmH7m8_EI0LjkcQHAYq5cr",
      "key": "qiniu.m3u8",
      "returnOld": 0
    },
    {
      "cmd": "vframe/jpg/offset/10|saveas/aWYtYmM6cWluaXUuanBn",
      "code": 0,
      "desc": "The fop was completed successfully",
      "hash": "FgxUIsJLHfGt1pY7uFnA-Nxf6TgN",
      "key": "qiniu.jpg",
      "returnOld": 0
    }
  ],
  "pipeline": "1380340116.jemy",
  "reqid": "MDkAAPYdMAF-AXIU"
}
```