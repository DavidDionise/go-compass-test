package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

var (
	CAL_START = []byte{0x00, 0xC0}
	CAL_END   = []byte{0x00, 0xC1}
	GET_ANGLE = []byte{0x00, 0x31}
)

func main() {

	_, err := host.Init()

	if err != nil {
		log.Fatal(err)
	}

	if os.Args[1] == "GET_ANGLE_CONTINUED" {
		for {
			command := GET_ANGLE
			process(command)
			time.Sleep(1 * time.Second)
		}
	} else {
		command, err := generateCommand(os.Args[1])

		if err != nil {
			log.Fatalln(err)
		}

		process(command)
	}
}
func process(command []byte) {

	bus, err := i2creg.Open("")

	if err != nil {
		log.Fatal(err)
	}

	dev := i2c.Dev{Bus: bus, Addr: 0x70}
	write := command
	buf := make([]byte, 8)

	err = dev.Tx(write, buf)
	if err != nil {
		log.Fatal("error writing\n", err)
	}

	// fmt.Println(buf)
	printBuffer(buf)
}

func printBuffer(buf []byte) {

	fmt.Print("[ ")
	for i, b := range buf {
		fmt.Printf("%x", b)
		if i != len(buf)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println(" ]")
}

func generateCommand(arg string) ([]byte, error) {
	switch arg {
	case "CAL_START":
		return CAL_START, nil
	case "CAL_END":
		return CAL_END, nil
	case "GET_ANGLE":
		return GET_ANGLE, nil
	case "SET_DEC_LOW":
		return []byte{0x04, 0x03}, nil
	case "FACTORY_DEFAULTS":
		return []byte{0xA0, 0xAA, 0xA5, 0xC5}, nil
	case "GET_TEMP":
		return []byte{0x00, 0x35}, nil
	default:
		return []byte{}, errors.New("unkown arg: " + arg)
	}
}
