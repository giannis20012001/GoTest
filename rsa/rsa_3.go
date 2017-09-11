package rsa

/**
 * Created by John Tsantilis
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 1/9/2017.
 */

import (
	"encoding/pem"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"crypto/rand"
	b64 "encoding/base64"

)

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
	fmt.Println()
	//==================================================================================================================
	//==================================================================================================================
	privateKeyBlock, _ := pem.Decode(privateKeyData)
	var pri *rsa.PrivateKey

	pri, parseErr := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if parseErr != nil {
		//fmt.Println("Load private key error")
		fmt.Println(parseErr)

	}

	encryptedData := uEnc
	decodedData, err := b64.URLEncoding.DecodeString(encryptedData)
	if err != nil {
		fmt.Println(err)

	}

	decryptedData, decryptErr := rsa.DecryptPKCS1v15(rand.Reader, pri, decodedData)
	//decryptedData, decryptErr := rsa.DecryptOAEP(sha1.New(), rand.Reader, pri, decodedData, nil)
	if decryptErr != nil {
		//fmt.Println("Decrypt data error")
		fmt.Println(decryptErr)

	}

	fmt.Println(string(decryptedData))

}

var privateKeyData = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEArKJkTUdIQ5y6AhOnrX8dYyQsxEapkTIiRNIvJ1ZIRctXDlgJ
PNgyZDwj/sjd8Xz+NKc8PYt0kzM5mTS3MSGHpUkvDjpItzdVrQpYLsshaPIY/9qt
MFRNioOLourNI2+vrg8cHPXOP98KcUHrovhoAQ6BCW80nzKsy1pVebKnrxeHyRhs
kM32lGRpXSjevVZVagFDzsqlM0REkBOhoizuOR47KZMDhvjz7yLyOonaaeOrKW2t
JYRf8fxUgNOXraTB2vEVVMMLpBD3Knu2kzYxgfqHEMMucyp/6M3rQd33LtkedOt3
KpJ0iCtD7O941/FFXH51lnkoLkKt+Zez3jlpzwIDAQABAoIBACgdkEs43jz/5WVo
JdW2LSEFnfV6KhsYWEg/wz9T2HWHe9JUeMkPwB30r7Sb/p8IGZYoqmHuzwcJpz/H
tS4QiZHKAMpAlvckP593QOiWwUu+vjpuGUKaxG4EhWU1RLgQUvWSg0jjgarr7GRH
GjeDm79rHGcR1VuNDgQvGJ+v+TFBVpzgMR9uBuRAISiZ1pb/eOTP/SbfcDNO4EXv
0ZIcgtzQYfEjrmV1QX252rWieqnxdSeQBPTE7nSFPPfJNHk+zIdtaupe+SkHg6cH
uML0SPGyc3mIUCQKWj2D1wDstYfUDz9GcgvUv2x1lRzIYrUoXWOZPFxIUXDSEhIU
hnA+FNkCgYEA9hH7KgR+aBn2J199ve9U83nzRO2R8EBdbHLR9wd1hwxIZELMEZ+1
YQ41p5mLkQhFDMmJDr0UKYtchDSO460ZP3gkTVxWe66ynIBL0SEtjX+Q7ZJKwjbr
9GikJp/LFiiQQgK8q/ZWSzwpP3ePSJRb8EhVpaQKeVmvF9p3JYb9N6sCgYEAs5nI
uOn0+ouX5/pKKidiIEos3AShXWgMsMx7ldQN1zVjMvpiS1SV4GPWcHoTLS8NkC/D
U4GJoCDeJ2PqzHhP1mYP3T0YTTBglBQmuLPcbpZic9KxNmuj6+iNV46HTDI/7whE
l0o+KE031WJfPSo7im81eI9KzFRkM2TLdHUMIm0CgYEAlV5Y3iqYKM0JlODsTS1I
nfs65m5nljUcAMa6GX/XocCc+O9rPI975IxhmXklNUaV06tKnT29TXKhGEdnLWFX
4CUntCYHAVEMbt+aJjpDko8LBIs3AimglV4ntqJs/uhkmQ1cRe5kd4rvZu1GdsJI
BNWD1+Z5zSvZk1sz0/3bG70CgYBvY2Nv+/oqIcWW5E53EnHzYM2tr1l9Gvkz9b54
UEo9PSlHBq9L1SwXPRRUMgIOte4NjPHxqpd2rqvZdki7g6rQaABS0H9v8B8V9+GE
EFkYZUCuYO/ztpK1z0dKFGWpRkkMsO4JfsxKJooIV7iFsNm4o/xfx082puh2T383
vRhHiQKBgElAHlkFKpOkp5bQL9glnG2oiuPzVFzDUpB4KWpRlFnf6VnrZtVB+Iha
S0vgFOz4hm34b9vi/DtkvjJstoJnyw068vxAH8Ot/mZHHPlj1lYhIZO88GpPT0OV
L4hG3NP5KzZA4QmsyUMg9xlK7Ph2Eld7y37Wck2lIi5pAVaDZFus
-----END RSA PRIVATE KEY-----`)

var publickKeyData = []byte(`-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArKJkTUdIQ5y6AhOnrX8d
YyQsxEapkTIiRNIvJ1ZIRctXDlgJPNgyZDwj/sjd8Xz+NKc8PYt0kzM5mTS3MSGH
pUkvDjpItzdVrQpYLsshaPIY/9qtMFRNioOLourNI2+vrg8cHPXOP98KcUHrovho
AQ6BCW80nzKsy1pVebKnrxeHyRhskM32lGRpXSjevVZVagFDzsqlM0REkBOhoizu
OR47KZMDhvjz7yLyOonaaeOrKW2tJYRf8fxUgNOXraTB2vEVVMMLpBD3Knu2kzYx
gfqHEMMucyp/6M3rQd33LtkedOt3KpJ0iCtD7O941/FFXH51lnkoLkKt+Zez3jlp
zwIDAQAB
-----END RSA PUBLIC KEY-----`)