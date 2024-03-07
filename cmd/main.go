package main

import (
	"fmt"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	png, err := qrcode.Encode("https://www.google.com", qrcode.Medium, 256)
	//byte to png

	if err != nil {
		panic(err)
	}

	err = os.WriteFile("", png, 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println(png)

}
