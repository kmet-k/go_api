package utils

import (
	"fmt"
	"free_credit/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/subosito/gotenv"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	gotenv.Load()
}

func GenerateToken(member models.Member) (string, string) {
	atClaims := jwt.MapClaims{}
	atClaims["id"] = member.ID
	atClaims["firstname"] = member.Firstname
	atClaims["lastname"] = member.Lastname
	atClaims["username"] = member.Username
	atClaims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	access_token, _ := at.SignedString([]byte(os.Getenv("ACCESS_TOKEN")))

	rtClaims := jwt.MapClaims{}
	rtClaims["id"] = member.ID
	rtClaims["firstname"] = member.Firstname
	rtClaims["lastname"] = member.Lastname
	rtClaims["username"] = member.Username
	rtClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refresh_token, _ := rt.SignedString([]byte(os.Getenv("REFRESH_TOKEN")))

	return access_token, refresh_token
}

func ValidAccessToken(accessToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an Error")
			}
			return []byte(os.Getenv("ACCESS_TOKEN")), nil
		})
	return token, err
}

func ValidRefreshToken(accessToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an Error")
			}
			return []byte(os.Getenv("REFRESH_TOKEN")), nil
		})
	return token, err
}

func Encrypt(text *string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(*text), 14)
	*text = string(bytes)
}

func Decrypt(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func ClaimsID(Token string) string {
	token, _ := jwt.Parse(Token, nil)
	claims := token.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
