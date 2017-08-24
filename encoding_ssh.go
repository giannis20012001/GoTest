package main

/**
 * Created by John Tsantilis 
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 23/8/2017.
 */

import (
	"os"
	"bufio"
	"bytes"
	"io"
	"strconv"
	"crypto/x509"
	"encoding/pem"
	"github.com/ianmcmahon/encoding_ssh"
	log "github.com/Sirupsen/logrus"

)

func main() {
	//Read in public key from file
	log.Info("Calling readFileWithReadLine.....................................")
	line, err := readFileWithReadLine(os.Getenv("HOME") + "/.ssh/authorized_keys")

	if err == io.EOF {
		//Do nothing

	}else {
		log.Error(err); os.Exit(1)

	}

	pub_key, err := ssh.DecodePublicKey(line)
	if err != nil { log.Info("%v\n", err); os.Exit(1) }
	// pub_key is of type *rsa.PublicKey

	// Marshal to ASN.1 DER encoding
	pkix, err := x509.MarshalPKIXPublicKey(pub_key)
	if err != nil { log.Info("%v\n", err); os.Exit(1) }

	// Encode to PEM format
	pem := string(pem.EncodeToMemory(&pem.Block{
		Type: "RSA PUBLIC KEY",
		Bytes: pkix,
	}))

	log.Info("%s", pem)

}

func readFileWithReadLine(fn string) (string, error) {
	var counter int
	var line []string

	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
		return line[0], err

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

	return line[0], err

}