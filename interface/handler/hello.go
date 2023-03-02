package handler

import (
	"echogorm/infrastructure/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		database.Connect()
		db, _ := database.DB.DB()
		defer db.Close()
		err := db.Ping()
		if err != nil {
			return c.String(http.StatusInternalServerError, "データベース接続失敗")
		} else {
			return c.String(http.StatusOK, "Hello, World!")
		}
	}
}
