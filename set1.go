package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)

func main() {}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Challenge 1.1
func HexToBase64(hexval string) string {
	hexbytes, _ := hex.DecodeString(hexval)
	return base64.StdEncoding.EncodeToString(hexbytes)
}

//Challenge 1.2
func FixedXOR(bytes1 []byte, bytes2 []byte) []byte {
	if len(bytes1) != len(bytes2) {
		panic("FixedXOR cannot handle inputs of different lengths.")
	}
	//var bytes3 []byte
	bytes3 := make([]byte, len(bytes1))
	for i := 0; i < len(bytes1); i++ {
		bytes3[i] = bytes1[i] ^ bytes2[i]
	}
	return bytes3
}

//Challenge 1.3
func XORBreaker(hash []byte) (string, string, int) {
	topkey := ""
	topplain := ""
	topscore := 0
	for i := 31; i < 127; i++ {
		keychar := string(rune(i))
		key := strings.Repeat(keychar, len(hash)/len(keychar))
		plaintext := FixedXOR(hash, []byte(key))
		score := ScorePlaintext(plaintext)
		if score > topscore {
			topscore = score
			topkey = keychar
			topplain = string(plaintext)
		}
	}
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
	var score int
	for i := 0; i < len(plaintext); i++ {
		//score += BoolToInt(strings.ContainsAny(string(plaintext[i]), TopEngChar))
		if strings.ContainsAny(string(plaintext[i]), TopEngChar) {
			for j := 0; j < len(TopEngChar); j++ {
				score += (len(TopEngChar) - j) * BoolToInt(strings.Contains(string(plaintext[i]), string(TopEngChar[j])))
			}
		}
	}
	return score
}

//Challenge 1.5
func RepeatXOR(plaintext []byte, key []byte) []byte {
	key = bytes.Repeat(key, 1+int(math.Ceil(float64(len(plaintext)/len(key)))))
	key = key[0:len(plaintext)]
	return FixedXOR(plaintext, key)
}

//Challenge 1.6
func HamDist(bytes1 []byte, bytes2 []byte) int {
	dist := 0
	if len(bytes1) != len(bytes2) {
		panic("HamDist cannot handle inputs of different lengths.")
	}
	for i := 0; i < len(bytes1); i++ {
		dist += strings.Count(fmt.Sprintf("%b", bytes1[i]^bytes2[i]), "1")
	}
	return dist
}
func XORCrusher(hashbytes []byte) string {
	var lowNormdist float64
	var lowKeysize int
	for i := 2; i <= 40; i++ {
		k := 2*i + 2
		normdist1 := HamDist(hashbytes[0:i], hashbytes[i+1:2*i+1])
		normdist2 := HamDist(hashbytes[k:k+i], hashbytes[k+i+1:k+1+2*i])
		//fmt.Println(0, i, "-", i+1, 2*i+1, "-", k, k+i, "-", k+i+1, k+1+2*i)
		normdist := float64(normdist1+normdist2) / float64(i)
		if lowKeysize == 0 || normdist < lowNormdist {
			lowNormdist = normdist
			lowKeysize = i
		}
		//fmt.Println(i, "-", normdist, lowNormdist, lowKeysize)
	}
	fmt.Printf("Keysize: %d\n", lowKeysize)
	keyparts := make([]string, lowKeysize)
	for i := 0; i < lowKeysize; i++ {
		hashparts := make([]byte, len(hashbytes)/lowKeysize)
		for j := 0; j < len(hashbytes)/lowKeysize; j++ {
			if lowKeysize*j+i < len(hashbytes) {
				hashparts[j] = hashbytes[lowKeysize*j+i]
			}
		}
		probKeychar, _, _ := XORBreaker(hashparts)
		keyparts[i] = probKeychar
	}
	fmt.Println(keyparts)
	return strings.Join(keyparts, "")
}
