package main

import (
	"math/big"
	"strconv"
	"fmt"
	"crypto/rsa"
	"encoding/hex"
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
	"crypto/rand"
	"errors"
	"encoding/asn1"
	b64 "encoding/base64"

)

/**
 * Created by John Tsantilis 
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 31/8/2017.
 */

/*
show by command prompt
openssl genrsa -out key.pem
openssl rsa -in key.pem  -pubout > key-pub.pem

echo polaris@studygolang.com | openssl rsautl -encrypt -pubin -inkey key-pub.pem > cipher.txt

cat cipher.txt | openssl rsautl -decrypt -inkey key.pem

** OR encoding by base64 **
echo polaris@studygolang.com | openssl rsautl -encrypt -pubin -inkey key-pub.pem | openssl base64

openssl base64 -d | openssl rsautl -decrypt -inkey key.pem
*/

type pkcs8Key struct {
	Version             int
	PrivateKeyAlgorithm []asn1.ObjectIdentifier
	PrivateKey          []byte

}

func main() {
	msg := "15750035-3e0e"
	data, err := RsaEncrypt([]byte(msg))
	//fmt.Printf("PKCS1v15 encrypted [%s] to \n[%x]\n", string(msg), data)
	ioutil.WriteFile("encrypted.txt", data, 0644)
	if err != nil {
		fmt.Println(err)

	}

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	/*uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))*/

	origData, err := RsaDecrypt(data)
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println("origData >> ", string(origData))
	//cipherText, _ := hex.DecodeString("b6ee3caf14430003a20625ba1ea9ad31560ad203f7ecee46dd8e31f2dc47d278f3248bc0180e03571fdbf34a60aad7310468e6d6013fcfd6b785d1562411b44e089281adcc275a2037db3dec8b447b91162c859ab97372081c1bcb22a1fb33b1f72a06a54b1784d9f733aa1e869c6d64d45a7a78534714a773920ef7219b31f89092fc54f87ff371aeae5c3e59cdaad3fa05c24e781e06fcd46b35127a431bd85f62bafded95e3d31127159a0b5d13b77f11ecef94a037ac1d2f2c32fc0e6623cfe056127457f8f82631c33139a50fcd16c17e577b12f853cd55ffb16e099097dd76a21d987c536ac102b470e36881fc86f1667b505120a531458a116ca285b7")
	cipherText, _ := hex.DecodeString(RsaEncryptByPreSetPubKey())
	origData2, err := RsaDecrypt(cipherText)
	if err != nil {
		panic(err)

	}

	fmt.Println("origData2 >> ", string(origData2))

}

func rsa2pkcs8(key *rsa.PrivateKey) ([]byte, error) {
	var pkey pkcs8Key
	pkey.Version = 0
	pkey.PrivateKeyAlgorithm = make([]asn1.ObjectIdentifier, 1)
	pkey.PrivateKeyAlgorithm[0] = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
	pkey.PrivateKey = x509.MarshalPKCS1PrivateKey(key)

	return asn1.Marshal(pkey)

}

func RsaEncryptByPreSetPubKey() string {
	n, _ := new(big.Int).SetString("24349343452348953201209477858721354875245881458202672294652984377378513954748002477250933828219774703952578332297494223229725595176463711802920124930360492553186030821158773846902662847263120685557322462156596316871394035160273640449724455863863094140814233064652945361596472111169159061323006507670749392076044355771083774400487999226532334510138900864338047649454583762051951010712101235391104817996664455285600818344773697074965056427233256586264138950003914735074112527568699379597208762648078763602593269860453947862814755877433560650621539845829407336712267915875159364773551462882284084578152070138814976772753", 10)
	e, _ := strconv.ParseInt("10001", 16, 0)
	fmt.Printf("##RsaEncrypt2 n %x\n", n)
	fmt.Printf("##RsaEncrypt2 e %x\n", e)
	pubKey := rsa.PublicKey{n, int(e)}
	data, _ := rsa.EncryptPKCS1v15(rand.Reader, &pubKey, []byte("it's great for rsa"))

	return hex.EncodeToString(data)

}

func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")

	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err

	}

	pub := pubInterface.(*rsa.PublicKey)
	/*fmt.Println("Modulus : ", pub.N.String())
	fmt.Println(">>> ", pub.N)
	fmt.Printf("Modulus(Hex) : %X\n", pub.N)
	fmt.Println("Public Exponent : ", pub.E)*/

	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)

}

func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")

	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return nil, err

	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)

}

