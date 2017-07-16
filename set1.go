package main

import (
	"encoding/base64"
	"encoding/hex"
)

func main() {
}

//Challenge 1.1
func HexToBase64(hexval string) string {
	hexbytes, _ := hex.DecodeString(hexval)
	return base64.StdEncoding.EncodeToString(hexbytes)
}

//Challenge 1.2
func FixedXOR(hex1 string, hex2 string) string {
	bytes1, _ := hex.DecodeString(hex1)
	bytes2, _ := hex.DecodeString(hex2)
	if len(bytes1) != len(bytes2) {
		panic("FixedXOR cannot handle inputs of different lengths.")
	}
	//var bytes3 []byte
	bytes3 := make([]byte, len(bytes1))
	for i := 0; i < len(bytes1); i++ {
		bytes3[i] = bytes1[i] ^ bytes2[i]
	}
	return hex.EncodeToString(bytes3)
}
