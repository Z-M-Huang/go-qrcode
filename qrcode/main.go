// go-qrcode
// Copyright 2014 Tom Harwood

package main

import (
	"fmt"
	"image"
	"os"

	qrcode "github.com/Z-M-Huang/go-qrcode"
)

func main() {
	outFile := "output.png"
	size := 256
	negative := false
	disableBorder := false
	content := "content to qrcode"

	logoImg, _ := os.Open("google.png")
	defer logoImg.Close()
	logo, _, _ := image.Decode(logoImg)

	var err error
	var q *qrcode.QRCode
	q, err = qrcode.NewWithLogo(content, qrcode.Highest, logo)
	checkError(err)

	if disableBorder {
		q.DisableBorder = true
	}

	if negative {
		q.ForegroundColor, q.BackgroundColor = q.BackgroundColor, q.ForegroundColor
	}

	var png []byte
	png, err = q.PNG(size)
	checkError(err)

	if outFile == "" {
		os.Stdout.Write(png)
	} else {
		var fh *os.File
		fh, err = os.Create(outFile + ".png")
		checkError(err)
		defer fh.Close()
		fh.Write(png)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
