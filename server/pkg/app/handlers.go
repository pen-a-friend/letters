package app

import (
	"fmt"
	"net/http"
	"pen-a-friend/pkg/database"
	"pen-a-friend/pkg/email"

	"github.com/labstack/echo/v4"
)


func CreateMail(c echo.Context) error{
	emailPayload := new(email.EmailPayload)
	if err:=c.Bind(emailPayload); err!= nil{
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}

	newEmail := database.Mail{To: emailPayload.To,From: emailPayload.From,Subject: emailPayload.Subject,Body: emailPayload.Body}
	db:= database.GetDB()
	response := db.Create(&newEmail)
	fmt.Println(newEmail.ID)

	if response.Error !=nil{
			return echo.NewHTTPError(http.StatusBadRequest,response.Error.Error())
	}

	resp:=email.CreateEmail(*emailPayload)
	newEmail.Status=resp.Status
	newEmail.Message=resp.Message
	db.Save(&newEmail)

	return c.JSON(http.StatusAccepted,map[string]string{
		"message":"ok",
	})
}