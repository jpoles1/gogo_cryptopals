package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestMisc(t *testing.T) {
	main()
}
func Test1(t *testing.T) {
	teststr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	validstr := ("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")
	if HexToBase64(teststr) != validstr {
		t.Fail()
	}
}
func Test2(t *testing.T) {
	str1 := "1c0111001f010100061a024b53535009181c"
	str2 := "686974207468652062756c6c277320657965"
	validstr := "746865206b696420646f6e277420706c6179"
	if FixedXOR(str1, str2) != validstr {
		t.Fail()
	}
}
func Test3(t *testing.T) {
	hash := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	key, plaintext, _ := XORBreaker(hash)
	fmt.Printf("XOR Breaker Test\nKey: %s\nPlaintext: %s\n%s\n", key, plaintext, strings.Repeat("-", 40))
}
func Test5(t *testing.T) {
	plaintext := fmt.Sprintf("%2x", "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := fmt.Sprintf("%2x", "ICE")
	validhash := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	if RepeatXOR(plaintext, key) != validhash {
		t.Fail()
	}
}
func Test6(t *testing.T) {
	//Testing the Hamming Distance Algorithm
	str1 := fmt.Sprintf("%2x", "this is a test")
	str2 := fmt.Sprintf("%2x", "wokka wokka!!!")
	bytes1, _ := hex.DecodeString(str1)
	bytes2, _ := hex.DecodeString(str2)
	if HamDist(bytes1, bytes2) != 37 {
		t.Fail()
	}
	//Testing the XORCrusher
	dat, err := ioutil.ReadFile("6.txt")
	check(err)
	filebytes, err := base64.StdEncoding.DecodeString(string(dat))
	check(err)
	XORCrusher(filebytes)
}
