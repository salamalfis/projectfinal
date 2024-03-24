package helper

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	SECRET_JWT = "mysecretjwtdontsharethistoanyoneelse"
)

func GenerateToken(claim any) (token string, err error) {
	jwtClaim := jwt.MapClaims{}
	b, err := json.Marshal(claim)
	if err != nil {
		log.Println("cannot marshal claim payload")
		return
	}
	err = json.Unmarshal(b, &jwtClaim)
	if err != nil {
		log.Println("cannot mapping claim to jwt claim")
		return
	}
	// prepare
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwtClaim)
	// generate token
	token, err = parseToken.SignedString([]byte(SECRET_JWT))
	if err != nil {
		log.Println("cannot generate token", err.Error())
		return
	}
	return
}

func ValidateToken(token string) (claim jwt.MapClaims, err error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(SECRET_JWT), nil
	})
	if err != nil {
		log.Println("error validating jwt token", err.Error())
		return
	}

	// translate claim
	claim, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("error translate claim")
		return
	}
	return
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateHash(in string) (out string, err error) {
	outByte, err := bcrypt.GenerateFromPassword([]byte(in), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error generate hash password", err.Error())
		return
	}
	return string(outByte), err
}

func GetUserID(ctx *gin.Context) (uint64, error) {
	userID, exist := ctx.Get("user_id")
	if !exist {
		return 0, errors.New("cannot get payload in access token")
	}
	Iduser := userID.(uint64)
	user := uint64(Iduser)
	if user == 0 {
		return 0, errors.New("cannot get user id")
	}
	return user, nil
}
