package util

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
)

//​func main() {
//	//key的长度必须都是8位
//	var key = "12345678"
//	var info = "110 119 120 122"
//
//	Enc_str := EncryptDES_CBC(info, key)
//	fmt.Println(Enc_str)
//	Dec_str := DecryptDES_CBC(Enc_str, key)
//	fmt.Println(Dec_str)
//
//	Enc_str = EncryptDES_ECB(info, key)
//	fmt.Println(Enc_str)
//	Dec_str = DecryptDES_ECB(Enc_str, key)
//	fmt.Println(Dec_str)
//}

//CBC加密
func DESEncryptByCBC(src, key string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	data = PKCS5Padding(data, block.BlockSize())
	//获取CBC加密模式
	iv := keyByte //用密钥作为向量(不建议这样使用)
	mode := cipher.NewCBCEncrypter(block, iv)
	out := make([]byte, len(data))
	mode.CryptBlocks(out, data)
	return fmt.Sprintf("%X", out)
}

//CBC解密
func DESDecryptByCBC(src, key string) string {
	keyByte := []byte(key)
	data, err := hex.DecodeString(src)
	if err != nil {
		panic(err)
	}
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	iv := keyByte //用密钥作为向量(不建议这样使用)
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(data))
	mode.CryptBlocks(plaintext, data)
	plaintext = PKCS5UnPadding(plaintext)
	return string(plaintext)
}

//ECB加密
func DESEncryptByECB(src, key string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte[:8])
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	//对明文数据进行补码
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		panic("Need a multiple of the block size")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//对明文按照blocksize进行分块加密
		//必要时可以使用go关键字进行并行加密
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return fmt.Sprintf("%X", out)
}

//ECB解密
func DESDecryptByECB(src, key string) string {
	data, err := hex.DecodeString(src)
	if err != nil {
		panic(err)
	}
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return string(out)
}

//明文补码算法
func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

//明文减码算法
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

//DES 和 3DES加密区别
//前者 加密  密钥必须是8byte
//后者加密 解密 再加密  密钥必须是24byte
//func main() {
//	//定义密钥，必须是24byte
//	key := []byte("123456789012345678901234")
//	//定义明文
//	origData := []byte("hello world")
//	​
//	//加密
//	en := ThriDESEnCrypt(origData,key)
//	​
//	//解密
//	de := ThriDESDeCrypt(en,key)
//	fmt.Println(string(de))
//}
//解密

func ThreeDESDeCrypt(src, key string) []byte {
	var crypt = []byte(src)
	var keyByte = []byte(key)

	//获取block块
	block, err := des.NewTripleDESCipher(keyByte)
	if err != nil {
		panic(err)
	}
	//创建切片
	context := make([]byte, len(crypt))
	//设置解密方式
	blockMode := cipher.NewCBCDecrypter(block, keyByte[:8])
	//解密密文到数组
	blockMode.CryptBlocks(context, crypt)
	//去补码
	context = PKCSUnPadding(context)
	return context
}

//去补码
func PKCSUnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:length-unPadding]
}

//加密
func ThreeDESEnCrypt(src, key string) []byte {
	//获取block块
	var origData = []byte(src)
	var keyByte = []byte(key)

	block, _ := des.NewTripleDESCipher(keyByte)
	//补码
	origData = PKCSPadding(origData, block.BlockSize())
	//设置加密方式为 3DES  使用3条56位的密钥对数据进行三次加密
	blockMode := cipher.NewCBCEncrypter(block, keyByte[:8])
	//创建明文长度的数组
	crypt := make([]byte, len(origData))
	//加密明文
	blockMode.CryptBlocks(crypt, origData)
	return crypt
}

//补码
func PKCSPadding(origData []byte, blockSize int) []byte {
	//计算需要补几位数
	padding := blockSize - len(origData)%blockSize
	//在切片后面追加char数量的byte(char)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(origData, padText...)
}
