package main

import (
	"time"

	hdl "./handler"
	usr "./services/user"
)

func main() {
	jwtConfig := usr.JWTConfig{
		SecretKey:       "123456",
		HoursTillExpire: time.Hour * 24 * 7,
	}

	configData := hdl.HTTPServer{
		Port:      ":5151",
		JWTConfig: jwtConfig,
		Auth: &usr.Authenticator{
			JWTConfig: jwtConfig,
		},
	}

	configData.Run()
}
