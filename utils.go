package main

import (
	"fmt"
	"strconv"
)

func hexToInt(bytes [2]byte) (int64, error) {

	return strconv.ParseInt(
		fmt.Sprintf("%d", bytes[0]*16*16+bytes[1]),
		16,
		32,
	)
}
