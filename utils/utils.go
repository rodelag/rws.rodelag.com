package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/smtp"
	"os"
)

func Configuracion() {
	viper.SetConfigName("configuracion")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func LogError(m string, e error) {
	Configuracion()
	f, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	if e != nil {
		log.Printf(m+" %v", e)
		notificacion(m, e)
	}
}

func notificacion(mensaje string, error error) {
	var (
		servidor = viper.GetString("correo.host")
		puerto   = viper.GetString("correo.port")
		usuario  = viper.GetString("correo.user")
		password = viper.GetString("correo.password")
		correoDe = viper.GetString("correo.mail")
		correoA  = viper.GetString("correo.to")
		asunto   = viper.GetString("correo.asunto")
	)

	msg := "From: " + correoDe + "\n" + "To: " + correoA + "\n" + "Subject: " + asunto + "\n\n" + fmt.Sprint(mensaje, error)

	err := smtp.SendMail(servidor+":"+puerto, smtp.PlainAuth("", usuario, password, servidor),
		correoDe, []string{correoA}, []byte(msg))

	if err != nil {
		log.Println("Problemas al enviar la notificación: ", err)
	}
}
