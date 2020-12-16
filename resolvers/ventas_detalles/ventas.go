package ventas_detalles

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
)

type Venta struct {
	RegistroFecha,
	RegistroSucursal,
	RegistroFactura,
	RegistroVendedor,
	RegistroTipoVendedor,
	RegistroEmpresa,
	RegistroNombre,
	RegistroApellido,
	RegistroDepto,
	RegistroCategory_L2,
	RegistroCategory_L3,
	RegistroCodigoID,
	RegistroCodigo,
	RegistroDescripcion,
	RegistroDescripLarga,
	RegistroMarca,
	RegistroParte,
	RegistroUnidades,
	RegistroCosto,
	RegistroVenta,
	RegistroUtilidad,
	RegistroMargen,
	RegistroProveedor,
	RegistroListaPrecio string
}

func conexionVentas() *sql.DB {
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

func ListarVentas(inicio, fin string) []Venta {
	connMySQL := conexionVentas()
	defer connMySQL.Close()

	rows, err := connMySQL.Query(consulta(inicio, fin))
	utils.LogError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	con, cons := Venta{}, []Venta{}

	for rows.Next() {
		err := rows.Scan(
			&con.RegistroFecha,
			&con.RegistroSucursal,
			&con.RegistroFactura,
			&con.RegistroVendedor,
			&con.RegistroTipoVendedor,
			&con.RegistroEmpresa,
			&con.RegistroNombre,
			&con.RegistroApellido,
			&con.RegistroDepto,
			&con.RegistroCategory_L2,
			&con.RegistroCategory_L3,
			&con.RegistroCodigoID,
			&con.RegistroCodigo,
			&con.RegistroDescripcion,
			&con.RegistroDescripLarga,
			&con.RegistroMarca,
			&con.RegistroParte,
			&con.RegistroUnidades,
			&con.RegistroCosto,
			&con.RegistroVenta,
			&con.RegistroUtilidad,
			&con.RegistroMargen,
			&con.RegistroProveedor,
			&con.RegistroListaPrecio,
		)
		utils.LogError("Problemas leer los datos: ", err)
		cons = append(cons, Venta{
			RegistroFecha:        con.RegistroFecha,
			RegistroSucursal:     con.RegistroSucursal,
			RegistroFactura:      con.RegistroFactura,
			RegistroVendedor:     con.RegistroVendedor,
			RegistroTipoVendedor: con.RegistroTipoVendedor,
			RegistroEmpresa:      con.RegistroEmpresa,
			RegistroNombre:       con.RegistroNombre,
			RegistroApellido:     con.RegistroApellido,
			RegistroDepto:        con.RegistroDepto,
			RegistroCategory_L2:  con.RegistroCategory_L2,
			RegistroCategory_L3:  con.RegistroCategory_L3,
			RegistroCodigoID:     con.RegistroCodigoID,
			RegistroCodigo:       con.RegistroCodigo,
			RegistroDescripcion:  con.RegistroDescripcion,
			RegistroDescripLarga: con.RegistroDescripLarga,
			RegistroMarca:        con.RegistroMarca,
			RegistroParte:        con.RegistroParte,
			RegistroUnidades:     con.RegistroUnidades,
			RegistroCosto:        con.RegistroCosto,
			RegistroVenta:        con.RegistroVenta,
			RegistroUtilidad:     con.RegistroUtilidad,
			RegistroMargen:       con.RegistroMargen,
			RegistroProveedor:    con.RegistroProveedor,
			RegistroListaPrecio:  con.RegistroListaPrecio,
		})
	}
	return cons
}

func consulta(inicio, fin string) string {
	consulta := `
		SELECT
			registroFecha,
			registroSucursal,
			registroFactura,
			registroVendedor,
			registroTipoVendedor,
			registroEmpresa,
			registroNombre,
			registroApellido,
			registroDepto,
			registroCategory_L2,
			registroCategory_L3,
			registroCodigoID,
			registroCodigo,
			registroDescripcion,
			registroDescripLarga,
			registroMarca,
			registroParte,
			registroUnidades,
			registroCosto,
			registroVenta,
			registroUtilidad,
			registroMargen,
			registroProveedor,
			registroListaPrecio
		FROM
			rodelag_indicadores.ventas_detalle
		WHERE
			DATE(registroFecha) BETWEEN '%s' AND '%s'
			ORDER BY registroFecha ASC
	`
	return fmt.Sprintf(consulta, inicio, fin)
}
