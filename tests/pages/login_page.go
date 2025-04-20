package pages

import (
	"context"

	"example.com/v2/tests/base"
	"github.com/chromedp/chromedp"
)

const (
	LoginURL      = "https://www.saucedemo.com/v1/"
	UsernameInput = `#user-name`
	PasswordInput = `#password`
	LoginButton   = `#login-button`
)

type LoginPage struct {
	ctx context.Context
}

func NewLoginPage(ctx context.Context) *LoginPage {
	return &LoginPage{ctx}
}

func (p *LoginPage) Navigate() error {
	err := chromedp.Run(p.ctx, chromedp.Navigate(LoginURL))
	if err != nil {
		base.TakeScreenshot(p.ctx, "login_navigation_error.png")
	}
	return err
}

func (p *LoginPage) SendUsername(username string) error {
	err := chromedp.Run(p.ctx, chromedp.SendKeys(UsernameInput, username))
	if err != nil {
		base.TakeScreenshot(p.ctx, "send_username_error.png")
	}
	return err
}

func (p *LoginPage) SendPassword(password string) error {
	err := chromedp.Run(p.ctx, chromedp.SendKeys(PasswordInput, password))
	if err != nil {
		base.TakeScreenshot(p.ctx, "send_password_error.png")
	}
	return err
}

func (p *LoginPage) ClickLogin() error {
	err := chromedp.Run(p.ctx, chromedp.Click(LoginButton))
	if err != nil {
		base.TakeScreenshot(p.ctx, "click_login_error.png")
	}
	return err
}

func (p *LoginPage) Login(username, password string) error {
	if err := p.SendUsername(username); err != nil {
		return err
	}
	if err := p.SendPassword(password); err != nil {
		return err
	}
	return p.ClickLogin()
}
