package utils

import (
	m "back/models"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"strings"
)

// PhoneNumber -
func PhoneNumber(number string) string {
	//Check Country Code and remove it
	if strings.HasPrefix(number, "959") {
		number = number[3:]
	} else if strings.HasPrefix(number, "09") {
		number = number[2:]
	}
	return number
}

// Int32ToString -
func Int32ToString(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

// GenerateToken - User to token string
func GenerateToken(phone string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"msisdn": "",
		"phone":  phone,
		"status": "",
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}

// TokenString -
func TokenString(value string) (m.Token, error) {
	t := m.Token{}
	token, _ := jwt.Parse(value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		mapstructure.Decode(claims, &t)
		return t, nil
	} else {
		return m.Token{}, errors.New("Invalid authorization token")
	}
}

//simple token for admin
func CreateToken(password string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["userName"] = ""
	atClaims["password"] = password
	//atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return token, nil
}
