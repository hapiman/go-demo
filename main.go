package main

import (
	crand "crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"net/http"
	"regexp"
	"runtime"
	"strings"
	"time"
)

func GenRealRandomNum(len int64) int64 {
	result, _ := crand.Int(crand.Reader, big.NewInt(len))
	return result.Int64()
}

func tarPrint() string {
	time.Sleep(time.Second * 2)
	return "hahh"
}

func main() {

	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())

	flysnowRegexp := regexp.MustCompile(`^http://www.flysnow.org/([\d]{4})/([\d]{2})/([\d]{2})/([\w-]+).html$`)
	params := flysnowRegexp.FindStringSubmatch("http://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html")
	// 返回[]string{}数据类型
	for _, param := range params {
		fmt.Println(param)
	}

	ch := make(chan string)
	fmt.Println("xxxx")
	go func() {
		select {
		case <-time.After(time.Second):
			fmt.Println("1111")
		case ch <- tarPrint():
			fmt.Println("xxx")
			ch <- "22"
		}
	}()
	fmt.Println("yyyy")
	fmt.Println("ch string =>", <-ch)

	xx := []string{"a", "b", "c", "d"}
	for k, v := range xx {
		fmt.Println("k =>", k, v)
	}

	fmt.Println("====>", 0xf007)

	sourceStr := "gitlab.miliantech.com/go/push_center_server/controller/handler.(*Promo).ComputeTargetPerson.func2"
	_fp := sourceStr[len(sourceStr)-5 : len(sourceStr)-1]
	fmt.Println("==>", 0x000A, strings.LastIndex(sourceStr, ".func2"), string([]byte(sourceStr)[:]), _fp)

	//c := time.Tick(5 * time.Second)
	//for now := range c {
	//	fmt.Printf("%v \n", now)
	//}

	//ss := time.Duration(10 * time.Hour)
	//dd, err := time.ParseDuration("-10h")
	//fmt.Println("dd =>", dd, err)

	//year := time.Now().Year() //年
	//
	//fmt.Println(year)
	//
	//month := time.Now().Month() //月
	//fmt.Println(month)
	//day := time.Now().Day() //日
	//fmt.Println(day)
	//
	//hour := time.Now().Hour() //小时
	//fmt.Println(hour)
	//minute := time.Now().Minute() //分钟
	//fmt.Println(minute)
	//second := time.Now().Second() //秒
	//fmt.Println(second)
	//nanosecond := time.Now().Nanosecond() //纳秒
	//fmt.Println(nanosecond)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local)
	name, offset := t.Zone()
	fmt.Println("t ===", name, offset, time.Now())

	const timeLayout string = "2006-01-02 15:04:05"

	h, m, s := t.Clock()
	fmt.Println(h, m, s)

	l, _ := time.LoadLocation("Asia/Shanghai")
	todayEnd, _ := time.ParseInLocation(timeLayout, fmt.Sprintf("%s %s", time.Now().Format("2006-01-02"), "23:59:59"), l)
	todayStart := time.Now()
	validSeds := todayEnd.Sub(todayStart).Minutes()
	fmt.Println("validSeds: ", int(validSeds))

	num := runtime.NumCPU()

	fmt.Println(num)

	str := "\345\214\227\344\272\254\345\270\202"
	str2 := "\351\224\205\344\273\224"
	fmt.Println(str, str2)

	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		fmt.Println("nums :", len(files))
		for _, file := range files {
			log.Println(file.Filename)

			// Upload the file to specific dst.
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8080")
}
