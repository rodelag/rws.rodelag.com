package formularios

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type SolicitudTarjetaRodelag struct {
	ID int
	Sucursal,
	Nombre,
	Apellido,
	FechaNacimiento,
	LugarNacimiento,
	Nacionalidad,
	Cedula,
	FotoCedula,
	FotoFicha,
	FotoRecibo,
	EstadoCivil,
	Correo,
	DireccionResidencia,
	Barrio,
	Provincia,
	TelefonoResidencia,
	Celular,
	Credito,
	Educacion,
	NombreEmpresa,
	TipoNegocio,
	Cargo,
	TiempoLaboral,
	DireccionTrabajo,
	TelefonoTrabajo,
	Extension,
	SalarioMensual,
	FuentesIngreso,
	MontoFuentesIngreso,
	DetalleFuentesIngreso,
	NombreReferenciaUno,
	TelefonoReferenciaUno,
	NombreReferenciaDos,
	TelefonoReferenciaDos,
	NombreReferenciaTres,
	TelefonoReferenciaTres,
	FechaRegistro string
}

func conexionSolicitudTarjetaRodelag() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.rodelag.user"),
		viper.GetString("basedatos.mysql.rodelag.password"),
		viper.GetString("basedatos.mysql.rodelag.server"),
		viper.GetString("basedatos.mysql.rodelag.database"),
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: ", errMySQL)
	}
	return connMySQL
}

func ListarSolicitudTarjetaRodelag() []SolicitudTarjetaRodelag {
	connMySQL := conexionSolicitudTarjetaRodelag()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT * FROM formulario_solicitudtarjeta;")
	utils.LogError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	str := SolicitudTarjetaRodelag{}
	strs := []SolicitudTarjetaRodelag{}

	for rows.Next() {
		err := rows.Scan(
			&str.ID,
			&str.Sucursal,
			&str.Nombre,
			&str.Apellido,
			&str.FechaNacimiento,
			&str.LugarNacimiento,
			&str.Nacionalidad,
			&str.Cedula,
			&str.FotoCedula,
			&str.FotoFicha,
			&str.FotoRecibo,
			&str.EstadoCivil,
			&str.Correo,
			&str.DireccionResidencia,
			&str.Barrio,
			&str.Provincia,
			&str.TelefonoResidencia,
			&str.Celular,
			&str.Credito,
			&str.Educacion,
			&str.NombreEmpresa,
			&str.TipoNegocio,
			&str.Cargo,
			&str.TiempoLaboral,
			&str.DireccionTrabajo,
			&str.TelefonoTrabajo,
			&str.Extension,
			&str.SalarioMensual,
			&str.FuentesIngreso,
			&str.MontoFuentesIngreso,
			&str.DetalleFuentesIngreso,
			&str.NombreReferenciaUno,
			&str.TelefonoReferenciaUno,
			&str.NombreReferenciaDos,
			&str.TelefonoReferenciaDos,
			&str.NombreReferenciaTres,
			&str.TelefonoReferenciaTres,
			&str.FechaRegistro,
		)
		utils.LogError("Problemas leer los datos: ", err)
		strs = append(strs, SolicitudTarjetaRodelag{
			ID:                     str.ID,
			Sucursal:               str.Sucursal,
			Nombre:                 str.Nombre,
			Apellido:               str.Apellido,
			FechaNacimiento:        str.FechaNacimiento,
			LugarNacimiento:        str.LugarNacimiento,
			Nacionalidad:           str.Nacionalidad,
			Cedula:                 str.Cedula,
			FotoCedula:             str.FotoCedula,
			FotoFicha:              str.FotoFicha,
			FotoRecibo:             str.FotoRecibo,
			EstadoCivil:            str.EstadoCivil,
			Correo:                 str.Correo,
			DireccionResidencia:    str.DireccionResidencia,
			Barrio:                 str.Barrio,
			Provincia:              str.Provincia,
			TelefonoResidencia:     str.TelefonoResidencia,
			Celular:                str.Celular,
			Credito:                str.Credito,
			Educacion:              str.Educacion,
			NombreEmpresa:          str.NombreEmpresa,
			TipoNegocio:            str.TipoNegocio,
			Cargo:                  str.Cargo,
			TiempoLaboral:          str.TiempoLaboral,
			DireccionTrabajo:       str.DireccionTrabajo,
			TelefonoTrabajo:        str.TelefonoTrabajo,
			Extension:              str.Extension,
			SalarioMensual:         str.SalarioMensual,
			FuentesIngreso:         str.FuentesIngreso,
			MontoFuentesIngreso:    str.MontoFuentesIngreso,
			DetalleFuentesIngreso:  str.DetalleFuentesIngreso,
			NombreReferenciaUno:    str.NombreReferenciaUno,
			TelefonoReferenciaUno:  str.TelefonoReferenciaUno,
			NombreReferenciaDos:    str.NombreReferenciaDos,
			TelefonoReferenciaDos:  str.TelefonoReferenciaDos,
			NombreReferenciaTres:   str.NombreReferenciaTres,
			TelefonoReferenciaTres: str.TelefonoReferenciaTres,
			FechaRegistro:          str.FechaRegistro,
		})
	}
	return strs
}

