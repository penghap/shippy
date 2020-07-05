package handler

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	pb "github.com/penghap/shippy/service-user/proto/user"
	"github.com/penghap/shippy/service-user/repository"
)

var (
	key      = []byte("token")
	duration = time.Hour * 72
)

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type AuthAble interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
	repo repository.Repository
}

// sting -> CustomClaims
func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return key, err
	})

	// validate token
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (srv *TokenService) Encode(user *pb.User) (string, error) {
	expireToken := time.Now().Add(duration).Unix()

	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "go.micro.srv.user",
		},
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
