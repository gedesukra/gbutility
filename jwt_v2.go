package gbutility

import (

	// "fmt"
	"time"

	"github.com/gbrlsnchs/jwt/v2"
	"github.com/twinj/uuid"
)

type Token struct {
	*jwt.JWT
	IsLoggedIn  bool   `json:"isLoggedIn"`
	CustomField string `json:"customField,omitempty"`
}

func GenJwtToken(username string) string {
	now := time.Now()
	// Define a signer.
	hs256 := jwt.NewHS256("s3cr3t-@pp-###")
	u4 := uuid.NewV4()
	// randomId := (hex.EncodeToString(uuid.NewV4()))
	randomId := u4.String()

	jot := &Token{
		JWT: &jwt.JWT{
			Issuer:         "gbapp",
			Subject:        "jwtappkey",
			Audience:       username,
			ExpirationTime: now.Add(24 * 30 * 12 * time.Hour).Unix(),
			NotBefore:      now.Add(30 * time.Minute).Unix(),
			IssuedAt:       now.Unix(),
			ID:             randomId,
		},
		IsLoggedIn:  true,
		CustomField: username,
	}

	jot.SetAlgorithm(hs256)
	jot.SetKeyID("kid")

	bolRes := true

	payload, err := jwt.Marshal(jot)
	if err != nil {
		// logger.Println("Error marshall jwt = %s", err.Error())
		bolRes = false
	}

	if bolRes {
		token, err := hs256.Sign(payload)
		if err != nil {
			// logger.Println("Error sign jwt = %s", err.Error())
			bolRes = false
		}
		if bolRes {
			// logger.Println("new jwt = %s", string(token))
			return string(token)
		} else {
			return ""
		}
	} else {
		return ""
	}
}

func VerJwtToken(token string) (bool, string) {
	res := true
	msg := ""

	// Timestamp the beginning.
	now := time.Now()
	// Define a signer.
	hs256 := jwt.NewHS256("s3cr3t-@pp-###")

	payload, sig, err := jwt.Parse(token)
	if err != nil {
		res = false
		msg = "Unable parse jwt"
		return res, msg
	}

	if err = hs256.Verify(payload, sig); err != nil {
		res = false
		msg = "Unable verify jwt"
		return res, msg
	}

	var jot Token
	if err = jwt.Unmarshal(payload, &jot); err != nil {
		res = false
		msg = "Unable unmarshall token"
		return res, msg
	}

	// Validate fields.
	iatValidator := jwt.IssuedAtValidator(now)
	expValidator := jwt.ExpirationTimeValidator(now)
	audValidator := jwt.AudienceValidator(jot.CustomField)
	if err = jot.Validate(iatValidator, expValidator, audValidator); err != nil {
		switch err {
		case jwt.ErrIatValidation:
			// handle "iat" validation error
			res = false
			msg = "Unable validate jwt iat"
			return res, msg
		case jwt.ErrExpValidation:
			// handle "exp" validation error
			res = false
			msg = "expired"
			return res, msg
		case jwt.ErrAudValidation:
			// handle "aud" validation error
			res = false
			msg = "Unable validate jwt aud"
			return res, msg
		}
	}

	return true, jot.CustomField
}
