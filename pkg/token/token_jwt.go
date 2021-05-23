package token

import (
	"context"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/pkg/errors"
)

const (
	claimUser string = "user"
	claimUID  string = "uid"
	claimRole string = "user_role"
)

type Manager interface {
	Create(username string, uid, tokenExpTime int64, userRoleID int32) (string, error)
	Check(context.Context, string) (bool, error)
	FindUserData(string) (uid int64, username string, userRoleID int32, err error)
}

type JwtToken struct {
	Secret []byte
}

func NewJwtToken(secret string) (*JwtToken, error) {
	return &JwtToken{Secret: []byte(secret)}, nil
}

func (tk *JwtToken) Create(username string, uid, tokenExpTime int64, userRoleID int32) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims[claimUser] = username
	claims[claimUID] = uid
	claims[claimRole] = userRoleID
	claims["exp"] = tokenExpTime
	return token.SignedString(tk.Secret)
}

func (tk *JwtToken) Check(ctx context.Context, inputToken string) (bool, error) {
	token, err := jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return tk.Secret, nil
	})
	if err != nil {
		return false, fmt.Errorf("can't parse jwt token: %v", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// If token is valid
		_, ok := claims[claimUID].(float64)
		if !ok {
			return false, errors.Errorf("token doesn't contains claim %s", claimUID)
		}
		// err = tk.repo.CheckUserByID(ctx, int64(userID))
		// if err != nil {
		// 	log.Errorf("repo.CheckUserByID(ctx, %d) error, %v", userID, err)
		// 	return false, errors.Errorf("user with id=%.f not found in the db", userID)
		// }
		return true, nil
	}
	return false, errors.Errorf("invalid token")
}

func (tk *JwtToken) FindUserData(inputToken string) (uid int64, username string, userRoleID int32, err error) {
	token, er := jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return tk.Secret, nil
	})
	if er != nil {
		err = er
		return
	}
	if !token.Valid {
		err = errors.New("invalid token")
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("invalid token")
		return
	}
	// If token is valid
	// uid
	f64UID, ok := claims[claimUID].(float64)
	if !ok {
		err = errors.Errorf("token doesn't contains claim %s", claimUID)
		return
	}
	uid = int64(f64UID)
	// username
	username, ok = claims[claimUser].(string)
	if !ok {
		err = errors.Errorf("token doesn't contains claim %s", claimUser)
		return
	}
	// userRoleID
	f64roleID, ok := claims[claimRole].(float64)
	if !ok {
		err = errors.Errorf("token doesn't contains claim %s", claimRole)
		return
	}
	userRoleID = int32(f64roleID)
	return uid, username, userRoleID, nil
}

func GetToken(r *http.Request) (string, error) {
	return request.AuthorizationHeaderExtractor.ExtractToken(r)
}
