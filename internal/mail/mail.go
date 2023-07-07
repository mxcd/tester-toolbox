package mail

import (
	"crypto/tls"

	"github.com/mxcd/go-config/config"
	"github.com/rs/zerolog/log"

	gomail "gopkg.in/gomail.v2"
)

func SendMail(targetAddress string) error {
	log.Info().Msgf("Sending test mail to %s", targetAddress)

	m := gomail.NewMessage()
	m.SetHeader("From", config.Get().String("FROM_ADDRESS"))
	m.SetHeader("To", targetAddress)
	m.SetHeader("Subject", "Test mail")
	m.SetBody("text/plain", "This is a test mail")

	d := gomail.NewDialer(
		config.Get().String("SMTP_HOST"),
		config.Get().Int("SMTP_PORT"),
		config.Get().String("SMTP_USERNAME"),
		config.Get().String("SMTP_PASSWORD"),
	)

	if !config.Get().Bool("SMTP_TLS") {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	err := d.DialAndSend(m)
	if err != nil {
		log.Error().Err(err).Msg("Error sending message")
		return err
	}
	return nil
}
