package user

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTConfig are custom claims extending default ones.
type JWTConfig struct {
	SecretKey       string
	HoursTillExpire time.Duration
}

// JwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

// Authenticator is struct to run authenticator
type Authenticator struct {
	JWTConfig JWTConfig
}

// Run anthenticator
func (a *Authenticator) Run(email, password string) (string, error) {

	jwt, err := authenticate(email, password, a.JWTConfig)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func authenticate(username string, password string, cfg JWTConfig) (jwtToken string, err error) {
	// Set custom claims
	claims := &JwtCustomClaims{
		username,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(cfg.HoursTillExpire).UTC().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	jwtToken, err = token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
