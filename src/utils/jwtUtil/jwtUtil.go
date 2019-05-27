package jwtUtil

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"ginserver/config"
	"reflect"
	"time"
)

var JWTKEY = config.InitBaseConfig().TOKEN_SECRET

type Token struct {
	Token string `json:"token"`
}


//加密方法：
func CreateToken(uid string) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Minute),
	})

	fmt.Println("token==", token)

	tokenString, err := token.SignedString([]byte(JWTKEY))

	if err != nil {
		panic(err)
	}

	return tokenString
}

//解密方法
func ParsingToken(tokenString string) (jwt.MapClaims, error){

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {


		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JWTKEY), nil

	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid || err != nil{
		return claims, fmt.Errorf("Token valid faile ", err)
	}


	fmt.Println("type:", reflect.TypeOf(claims))

	return claims,nil		// map[exp:2019-04-09T18:07:44.236967+08:00 uid:100061]

}
