package main

import (
	"fmt"
	"log"

	"github.com/d2r2/go-i2c"
)

func main() {
	i2c, err := i2c.NewI2C(0x70, 1)
	if err != nil {
		log.Fatalln(err)
	}

	defer i2c.Close()

	_, err = i2c.WriteBytes([]byte{0x00, 0x31})
	if err != nil {
		log.Fatalln(err)
	}

	buf := make([]byte, 8)

	_, err = i2c.ReadBytes(buf)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(buf)
}
