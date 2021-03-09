package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ecnuvj/vhoj_api/model/entity"
	"strconv"
	"time"
)

var jwtSecret = []byte("ecnuvj-tcg")

type Claims struct {
	jwt.StandardClaims
	UserId   string   `json:"userId"`
	Username string   `json:"username"`
	Roles    []string `json:"role"`
}

func GenerateToken(user *entity.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(72 * time.Hour)
	roles := make([]string, len(user.Roles))
	for i, r := range user.Roles {
		roles[i] = r.RoleName
	}
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ecnuvj",
		},
		UserId:   strconv.Itoa(int(user.UserId)),
		Username: user.Username,
		Roles:    roles,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
