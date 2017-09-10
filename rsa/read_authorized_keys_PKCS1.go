package rsa

/**
 * Created by John Tsantilis 
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 23/8/2017.
 */

import (
	"os"
	"bufio"
	"bytes"
	"strconv"
	"io"
	"io/ioutil"
	"fmt"
	"crypto/x509"
	"encoding/pem"
	"github.com/ianmcmahon/encoding_ssh"
	log "github.com/Sirupsen/logrus"

)

func main() {
	//Read in public key from file
	//log.Info("Calling readFileWithReadLine.....................................")
	line, err := ReadFileWithReadLine(os.Getenv("HOME") + "/Desktop/authorized_keys")
	if err == io.EOF {
		//Do nothing

	}else {
		log.Error(err); os.Exit(1)

	}

	// decode string ssh-rsa format to native type
	// pub_key is of type *rsa.PublicKey
	pub_key, err := ssh.DecodePublicKey(line)
	if err != nil {
		log.Error(err); os.Exit(1)

	}

	// Marshal to ASN.1 DER encoding
	pkix, err := x509.MarshalPKIXPublicKey(pub_key)
	if err != nil {
		log.Error(err); os.Exit(1)

	}

	// Encode to PEM format
	pemStructPublic := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PUBLIC KEY",
		Bytes: pkix,
	})

	//log.Info(string(pem))
	fmt.Println(string(pemStructPublic))
	fmt.Println()

	//==================================================================================================================
	//==================================================================================================================
	//Read in private key from file
	bytes, err := ioutil.ReadFile(os.Getenv("HOME") + "/Desktop/test.pem")
	if err != nil { fmt.Printf("%v\n", err); os.Exit(1) }

	// decode PEM encoding to ANS.1 PKCS1 DER
	block, _ := pem.Decode(bytes)
	if block == nil { fmt.Printf("No Block found in keyfile\n"); os.Exit(1) }
	if block.Type != "RSA PRIVATE KEY" { fmt.Printf("Unsupported key type"); os.Exit(1) }

	// parse DER format to a native type
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	// Encode to PEM format
	pemStructPivate := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})

	fmt.Println(string(pemStructPivate))

	// encode the public key portion of the native key into ssh-rsa format
	// second parameter is the optional "comment" at the end of the string (usually 'user@host')
	ssh_rsa, err := ssh.EncodePublicKey(key.PublicKey, "")

	fmt.Printf("%s\n", ssh_rsa)

}

func ReadFileWithReadLine(fn string) (string, error) {
	var counter int
	var line []string

	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
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
				break

			}

		}

		if err == io.EOF {
			break

		}

		line = append(line, buffer.String())
		log.Debug("Read " + strconv.Itoa(len(line[counter])) + " characters\n")

		// Process the line here.
		log.Debug(line[counter])
		counter++

	}

	return line[1], err

}