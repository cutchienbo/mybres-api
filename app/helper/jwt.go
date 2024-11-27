package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type UserJWTSubject struct {
	Id 		int64 	`json:"id"`
	Name 	string	`json:"name"`
}

type Header struct {
	Alg string
	Typ string
}

type Payload struct {
	Sub UserJWTSubject
	Exp string
}

type Signature struct {
	HeaderEncode	string
	PayloadEncode	string
}

func GetRefreshTokenSub(jwt string) UserJWTSubject{
	var jwtElement = strings.Split(strings.Trim(jwt, "Bearer "), ".")

	var payload Payload

	payloadJson, _ := base64.RawURLEncoding.DecodeString(jwtElement[1])

	json.Unmarshal(payloadJson, &payload)

	return payload.Sub
}

func CheckJWT(jwt string, tokenType string) error {
	var secretKey string

	if tokenType == "access" {
		secretKey = os.Getenv("SECRET_KEY_ACCESS")
	} else{
		secretKey = os.Getenv("SECRET_KEY_REFRESH")
	}

	var jwtElement = strings.Split(strings.Trim(jwt, "Bearer "), ".")

	if jwt == "" {
		return errors.New("Token not found")
	}
	
	// Start check valid token
	var signature Signature = Signature{
		HeaderEncode: jwtElement[0],
		PayloadEncode: jwtElement[1],
	}

	signatureJson, _ := json.Marshal(signature)

	h := hmac.New(sha256.New, []byte(secretKey))

	h.Write(signatureJson)

	signatureHmac := h.Sum(nil)

	signatureEncode := base64.RawURLEncoding.EncodeToString(signatureHmac)

	if check := signatureEncode == jwtElement[2]; !check {
		return errors.New("Token not valid")
	}
	// End check valid token
	
	// Start check exp token
	payloadJson, _ := base64.RawURLEncoding.DecodeString(jwtElement[1])

	var payload Payload

	json.Unmarshal(payloadJson, &payload)

	exp, _ := time.Parse("02-01-2006 15:04:05", payload.Exp)

	currentTime := GetCurrentTimeVN()

	if checkTime := currentTime.Before(exp); !checkTime {
		return errors.New("Token expired")
	}
	// End check exp token

	return nil
}

func GenerateAccessToken(user UserJWTSubject) string {
	var secretKey = os.Getenv("SECRET_KEY")

	var header Header = Header{
		Alg : "sha256",
		Typ : "jwt",
	}

	headerJson, _ := json.Marshal(header)

	var headerEncode = base64.RawURLEncoding.EncodeToString(headerJson)

	currentTime := GetCurrentTimeVN()

	tokenExp := currentTime.Add(time.Hour * 1).Format("02-01-2006 15:04:05")

	var payload Payload = Payload{
		Sub : user,
		Exp : tokenExp,
	}

	payloadJson, _ := json.Marshal(payload)

	var payloadEncode = base64.RawURLEncoding.EncodeToString(payloadJson)

	var signature Signature = Signature{
		HeaderEncode: headerEncode,
		PayloadEncode: payloadEncode,
	}

	signatureJson, _ := json.Marshal(signature)

	h := hmac.New(sha256.New, []byte(secretKey))

	h.Write(signatureJson)

	signatureHmac := h.Sum(nil)	

	signatureEncode := base64.RawURLEncoding.EncodeToString(signatureHmac)

	token := fmt.Sprintf("%s.%s.%s", headerEncode, payloadEncode, signatureEncode)

	return token
}


func GenerateRefreshToken(user UserJWTSubject) string {
	var secretKey = os.Getenv("SECRET_KEY_REFRESH")

	var header Header = Header{
		Alg : "sha256",
		Typ : "jwt",
	}

	headerJson, _ := json.Marshal(header)

	var headerEncode = base64.RawURLEncoding.EncodeToString(headerJson)

	currentTime := GetCurrentTimeVN()

	tokenExp := currentTime.Add(time.Hour * 24 * 7).Format("02-01-2006 15:04:05")

	var payload Payload = Payload{
		Sub : user,
		Exp : tokenExp,
	}

	payloadJson, _ := json.Marshal(payload)

	var payloadEncode = base64.RawURLEncoding.EncodeToString(payloadJson)

	var signature Signature = Signature{
		HeaderEncode: headerEncode,
		PayloadEncode: payloadEncode,
	}

	signatureJson, _ := json.Marshal(signature)

	h := hmac.New(sha256.New, []byte(secretKey))

	h.Write(signatureJson)

	signatureHmac := h.Sum(nil)	

	signatureEncode := base64.RawURLEncoding.EncodeToString(signatureHmac)

	token := fmt.Sprintf("%s.%s.%s", headerEncode, payloadEncode, signatureEncode)

	return token
}