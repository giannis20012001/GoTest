package main

/**
 * Created by John Tsantilis
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 7/9/2017.
 */

import (
	"fmt"
	"time"
	"bytes"
	"strings"
	"net/http"
	"os/user"
	"io/ioutil"
	"crypto/rsa"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/giannis20012001/GoTest/util"
	b64 "encoding/base64"
	log "github.com/Sirupsen/logrus"

)

//RSA KEYS AND INITIALISATION
const (
	privKeyPath = "/home/lumi/Desktop/ssh_keys/newlumimainkeypair.pem"
	pubKeyPath = "/home/lumi/Desktop/ssh_keys/newlumimainkeypair.pub"

)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey

)

//STRUCT DEFINITIONS
type UserCredentials struct {
	Username	string  `json:"username"`
	Password	string	`json:"password"`

}

type User struct {
	ID			int 	`json:"id"`
	Name		string  `json:"name"`
	Username	string  `json:"username"`
	Password	string	`json:"password"`

}

type Response struct {
	Data	string	`json:"data"`

}

type Token struct {
	Token 	string    `json:"token"`

}

//=====================================================================================================================
//=====================================================================================================================
func main() {
	//initKeys()

	url := "http://arcadia-sc.euprojects.net/api/v1/node/497/config"
	fmt.Println("URL:>", url)

	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)

	}

	nid := "497"
	publicKey, _ := util.GetPublicKey(usr.HomeDir + "/Desktop/authorized_keys")
	arrPublicKey, _ := util.ParseRsaPublicKeyFromPemStr(publicKey)
	arrNid := []byte(nid)

	uEnc, err := util.RsaEncrypt(*arrPublicKey, arrNid)
	if err != nil {
		log.Error(err)

	}

	authorizationKey := b64.URLEncoding.EncodeToString(uEnc)

	var jsonStr = []byte("parameters")
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + authorizationKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	//StartServer()


}
//=====================================================================================================================
//=====================================================================================================================

func fatal(err error) {
	if err != nil {
		log.Fatal(err)

	}

}

func initKeys(){
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)

}

//SERVER ENTRY POINT
func StartServer(){
	// Non-Protected Endpoint(s)
	http.HandleFunc("/login", LoginHandler)

	// Protected Endpoints
	http.Handle("/resource", negroni.New(
		negroni.HandlerFunc(ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(ProtectedHandler)),
	))

	log.Println("Now listening...")
	http.ListenAndServe(":8080", nil)

}

//////////////////////////////////////////
/////////////ENDPOINT HANDLERS////////////
/////////////////////////////////////////
func ProtectedHandler(w http.ResponseWriter, r *http.Request){
	response := Response{"Gained access to protected resource"}
	JsonResponse(response, w)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user UserCredentials

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")

		return

	}

	if strings.ToLower(user.Username) != "someone" {
		if user.Password != "p@ssword" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Error logging in")
			fmt.Fprint(w, "Invalid credentials")

			return

		}

	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error extracting the key")
		fatal(err)

	}

	tokenString, err := token.SignedString(signKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		fatal(err)

	}

	response := Token{tokenString}
	JsonResponse(response, w)

}

//AUTH TOKEN VALIDATION
func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}

}

//HELPER FUNCTIONS
func JsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return

	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}