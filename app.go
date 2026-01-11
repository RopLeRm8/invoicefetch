package main

import (
	"context"
	"fmt"
	oauth "invoice_fetch/oauth"
)

type AuthResult struct {
	OK    bool   `json:"ok"`
	Url   string `json:"url"`
	Error string `json:"error,omitempty"`
}

type ActiveEmail struct {
	Email string `json:"email"`
	Error string `json:"error,omitempty"`
}

type Logout struct {
	Error *string `json:"error,omitempty"`
}

type Verify struct {
	Email *string `json:"email"`
	Error *string `json:"error,omitempty"`
}

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Auth(provider string) AuthResult {

	url, err := oauth.RunAuth(a.ctx)

	var result AuthResult = AuthResult{OK: true, Error: ""}

	if err != nil {
		result.OK = false
		result.Error = err.Error()
		return result
	}

	result.Url = url
	return result
}

func (a *App) GetActiveEmail() ActiveEmail {
	email, err := oauth.GetIdentity()
	if err != nil {
		return ActiveEmail{Error: err.Error(), Email: ""}
	}
	return ActiveEmail{Email: email, Error: ""}
}

func (a *App) Logout() Logout {
	err := oauth.Logout(a.ctx)
	if err != nil {
		msg := err.Error()
		return Logout{Error: &msg}
	}
	return Logout{Error: nil}

}

func (a *App) Verify() Verify {
	m, err := oauth.GetIdentity()
	if err != nil {
		msg := err.Error()
		return Verify{Error: &msg}
	}
	return Verify{Email: &m}
}
