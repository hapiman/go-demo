package jwt_demo

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(memberId int64) (tokenss string, err error) {
	//自定义claim
	claim := jwt.MapClaims{
		"memberId": memberId,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * time.Duration(3000)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 使用viper获取secret viper.GetString("token.secret")
	tokenss, err = token.SignedString([]byte("this_is_jwt_ghb"))
	return
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("this_is_jwt_ghb"), nil
		// return []byte(viper.GetString("this_is_jwt_secret")), nil
	}
}

func ParseToken(tokenss string) (memberId int64, err error) {
	token, err := jwt.Parse(tokenss, secret())
	if err != nil {
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}

	return int64(claim["memberId"].(float64)), err
}
