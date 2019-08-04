package token

import (
	"encoding/json"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	clientInfo = "clientInfo"
)

type TokenClientInfo struct {
	ClientId string
	CorpCode string
	Scopes   []string
}

// CreateSignedJwtToken returns a JWT signed token with the given info.
func CreateSignedJwtToken(signingKey string, c *TokenClientInfo, expirationTimeInSec int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["alg"] = "HS256"
	token.Header["typ"] = "JWT"

	b, err := json.Marshal(*c)
	if err != nil {
		return "", err
	}
	token.Claims[clientInfo] = string(b)

	token.Claims["iat"] = time.Now().Unix()
	token.Claims["exp"] = time.Now().Add(time.Second * time.Duration(expirationTimeInSec)).Unix()

	return token.SignedString([]byte(signingKey))
}

// ValidateSignedJwtToken validates and parses the give JWT signed token.
// If the token is valid, returns info parsed and nil error.
func ValidateSignedJwtToken(tokenStr string, signingKey string) (*TokenClientInfo, error) {
	token, err := jwt.Parse(tokenStr, func(jt *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	info := TokenClientInfo{}

	if str, ok := token.Claims[clientInfo].(string); ok {
		err := json.Unmarshal([]byte(str), &info)
		if err != nil {
			return nil, err
		}
		return &info, nil
	}

	return nil, fmt.Errorf("Failed to parse client info from token")
}
