package resolvers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type PagoACH struct {
	ID              int
	Nombre          string
	Apellido        string
	TitularCuenta   string
	Cedula          string
	Correo          string
	Telefono        string
	CompraOrigen    string
	NumeroOrden     string
	FotoComprobante string
	FechaRegistro   string
}

func configuracion() {
	viper.SetConfigName("configuracion")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func logError(m string, e error) {
	configuracion()
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

func conexion() *sql.DB {
	configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.user"),
		viper.GetString("basedatos.mysql.password"),
		viper.GetString("basedatos.mysql.server"),
		viper.GetString("basedatos.mysql.database"),
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		logError("Problemas con la conexion a mysql: ", errMySQL)
	}
	return connMySQL
}

func ListarPagosACH() []PagoACH {
	rows, err := conexion().Query("SELECT * FROM ovnicom_pagosach;")
	logError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	pagoACH := PagoACH{}
	pagosACH := []PagoACH{}

	for rows.Next() {
		err := rows.Scan(&pagoACH.ID, &pagoACH.Nombre, &pagoACH.Apellido, &pagoACH.TitularCuenta, &pagoACH.Cedula, &pagoACH.Correo, &pagoACH.Telefono, &pagoACH.CompraOrigen, &pagoACH.NumeroOrden, &pagoACH.FotoComprobante, &pagoACH.FechaRegistro)
		logError("Problemas leer los datos: ", err)
		pagosACH = append(pagosACH, PagoACH{
			ID:              pagoACH.ID,
			Nombre:          pagoACH.Nombre,
			Apellido:        pagoACH.Apellido,
			TitularCuenta:   pagoACH.TitularCuenta,
			Cedula:          pagoACH.Cedula,
			Correo:          pagoACH.Correo,
			Telefono:        pagoACH.Telefono,
			CompraOrigen:    pagoACH.CompraOrigen,
			NumeroOrden:     pagoACH.NumeroOrden,
			FotoComprobante: pagoACH.FotoComprobante,
			FechaRegistro:   pagoACH.FechaRegistro,
		})
	}
	return pagosACH
}

func CrearPagoACH(nombre string, apellido string, titularCuenta string, cedula string, correo string, telefono string, compraOrigen string, numeroOrden string, fotoComprobante string) PagoACH {
	pagosACH := PagoACH{
		Nombre:          nombre,
		Apellido:        apellido,
		TitularCuenta:   titularCuenta,
		Cedula:          cedula,
		Correo:          correo,
		Telefono:        telefono,
		CompraOrigen:    compraOrigen,
		NumeroOrden:     numeroOrden,
		FotoComprobante: fotoComprobante,
		FechaRegistro:   time.Now().UTC().Format("2006-01-02 15:04:05"),
	}

	conn, err := conexion().Prepare("INSERT INTO ovnicom_pagosach (nombre, apellido, titularCuenta, cedula, correo, telefono, compraOrigen, numeroOrden, fotoComprobante, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	logError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(pagosACH.Nombre, pagosACH.Apellido, pagosACH.TitularCuenta, pagosACH.Cedula, pagosACH.Correo, pagosACH.Telefono, pagosACH.CompraOrigen, pagosACH.NumeroOrden, pagosACH.FotoComprobante, pagosACH.FechaRegistro)

	return pagosACH
}
