package main

import (
	"fmt"
	"testing"
)

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
	fmt.Println(XORBreaker(hash))
}
