package handler

import (
	"echogorm/infrastructure/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id    int    `json:"id" param:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func IndexHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		users := []User{}
		database.DB.Find(&users)
		return c.JSON(http.StatusOK, users)
	}
}

func ShowHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := User{}
		if err := c.Bind(&user); err != nil {
			return err
		}
		database.DB.Take(&user)
		return c.JSON(http.StatusOK, user)
	}
}

func StoreHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := User{}
		if err := c.Bind(&user); err != nil {
			return err
		}
		database.DB.Create(&user)
		return c.JSON(http.StatusCreated, user)
	}
}

func UpdateHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := User{}
		if err := c.Bind(&user); err != nil {
			return err
		}
		database.DB.Save(&user)
		return c.JSON(http.StatusOK, user)
	}
}

func DeleteHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		database.DB.Delete(&User{}, id)
		return c.NoContent(http.StatusNoContent)
	}
}
