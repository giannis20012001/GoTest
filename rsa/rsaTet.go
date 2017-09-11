package rsa

/**
 * Created by John Tsantilis
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 5/7/2017.
 */

import (
	"encoding/pem"
	"crypto/x509"
	"errors"
	"crypto/rsa"
	"crypto/rand"
	b64 "encoding/base64"

	"fmt"
)

var data64 string = "RKWQoWfk2Z6da5ckU6cxP+7a5u896Y6ZU6EYABsRo6RCSDWQRc6Er1/mULV92MXIkJzcJlNmawCwopRgJYAZF+UlOp6QDqgCE8eP9fD0tHxe2QLPCMz2c3zRpaaVTfs6IxMqqrbLIIziXtODgxIVrlohQ7+Hn6WPxi1sEvxlnfdtMC2jttGUafgcpN3oyXqjRQUnRy+q/hbG+tMbCto+oRy+DCamP1rxjMr0gpsZcPYciWzIS7aLx1Tqvyuid4tGJQ7HTJo3GmTN4+7Hhsm4TdfbL0rIbOJ04fRRVAr7XPtNykPVhfgLaJ2wILYUCYKmz55P9UJLs5jE2/AvXvwJRw=="

func main() {
	data, err := b64.StdEncoding.DecodeString(data64)
	if err != nil {
		fmt.Println(err)

	}

	origData, err := RsaDecrypt(data, privateKey)
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println("rsa-------" + string(origData))

}

func RsaEncrypt(origData []byte, pubKey []byte) ([]byte, error) {
	block, _ := pem.Decode(pubKey)
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

func RsaDecrypt(ciphertext []byte, privKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privKey)
	if block == nil {
		return nil, errors.New("private key error!")

	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err

	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)

}

var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
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

var publicKey = []byte(`-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArKJkTUdIQ5y6AhOnrX8d
YyQsxEapkTIiRNIvJ1ZIRctXDlgJPNgyZDwj/sjd8Xz+NKc8PYt0kzM5mTS3MSGH
pUkvDjpItzdVrQpYLsshaPIY/9qtMFRNioOLourNI2+vrg8cHPXOP98KcUHrovho
AQ6BCW80nzKsy1pVebKnrxeHyRhskM32lGRpXSjevVZVagFDzsqlM0REkBOhoizu
OR47KZMDhvjz7yLyOonaaeOrKW2tJYRf8fxUgNOXraTB2vEVVMMLpBD3Knu2kzYx
gfqHEMMucyp/6M3rQd33LtkedOt3KpJ0iCtD7O941/FFXH51lnkoLkKt+Zez3jlp
zwIDAQAB
-----END RSA PUBLIC KEY-----`)