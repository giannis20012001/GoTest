package main

import (
	"os"
	"io"
	"fmt"
	"bufio"
	"bytes"
	"strconv"
	ssh2pem "github.com/ssh-vault/ssh2pem"
	log "github.com/Sirupsen/logrus"

	"encoding/pem"

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

	/*test, _ := ssh2pem.DecodePublicKey(line)
	fmt.Println(test)*/
	yolo, _ := ssh2pem.GetPublicKeyPem(line)

	// Encode to PEM format
	pem := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PUBLIC KEY",
		Bytes: yolo,

	})

	//log.Info(string(pem))
	fmt.Println(string(pem))

}