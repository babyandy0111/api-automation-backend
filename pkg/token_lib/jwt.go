package token_lib

import (
	"api-automation-backend/config"
	"api-automation-backend/pkg/er"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Generate jwt token
func GenToken(accId, account string) (string, time.Time, error) {
	salt := os.Getenv("JWT_SALT")
	secret := []byte(salt)

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	exp := time.Now().AddDate(0, 0, 30).UTC().Unix()
	// development exp 10 mins
	if config.GetEnvironment() == config.EnvDevelopment {
		exp = time.Now().Add(time.Hour*time.Duration(0) +
			time.Minute*time.Duration(10) +
			time.Second*time.Duration(0)).UTC().Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "Api Automation",
		// Expiration Time,
		"exp":        exp,
		"iat":        time.Now().UTC().Unix(),
		"account_id": accId,
		"account":    account,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)

	return tokenString, time.Unix(exp, 0).UTC(), err
}

// Parse jwt token
func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	salt := os.Getenv("JWT_SALT")
	return parseToken(tokenStr, salt)
}

func parseToken(tokenStr, salt string) (jwt.MapClaims, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(salt), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func ParseClaims(claims interface{}) jwt.MapClaims {
	if c, ok := claims.(jwt.MapClaims); ok {
		return c
	}

	return nil
}

func CheckJWTAccId(c *gin.Context, accId int64) error {
	claims, _ := c.Get("claims")

	claimsMap := ParseClaims(claims)

	if strconv.FormatInt(accId, 10) != claimsMap["account_id"] {
		err := er.NewAppErr(401, er.UnauthorizedError, "Token is not valid", nil)
		return err
	}

	return nil
}
