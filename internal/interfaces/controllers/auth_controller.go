package controllers

import (
	"net/http"
	"os"

	"github.com/YasuhiroOsajima/go-auth-api/internal/model"
	"github.com/YasuhiroOsajima/go-auth-api/internal/usecase"
)

var domain = os.Getenv("WEB_SERVER_COOKIE_DOMAIN")

// Input
type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserNameInput struct {
	Username string `json:"username" binding:"required"`
}

// Result
type DataResult struct {
	Data any `json:"data"`
}

type TokenResult struct {
	Message string `json:"message"`
}

type ErrorResult struct {
	Error string `json:"error"`
}

// Controller
type AuthController struct {
	authInteractor *usecase.AuthInteractor
	userInteractor *usecase.UserInteractor
}

func NewAuthController(authInteractor *usecase.AuthInteractor, userInteractor *usecase.UserInteractor) *AuthController {
	return &AuthController{authInteractor, userInteractor}
}

func (c *AuthController) Register(ctx Context) {
	var input UserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		if err.Error() == "EOF" {
			ctx.JSON(http.StatusBadRequest, ErrorResult{Error: "username and password are required"})
			return
		}

		ctx.JSON(http.StatusBadRequest, ErrorResult{Error: err.Error()})
		return
	}

	user := model.NewUser(-1, input.Username, input.Password, "", false)
	savedUser, err := c.userInteractor.Register(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResult{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, DataResult{Data: savedUser.PrepareOutput()})
}

func (c *AuthController) GetToken(ctx Context, tokenCookieName string) {
	var input UserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		if err.Error() == "EOF" {
			ctx.JSON(http.StatusBadRequest, ErrorResult{Error: "username and password are required"})
			return
		}

		ctx.JSON(http.StatusBadRequest, ErrorResult{Error: err.Error()})
		return
	}

	user := model.NewUser(-1, input.Username, input.Password, "", false)
	token, err := c.authInteractor.GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResult{Error: err.Error()})
		return
	}

	ctx.SetCookie(tokenCookieName, token, 3600, "/", domain, false, true)
	ctx.JSON(http.StatusOK, TokenResult{Message: "Token was generated"})
}

func (c *AuthController) UnSetToken(ctx Context, tokenCookieName string) {
	ctx.SetCookie(tokenCookieName, "", 3600, "/", domain, false, true)
	ctx.JSON(http.StatusOK, TokenResult{Message: "Token was deleted"})
}

func (c *AuthController) GetUserInfo(ctx Context, token string) {
	authenticatedUser, err := c.authInteractor.GetUserInfoByToken(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResult{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, DataResult{Data: authenticatedUser.PrepareOutput()})
}

func (c *AuthController) Enable(ctx Context) {
	var input UserNameInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		if err.Error() == "EOF" {
			ctx.JSON(http.StatusBadRequest, ErrorResult{Error: "username and password are required"})
			return
		}

		ctx.JSON(http.StatusBadRequest, ErrorResult{Error: err.Error()})
		return
	}

	user := model.NewUser(-1, input.Username, "", "", false)
	authenticatedUser, err := c.authInteractor.GetUserInfoByName(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResult{Error: err.Error()})
		return
	}

	enabledUser, err := c.userInteractor.Enable(authenticatedUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResult{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, DataResult{Data: enabledUser.PrepareOutput()})
}

func (c *AuthController) Disable(ctx Context) {
	var input UserNameInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		if err.Error() == "EOF" {
			ctx.JSON(http.StatusBadRequest, ErrorResult{Error: "username and password are required"})
			return
		}

		ctx.JSON(http.StatusBadRequest, ErrorResult{Error: err.Error()})
		return
	}

	user := model.NewUser(-1, input.Username, "", "", false)
	authenticatedUser, err := c.authInteractor.GetUserInfoByName(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResult{Error: err.Error()})
		return
	}

	disabledUser, err := c.userInteractor.Disable(authenticatedUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResult{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, DataResult{Data: disabledUser.PrepareOutput()})
}
