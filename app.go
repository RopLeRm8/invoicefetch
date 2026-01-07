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
