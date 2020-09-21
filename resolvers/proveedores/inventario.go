package proveedores

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"strings"
)

type ProveedorInventario struct {
	Consecutivo int
	NomSuc,
	Depto,
	Categ,
	Parte,
	Codigo,
	Descripcion,
	PreReg,
	Oferta,
	Exist,
	Marca,
	OfeIni,
	OfeFin,
	ProvIDElconix,
	NomProv,
	ProvID,
	CategID string
}

func conexionProveedorInventario() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.elconix.user"),
		viper.GetString("basedatos.mysql.elconix.password"),
		viper.GetString("basedatos.mysql.elconix.server"),
		viper.GetString("basedatos.mysql.elconix.database"),
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: ", errMySQL)
	}
	return connMySQL
}

func ListarProveedorInventario(sucursal, proveedor, proveedorID, campo, condicion string) []ProveedorInventario {
	connMySQL := conexionProveedorInventario()
	defer connMySQL.Close()

	rows, err := connMySQL.Query(consulta(sucursal, proveedor, proveedorID, campo, condicion))
	utils.LogError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	pi, psi := ProveedorInventario{}, []ProveedorInventario{}

	for rows.Next() {
		err := rows.Scan(&pi.Consecutivo, &pi.NomSuc, &pi.Depto, &pi.Categ, &pi.Parte, &pi.Codigo, &pi.Descripcion, &pi.PreReg, &pi.Oferta, &pi.Exist, &pi.Marca, &pi.OfeIni, &pi.OfeFin, &pi.ProvIDElconix, &pi.NomProv, &pi.ProvID, &pi.CategID)
		utils.LogError("Problemas leer los datos: ", err)
		psi = append(psi, ProveedorInventario{
			Consecutivo:   pi.Consecutivo,
			NomSuc:        pi.NomSuc,
			Depto:         pi.Depto,
			Categ:         pi.Categ,
			Parte:         pi.Parte,
			Codigo:        pi.Codigo,
			Descripcion:   strings.TrimSpace("\t" + pi.Descripcion + "\n"),
			PreReg:        pi.PreReg,
			Oferta:        pi.Oferta,
			Exist:         pi.Exist,
			Marca:         pi.Marca,
			OfeIni:        pi.OfeIni,
			OfeFin:        pi.OfeFin,
			ProvIDElconix: pi.ProvIDElconix,
			NomProv:       strings.TrimSpace("\t" + pi.NomProv + "\n"),
			ProvID:        strings.TrimSpace("\t" + pi.ProvID + "\n"),
			CategID:       pi.CategID,
		})
	}
	return psi
}

//Esta función más adelante puede cambiarse por un llamado a una API
func consulta(sucursal, proveedor, proveedorID, campo, condicion string) string {
	consulta := `
		SELECT
			(@row_number := @row_number + 1) AS Consecutivo,
			b.Sucursal AS NomSuc,
			a.Category AS Depto,
			a.Category_L2 AS Categ,
			IFNULL(Codigo_Externo, Part_Number) AS Parte,
			Item_Number AS Codigo,
			a.Nombre AS Descripcion,
			GET_PRICE(a.id, '') AS PreReg,
			IFNULL((SELECT NewPrice FROM promotions AS p JOIN promotions_products AS pp ON p.id = pp.Promotion WHERE CodeType = a.id AND pricelist = 'PRECIO REGULAR' AND expira > CURDATE() ORDER BY expira DESC LIMIT 1), 0) AS Oferta,
			b.Instock AS Exist,
			Marca AS Marca,
			IFNULL((SELECT Fecha_inicio FROM promotions AS p JOIN promotions_products AS pp ON p.id = pp.Promotion WHERE CodeType = a.id AND pricelist = 'PRECIO REGULAR' AND expira > CURDATE() ORDER BY expira DESC  LIMIT 1), 0) AS OfeIni,
			IFNULL((SELECT Expira FROM promotions AS p JOIN promotions_products AS pp ON p.id = pp.Promotion WHERE CodeType = a.id AND pricelist = 'PRECIO REGULAR' AND expira > CURDATE() ORDER BY expira DESC LIMIT 1), 0) AS OfeFin,
			a.Proveedor_Principal AS ProvIDElconix,
			p3.Empresa AS NomProv,
			p3.Codigo_Inportacion AS ProvID,
			0 AS CategID
		FROM
			(SELECT @row_number := 0) cnsc
			JOIN products AS a
				INNER JOIN products_mview_instock_actualizado AS b
						   ON a.id = b.Item
				INNER JOIN providers AS p3
						   ON a.Proveedor_Principal = p3.id
		WHERE
				b.instock > 0
		  AND
				b.Sucursal = '%s'
		  AND
				%s LIKE '%s'
		  AND
				%s LIKE '%s'
	`
	return fmt.Sprintf(consulta, sucursal, proveedor, proveedorID, campo, condicion)
}
