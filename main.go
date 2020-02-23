package main

import (
	crand "crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"net/http"
	"runtime"
)

func GenRealRandomNum(len int64) int64 {
	result, _ := crand.Int(crand.Reader, big.NewInt(len))
	return result.Int64()
}

func main() {

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
