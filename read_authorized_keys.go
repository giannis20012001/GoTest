package main

import (
	"os/user"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

/**
 * Created by John Tsantilis 
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 23/8/2017.
 */

func main()  {
	var publicKey ssh.PublicKey
	usr, err := user.Current()
	if err != nil {
		panic(err)

	}

	publicKey = publicKeyFile(usr.HomeDir + "/.ssh/authorized_keys")

	fmt.Println(publicKey)

}

func publicKeyFile(file string) (ssh.PublicKey) {
	fmt.Println(file)
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil

	}

	key, err := ssh.ParsePublicKey(buffer)
	if err != nil {
		return nil

	}

	return key

}