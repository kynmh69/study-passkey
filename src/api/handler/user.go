package handler

import (
	"encoding/base64"
	"encoding/json"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/kynmh69/study-passkey/consts"
	"github.com/kynmh69/study-passkey/domain"
	"github.com/kynmh69/study-passkey/dto"
	"github.com/kynmh69/study-passkey/prisma/db"
	"github.com/kynmh69/study-passkey/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetUserById is a function that returns a handler function that gets a user by id.
func GetUserById() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the id from the request parameter.
		id := c.Request().Header.Get("Authorization")

		c.Logger().Debug("connected to database")
		// Get the user by id.
		user, err := utils.Client.User.FindUnique(
			db.User.ID.Equals(id),
		).Exec(c.Request().Context())
		// If an error occurs, return a 404 status code.
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		// Return the user.
		domainUser := domain.User{
			ID:    user.ID,
			Name:  user.Username,
			Email: user.Email,
		}
		return c.JSON(http.StatusOK, domainUser)
	}
}

// BeginRegistration is a function that returns a handler function that creates a user.
func BeginRegistration() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			newUser dto.CreateUserInput
		)
		// bind the request body to the newUser variable.
		if err := c.Bind(&newUser); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// Check if the username and email already exist.
		existingUser, err := utils.Client.User.FindFirst(
			db.User.Username.Equals(newUser.UserName),
		).Exec(c.Request().Context())
		if err == nil && existingUser != nil {
			return echo.NewHTTPError(http.StatusConflict, "username already exists")
		} else if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		existingEmail, err := utils.Client.User.FindFirst(
			db.User.Email.Equals(newUser.Email),
		).Exec(c.Request().Context())
		if err == nil && existingEmail != nil {
			return echo.NewHTTPError(http.StatusConflict, "email already exists")
		} else if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		// Create a new WebAuthn instance.
		domainUser := domain.User{
			ID:    base64.StdEncoding.EncodeToString([]byte(newUser.UserName)),
			Name:  newUser.UserName,
			Email: newUser.Email,
		}

		// 空の認証情報リストをJSON形式で作成
		emptyCredentials := domain.CredentialsList{
			Credentials: []webauthn.Credential{},
		}
		credentialsJSON, err := json.Marshal(emptyCredentials)

		// Create a user.
		_, err = utils.Client.User.CreateOne(
			db.User.ID.Set(domainUser.ID),
			db.User.Email.Set(domainUser.Email),
			db.User.Username.Set(domainUser.Name),
			db.User.Credentials.Set(credentialsJSON),
		).Exec(c.Request().Context())
		// Begin the registration process.
		requireResidentKey := true
		options, sessionData, err := utils.WebAuthn.BeginRegistration(
			&domainUser,
			webauthn.WithAuthenticatorSelection(
				protocol.AuthenticatorSelection{
					RequireResidentKey: &requireResidentKey,
					UserVerification:   protocol.VerificationRequired,
				},
			))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// Create a session.
		sessionDataJson, err := json.Marshal(sessionData)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		tempToken, err := utils.Sessions.CreateSession(
			domainUser.Name,
			map[string]interface{}{
				"sessionData": string(sessionDataJson),
				"type":        "registration",
			},
		)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(
			http.StatusOK,
			map[string]interface{}{
				"options":   options,
				"tempToken": tempToken,
			},
		)
	}
}

// CompleteRegistration is a function that returns a handler function that completes the registration process.
func CompleteRegistration() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req struct {
			Username  string `json:"username"`
			TempToken string `json:"tempToken"`
		}
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		sessionData, err := utils.Sessions.GetSession(req.TempToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		if sessionData["type"] != consts.REGISTRATION {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid session type")
		}

		// get the user by username.
		user, err := utils.Client.User.FindFirst(
			db.User.Username.Equals(req.Username),
		).Exec(c.Request().Context())
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		var webAuthnSessionData webauthn.SessionData
		err = json.Unmarshal([]byte(sessionData["sessionData"].(string)), &webAuthnSessionData)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		domainUser, err := domain.NewWebAuthnUser(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		credential, err := utils.WebAuthn.FinishRegistration(domainUser, webAuthnSessionData, c.Request())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		var currentCredentials domain.CredentialsList
		if err := json.Unmarshal(user.Credentials, &currentCredentials); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		currentCredentials.Credentials = append(currentCredentials.Credentials, *credential)
		updateCredentialJson, err := json.Marshal(currentCredentials)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		_, err = utils.Client.User.FindUnique(
			db.User.ID.Equals(user.ID),
		).Update(
			db.User.Credentials.Set(updateCredentialJson),
		).Exec(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		err = utils.Sessions.DeleteSession(req.TempToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		token, err := utils.Sessions.CreateSession(user.ID, map[string]interface{}{
			"username": user.Username,
			"email":    user.Email,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "registration completed",
			"token":   token,
		})
	}
}

func Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionId := c.Request().Header.Get("Authorization")
		if err := utils.Sessions.DeleteSession(sessionId); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusOK)
	}
}
