package handler

import (
	"net/http"

	usr "../services/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*usr.JwtCustomClaims)
	name := claims.Name

	return c.JSON(http.StatusOK, echo.Map{
		"name": name,
	})
}

// HTTPServer create a service to echo server
type HTTPServer struct {
	Port      string
	Auth      *usr.Authenticator
	JWTConfig usr.JWTConfig
}

// Run this function of start the server
func (h *HTTPServer) Run() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Unauthenticated route
	e.GET("/", accessible)

	// Login route
	ah := &AuthHandler{
		signin: h.Auth.Run,
	}

	e.POST("/login", ah.login)

	// Authenticate group
	r := e.Group("/api")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &usr.JwtCustomClaims{},
		SigningKey: []byte(h.JWTConfig.SecretKey),
	}

	r.Use(middleware.JWTWithConfig(config))

	r.GET("/user", restricted)

	e.Logger.Fatal(e.Start(h.Port))
}
