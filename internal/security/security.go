package security

import (
	"backend/internal/models"
	userrepo "backend/internal/mongodb/user"
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	salt       = "dfhvbdfb34h5b34c543g"
	signingKey = "34vhv34vgc53hfb56vv7"
	tokenTTL   = 1 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId primitive.ObjectID `json:"user_id"`
}

func CreateUser(user models.User) (id primitive.ObjectID, err error) {
	user.Password = HashPassword(user.Password)
	_, err = userrepo.FindOneUser(user.Username, user.Password)
	if err == nil {
		return id, errors.New("username is not free")
	}
	err = userrepo.CreateUser(context.Background(), user)
	if err != nil {
		return id, err
	}
	return user.ID, err
}

func GetUser(id primitive.ObjectID) (models.User, error) {
	return userrepo.FindByIdUser(context.Background(), id)
}

func GenerateToken(username, password string) (string, error) {
	user, err := userrepo.FindOneUser(username, HashPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func ParseToken(accessToken string) (id primitive.ObjectID, err error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return id, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return id, errors.New("token claims not the *tokenClaims")
	}
	return claims.UserId, err
}

func HashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