func CrearSolicitudTarjetaRodelag(
	sucursal,
	nombre,
	apellido,
	fechaNacimiento,
	lugarNacimiento,
	nacionalidad,
	cedula,
	fotoCedula,
	fotoFicha,
	fotoRecibo,
	estadoCivil,
	correo,
	direccionResidencia,
	barrio,
	provincia,
	telefonoResidencia,
	celular,
	credito,
	educacion,
	nombreEmpresa,
	tipoNegocio,
	cargo,
	tiempoLaboral,
	direccionTrabajo,
	telefonoTrabajo,
	extension,
	salarioMensual,
	fuentesIngreso,
	montoFuentesIngreso,
	detalleFuentesIngreso,
	nombreReferenciaUno,
	telefonoReferenciaUno,
	nombreReferenciaDos,
	telefonoReferenciaDos,
	nombreReferenciaTres,
	telefonoReferenciaTres string,
) SolicitudTarjetaRodelag {
	sec := SolicitudTarjetaRodelag{
		Sucursal:               sucursal,
		Nombre:                 nombre,
		Apellido:               apellido,
		FechaNacimiento:        fechaNacimiento,
		LugarNacimiento:        lugarNacimiento,
		Nacionalidad:           nacionalidad,
		Cedula:                 cedula,
		FotoCedula:             fotoCedula,
		FotoFicha:              fotoFicha,
		FotoRecibo:             fotoRecibo,
		EstadoCivil:            estadoCivil,
		Correo:                 correo,
		DireccionResidencia:    direccionResidencia,
		Barrio:                 barrio,
		Provincia:              provincia,
		TelefonoResidencia:     telefonoResidencia,
		Celular:                celular,
		Credito:                credito,
		Educacion:              educacion,
		NombreEmpresa:          nombreEmpresa,
		TipoNegocio:            tipoNegocio,
		Cargo:                  cargo,
		TiempoLaboral:          tiempoLaboral,
		DireccionTrabajo:       direccionTrabajo,
		TelefonoTrabajo:        telefonoTrabajo,
		Extension:              extension,
		SalarioMensual:         salarioMensual,
		FuentesIngreso:         fuentesIngreso,
		MontoFuentesIngreso:    montoFuentesIngreso,
		DetalleFuentesIngreso:  detalleFuentesIngreso,
		NombreReferenciaUno:    nombreReferenciaUno,
		TelefonoReferenciaUno:  telefonoReferenciaUno,
		NombreReferenciaDos:    nombreReferenciaDos,
		TelefonoReferenciaDos:  telefonoReferenciaDos,
		NombreReferenciaTres:   nombreReferenciaTres,
		TelefonoReferenciaTres: telefonoReferenciaTres,
		FechaRegistro:          time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionSolicitudTarjetaRodelag()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_solicitudtarjeta (sucursal, nombre, apellido, fechaNacimiento, lugarNacimiento, nacionalidad, cedula, fotoCedula, fotoFicha, fotoRecibo, estadoCivil, correo, direccionResidencia, barrio, provincia, telefonoResidencia, celular, credito, educacion, nombreEmpresa, tipoNegocio, cargo, tiempoLaboral, direccionTrabajo, telefonoTrabajo, extension, salarioMensual, fuentesIngreso, montoFuentesIngreso, detalleFuentesIngreso, nombreReferenciaUno, telefonoReferenciaUno, nombreReferenciaDos, telefonoReferenciaDos, nombreReferenciaTres, telefonoReferenciaTres, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(
		sec.Sucursal,
		sec.Nombre,
		sec.Apellido,
		sec.FechaNacimiento,
		sec.LugarNacimiento,
		sec.Nacionalidad,
		sec.Cedula,
		sec.FotoCedula,
		sec.FotoFicha,
		sec.FotoRecibo,
		sec.EstadoCivil,
		sec.Correo,
		sec.DireccionResidencia,
		sec.Barrio,
		sec.Provincia,
		sec.TelefonoResidencia,
		sec.Celular,
		sec.Credito,
		sec.Educacion,
		sec.NombreEmpresa,
		sec.TipoNegocio,
		sec.Cargo,
		sec.TiempoLaboral,
		sec.DireccionTrabajo,
		sec.TelefonoTrabajo,
		sec.Extension,
		sec.SalarioMensual,
		sec.FuentesIngreso,
		sec.MontoFuentesIngreso,
		sec.DetalleFuentesIngreso,
		sec.NombreReferenciaUno,
		sec.TelefonoReferenciaUno,
		sec.NombreReferenciaDos,
		sec.TelefonoReferenciaDos,
		sec.NombreReferenciaTres,
		sec.TelefonoReferenciaTres,
		sec.FechaRegistro,
	)

	return sec
}