package util

/**
 * Created by John Tsantilis
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 5/9/2017.
 */

import (
	"os"
	"bufio"
	"bytes"
	"io"
	"crypto/rsa"
	"crypto/x509"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"errors"
	"github.com/ianmcmahon/encoding_ssh"

)

func GetPrivateKeyPemStr(path string) (string, error) {
	//Read in private key from file
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%v\n", err); os.Exit(1)

	}

	// decode PEM encoding to ANS.1 PKCS1 DER
	block, _ := pem.Decode(bytes)
	if block == nil {
		fmt.Println("No Block found in keyfile")

		return "", nil

	}

	if block.Type != "RSA PRIVATE KEY" {
		fmt.Println("Unsupported key type")

		return "", nil

	}

	// parse DER format to a native type
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	privateKey := ExportRsaPrivateKeyAsPemStr(key)

	return privateKey, err

}

func GetPrivateKeyPem(path string)  ([]byte) {
	//Read in private key from file
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)

	}

	// decode PEM encoding to ANS.1 PKCS1 DER
	block, _ := pem.Decode(bytes)
	if block == nil {
		fmt.Println("No Block found in keyfile")

	}

	if block.Type != "RSA PRIVATE KEY" {
		fmt.Println("Unsupported key type")

	}

	// parse DER format to a native type
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	// Encode to PEM format
	pemStructPivate := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})

	fmt.Println(string(pemStructPivate))

	return pemStructPivate

}

func GetPublicKeyPemStr(path string) (string, error) {
	//Read in public key from file
	fmt.Println("Calling readFileWithReadLine.....................................")
	line, err := readFileWithReadLine(path)
	if err == io.EOF {
		//Do nothing

	}else {
		fmt.Println(err)

		return "", err

	}

	// decode string ssh-rsa format to native type
	// pub_key is of type *rsa.PublicKey
	pub_key, err := ssh.DecodePublicKey(line)
	if err != nil {
		fmt.Println(err)

		return "", err

	}

	publicKey, err := ExportRsaPublicKeyAsPemStr(pub_key.(*rsa.PublicKey))

	if err != nil {
		fmt.Println(err)

		return "", err

	}

	return publicKey, err

}

func GetPublicKeyPem(path string) ([]byte) {
	//Read in public key from file
	fmt.Println("Calling readFileWithReadLine.....................................")
	line, err := readFileWithReadLine(path)
	if err == io.EOF {
		//Do nothing

	}else {
		fmt.Println(err); os.Exit(1)

	}

	// decode string ssh-rsa format to native type
	// pub_key is of type *rsa.PublicKey
	pub_key, err := ssh.DecodePublicKey(line)
	if err != nil {
		fmt.Println(err)

	}

	yolo := pub_key.(*rsa.PublicKey)

	// Marshal to ASN.1 DER encoding
	pkix, err := x509.MarshalPKIXPublicKey(yolo)
	if err != nil {
		fmt.Println(err)

	}

	// Encode to PEM format
	pemStructPublic := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PUBLIC KEY",
		Bytes: pkix,
	})

	//fmt.Println(string(pemStructPublic))

	return pemStructPublic

}

func RsaEncrypt(publicKey []byte, origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		err := errors.New("public key error")
		fmt.Println(err)

		return nil, err

	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)

		return nil, err

	}

	pub := pubInterface.(*rsa.PublicKey)

	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)

}

func RsaDecrypt(privateKey []byte, ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		err := errors.New("private key error")
		fmt.Println(err)

		return nil, err

	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)

		return nil, err

	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)

}

func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, _ := rsa.GenerateKey(rand.Reader, 2048)

	return privkey, &privkey.PublicKey

}

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		fmt.Println(err)

		return "", err

	}

	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return string(pubkey_pem), nil

}

func ExportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)

	return string(privkey_pem)

}

func ParseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		err := errors.New("failed to parse PEM block containing the key")
		fmt.Println(err)

		return nil, err

	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)

		return nil, err

	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil

	default:
		break // fall through

	}

	err = errors.New("failed to parse PEM block containing the key")
	fmt.Println(err)

	return nil, err

}

func ParseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		err := errors.New("failed to parse PEM block containing the key")
		fmt.Println(err)

		return nil, err

	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)

		return nil, err

	}

	return priv, nil
}

func readFileWithReadLine(fn string) (string, error) {
	var counter int
	var line []string

	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
		fmt.Println(err)

		return line[1], err

	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	for {
		var buffer bytes.Buffer
		var l []byte
		var isPrefix bool

		for {
			l, isPrefix, err = reader.ReadLine()
			buffer.Write(l)

			// If we've reached the end of the line, stop reading.
			if !isPrefix {
				break
			}

			// If we're just at the EOF, break
			if err != nil {
				fmt.Println(err)

				break

			}

		}

		if err == io.EOF {
			break

		}

		line = append(line, buffer.String())
		//fmt.Println("Read " + strconv.Itoa(len(line[counter])) + " characters\n")

		// Process the line here.
		//fmt.Println(line[counter])
		counter++

	}

	return line[1], err

}