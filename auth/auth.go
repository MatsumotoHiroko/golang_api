package auth

import (
    "net/http"
    "os"
    "time"

    jwtmiddleware "github.com/auth0/go-jwt-middleware"
    jwt "github.com/dgrijalva/jwt-go"
)

// GetTokenHandler get token
//var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
    // set header
    token := jwt.New(jwt.SigningMethodHS256)

    // set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["admin"] = true
    claims["sub"] = "54546557354"
    claims["name"] = "testuser"
    claims["iat"] = time.Now()
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    // e-signature
    tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

    // return JWT
    w.Write([]byte(tokenString))
}

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
    ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("SIGNINGKEY")), nil
    },
    SigningMethod: jwt.SigningMethodHS256,
})
