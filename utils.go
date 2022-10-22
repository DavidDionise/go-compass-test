package main

func hexToInt(bytes [2]byte) int {
	return (int(bytes[0]) << 8) + int(bytes[1])
}
