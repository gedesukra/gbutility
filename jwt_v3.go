package gbutility

import (
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/twinj/uuid"
)

var accessTokenV3 = "accessToken"
var refreshTokenV3 = "refreshToken"

type TokenPayloadV3 struct {
	jwt.Payload
	Userid    string `json:"userid"`
	Username  string `json:"username"`
	TokenType string `json:"tokentype"`
	TokenRef  string `json:"tokenref,omitempty"`
}

type TokenRefreshV3 struct {
	Result   bool
	Err      error
	Msg      string
	Userid   string
	Username string
	TokenRef string
}

type TokenAccessV3 struct {
	Result   bool
	Err      error
	Msg      string
	Userid   string
	Username string
}

func GetTokenAccessV3(audience []string, userid string, username string, expired time.Duration) (string, error) {
	now := time.Now()
	var hs = jwt.NewHS256([]byte("s3cr3t-@pp-" + accessTokenV3))
	u4 := uuid.NewV4()
	// randomId := (hex.EncodeToString(uuid.NewV4()))
	randomId := u4.String()

	tokenPayload := TokenPayloadV3{
		Payload: jwt.Payload{
			Issuer:         "greatbold",
			Subject:        "jwtappkey",
			Audience:       jwt.Audience(audience),
			ExpirationTime: jwt.NumericDate(now.Add(expired)),
			NotBefore:      jwt.NumericDate(now.Add(expired)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          randomId,
		},
		Userid:    userid,
		Username:  username,
		TokenType: accessTokenV3,
	}

	token, err := jwt.Sign(tokenPayload, hs)
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func GetTokenRefreshV3(audience []string, userid string, username string, tokenRef string, expired time.Duration) (string, error) {
	now := time.Now()
	var hs = jwt.NewHS256([]byte("s3cr3t-@pp-" + refreshTokenV3))
	u4 := uuid.NewV4()
	// randomId := (hex.EncodeToString(uuid.NewV4()))
	randomId := u4.String()

	tokenPayload := TokenPayloadV3{
		Payload: jwt.Payload{
			Issuer:         "greatbold",
			Subject:        "jwtappkey",
			Audience:       jwt.Audience(audience),
			ExpirationTime: jwt.NumericDate(now.Add(expired)),
			NotBefore:      jwt.NumericDate(now.Add(expired)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          randomId,
		},
		Userid:    userid,
		Username:  username,
		TokenType: refreshTokenV3,
		TokenRef:  tokenRef,
	}

	token, err := jwt.Sign(tokenPayload, hs)
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func VerifyTokenAccessV3(audience []string, token string) TokenAccessV3 {
	var result TokenAccessV3

	var hs = jwt.NewHS256([]byte("s3cr3t-@pp-" + accessTokenV3))

	var (
		now = time.Now()
		aud = jwt.Audience(audience)
		// Validate claims "iat", "exp" and "aud".
		iatValidator = jwt.IssuedAtValidator(now)
		expValidator = jwt.ExpirationTimeValidator(now)
		audValidator = jwt.AudienceValidator(aud)

		// Use jwt.ValidatePayload to build a jwt.VerifyOption.
		// Validators are run in the order informed.
		tokenPayload    TokenPayloadV3
		validatePayload = jwt.ValidatePayload(&tokenPayload.Payload, iatValidator, expValidator, audValidator)
	)

	_, err := jwt.Verify([]byte(token), hs, &tokenPayload, validatePayload)
	if err != nil {
		switch err {
		case jwt.ErrIatValidation:
			result = TokenAccessV3{false, err, "Unable validate iat", "", ""}
			return result
		case jwt.ErrExpValidation:
			result = TokenAccessV3{false, err, "Expired token", "", ""}
			return result
		case jwt.ErrAudValidation:
			result = TokenAccessV3{false, err, "Unable validate jwt au", "", ""}
			return result
		default:
			result = TokenAccessV3{false, err, "Unknown error", "", ""}
			return result
		}
	}

	result = TokenAccessV3{true, err, "", tokenPayload.Userid, tokenPayload.Username}

	return result
}

func VerifyTokenRefreshV3(audience []string, token string) TokenRefreshV3 {
	var result TokenRefreshV3

	var hs = jwt.NewHS256([]byte("s3cr3t-@pp-" + refreshTokenV3))

	var (
		now = time.Now()
		aud = jwt.Audience(audience)
		// Validate claims "iat", "exp" and "aud".
		iatValidator = jwt.IssuedAtValidator(now)
		expValidator = jwt.ExpirationTimeValidator(now)
		audValidator = jwt.AudienceValidator(aud)

		// Use jwt.ValidatePayload to build a jwt.VerifyOption.
		// Validators are run in the order informed.
		tokenPayload    TokenPayloadV3
		validatePayload = jwt.ValidatePayload(&tokenPayload.Payload, iatValidator, expValidator, audValidator)
	)

	_, err := jwt.Verify([]byte(token), hs, &tokenPayload, validatePayload)
	if err != nil {
		switch err {
		case jwt.ErrIatValidation:
			result = TokenRefreshV3{false, err, "Unable validate iat", "", "", ""}
			return result
		case jwt.ErrExpValidation:
			result = TokenRefreshV3{false, err, "Expired token", "", "", ""}
			return result
		case jwt.ErrAudValidation:
			result = TokenRefreshV3{false, err, "Unable validate jwt aud", "", "", ""}
			return result
		default:
			result = TokenRefreshV3{false, err, "Unknown error", "", "", ""}
			return result
		}
	}

	result = TokenRefreshV3{true, err, "ok", tokenPayload.Userid, tokenPayload.Username, tokenPayload.TokenRef}

	return result
}

func VerifyExpiredTokenAccessV3(audience []string, token string) (bool, error, string) {
	var (
		res bool
		err error
		msg string
	)

	var hs = jwt.NewHS256([]byte("s3cr3t-@pp-" + accessTokenV3))

	var (
		now = time.Now()
		aud = jwt.Audience(audience)
		// Validate claims "iat", "exp" and "aud".
		iatValidator = jwt.IssuedAtValidator(now)
		expValidator = jwt.ExpirationTimeValidator(now)
		audValidator = jwt.AudienceValidator(aud)

		// Use jwt.ValidatePayload to build a jwt.VerifyOption.
		// Validators are run in the order informed.
		tokenPayload    TokenPayloadV3
		validatePayload = jwt.ValidatePayload(&tokenPayload.Payload, iatValidator, expValidator, audValidator)
	)

	_, err = jwt.Verify([]byte(token), hs, &tokenPayload, validatePayload)
	if err != nil {
		switch err {
		case jwt.ErrIatValidation:
			res = false
			msg = "Unable validate iat"
			return res, err, msg
		case jwt.ErrExpValidation:
			res = true
			msg = "Expired token"
			return res, err, msg
		case jwt.ErrAudValidation:
			res = false
			msg = "Unable validate jwt aud"
			return res, err, msg
		default:
			res = false
			msg = "Unknown error"
			return res, err, msg
		}
	}

	res = false
	msg = "Unable refresh token,because token not expired"
	return res, err, msg
}
