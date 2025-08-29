package utils

import (
    "time"
    "github.com/dgrijalva/jwt-go"
)

var JWTSecret = []byte("your-secret-key-change-in-production")

type Claims struct {
    UserID   int    `json:"user_id"`
    Username string `json:"username"`
    jwt.StandardClaims
}

func GenerateJWT(userID int, username string) (string, error) {
    claims := Claims{
        UserID:   userID,
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
            Issuer:    "e-ticketing",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(JWTSecret)
}

func ValidateJWT(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return JWTSecret, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }

    return nil, err
}