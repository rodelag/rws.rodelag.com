package central

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
)

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

func CentralTelefonica(consulta string) Colaborador {
	connRodelag := conexionCentralTelefonica()
	defer connRodelag.Close()

	colaborador := Colaborador{}

	err := connRodelag.QueryRow(consultaExtension(consulta)).Scan(
		&colaborador.RegistroNombre,
		&colaborador.RegistroExtension,
		nil,
	)
	utils.LogError("Problemas al leer registro: ", err)

	return colaborador
}

func consultaExtension(consulta string) string {
	c := `
		SELECT
			registroNombre,
			registroExtension,
			MATCH (registroNombre) AGAINST ('+%s' IN NATURAL LANGUAGE MODE WITH QUERY EXPANSION) AS score
		FROM
			rodelag_centraltelefonica.central
		WHERE
			MATCH (registroNombre) AGAINST ('+%s' IN NATURAL LANGUAGE MODE WITH QUERY EXPANSION) HAVING score > 5
			LIMIT 1;
	`
	return fmt.Sprintf(c, consulta, consulta)
}
