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

	url := "http://arcadia-sc.euprojects.net/api/v1/node/498/config"
	fmt.Println("URL:>", url)

	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)

	}

	nid := "15750035-3e0e"
	publicKey, _ := util.GetPublicKey(usr.HomeDir + "/Desktop/authorized_keys")
	arrPublicKey, yolo, _ := util.ParseRsaPublicKeyFromPemStr(publicKey)
	fmt.Println(arrPublicKey)
	arrNid := []byte(nid)

	/*pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	pub := pubInterface.(*rsa.PublicKey)*/
	uEnc, err := util.RsaEncrypt(yolo, arrNid)
	if err != nil {
		log.Error(err)

	}

	authorizationKey := b64.URLEncoding.EncodeToString(uEnc)

	var jsonStr = []byte("parameters")
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorizationKey)

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

var privateKeyData = []byte(`-----BEGIN RSA PRIVATE KEY-----
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

var publickKeyData = []byte(`-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6yUrWecafi6dBgBjoCeV
ZY7AM0sRqdWcx8UV8VFEpbsManyMHIb8YC6FREXAEKgafCqU2s3j2NJudHV/BHZQ
KILFJxxEOo5jwtNHwzkwaOX62RaeLhkMf+aT79g/YN2o1XghAt5TRgWeN9FRJH49
09bb3ebamIQh7V+IqpWO2zJHZqx2Cns20LpU/ep/WE5npYTALDlTOCbULSPI9isI
RGH44Ucspg3vazoSXJx7iYc9z4t+sFuJ9o7PG/LaBq97lKGxgESgL7lVSjvU+ZAX
fo1A7BLZ7pW/W9TgNwJ8g/PHa25cbPbK/gtU0Q57GOtfLtSOskljPLGYR+AQKLsC
FQIDAQAB
-----END RSA PUBLIC KEY-----`)