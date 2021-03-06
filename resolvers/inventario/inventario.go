package inventario

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
)

type Inventario struct {
	Consecutivo,
	Suc,
	Departamento,
	Codigo,
	Descripcion,
	Precio,
	Cant,
	Parte,
	Marca,
	Oferta,
	FecIni,
	FecFin string
}

func conexionInventario() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.elconix.user"),
		viper.GetString("basedatos.mysql.elconix.password"),
		viper.GetString("basedatos.mysql.elconix.server"),
		viper.GetString("basedatos.mysql.elconix.database"),
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: (Inventario) ", true, errMySQL)
	}
	return connMySQL
}

func ListarInventario(busqueda string, cursor, limite int) []Inventario {
	connMySQL := conexionInventario()
	defer connMySQL.Close()

	rows, err := connMySQL.Query(consulta(busqueda, cursor, limite))
	utils.LogError("Problemas al listar los registros de la base de datos: (Inventario) ", false, err)
	defer rows.Close()

	con, cons := Inventario{}, []Inventario{}

	for rows.Next() {
		err := rows.Scan(
			&con.Consecutivo,
			&con.Suc,
			&con.Departamento,
			&con.Codigo,
			&con.Descripcion,
			&con.Precio,
			&con.Cant,
			&con.Parte,
			&con.Marca,
			&con.Oferta,
			&con.FecIni,
			&con.FecFin,
		)
		utils.LogError("Problemas leer los datos: (Inventario) ", false, err)
		cons = append(cons, Inventario{
			Consecutivo:  con.Consecutivo,
			Suc:          con.Suc,
			Departamento: con.Departamento,
			Codigo:       con.Codigo,
			Descripcion:  con.Descripcion,
			Precio:       con.Precio,
			Cant:         con.Cant,
			Parte:        con.Parte,
			Marca:        con.Marca,
			Oferta:       con.Oferta,
			FecIni:       con.FecIni,
			FecFin:       con.FecFin,
		})
	}
	return cons
}

func consulta(busqueda string, cursor, limite int) string {
	// TODO: Estar pendiente del rendimiento de esta consulta... se le puso un límite de 100 y trabajar más adelante en una paginación.
	consulta := `
		SELECT
		   @cursor AS Consecutivo,
		   Case When a.WareHouse ='HERMINIA' Then 'Bodega-CEDI' Else a.WareHouse End AS Suc,
		   Category AS Departamento,
		   b.Item_Number AS Codigo,
		   b.Nombre AS Descripcion,
		   Get_Price(b.id, '') AS Precio,
		   Get_avialable(b.ID,a.WareHouse,'') Cant,  
		   IFNULL(part_number, codigo_externo) AS Parte,
		   Marca,
		   IFNULL((select NewPrice from promotions p join promotions_products pp on p.id =pp.Promotion  where CodeType =b.id and pricelist ='PRECIO REGULAR' and expira >CURDATE() order by expira desc limit 1)  ,0) AS Oferta,
		   IFNULL((select Fecha_Inicio from promotions p join promotions_products pp on p.id =pp.Promotion where CodeType =b.id and pricelist ='PRECIO REGULAR' and expira >CURDATE() order by expira desc limit 1)  ,'') AS FecIni,
		   IFNULL((select Expira from promotions p join promotions_products pp on p.id =pp.Promotion where CodeType =b.id and pricelist ='PRECIO REGULAR' and expira >CURDATE() order by expira desc limit 1)  ,'') AS FecFin
		FROM
			(SELECT @cursor := 0) c,
		   enx_rodelag.products_mview_instock_actualizado AS a
		   INNER JOIN products AS b ON a.Item = b.ID
		   inner join warehouses w2 on a.WareHouse =w2.nombre
		WHERE a.WareHouse not like('bodega%%') and WareHouse not like ('inco%%') and WareHouse not like ('%%out%%') and WareHouse not like ('obso%%') and      
			 CONCAT(replace(replace(replace(REPLACE(b.nombre,'  ',' '),'  ',' '),'  ',' '),'"',''), b.Item_Number, IFNULL(part_number, codigo_externo)) LIKE '%%%s%%' AND b.Status ='ACTIVO'
			AND
			(@cursor := @cursor + 1) > %d LIMIT %d;
	`
	return fmt.Sprintf(consulta, busqueda, cursor, limite)
}
