package gbutility

import (
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/twinj/uuid"
)

var accessTokenGB = "accessToken"
var refreshTokenGB = "refreshToken"

type TokenPayloadGB struct {
	jwt.Payload
	Userid    string `json:"userid"`
	Username  string `json:"username"`
	TokenType string `json:"tokentype"`
	TokenRef  string `json:"tokenref,omitempty"`
}

type TokenRefreshGB struct {
	Result      bool
	Err         error
	Msg         string
	Userid      string
	Username    string
	TokenRef    string
	Issuedtime  time.Time
	Expiredtime time.Time
}

type TokenAccessGB struct {
	Result      bool
	Err         error
	Msg         string
	Userid      string
	Username    string
	Issuedtime  time.Time
	Expiredtime time.Time
}

func GetTokenAccessGB(audience []string, userid string, username string, expired time.Duration) (string, error) {
	now := time.Now()
	var hs = jwt.NewHS256([]byte("s3cr3t-@pp-" + accessTokenGB))
	u4 := uuid.NewV4()
	// randomId := (hex.EncodeToString(uuid.NewV4()))
	randomId := u4.String()

	tokenPayload := TokenPayloadGB{
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
		TokenType: accessTokenGB,
	}

	token, err := jwt.Sign(tokenPayload, hs)
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func GetTokenRefreshGB(audience []string, userid string, username string, tokenRef string, expired time.Duration) (string, error) {
	now := time.Now()
	var hs = jwt.NewHS256([]byte("s3cr3t-@pp-" + refreshTokenGB))
	u4 := uuid.NewV4()
	// randomId := (hex.EncodeToString(uuid.NewV4()))
	randomId := u4.String()

	tokenPayload := TokenPayloadGB{
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
		TokenType: refreshTokenGB,
		TokenRef:  tokenRef,
	}

	token, err := jwt.Sign(tokenPayload, hs)
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func VerifyTokenAccessGB(audience []string, token string) TokenAccessGB {
	var result TokenAccessGB

	var hs = jwt.NewHS256([]byte("s3cr3t-@pp-" + accessTokenGB))

	var (
		now = time.Now()
		aud = jwt.Audience(audience)
		// Validate claims "iat", "exp" and "aud".
		iatValidator = jwt.IssuedAtValidator(now)
		expValidator = jwt.ExpirationTimeValidator(now)
		audValidator = jwt.AudienceValidator(aud)

		// Use jwt.ValidatePayload to build a jwt.VerifyOption.
		// Validators are run in the order informed.
		tokenPayload    TokenPayloadGB
		validatePayload = jwt.ValidatePayload(&tokenPayload.Payload, iatValidator, expValidator, audValidator)
	)

	_, err := jwt.Verify([]byte(token), hs, &tokenPayload, validatePayload)
	if err != nil {
		switch err {
		case jwt.ErrIatValidation:
			result = TokenAccessGB{false, err, "Unable validate iat", "", "", tokenPayload.IssuedAt.Time, tokenPayload.ExpirationTime.Time}
			return result
		case jwt.ErrExpValidation:
			result = TokenAccessGB{false, err, "Expired token", "", "", tokenPayload.IssuedAt.Time, tokenPayload.ExpirationTime.Time}
			return result
		case jwt.ErrAudValidation:
			result = TokenAccessGB{false, err, "Unable validate jwt au", "", "", tokenPayload.IssuedAt.Time, tokenPayload.ExpirationTime.Time}
			return result
		default:
			result = TokenAccessGB{false, err, "Unknown error", "", "", tokenPayload.IssuedAt.Time, tokenPayload.ExpirationTime.Time}
			return result
		}
	}

	result = TokenAccessGB{true, err, "", tokenPayload.Userid, tokenPayload.Username, tokenPayload.IssuedAt.Time, tokenPayload.ExpirationTime.Time}

	return result
}

func VerifyTokenRefreshGB(audience []string, token string) TokenRefreshGB {
	var result TokenRefreshGB

	var hs = jwt.NewHS256([]byte("s3cr3t-@pp-" + refreshTokenGB))

	var (
		now = time.Now()
		aud = jwt.Audience(audience)
		// Validate claims "iat", "exp" and "aud".
		iatValidator = jwt.IssuedAtValidator(now)
		expValidator = jwt.ExpirationTimeValidator(now)
		audValidator = jwt.AudienceValidator(aud)

		// Use jwt.ValidatePayload to build a jwt.VerifyOption.
		// Validators are run in the order informed.
		tokenPayload    TokenPayloadGB
		validatePayload = jwt.ValidatePayload(&tokenPayload.Payload, iatValidator, expValidator, audValidator)
	)

	_, err := jwt.Verify([]byte(token), hs, &tokenPayload, validatePayload)
	if err != nil {
		switch err {
		case jwt.ErrIatValidation:
			result = TokenRefreshGB{false, err, "Unable validate iat", "", "", "", tokenPayload.IssuedAt.Time, tokenPayload.ExpirationTime.Time}
			return result
		case jwt.ErrExpValidation:
			result = TokenRefreshGB{false, err, "Expired token", "", "", "", tokenPayload.IssuedAt.Time, tokenPayload.ExpirationTime.Time}
			return result
		case jwt.ErrAudValidation:
			result = TokenRefreshGB{false, err, "Unable validate jwt aud", "", "", "", tokenPayload.IssuedAt.Time, tokenPayload.ExpirationTime.Time}
			return result
		default:
			result = TokenRefreshGB{false, err, "Unknown error", "", "", "", tokenPayload.IssuedAt.Time, tokenPayload.ExpirationTime.Time}
			return result
		}
	}

	result = TokenRefreshGB{true, err, "ok", tokenPayload.Userid, tokenPayload.Username, tokenPayload.TokenRef, tokenPayload.IssuedAt.Time, tokenPayload.ExpirationTime.Time}

	return result
}

func VerifyExpiredTokenAccessGB(audience []string, token string) (bool, error, string) {
	var (
		res bool
		err error
		msg string
	)

	var hs = jwt.NewHS256([]byte("s3cr3t-@pp-" + accessTokenGB))

	var (
		now = time.Now()
		aud = jwt.Audience(audience)
		// Validate claims "iat", "exp" and "aud".
		iatValidator = jwt.IssuedAtValidator(now)
		expValidator = jwt.ExpirationTimeValidator(now)
		audValidator = jwt.AudienceValidator(aud)

		// Use jwt.ValidatePayload to build a jwt.VerifyOption.
		// Validators are run in the order informed.
		tokenPayload    TokenPayloadGB
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
