package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phongnd2802/ezy-mark/internal/models"
	"github.com/phongnd2802/ezy-mark/internal/response"
	"github.com/phongnd2802/ezy-mark/internal/services"
)

type authController struct{}

var Auth = new(authController)

// Login godoc
// @Summary      User Login
// @Description  Authenticate user with email and password
// @Tags         Authentication Management
// @Accept       json
// @Produce      json
// @Param        payload body models.LoginRequest true "payload"
// @Success      200  {object}  response.Response
// @Router       /auth/signin [post]
func (c *authController) Login(ctx *fiber.Ctx) error {
	params := new(models.LoginRequest)
	if err := ctx.BodyParser(params); err != nil {
		return response.ErrorResponse(ctx, response.ErrCodeInvalidParams, err)
	}

	code, data, err := services.AuthService().Login(ctx.UserContext(), params)
	if err != nil {
		return response.ErrorResponse(ctx, code, err)
	}
	return response.SuccessResponse(ctx, code, data)
}

// Register godoc
// @Summary      User Register
// @Description  Register user with email and password
// @Tags         Authentication Management
// @Accept       json
// @Produce      json
// @Param        payload body models.RegisterRequest true "payload"
// @Success      200  {object}  response.Response
// @Router       /auth/signup [post]
func (c *authController) Register(ctx *fiber.Ctx) error {
	params := new(models.RegisterRequest)
	if err := ctx.BodyParser(params); err != nil {
		return response.ErrorResponse(ctx, response.ErrCodeInvalidParams, err)
	}
	/// Call Auth Service
	code, data, err := services.AuthService().Register(ctx.UserContext(), params)
	if err != nil {
		return response.ErrorResponse(ctx, code, err)
	}
	return response.SuccessResponse(ctx, code, data)
}

// VerifyOTP godoc
// @Summary      User Verify OTP
// @Description  Authenticate OTP
// @Tags         Authentication Management
// @Accept       json
// @Produce      json
// @Param        payload body models.VerifyOTPReq true "payload"
// @Success      200  {object}  response.Response
// @Router       /auth/verify-otp [post]
func (c *authController) VerifyOTP(ctx *fiber.Ctx) error {
	params := new(models.VerifyOTPReq)
	if err := ctx.BodyParser(params); err != nil {
		return response.ErrorResponse(ctx, response.ErrCodeInvalidParams, err)
	}
	code, data, err := services.AuthService().VerifyOTP(ctx.UserContext(), params)
	if err != nil {
		return response.ErrorResponse(ctx, code, err)
	}
	return response.SuccessResponse(ctx, code, data)
}

// GetTTLOtp godoc
// @Summary      Get TTL OTP
// @Description  Get Time To Live OTP
// @Tags         Authentication Management
// @Accept       json
// @Produce      json
// @Param        token query string true "Token"
// @Success      200  {object}  response.Response
// @Router       /auth/verify-otp [get]
func (c *authController) GetTTLOtp(ctx *fiber.Ctx) error {
	token := ctx.Query("token")
	if token == "" {
		return response.ErrorResponse(ctx, response.ErrCodeExpiredSession, nil)
	}
	code, data, err := services.AuthService().GetTTLOtp(ctx.UserContext(), token)
	if err != nil {
		return response.ErrorResponse(ctx, code, err)
	}
	return response.SuccessResponse(ctx, code, data)
}

// ResendOTP godoc
// @Summary      Resend OTP
// @Description  Resend OTP
// @Tags         Authentication Management
// @Accept       json
// @Produce      json
// @Param        payload body models.ResendOTPReq true "payload"
// @Success      200  {object}  response.Response
// @Router       /auth/resend-otp [post]
func (c *authController) ResendOTP(ctx *fiber.Ctx) error {
	params := new(models.ResendOTPReq)
	if err := ctx.BodyParser(params); err != nil {
		return response.ErrorResponse(ctx, response.ErrCodeInvalidParams, err)
	}

	code, err := services.AuthService().ResendOTP(ctx.UserContext(), params)
	if err != nil {
		return response.ErrorResponse(ctx, code, err)
	}
	return response.SuccessResponse(ctx, code, nil)
}
