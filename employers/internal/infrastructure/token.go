package infrastructure

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/qaZar1/HHforURFU/employers/internal/models"
)

func (infra *Infrastructure) GetClaims(username string) models.Claims {
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(time.Duration(infra.ExpiresIn) * time.Second)

	return models.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    infra.issuer,
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			Audience: jwt.ClaimStrings{
				resourceHelloWorld,
				resourcePingPong,
			},
		},

		Username:  username,
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
		ExpiresIn: infra.ExpiresIn,
		Type:      tokenType,
		Resources: infra.resources,
	}
}

func (infra *Infrastructure) GetSignedToken(claims models.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(infra.secret)
}

func (infra *Infrastructure) GetTokenClaims(auth string) (*models.Claims, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(
		auth,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return infra.secret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, TokenError
	}

	return claims, nil
}
