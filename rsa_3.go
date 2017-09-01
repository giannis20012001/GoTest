package main

import (
	"crypto/sha1"
	"encoding/pem"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"crypto/rand"
	b64 "encoding/base64"

)

/**
 * Created by John Tsantilis 
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 1/9/2017.
 */

func main() {
	msg := "15750035-3e0e"

	block, _ := pem.Decode(publickKeyData)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	pub := pubInterface.(*rsa.PublicKey)
	/*fmt.Println("Modulus : ", pub.N.String())
	fmt.Println(">>> ", pub.N)
	fmt.Printf("Modulus(Hex) : %X\n", pub.N)
	fmt.Println("Public Exponent : ", pub.E)*/
	data, _ := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(msg))
	//fmt.Printf("PKCS1v15 encrypted [%s] to \n[%x]\n", string(msg), data)
	if err != nil {
		fmt.Println(err)

	}

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	//golang encrypted data
	encryptedData := uEnc
	privateKeyBlock, _ := pem.Decode(privateKeyData)
	var pri *rsa.PrivateKey

	pri, parseErr := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if parseErr != nil {
		fmt.Println("Load private key error")
		panic(parseErr)

	}

	decodedData, _ := b64.URLEncoding.DecodeString(encryptedData)
	decryptedData, decryptErr := rsa.DecryptOAEP(sha1.New(), rand.Reader, pri, decodedData, nil)
	if decryptErr != nil {
		fmt.Println("Decrypt data error")
		panic(decryptErr)

	}

	fmt.Println(string(decryptedData))

}

var privateKeyData = []byte(`-----BEGIN RSA PRIVATE KEY-----
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

var publickKeyData = []byte(`-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsdoMC1ek1ysJt1xvpl3Z
ZSTta8oEzUvzENHPHN/b/o94+y3r+OMSn6S0CizJk92RicGz5L1yIAhNhbmfTkzr
CacQulpXe1olelJjdZbGIevFp7fyeiiRmGl0zrTA1bLVUKM5oqCaVtfVRH1hVeRv
WQy4phszMCVVSP5c6WdNrLRasPSWycuk8dMZrxwBGzK/jgeDMxbCTzxA+15cekoC
z8Qj+OXYsLAuyg62Ra9k6qKXZm/JIAW36J5U/rd9140ItXpaAFWmD1GlJUPFLh9h
04FMbIINnHSBd9Via/ptV95XaAQD5n3fMySDcnawTTTREeWYnskoji+zC8K3NrUF
bwIDAQAB
-----END RSA PUBLIC KEY-----`)

var privateKeyData_OLD = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAygGoUiTD+LjwZIgwFZyjiibWNQ2LM9xZ2pjKQGP8iUBtAuAW
629/Ofw8qxToMyixPrG4A7j8+KOPwYrWPGV6Og//4zm3cG+1hQvnNUWtMjHHBY8O
ByUPQ6/T8XHER1DxFBfnWfFLZ1yFX6oNNuvtLgOreI6ehehJd5IB/4mOjMvFEBgO
Eejado2n55VNdcFpdQ3RcvGV+f/rl/lsIM08QvL3lc5gqawj53sW9YZi1DL/uN48
R+ghvAYhtx2jpHDBvlH1NCF1rU6CynYsgV9QIksv0ihwl4T+k5F9ir0uv0WIS6kK
KS1SRpAprRKunos4PlE8l2+jC6LaJUPhDZlj/wIDAQABAoIBAHIcX5YPeLie2AUi
PW9n7aYT7DtJ7FGebw+h8dZP5Q8vWqUeKzRR5p+90hOemtCTcxSEVfucWyKlWoat
Q/oYJOR5t0YHi40zPWnr4G7ibkUFg3Sra/QzRh0pTON+La9PlO+R1TmkqcC4rgrt
R8u3mGK+5fUTM49XOXEXBJPyg5kaXQpiA4BoIRdRnCSitNxWA8kxMkQYJYlwAYab
cKo4Ik/J6+YGG7m2FtrUAWpWVUMBzEYOmGJ7JhSJ1u0UC/Oh1HOS1xlGopkmexbd
EygY3hTNWzHmYaYcYQs0f+8aVcVL64Gm0dtqvAHNnBvudMThhQgdYPc39mNLbrwI
ks4uS8ECgYEA9XfvcGKsNrHA0nqoPUPMT0Nfvv/4XCaKOYk25brH4LbqJPm6CiU6
uNlKFQsxzHPmx7OEK7EYVVZCbSO9s4t/xCzDVNbOZ9kDL6bkTX9DArLE4d6IRF/1
WW/AlNPuwVgxl0kcJILFtLqA1WoC5UWMhbYe2YB/Q3rCozmn0AiwyqECgYEA0qxd
KClKAMIsrB0WJ9gZEsJOpFi4q4g6T1BwT40Xj6Ul6o6DHi6hFhPgZAstqmnY0ANz
ezQ2yxtIm7zSy7S+nwDUycjY9riJcomc/YQZNA2QVM16hEv84VLwH1MVV2wkTb41
DWjbcg/ZNofZHl9AQIw7es+R3mmtDN+8BZOZSp8CgYBHtwmaUQm1VQtbswAyHfuz
8KApgklCSvQ5SRBj38UDrw0LTnZ+/k+Ar+MH8ORUskvrblQgG7ZbQD9Z+YYzzX6/
hsBuqe9Vwb4/jsfGqHagdDA3OTegmlRpE9A06xInJKggZfi15gry+UYok7dS2pXq
fsHWk8capOP2oiKYEeHs4QKBgF2KcLaDVrtte/5Tz+GTHtbodZidWCm5jAJpeeSo
hfye3G4AJxHArH+sBacGG5md88mwrpbWwTl/fMbBmWsfbsAU02ZhCozJtSWpGo6q
F7K4DwzIS4zwXHEDrWCLOF+fwaLPQKkalM1ZYh3HRc0ph9LhMQu/nEn/6/laYhar
yZWLAoGASvCrpFKn0qllMKNUetBmYFpgtjmnNuW7l0xT2UftkW6AuFjU19gKgXhe
I+uZciHQ8kIUHfNLYBbhETsF3iqsklKfeoIr23zYHLE5GpoC151IpKf4guoPbCHX
a1oCDuZm//f5HMePb9juJN0WR//d5jWuizAycZf41XoEd8Bqydg=
-----END RSA PRIVATE KEY-----`