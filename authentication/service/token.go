package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-todo-list/authentication/repository"
	"time"
)

type (
	TokenService interface {
		CreateJWT(userID int) (JWTToken, error)
	}

	tokenService struct {
		refreshTTL            int
		accessTTL             int
		secret                string
		accessTokenRepository repository.AccessTokenRepositoryInterface
	}

	JWTToken struct {
		AccessToken  string
		RefreshToken string
	}

	JWtClaims struct {
		jwt.StandardClaims
		TokenType string `json:"token_type"`
		UserID    int    `json:"uid"`
	}
)


func NewTokenService(accessTokenRepository repository.AccessTokenRepositoryInterface) TokenService {
	secret := viper.GetString("application.jwt.secret")
	accessTTL := viper.GetInt("application.jwt.tokenTtl")
	refreshTTL := viper.GetInt("application.jwt.refreshTtl")

	return &tokenService{
		refreshTTL,
		accessTTL,
		secret,
		accessTokenRepository,
	}
}


func (t *tokenService) CreateJWT(userID int) (JWTToken, error) {

	jwtToken := JWTToken{}
	now := time.Now()

	accessTokenDuration := time.Duration(t.accessTTL) * time.Hour
	expiredAt := now.Add(accessTokenDuration)

	refreshTokenDuration := time.Duration(t.refreshTTL) * time.Hour
	refreshTokenExpiration := now.Add(refreshTokenDuration)

	token, err := t.generateJWT(userID, expiredAt)
	refreshToken, _ := t.generateRefreshToken(userID, refreshTokenExpiration)
	logrus.Info("save token")
	err = t.accessTokenRepository.Save(token, userID, expiredAt, refreshToken)

	if err != nil {
		logrus.Error(err)
		return jwtToken, err
	}

	jwtToken.AccessToken = token
	jwtToken.RefreshToken = refreshToken

	return jwtToken, err
}


func (t *tokenService) generateRefreshToken(userID int, expiration time.Time) (string, error) {

	claims := JWtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "SmartRetail",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expiration.Unix(),
		},
		TokenType: "refresh",
		UserID:    userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(t.secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func (t *tokenService) generateJWT(uid int, expiration time.Time) (string, error) {

	claims := JWtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "SmartRetail",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expiration.Unix(),
		},
		TokenType: "token",
		UserID:    uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(t.secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
