package encrypts

import (
	"fmt"
	"strconv"
	"testing"
)

func TestEncrypt(t *testing.T) {
	plain := "17"
	var plainInt int64
	plainInt = 17
	// AES 规定有3种长度的key: 16, 24, 32分别对应AES-128, AES-192, or AES-256
	key := "awiugdhrwuiafaoaiuywfhbg"
	// 加密
	cipherByte, err := Encrypt(plain, key)
	cipherInt, err := EncryptInt64(plainInt, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s ==> %s\n", plain, cipherByte)
	fmt.Printf("%d ==> %s\n", plainInt, cipherInt)
	// 解密
	//cipherByte := "38c6"

	plainText, err := Decrypt(cipherByte, key)
	if err != nil {
		fmt.Println(err)
	}
	plainCode, _ := strconv.ParseInt(plainText, 10, 64)
	fmt.Printf("%s ==> %s\n", cipherByte, plainText)
	fmt.Printf("plainCode: %d\n", plainCode)
}
