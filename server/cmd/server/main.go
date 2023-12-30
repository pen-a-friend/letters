package main

import (
	"pen-a-friend/pkg/app"
	"pen-a-friend/pkg/database"
	"pen-a-friend/pkg/webhook"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)
func main(){
	godotenv.Load()

	database.GetDB()

	e := echo.New()

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		he, ok := err.(*echo.HTTPError)
		if ok {
			if he.Internal != nil {
				if herr, ok := he.Internal.(*echo.HTTPError); ok {
					he = herr
				}
			}
		}

		// Call the default handler to return the HTTP response
		e.DefaultHTTPErrorHandler(err, c)
	}

		e.GET("/health", app.Health)
		apiRouter := e.Group("/api/v1")
		apiRouter.POST("/webhook", webhook.CreatedUserWebhook)
		apiRouter.POST("/send",app.CreateMail)

		e.Logger.Fatal(e.Start(":1323"))
	}


