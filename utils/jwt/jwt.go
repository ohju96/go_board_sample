package jwt

import (
	"context"
	"errors"
	"fmt"
	"ginSample/config"
	"ginSample/config/db"
	custom "ginSample/handler/err"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

type CustomClaims struct {
	Email string
	Name  string
	Role  string
}

type Token struct {
	Access  string
	Refresh string
}

const (
	ROLE_ADMIN = "ADMIN"
	ROLE_USER  = "USER"
)

func CreateAccessToken(customClaims CustomClaims, jwtSecretKey string) (string, *custom.ErrRes) {
	fmt.Println("#jwt# create access token start")

	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["email"] = customClaims.Email
	claims["name"] = customClaims.Name
	claims["role"] = customClaims.Role
	claims["expiresAt"] = time.Now().Add(time.Minute * 1).Unix()

	res, err := accessToken.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", custom.NewErrRes(custom.ERR_JWT_CREATE_FAIL)
	}
	return res, nil
}

func CreateRefreshToken(ctx context.Context, customClaims CustomClaims, jwtSecretKey string) (string, *custom.ErrRes) {
	fmt.Println("#jwt# create refresh token start")
	unixTime := time.Now().Add(time.Hour * 240).Unix()
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claims := refreshToken.Claims.(jwt.MapClaims)
	claims["email"] = customClaims.Email
	claims["name"] = customClaims.Name
	claims["role"] = customClaims.Role
	claims["expiresAt"] = unixTime
	res, err := refreshToken.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", custom.NewErrRes(custom.ERR_JWT_CREATE_FAIL)
	}

	// unixTime convert to utcTime
	utcTime := time.Unix(unixTime, 0)

	// set redis
	if err := db.Redis.Set(ctx, res, customClaims.Email, utcTime.Sub(time.Now())); err == nil {
		fmt.Println(err)
		return "", custom.NewErrRes(custom.ERR_REDIS_SET_FAIL)
	}

	return res, nil
}

func CreateAccessAndRefreshToken(ctx context.Context, customClaims CustomClaims, jwtSecretKey string) (Token, *custom.ErrRes) {
	fmt.Println("#jwt# create access and refresh token start")

	accessToken, res := CreateAccessToken(customClaims, jwtSecretKey)
	if res != nil {
		return Token{}, res
	}

	refreshToken, res := CreateRefreshToken(ctx, customClaims, jwtSecretKey)
	if res != nil {
		return Token{}, res
	}
	return Token{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func TokenParsing(req *http.Request, toml config.Config) (*CustomClaims, *custom.ErrRes) {
	fmt.Println("#jwt# token parsing start")

	token, err := verifyToken(req, toml)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		exp, ok := claims["expiresAt"].(float64)
		if !ok {
			return nil, custom.NewErrRes(custom.ERR_JWT_INVALID)
		}
		if exp < float64(time.Now().Unix()) {
			return nil, custom.NewErrRes(custom.ERR_JWT_ACCESS_EXPIRED)
		}
		email, ok := claims["email"].(string)
		if !ok {
			return nil, custom.NewErrRes(custom.ERR_JWT_INVALID)
		}
		name, ok := claims["name"].(string)
		if !ok {
			return nil, custom.NewErrRes(custom.ERR_JWT_INVALID)
		}
		role, ok := claims["role"].(string)
		if !ok {
			return nil, custom.NewErrRes(custom.ERR_JWT_INVALID)
		}
		return &CustomClaims{
			Email: email,
			Name:  name,
			Role:  role,
		}, nil
	}
	return nil, custom.NewErrRes(custom.ERR_JWT_INVALID)
}

func verifyToken(req *http.Request, toml config.Config) (*jwt.Token, *custom.ErrRes) {
	fmt.Println("#jwt# verify token start")

	tokenString := extractToken(req)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("jwt invalid")
		}
		return []byte(toml.Local.JwtSecretKey), nil
	})
	if err != nil {
		return nil, custom.NewErrRes(custom.ERR_JWT_INVALID)
	}
	return token, nil
}

func extractToken(req *http.Request) string {
	fmt.Println("#jwt# extract token start")

	getToken := req.Header.Get("Authorization")
	tokenArr := strings.Split(getToken, " ")
	if len(tokenArr) == 2 {
		return tokenArr[1]
	}
	return ""
}

func RefreshTokenParsing(req string, jwtSecretKey string) (*CustomClaims, *custom.ErrRes) {
	fmt.Println("#jwt# token parsing start")

	token, err := refreshVerifyToken(req, jwtSecretKey)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		exp, ok := claims["expiresAt"].(float64)
		if !ok {
			return nil, custom.NewErrRes(custom.ERR_JWT_INVALID)
		}
		if exp < float64(time.Now().Unix()) {
			return nil, custom.NewErrRes(custom.ERR_JWT_REFRESH_EXPIRED)
		}
		email, ok := claims["email"].(string)
		if !ok {
			return nil, custom.NewErrRes(custom.ERR_JWT_INVALID)
		}
		name, ok := claims["name"].(string)
		if !ok {
			return nil, custom.NewErrRes(custom.ERR_JWT_INVALID)
		}
		role, ok := claims["role"].(string)
		if !ok {
			return nil, custom.NewErrRes(custom.ERR_JWT_INVALID)
		}
		return &CustomClaims{
			Email: email,
			Name:  name,
			Role:  role,
		}, nil
	}
	return nil, custom.NewErrRes(custom.ERR_JWT_INVALID)
}

func refreshVerifyToken(tokenString string, jwtSecretKey string) (*jwt.Token, *custom.ErrRes) {
	fmt.Println("#jwt# verify token start")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("jwt invalid")
		}
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, custom.NewErrRes(custom.ERR_JWT_INVALID)
	}
	return token, nil
}
