package main

import (
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func writePng(base64, filename string) {
	code, err := qr.Encode(base64, qr.L, qr.Unicode)
	if err != nil {
		log.Fatal(err)
	}

	if base64 != code.Content() {
		log.Fatal("data differs")
	}

	code, err = barcode.Scale(code, 512, 512)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(file, code)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func main() {
	// xxx 二维码扫出来的内容
	xxx := "https://baidu.com"
	// name 生成的文件名
	name := "test.png"
	writePng(xxx, name)
}
