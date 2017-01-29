package main

import (
	"os"
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"github.com/speps/go-hashids"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func main() {
	// 暗号化したい文字列
	plaintext := []byte("100000000000000000")

	// aesの暗号化文字列
	key_text := "astaxie12798aklj"

	fmt.Println(len(key_text))

	// 暗号化アルゴリズムaesを作成
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}

	// 暗号化文字列
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s=>%x\n", plaintext, ciphertext)

	// 復号文字列
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)

	arr := []int{}
	for _, n := range ciphertext {
		arr = append(arr, int(n))
	}

	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 32
	h := hashids.NewWithData(hd)
	e, _ := h.Encode(arr)
	fmt.Println(e)
	d, _ := h.DecodeWithError(e)
	fmt.Println(d)
}
