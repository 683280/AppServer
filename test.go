package main

import (
	"fmt"
	//"reflect"
	//"go/types"
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"errors"
	"crypto/rand"
	"time"
)


func main() {
	//instance := reflect.New(types.Elem())
	//ii := types.Elem()
	//i := instance.(ii.T)

	//t := Test{"123"}
	//i := *ts
	//
	//ts = append(i,t)
	//fmt.Println(reflect.ValueOf(&ts).Addr())
	//fmt.Println(reflect.TypeOf(in).Elem().Elem())
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)
	var s int = int(timestamp)
	fmt.Println(s)
	fmt.Println([]byte("\n"))
	code,_ := RsaEncrypt([]byte("你好"))
	fmt.Println(string(code))
	code,_ = RsaDecrypt(code)
	fmt.Println(string(code))
}
var publicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC519uh7Yp4AkzB2hQGgJ7Mvgcz
tPRvdjHQkpA+UMwA1n1HzYA06SHXF021gXUYilBxgfpzbqQaEkvaLrwqlxslDfK7
Al2mA3eM0EjusoFQF+v6VT65dC2TzpHoQeblC2b9xCwlyUXoH0uVIEcuAKSKZoMZ
Qfxr3ohFb3TL5zTyPwIDAQAB
-----END PUBLIC KEY-----
`
var privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQC519uh7Yp4AkzB2hQGgJ7MvgcztPRvdjHQkpA+UMwA1n1HzYA0
6SHXF021gXUYilBxgfpzbqQaEkvaLrwqlxslDfK7Al2mA3eM0EjusoFQF+v6VT65
dC2TzpHoQeblC2b9xCwlyUXoH0uVIEcuAKSKZoMZQfxr3ohFb3TL5zTyPwIDAQAB
AoGAYLSxxp5sWqyfspQ/rW6Ks/ICn2Z/d+ziWS2bP8IdliYHBTErkNzrzhiDSHr4
Ku/2kkpXwG+Hl0WEESIWqnb9GTPTmDXWhui9ZdoJ3wSf/B8JDfsMVR6R2uw2h7ig
Z/LSm0ItCBEej6dhmkxC8XJE29/+Q+DQVdKmyuQPZSG1DmkCQQDyhBN3DngZ1+m8
A5uQbiBOrH+u6CeMRWl2ZGHvjNbDAMJiGi+5jJeqNZpHf7WJ5Ou+eBxFqmUV1/jj
Xj+7ACarAkEAxC0dCvg/GdJVKBsCW4vGlmV+JagNuL4f1afr5K2pKKJret658Hn1
rfwIHp6Tj7S2I23UW+LgwVevqpm9faEyvQJANGUCi5NNsU+riNpCrsaMJlMwVsqD
WNPaQCDZ49ZKw+CTHnzH2M+eKMDh7xaRUxRpNkJe4VI5+qkpdX30SON0dwJAYQVe
w7oamw6nBvq0o8nxIRh41u7SOnftDqHJzIMGkg4h0datZv0qQC3RZjNPD1d0bPk4
eWkvdu+C9YCrcqJykQJBAI8fq3oDk121FCP0YjDj/FglDuEX9O+Xinv0FOBE1cbE
PL4vDFtaG5JRDdgTsIr87mKsssABd1DR/qZMxwCmDsc=
-----END RSA PRIVATE KEY-----
`
// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
