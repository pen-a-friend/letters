package email

import (
	"pen-a-friend/pkg/config"

	"github.com/resendlabs/resend-go"
)

type EmailPayload struct {
    To      string `json:"to"`
    From    string `json:"from"`
    Subject string `json:"subject"`
    Body    string `json:"body"`
}


type statusResponse struct {
	Status string
	Message string
}

func CreateEmail(emailParams EmailPayload) *statusResponse{
	apiKey:=config.Get().ResendApiKey
	client := resend.NewClient(apiKey)
	params := &resend.SendEmailRequest{
        From:    emailParams.From,
        To:      []string{emailParams.To},
        Subject: emailParams.Subject,
        Html:    emailParams.Body,
    }

    response , err := client.Emails.Send(params)
		status := new(statusResponse)
		if err != nil{
			status.Status ="Failed"
			status.Message=err.Error()
		}else{
			status.Status = "Succeeded"
			status.Message = response.Id
		}

		return status
}