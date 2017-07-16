package main

import (
	"encoding/base64"
	"encoding/hex"
)

func main() {
}
func HexToBase64(hexval string) string {
	hexbytes, _ := hex.DecodeString(hexval)
	return base64.StdEncoding.EncodeToString(hexbytes)
}
func FixedXOR(hex1 string, hex2 string) string {
	bytes1, _ := hex.DecodeString(hex1)
	bytes2, _ := hex.DecodeString(hex2)
	var bytes3 []byte
	for i := 0; i < len(bytes1); i++ {
		bytes3 = append(bytes3, bytes1[i]^bytes2[i])
	}
	return hex.EncodeToString(bytes3)
}
