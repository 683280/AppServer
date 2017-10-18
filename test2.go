package main

import (
	"crypto/aes"
	"crypto/cipher"
	"bytes"
	"fmt"
	"compress/zlib"
	"io"
	"time"
	"crypto/md5"
	"strconv"
	"encoding/base64"
	"encoding/hex"
	"crypto/rand"
)

func test() string {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}


//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func GetGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}
func main() {
	guid := GetGuid()
	fmt.Println(len(guid))
	fmt.Println(guid)
}
//func main() {
//	fmt.Println(len("8e1a188743c6077110da3c9778183031"))
//	token := test()
//	b,_ := AesEncrypt("你好123123123123123123123123123123123123",[]byte(token))
//	//fmt.Println(err)
//	//fmt.Println(len(b))
//	//b = DoZlibCompress(b)
//	//fmt.Println(len(b))
//	//b = DoZlibUnCompress(b)
//	//fmt.Println(len(b))
//	base := base64.NewEncoding()
//	b,_ = AesDecrypt(b,[]byte(token))
//	fmt.Println(string(b))
//}
//进行zlib压缩
func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

//进行zlib解压缩
func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

func AesEncrypt(data string, key []byte) ([]byte, error) {
	origData := []byte(data)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := aes.BlockSize
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	fmt.Println(len(key[:blockSize]))
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted ,key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := aes.BlockSize
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}