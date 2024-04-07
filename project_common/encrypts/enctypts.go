package encrypts

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strconv"
	"test.com/project_project/pkg/model"
)

func Md5(str string) string {
	hash := md5.New()
	_, _ = io.WriteString(hash, str)
	return hex.EncodeToString(hash.Sum(nil))
}

//用户ID加密

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func DecryptNoErr(cipherStr string) int64 {
	decrypt, _ := Decrypt(cipherStr, model.AESKey)
	parseInt, _ := strconv.ParseInt(decrypt, 10, 64)
	return parseInt
}

func EncryptInt64NoErr(id int64) string {
	str, _ := EncryptInt64(id, model.AESKey)
	return str
}

func EncryptInt64(id int64, keyText string) (cipherStr string, err error) {
	idStr := strconv.FormatInt(id, 10)
	return Encrypt(idStr, keyText)
}
func Encrypt(plainText string, keyText string) (cipherStr string, err error) {
	if plainText == "" || keyText == "" {
		return "", errors.New("plainText and keyText should not be empty")
	}

	plainByte := []byte(plainText)
	keyByte := []byte(keyText)

	c, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cipherByte := make([]byte, len(plainByte))
	cfb.XORKeyStream(cipherByte, plainByte)
	cipherStr = hex.EncodeToString(cipherByte)
	return
}

//func Decrypt(cipherStr string, keyText string) (plainText string, err error) {
//	// 转换成字节数据, 方便加密
//	keyByte := []byte(keyText)
//	// 创建加密算法aes
//	c, err := aes.NewCipher(keyByte)
//	if err != nil {
//		return "", err
//	}
//	// 解密字符串
//	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
//	cipherByte, _ := hex.DecodeString(cipherStr)
//	plainByte := make([]byte, len(cipherByte))
//	cfbdec.XORKeyStream(plainByte, cipherByte)
//	plainText = string(plainByte)
//	return
//}

func Decrypt(cipherStr string, keyText string) (plainText string, err error) {
	if cipherStr == "" || keyText == "" {
		return "", errors.New("cipherStr and keyText should not be empty")
	}

	keyByte := []byte(keyText)

	c, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}

	cfbdec := cipher.NewCFBDecrypter(c, commonIV)

	cipherByte, err := hex.DecodeString(cipherStr)
	if err != nil {
		return "", fmt.Errorf("failed to decode cipherStr: %v", err)
	}

	plainByte := make([]byte, len(cipherByte))
	cfbdec.XORKeyStream(plainByte, cipherByte)
	plainText = string(plainByte)
	return
}
