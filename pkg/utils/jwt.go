// github.com/bartmika/mulberry-server/pkg/utils/jwt.go
package utils

import (
    "time"

    jwt "github.com/dgrijalva/jwt-go"
)

// Generate the `access token` and `refresh token` for the secret key.
func GenerateJWTTokenPair(hmacSecret []byte, clientUuid string) (string, string, error) {
    //
    // Generate token.
    //
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["user_uuid"] = clientUuid
    claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

    tokenString, err := token.SignedString(hmacSecret)
    if err != nil {
        return "", "", err
    }

    //
    // Generate refresh token.
    //
    refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["user_uuid"] = clientUuid
	rtClaims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	refreshTokenString, err := refreshToken.SignedString(hmacSecret)
	if err != nil {
		return "", "", err
	}

    return tokenString, refreshTokenString, nil
}

// Validates either the `access token` or `refresh token` and returns either the
// `user_uuid` if success or error on failure.
func ProcessJWTToken(hmacSecret []byte, reqToken string) (string, error){
    token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
        return hmacSecret, nil
    })
    if err == nil && token.Valid {
        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            user_uuid := claims["user_uuid"].(string)
            // m["exp"] := string(claims["exp"].(float64))
            return user_uuid, nil
        } else {
            return "", err
        }

    } else {
        return "", err
    }
    return "", nil
}
