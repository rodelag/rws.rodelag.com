package resolvers

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/smtp"
)

func notificacion(mensaje string, error error) {
	var (
		servidor = viper.GetString("correo.host")
		puerto   = viper.GetString("correo.port")
		password = viper.GetString("correo.password")
		correoDe = viper.GetString("correo.mail")
		correoA  = viper.GetString("correo.to")
		asunto   = viper.GetString("correo.asunto")
	)

	msg := "From: " + correoDe + "\n" + "To: " + correoA + "\n" + "Subject: " + asunto + "\n\n" + fmt.Sprint(mensaje, error)

	err := smtp.SendMail(servidor+":"+puerto, smtp.PlainAuth("", correoDe, password, servidor),
		correoDe, []string{correoA}, []byte(msg))

	if err != nil {
		log.Println("Problemas al enviar la notificación: ", err)
	}
}
