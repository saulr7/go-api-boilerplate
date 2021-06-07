package auth

import (
	"api-boilerplate/models"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(data *models.Login) (string, error) {
	claim := models.Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "Rockstar Lab",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken .
func ValidateToken(t string) (models.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &models.Claim{}, verifyFunction)
	if err != nil {
		return models.Claim{}, err
	}
	if !token.Valid {
		return models.Claim{}, errors.New("token no válido")
	}

	claim, ok := token.Claims.(*models.Claim)
	if !ok {
		return models.Claim{}, errors.New("no se pudo obtener los claim")
	}

	return *claim, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
