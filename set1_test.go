package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
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
	bytes1, _ := hex.DecodeString(str1)
	bytes2, _ := hex.DecodeString(str2)
	validstr := "746865206b696420646f6e277420706c6179"
	if bytes.Compare(FixedXOR(bytes1, bytes2), []byte(validstr)) == 0 {
		t.Fail()
	}
}
func Test3(t *testing.T) {
	hash := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	hashbytes, _ := hex.DecodeString(hash)
	key, plaintext, _ := XORBreaker(hashbytes)
	fmt.Printf("XOR Breaker Test\nKey: %s\nPlaintext: %s\n%s\n", key, plaintext, strings.Repeat("-", 40))
}
func Test5(t *testing.T) {
	plaintext := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	validhash := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	if bytes.Compare(RepeatXOR([]byte(plaintext), []byte(key)), []byte(validhash)) == 0 {
		t.Fail()
	}
}
func Test6pt1(t *testing.T) {
	//lipsum := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut sollicitudin metus ac arcu vulputate, vulputate lacinia nisi suscipit. Nunc finibus diam ac quam blandit sodales. Nulla ullamcorper laoreet ante, et lacinia enim semper ac. Proin aliquet porta dui mollis vulputate. Vivamus fringilla purus eget aliquet fermentum. Vestibulum quis varius arcu."
	engsum := "Is Given so lesser deep had fourth they're stars fly have shall thing female gathering us in cattle heaven greater cattle give they're may heaven fowl likeness. Have first life green make fish own dry so may him green may fruit. Of of beginning called open saying fruitful very evening, rule saw. Moving. That kind. Creeping given beginning man moving grass. Can't. Upon rule. Had over form. First good moveth said that fly night. Blessed, seed over. Beginning You them spirit. Won't void. Third itself she'd kind. Days his great rule meat above heaven meat land called fifth under itself."
	plaintext := engsum //"Testin testin wan two tree")
	key := "wew"
	hash := RepeatXOR([]byte(plaintext), []byte(key))
	crushedKey := fmt.Sprintf("%2x", XORCrusher(hash))
	fmt.Println(crushedKey)
	fmt.Println(bytes.Compare([]byte(plaintext), RepeatXOR([]byte(hash), []byte(key))) == 0)
	tst, _ := hex.DecodeString(string(RepeatXOR([]byte(hash), []byte(key))))
	fmt.Println(string(tst))
	fmt.Println(strings.Repeat("-", 40))
	//tst, _ := hex.DecodeString(RepeatXOR([]byte(hash), []byte(key)))
	//fmt.Printf("%s\n", tst)
}

/*func Test6(t *testing.T) {
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
	//fmt.Println(string(filebytes))
	crushedKey := fmt.Sprintf("%2x", XORCrusher(filebytes))
	_ = crushedKey
	fmt.Println(string(RepeatXOR(filebytes, []byte(crushedKey))))
}*/
