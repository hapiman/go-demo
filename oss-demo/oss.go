package oss_demo

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"strings"
)

import "fmt"

const (
	gAccessKeyId     = ""
	gAccessKeySecret = ""
	gendpoint        = "http://oss-cn-hangzhou.aliyuncs.com" //"http://oss-cn-hangzhou.aliyuncs.com"
)

// "http://oss-cn-hangzhou-internal.aliyuncs.com" //"http://oss-cn-hangzhou.aliyuncs.com"

var client *oss.Client
var bTarget = "milian-static"

func init() {
	var err error
	client, err = oss.New(gendpoint, gAccessKeyId, gAccessKeySecret, oss.Timeout(10, 60))
	if err != nil {
		panic(err)
	}
}

func GetList() {
	lsRes, err := client.ListBuckets()
	if err != nil {
		panic(err)
	}
	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}

func GoPut() {
	bucket, err := client.Bucket(bTarget)
	if err != nil {
		panic(err)
	}

	filePth := "/Users/pengjian/Desktop/code/other/go-demo/oss-demo/avatar.png"
	f, err := os.Open(filePth)
	if err != nil {
		panic(err)
	}

	path := "uploads/small_team2/avatar_url/10101010/avatar.png"
	path = strings.TrimLeft(path, "/")
	fmt.Println("path ->", path)
	err = bucket.PutObject(path, f)
	if err != nil {
		panic(err)
	}
	fmt.Println("succeed")
}
