package formularios

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type Esiaa struct {
	ID int
	Nombre,
	Apellido,
	Cedula,
	Correo,
	Calificacion,
	Atencion,
	ResolverInstalacion,
	TiempoRazonable,
	Recomendacion,
	CalificacionManera,
	FechaRegistro string
}

func conexionEsiaa() *sql.DB {
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

func ListarEsiaa() []Esiaa {
	connMySQL := conexionEsiaa()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT * FROM formulario_esiaa;")
	utils.LogError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	ct := Esiaa{}
	cts := []Esiaa{}

	for rows.Next() {
		err := rows.Scan(
			&ct.ID,
			&ct.Nombre,
			&ct.Apellido,
			&ct.Cedula,
			&ct.Correo,
			&ct.Calificacion,
			&ct.Atencion,
			&ct.ResolverInstalacion,
			&ct.TiempoRazonable,
			&ct.Recomendacion,
			&ct.CalificacionManera,
			&ct.FechaRegistro,
		)
		utils.LogError("Problemas leer los datos: ", err)
		cts = append(cts, Esiaa{
			ID:                  ct.ID,
			Nombre:              ct.Nombre,
			Apellido:            ct.Apellido,
			Cedula:              ct.Cedula,
			Correo:              ct.Correo,
			Calificacion:        ct.Calificacion,
			Atencion:            ct.Atencion,
			ResolverInstalacion: ct.ResolverInstalacion,
			TiempoRazonable:     ct.TiempoRazonable,
			Recomendacion:       ct.Recomendacion,
			CalificacionManera:  ct.CalificacionManera,
			FechaRegistro:       ct.FechaRegistro,
		})
	}
	return cts
}

func CrearEsiaa(nombre, apellido, cedula, correo, calificacion, atencion, resolverInstalacion, tiempoRazonable, recomendacion, calificacionManera string) Esiaa {
	ct := Esiaa{
		Nombre:              nombre,
		Apellido:            apellido,
		Cedula:              cedula,
		Correo:              correo,
		Calificacion:        calificacion,
		Atencion:            atencion,
		ResolverInstalacion: resolverInstalacion,
		TiempoRazonable:     tiempoRazonable,
		Recomendacion:       recomendacion,
		CalificacionManera:  calificacionManera,
		FechaRegistro:       time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionEsiaa()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_esiaa (nombre, apellido, cedula, correo, calificacion, atencion, resolverInstalacion, tiempoRazonable, recomendacion, calificacionManera, FechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(
		ct.Nombre,
		ct.Apellido,
		ct.Cedula,
		ct.Correo,
		ct.Calificacion,
		ct.Atencion,
		ct.ResolverInstalacion,
		ct.TiempoRazonable,
		ct.Recomendacion,
		ct.CalificacionManera,
		ct.FechaRegistro,
	)

	return ct
}
