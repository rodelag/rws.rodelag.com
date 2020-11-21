package central

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
)

type Central struct {
	RegistroDetalle []Colaborador
}

type Colaborador struct {
	RegistroNombre,
	RegistroExtension string
}

func conexionCentralTelefonica() *sql.DB {
	utils.Configuracion()
	connStringRodelag := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.rodelag.user"),
		viper.GetString("basedatos.mysql.rodelag.password"),
		viper.GetString("basedatos.mysql.rodelag.server"),
		viper.GetString("basedatos.mysql.rodelag.database"),
	)
	connRodelag, errRodelag := sql.Open("mysql", connStringRodelag)
	if errRodelag != nil {
		utils.LogError("Problemas con la conexion a rodelag: ", errRodelag)
	}
	return connRodelag
}

func CentralTelefonica(consulta string) Central {
	connRodelag := conexionCentralTelefonica()
	defer connRodelag.Close()

	central := Central{
		RegistroDetalle: func() []Colaborador {
			rows, err := connRodelag.Query(consultaExtension(consulta))
			utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: ", err)
			defer rows.Close()

			colaborador, colaboradores := Colaborador{}, []Colaborador{}

			for rows.Next() {
				err := rows.Scan(&colaborador.RegistroNombre, &colaborador.RegistroExtension)
				utils.LogError("Problemas al iterar los registros: ", err)
				colaboradores = append(colaboradores, Colaborador{
					RegistroNombre:    colaborador.RegistroNombre,
					RegistroExtension: colaborador.RegistroExtension,
				})
			}
			return colaboradores
		}(),
	}
	return central
}

func consultaExtension(consulta string) string {
	c := `
		SELECT
			registroNombre,
			registroExtension
		FROM
			rodelag_centraltelefonica.central
		WHERE
			registroNombre LIKE '%%%s%%' OR registroExtension LIKE '%%%s%%';
	`
	return fmt.Sprintf(c, consulta, consulta)
}
