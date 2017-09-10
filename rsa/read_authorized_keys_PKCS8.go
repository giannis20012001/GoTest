package rsa

import (
	"os"
	"io"
	"fmt"
	"bufio"
	"bytes"
	"strconv"
	"encoding/pem"
	log "github.com/Sirupsen/logrus"
	ssh2pem "github.com/ssh-vault/ssh2pem"

)

/**
 * Created by John Tsantilis 
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 31/8/2017.
 */

func main() {
	//Read in public key from file
	//log.Info("Calling readFileWithReadLine.....................................")
	line, err := readFileWithReadLine(os.Getenv("HOME") + "/Desktop/authorized_keys")
	if err == io.EOF {
		//Do nothing

	}else {
		log.Error(err); os.Exit(1)

	}

	test, _ := ssh2pem.DecodePublicKey(line)
	fmt.Println(test)
	yolo, _ := ssh2pem.GetPublicKeyPem(line)

	// Encode to PEM format
	pem := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PUBLIC KEY",
		Bytes: yolo,

	})

	//log.Info(string(pem))
	fmt.Println(string(pem))

}

func readFileWithReadLine(fn string) (string, error) {
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