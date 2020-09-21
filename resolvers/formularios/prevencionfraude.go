package formularios

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type PrevencionFraude struct {
	ID int
	Nombre,
	Apellido,
	FechaNacimiento,
	LugarResidencia,
	Celular,
	FotoCedula,
	FotoTarjeta,
	FechaRegistro string
}

func conexionPrevencionFraude() *sql.DB {
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

func ListarPrevencionFraude() []PrevencionFraude {
	connMySQL := conexionPrevencionFraude()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT * FROM formulario_prevencionfraude;")
	utils.LogError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	pf := PrevencionFraude{}
	pfs := []PrevencionFraude{}

	for rows.Next() {
		err := rows.Scan(
			&pf.ID,
			&pf.Nombre,
			&pf.Apellido,
			&pf.FechaNacimiento,
			&pf.LugarResidencia,
			&pf.Celular,
			&pf.FotoCedula,
			&pf.FotoTarjeta,
			&pf.FechaRegistro,
		)
		utils.LogError("Problemas leer los datos: ", err)
		pfs = append(pfs, PrevencionFraude{
			ID:              pf.ID,
			Nombre:          pf.Nombre,
			Apellido:        pf.Apellido,
			FechaNacimiento: pf.FechaNacimiento,
			LugarResidencia: pf.LugarResidencia,
			Celular:         pf.Celular,
			FotoCedula:      pf.FotoCedula,
			FotoTarjeta:     pf.FotoTarjeta,
			FechaRegistro:   pf.FechaRegistro,
		})
	}
	return pfs
}

func CrearPrevencionFraude(nombre, apellido, fechaNacimiento, lugarResidencia, celular, fotoCedula, fotoTarjeta string) PrevencionFraude {
	pf := PrevencionFraude{
		Nombre:          nombre,
		Apellido:        apellido,
		FechaNacimiento: fechaNacimiento,
		LugarResidencia: lugarResidencia,
		Celular:         celular,
		FotoCedula:      fotoCedula,
		FotoTarjeta:     fotoTarjeta,
		FechaRegistro:   time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionPrevencionFraude()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_prevencionfraude (nombre, apellido, fechaNacimiento, lugarResidencia, celular, fotoCedula, fotoTarjeta, FechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(
		pf.Nombre,
		pf.Apellido,
		pf.FechaNacimiento,
		pf.LugarResidencia,
		pf.Celular,
		pf.FotoCedula,
		pf.FotoTarjeta,
		pf.FechaRegistro,
	)

	return pf
}
