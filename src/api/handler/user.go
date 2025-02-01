package handler

import (
	"github.com/kynmh69/study-passkey/logger"
	"github.com/kynmh69/study-passkey/prisma/db"
	"github.com/kynmh69/study-passkey/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// GetUserById is a function that returns a handler function that gets a user by id.
func GetUserById() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the id from the request parameter.
		id, _ := strconv.Atoi(c.Param("id"))

		client := db.NewClient()
		if err := client.Connect(); err != nil {
			logger.Logger.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, utils.NewHttpError(err.Error()))
		}

		defer func() {
			if err := client.Disconnect(); err != nil {
				logger.Logger.Error(err.Error())
			}
		}()

		c.Logger().Debug("connected to database")
		// Get the user by id.
		user, err := client.User.FindFirst(
			db.User.ID.Equals(id),
		).Exec(c.Request().Context())
		// If an error occurs, return a 404 status code.
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusNotFound, utils.NewHttpError(err.Error()))
		}
		// Return the user.
		return c.JSON(http.StatusOK, user)
	}
}

// CreateUser is a function that returns a handler function that creates a user.
func CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		client := db.NewClient()

		defer func() {
			if err := client.Disconnect(); err != nil {
				logger.Logger.Error(err.Error())
			}
		}()

		c.Logger().Debug("connected to database")
		// Create a user.
		user, err := client.User.CreateOne(
			db.User.Email.Set(""),
			db.User.Username.Set(""),
			db.User.Passkey.Set(""),
		).Exec(c.Request().Context())
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, utils.NewHttpError(err.Error()))
		}
		return c.JSON(http.StatusOK, user)
	}
}
