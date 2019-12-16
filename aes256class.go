package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

type AES struct {
	key []byte
	iv  []byte
}

func NewAES(key, iv []byte) *AES {
	return &AES{key: key, iv: iv}
}

func (a *AES) getKey() []byte {

	keyLen := len(a.key)
	if keyLen < 16 {
		panic("res key 长度不能小于16")
	}
	arrKey := a.key
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	}
	//取前16个字节
	return arrKey[:16]
}

//加密字符串
func (a *AES) Encrypt(in string) (string, error) {

	plantText := []byte(in)
	key := a.getKey()
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
		return "", err
	}
	plantText = a.pKCS7Padding(plantText, block.BlockSize())

	blockModel := cipher.NewCBCEncrypter(block, a.iv[:aes.BlockSize])

	cipherText := make([]byte, len(plantText))

	blockModel.CryptBlocks(cipherText, plantText)
	return string(cipherText), nil
}

//解密字符串
func (a *AES) Decrypt(src string) (strDesc string, err error) {

	defer func() {
		//错误处理
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	key := a.getKey()
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
		return "", err
	}
	blockModel := cipher.NewCBCDecrypter(block, a.iv[:aes.BlockSize])
	plantText := make([]byte, len(src))
	blockModel.CryptBlocks(plantText, []byte(src))
	plantText = a.pKCS7UnPadding(plantText, block.BlockSize())
	return string(plantText), nil
}

//补位
func (a *AES) pKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unPadding := int(plantText[length-1])
	return plantText[:(length - unPadding)]
}

//补位
func (a *AES) pKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}
