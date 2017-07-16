package main

import (
	"encoding/base64"
	"encoding/hex"
	"strconv"
	"strings"
)

func main() {
	FixedXOR("a", "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
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

//Challenge 1.3
func XORBreaker(hash string) (string, string, int) {
	topkey := ""
	topplain := ""
	topscore := 0
	for i := 0; i < 127; i++ {
		keychar := strconv.FormatInt(int64(i), 16)
		key := strings.Repeat(keychar, len(hash)/len(keychar))
		plaintext, _ := hex.DecodeString(FixedXOR(hash, key))
		score := ScorePlaintext(plaintext)
		if score > topscore {
			topscore = score
			topkey = keychar
			topplain = string(plaintext)
		}
	}
	z, _ := hex.DecodeString(topkey)
	topkey = string(z)
	return topkey, topplain, topscore
}
func BoolToInt(b bool) int {
	i := 0
	if b {
		i = 1
	}
	return i
}
func ScorePlaintext(plaintext []byte) int {
	TopEngChar := "ETAOINSHRDLU"
	var score int = 0
	for i := 0; i < len(plaintext); i++ {
		score += BoolToInt(strings.ContainsAny(string(plaintext[i]), TopEngChar))
	}
	return score
}
