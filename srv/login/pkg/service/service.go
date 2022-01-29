package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// LoginService describes the service.
type LoginService interface {
	Login(ctx context.Context, name, pwd string) (token string, err error)
	RefreshToken(ctx context.Context, oldToken string) (token string)
}

type basicLoginService struct{}

func (b *basicLoginService) Login(ctx context.Context, name string, pwd string) (token string, err error) {
	return b.login(ctx, name, pwd)
}

func (b *basicLoginService) login(ctx context.Context, name string, pwd string) (token string, err error) {
	token = "failed"
	if name == "admin" && pwd == "1" {
		token = fmt.Sprintf("%v||%v", name, time.Now().Unix())
		//token = string(bytes)
	}
	return token, err
}

func (b *basicLoginService) RefreshToken(ctx context.Context, oldToken string) (token string) {
	return b.refreshToken(ctx, oldToken)
}

// NewBasicLoginService returns a naive, stateless implementation of LoginService.
func NewBasicLoginService() LoginService {
	return &basicLoginService{}
}

// New returns a LoginService with all of the expected middleware wired in.
func New(middleware []Middleware) LoginService {
	var svc LoginService = NewBasicLoginService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicLoginService) refreshToken(ctx context.Context, oldToken string) (token string) {
	if oldToken == "failed" {
		return oldToken
	}

	stres := strings.Split(oldToken, "||")
	if len(stres) >= 2 {
		oldUnix := stres[1]
		name := stres[0]
		unix, err := strconv.ParseInt(oldUnix, 10, 64)
		if err != nil {
			return "failed"
		}
		if unix > time.Now().AddDate(0, 0, -1).Unix() {
			return GetToken(name)
		} else {
			return "token expire"
		}
	}
	return
}

func GetToken(name string) (token string) {
	token = fmt.Sprintf("%v||%v", name, time.Now().Unix())
	return
}
