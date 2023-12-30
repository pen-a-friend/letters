package webhook

import (
	"net/http"
	"pen-a-friend/pkg/database"

	"github.com/labstack/echo/v4"
)

type UserCreatedWebhook struct{
		Data struct {
		Birthday       string `json:"birthday"`
		CreatedAt      int64  `json:"created_at"`
		EmailAddresses []struct {
			EmailAddress string        `json:"email_address"`
			ID           string        `json:"id"`
			LinkedTo     []interface{} `json:"linked_to"`
			Object       string        `json:"object"`
			Verification struct {
				Status   string `json:"status"`
				Strategy string `json:"strategy"`
			} `json:"verification"`
		} `json:"email_addresses"`
		ExternalAccounts      []interface{} `json:"external_accounts"`
		ExternalID            string        `json:"external_id"`
		FirstName             string        `json:"first_name"`
		Gender                string        `json:"gender"`
		ID                    string        `json:"id"`
		LastName              string        `json:"last_name"`
		Locked                bool          `json:"locked"`
		LastSignInAt          int64         `json:"last_sign_in_at"`
		Object                string        `json:"object"`
		PasswordEnabled       bool          `json:"password_enabled"`
		PhoneNumbers          []interface{} `json:"phone_numbers"`
		PrimaryEmailAddressID string        `json:"primary_email_address_id"`
		PrimaryPhoneNumberID  interface{}   `json:"primary_phone_number_id"`
		PrimaryWeb3WalletID   interface{}   `json:"primary_web3_wallet_id"`
		PrivateMetadata       struct {
		} `json:"private_metadata"`
		ProfileImageURL string `json:"profile_image_url"`
		PublicMetadata  struct {
		} `json:"public_metadata"`
		TwoFactorEnabled bool `json:"two_factor_enabled"`
		UnsafeMetadata   struct {
		} `json:"unsafe_metadata"`
		UpdatedAt   int64         `json:"updated_at"`
		Username    interface{}   `json:"username"`
		Web3Wallets []interface{} `json:"web3_wallets"`
	} `json:"data"`
	Object string `json:"object"`
	Type   string `json:"type"`
}

func CreatedUserWebhook(c echo.Context) error{
	
	body := new(UserCreatedWebhook)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if body.Type !="user.created"{
		c.JSON(http.StatusAccepted,map[string]string{"message":"ok"})
	}

	db := database.GetDB()
	users:=database.Users{Email: body.Data.EmailAddresses[0].EmailAddress,FirstName: body.Data.FirstName,LastName: body.Data.LastName,ClerkId: body.Data.ID}
	err:=db.Create(&users)
	if err.Error!=nil{
		return echo.NewHTTPError(http.StatusBadRequest, "hello")
	}

	return	c.JSON(http.StatusAccepted,map[string]string{"message":"ok"})
}