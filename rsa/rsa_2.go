package rsa

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
		fmt.Println(err)

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
MIIEogIBAAKCAQEA6yUrWecafi6dBgBjoCeVZY7AM0sRqdWcx8UV8VFEpbsManyM
HIb8YC6FREXAEKgafCqU2s3j2NJudHV/BHZQKILFJxxEOo5jwtNHwzkwaOX62Rae
LhkMf+aT79g/YN2o1XghAt5TRgWeN9FRJH4909bb3ebamIQh7V+IqpWO2zJHZqx2
Cns20LpU/ep/WE5npYTALDlTOCbULSPI9isIRGH44Ucspg3vazoSXJx7iYc9z4t+
sFuJ9o7PG/LaBq97lKGxgESgL7lVSjvU+ZAXfo1A7BLZ7pW/W9TgNwJ8g/PHa25c
bPbK/gtU0Q57GOtfLtSOskljPLGYR+AQKLsCFQIDAQABAoIBAHLolQi38LBc4wr9
wbGP5mfcnHv0YUtXQeg1nYVxq1Q51dVry6RdKpNt6F32n+cD3v1yaY+LMZ7RJnzu
tgVeM8m4LdPw8j6TOpkIWndCSS/Zwv23GdF5GCygZwa77CzP5SN8MtWOA6+iKzEn
rrlgn6IXyxXjLEt/tnyjRMvnjgn61hMSehJjRQs1fMxGKQHgy59zVM5v7LvwTjOL
+7xEkiD8XdsC4B7uy7Vnm7mkEvTCMzvVkUT5JHv4bwpFq4GQN79abFG06QX0a7Nf
2ZTC1SjVjcSNvPdQx6KCcF3LQPWcgRfmJZ90Dq3BBc/eUE5A5wO552lQ8OT4S03Z
e58u7UECgYEA/UFZN0RBqPRn7i/RbN2H/M1qlrxVCf/1pug47mqe0JKDkzutOSsG
ecXMXkmNbzq8L0oBhZpkMNXdulNRwRsCNGCHVeTvc5m+/YFH/S/6jHP9EZ6wnJZk
iRpQ9QjOGSg4aPETJ0FwqZNK3wYthU1W4n8O9LxtTSjpzYdy2/1ohL0CgYEA7bGT
JvJexsVDrXAF0EbuRudEKF4oA3w5eRSu4Exlhpw35s+X6WTxmaW4xPZ1qfcMF4m8
tPoQya7v0nbFJlsXsXTlnosYl/KccB5sM5R41MKRaTTU5v1//QMZ9zua060YFoyW
aVKlj1hefsUjbsfIDQ9uSYk1jEEDoRKHkhJUhDkCgYAO+mJwZlULzQioiaN1MNTu
GqgHKjRNVxoMOQfE8gFajI/DkW/5RZYodY5UtTzsKykeEt5sLGloif2HG45mQVas
Cll2tweCasFk9NRxIPlMfT+mXyBK4oonoarQEyk9S6eqbTeYxsIHBXMUJaVjkONm
meUkjFBak+TgBvbAFAiucQKBgHOgw63Zj8NwKOkRKrLUHou9awmcLCjt4GOHbT7N
y0G9cvBEImk2YtVUqdqe7kRdWrOEnJkJYOtLv3yJrIRpIdCAxkbm8XLRYcqk8gvx
eQo/EE+2lK89uGpTfOkpRLseZC5r+6uGueVOnsFak08LvsSjsgnxxmvRILvVcL/d
TOH5AoGAD0Zg0HvSm3Lf4s8T8ANCLUttZkYm/2OOtcCR0yBc3hZSdtPo4la6DFKZ
G/+KXgI6psszh/MMl7uCvubDCsY4yBHuQT5KmPD73AFV6wyH3zuBgmNBw5oCsE15
0A4FQWcyiM6vpj+nmG45i3T26+83GyhGfghm7LT/alp84W5G/Wo=
-----END RSA PRIVATE KEY-----`)

var publicKey = []byte(`-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6yUrWecafi6dBgBjoCeV
ZY7AM0sRqdWcx8UV8VFEpbsManyMHIb8YC6FREXAEKgafCqU2s3j2NJudHV/BHZQ
KILFJxxEOo5jwtNHwzkwaOX62RaeLhkMf+aT79g/YN2o1XghAt5TRgWeN9FRJH49
09bb3ebamIQh7V+IqpWO2zJHZqx2Cns20LpU/ep/WE5npYTALDlTOCbULSPI9isI
RGH44Ucspg3vazoSXJx7iYc9z4t+sFuJ9o7PG/LaBq97lKGxgESgL7lVSjvU+ZAX
fo1A7BLZ7pW/W9TgNwJ8g/PHa25cbPbK/gtU0Q57GOtfLtSOskljPLGYR+AQKLsC
FQIDAQAB
-----END RSA PUBLIC KEY-----`)

var privateKeyData_OLD = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`)