var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpwIBAAKCAQEAsdoMC1ek1ysJt1xvpl3ZZSTta8oEzUvzENHPHN/b/o94+y3r
+OMSn6S0CizJk92RicGz5L1yIAhNhbmfTkzrCacQulpXe1olelJjdZbGIevFp7fy
eiiRmGl0zrTA1bLVUKM5oqCaVtfVRH1hVeRvWQy4phszMCVVSP5c6WdNrLRasPSW
ycuk8dMZrxwBGzK/jgeDMxbCTzxA+15cekoCz8Qj+OXYsLAuyg62Ra9k6qKXZm/J
IAW36J5U/rd9140ItXpaAFWmD1GlJUPFLh9h04FMbIINnHSBd9Via/ptV95XaAQD
5n3fMySDcnawTTTREeWYnskoji+zC8K3NrUFbwIDAQABAoIBAB+V4OO2ygColQ4q
bW23Zx9uYnftJdMr9Nv81vKC0zgvVMYPDsRh13Hgt1TDRX1sMSes7fzmpDkKIHJq
V995C5joQbFG0BiJFvTVYvKEN2XY0H6LzZViVREjUWpQcZrCKt4qTtcR/LcFl8uV
FM9DRX8kIjrlw75+MtpGykttXD8acMFL76kR0aEiH3phqVdCOoVXls6v5atpbbzI
1AkgYmA6rgxzlyGC3JGGEfylz2vk40UrSpEMzPjllfHUrOrVkkV+XXDiH88tv3Ni
ph/nqGBGgLgO0Yn26+Wff4+WpoJetXCH3TWHqKHv6+YZnbLGAwEG1ltdNmfS9YLb
pv5ZfnkCggCBAMlDMQGjLb4QqSHtrTNgkP2pQ85kJPVX4zkteLCHBDk8/f6iZt4R
srDrwetFIlKEpnoTTOLm5h94TeNOMLWL3eQ/2pZKlgVQwrDH/2dSyIsjIkKk63Iy
DOy1YpGTBIvQMKouBv2Su8fdv+wm8FWsnCVFoJ8KvsdjR2ytMkWk4WKlAoIAgQDi
OOUPWuNvSy9z/o3M+6L40xmdk1gznDLBgObzi4jtWIILDQpkgLmoRCq2IpTxzuqP
FOwMr1arpaPPBvmq/LB1tSdX3OjD7gLroR6XCBqWjY5Im2kWFizyB0jB3fLdkgiF
MLjHp5soCWNRouZSnNHRrCkCxIyNF0H8NMrHFhJvgwKCAIAxRU2+BOCx4wpE0BbS
uRVFxCZhZIrXtUhfOw6MyMDz8kaMC937B4IjZdP3cNlSMj8/K3JrKfO0SJNozj+G
clNECY3NtTy82GSROqT/+ng198fvYMhzEGWxnUYgNUmroLAJOXNkqINoDpVpjq7P
io4/YahMxyilT1yw8kWYXyD+MQKCAIBMvpJuwB2yDuadSjke1is2kJ77BxaAA7hw
TuOKscaVBlavmw/wgn75Z4651UeENPb77VQbBYy4lGva+vnOGQVgGUGsnAPXHDKf
bzxqblReyM0BMr/Wv1UmEnSLWFlg836yeJHnIaQQVAb+le3fDtZbdDVBfB5WaxJ2
lg5IteMxMwKCAIAlZqa90lVPd8DuI/s7R1ivOnPEC6bR+r3+JZSBUoyDJL+xq62h
0WGsCytx8T4v+1hs+JL+mKXzUGlpYJQ0NaqD0NP9ffy5kiiDgt51CpKFFof47XvP
emG7I+VLnczzn01O4kT+nnl/+RuqGbpCn4p6uqetwVAZYpC+aYW7oXkTGA==
-----END RSA PRIVATE KEY-----`)

var publicKey = []byte(`-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsdoMC1ek1ysJt1xvpl3Z
ZSTta8oEzUvzENHPHN/b/o94+y3r+OMSn6S0CizJk92RicGz5L1yIAhNhbmfTkzr
CacQulpXe1olelJjdZbGIevFp7fyeiiRmGl0zrTA1bLVUKM5oqCaVtfVRH1hVeRv
WQy4phszMCVVSP5c6WdNrLRasPSWycuk8dMZrxwBGzK/jgeDMxbCTzxA+15cekoC
z8Qj+OXYsLAuyg62Ra9k6qKXZm/JIAW36J5U/rd9140ItXpaAFWmD1GlJUPFLh9h
04FMbIINnHSBd9Via/ptV95XaAQD5n3fMySDcnawTTTREeWYnskoji+zC8K3NrUF
bwIDAQAB
-----END RSA PUBLIC KEY-----`)