package auth

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/spf13/viper"
)

func configuracion() {
	viper.SetConfigName("configuracion")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func ValidateToken(t string) (*jwt.Token, bool, error) {
	configuracion()
	var jwtSecret = []byte(viper.GetString("frasesecreta"))

	if t == "" {
		return nil, false, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
	}

	token, _ := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return jwtSecret, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, true, nil
	} else {
		return nil, false, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
	}
}
