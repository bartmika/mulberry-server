// github.com/bartmika/mulberry-server/pkg/utils/jwt.go
package utils

import (
    "time"

    jwt "github.com/dgrijalva/jwt-go"
)


func GenerateJWT(hmacSecret []byte, sessionUuid string, clientUuid string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)

    claims["authorized"] = true
    claims["session_uuid"] = sessionUuid
    claims["client_uuid"] = clientUuid
    claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

    tokenString, err := token.SignedString(hmacSecret)

    if err != nil {
        return "", err
    }
    return tokenString, nil
}

func ProcessJWT(hmacSecret []byte, reqToken string) (map[string]string, error){
    token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
        return hmacSecret, nil
    })
    if err == nil && token.Valid {
        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            m := make(map[string]string)
            m["session_uuid"] = claims["session_uuid"].(string)
            m["client_uuid"] = claims["client_uuid"].(string)
            // m["exp"] = claims["exp"]
            return m, nil
        } else {
            return nil, err
        }

    } else {
        return nil, err
    }
    return nil, nil
}
