package proveedores

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"strings"
)

type ProveedorVentas struct {
	Consecutivo,
	NombreSucursal,
	Departamento,
	Categoria,
	Parte,
	Codigo,
	Descripcion,
	PrecioRegular,
	PrecioOferta,
	Existencia,
	Cantidad,
	VentaTotal,
	Fecha,
	Marca,
	DescripcionLarga,
	ProveedorID,
	NombreProveedor,
	OfertaInicial,
	OfertaFinal,
	CategoriaID string
}

func conexionProveedorVentas() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.rodelag.user"),
		viper.GetString("basedatos.mysql.rodelag.password"),
		viper.GetString("basedatos.mysql.rodelag.server"),
		"rodelag_proveedores",
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: (ProveedorVentas) ", true, errMySQL)
	}
	return connMySQL
}

func ListarProveedorVentas(proveedor, proveedorID, sucursal, campo, condicion, campo2, condicion2, fechainicial, fechafinal string, cursor, limite int) []ProveedorVentas {
	connMySQL := conexionProveedorVentas()
	defer connMySQL.Close()

	rows, err := connMySQL.Query(consultaVentas(proveedor, proveedorID, sucursal, campo, condicion, campo2, condicion2, fechainicial, fechafinal, cursor, limite))
	utils.LogError("Problemas al listar los registros de la base de datos: (ProveedorVentas) ", false, err)
	defer rows.Close()

	pv, psv := ProveedorVentas{}, []ProveedorVentas{}

	for rows.Next() {
		err := rows.Scan(&pv.Consecutivo, &pv.NombreSucursal, &pv.Departamento, &pv.Categoria, &pv.Parte, &pv.Codigo, &pv.Descripcion, &pv.PrecioRegular, &pv.PrecioOferta, &pv.Existencia, &pv.Cantidad, &pv.VentaTotal, &pv.Fecha, &pv.Marca, &pv.DescripcionLarga, &pv.ProveedorID, &pv.NombreProveedor, &pv.OfertaInicial, &pv.OfertaFinal, &pv.CategoriaID)
		utils.LogError("Problemas leer los datos: (ProveedorVentas) ", false, err)
		psv = append(psv, ProveedorVentas{
			Consecutivo:      pv.Consecutivo,
			NombreSucursal:   pv.NombreSucursal,
			Departamento:     pv.Departamento,
			Categoria:        pv.Categoria,
			Parte:            pv.Parte,
			Codigo:           pv.Codigo,
			Descripcion:      strings.TrimSpace("\t" + pv.Descripcion + "\n"),
			PrecioRegular:    pv.PrecioRegular,
			PrecioOferta:     pv.PrecioOferta,
			Existencia:       pv.Existencia,
			Cantidad:         pv.Cantidad,
			VentaTotal:       pv.VentaTotal,
			Fecha:            pv.Fecha,
			Marca:            pv.Marca,
			DescripcionLarga: strings.TrimSpace("\t" + pv.DescripcionLarga + "\n"),
			ProveedorID:      strings.TrimSpace("\t" + pv.ProveedorID + "\n"),
			NombreProveedor:  pv.NombreProveedor,
			OfertaInicial:    pv.OfertaInicial,
			OfertaFinal:      pv.OfertaFinal,
			CategoriaID:      pv.CategoriaID,
		})
	}
	return psv
}

//Esta función más adelante puede cambiarse por un llamado a una API
func consultaVentas(proveedor, proveedorID, sucursal, campo, condicion, campo2, condicion2, fechainicial, fechafinal string, cursor, limite int) string {
	var consulta string

	if condicion2 != "" {
		// INFO: Esta consulta es para los usuarios que son fabricantes y que tienen hasta 2 marcas en su haber.
		consulta = `
			SELECT
					@cursor AS Consecutivo,
					registros.nombreSucursal,
					registros.departamento,
					registros.categoria,
					registros.parte,
					registros.codigo,
					registros.descripcion,
					registros.precioRegular,
					registros.precioOferta,
					registros.existencia,
					registros.cantidad,
					registros.ventaTotal,
					registros.fecha,
					registros.marca,
					registros.descripcionLarga,
					registros.proveedorID,
					registros.nombreProveedor,
					registros.ofertaInicial,
					registros.ofertaFinal,
					registros.categoriaID
				FROM
					(SELECT @cursor := 0) c,
					(
						 SELECT
							 NomSuc AS nombreSucursal,
							 Depto AS departamento,
							 Categ AS categoria,
							 Parte AS parte,
							 Codigo AS codigo,
							 Descripcion AS descripcion,
							 PreReg AS precioRegular,
							 Oferta AS precioOferta,
							 Exist AS existencia,
							 Cant AS cantidad,
							 VtaTOT AS ventaTotal,
							 Fecha AS fecha,
							 Marca AS marca,
							 'DescripcionLarga' AS descripcionLarga,
							 ProvID AS proveedorID,
							 NomProv AS nombreProveedor,
							 OfeIni AS ofertaInicial,
							 OfeFin AS ofertaFinal,
							 CategID AS categoriaID
						 FROM
							 rodelag_proveedores.proveedores_elconix
						 WHERE
								 %s LIKE '%%%s'
						   AND
								 NomSuc LIKE '%%%s'
						   AND
								 %s LIKE '%%%s'
						   AND
							 Fecha BETWEEN '%s' AND '%s'
					UNION
						 SELECT
							 NomSuc AS nombreSucursal,
							 Depto AS departamento,
							 Categ AS categoria,
							 Parte AS parte,
							 Codigo AS codigo,
							 Descripcion AS descripcion,
							 PreReg AS precioRegular,
							 Oferta AS precioOferta,
							 Exist AS existencia,
							 Cant AS cantidad,
							 VtaTOT AS ventaTotal,
							 Fecha AS fecha,
							 Marca AS marca,
							 'DescripcionLarga' AS descripcionLarga,
							 ProvID AS proveedorID,
							 NomProv AS nombreProveedor,
							 OfeIni AS ofertaInicial,
							 OfeFin AS ofertaFinal,
							 CategID AS categoriaID
						 FROM
							 rodelag_proveedores.proveedores_elconix
						 WHERE
								 %s LIKE '%%%s'
						   AND
								 NomSuc LIKE '%%%s'
						   AND
								 %s LIKE '%%%s'
						   AND
							 Fecha BETWEEN '%s' AND '%s') AS registros
				WHERE
					(@cursor := @cursor + 1) > %d LIMIT %d
		`
		if limite == 0 {
			limite = 20
		}

		return fmt.Sprintf(consulta, proveedor, proveedorID, sucursal, campo, condicion, fechainicial, fechafinal, campo2, condicion2, sucursal, campo, condicion, fechainicial, fechafinal, cursor, limite)
	} else {
		// INFO: Esta consulta es para los usuarios que son proveedores.
		consulta = `
			SELECT
				@cursor AS Consecutivo,
				NomSuc AS nombreSucursal,
				Depto AS departamento,
				Categ AS categoria,
				Parte AS parte,
				Codigo AS codigo,
				Descripcion AS descripcion,
				PreReg AS precioRegular,
				Oferta AS precioOferta,
				Exist AS existencia,
				Cant AS cantidad,
				VtaTOT AS ventaTotal,
				Fecha AS fecha,
				Marca AS marca,
				DescripcionLarga AS descripcionLarga,
				ProvID AS proveedorID,
				NomProv AS nombreProveedor,
				OfeIni AS ofertaInicial,
				OfeFin AS ofertaFinal,
				CategID AS categoriaID
			FROM
				(SELECT @cursor := 0) c,
				rodelag_proveedores.proveedores_elconix
			WHERE
				%s LIKE '%%%s'
				AND
				NomSuc LIKE '%%%s'
				AND
				%s LIKE '%%%s'
				AND
				Fecha BETWEEN '%s' AND '%s'
				AND
				(@cursor := @cursor + 1) > %d LIMIT %d
		`
		if limite == 0 {
			limite = 20
		}

		return fmt.Sprintf(consulta, proveedor, proveedorID, sucursal, campo, condicion, fechainicial, fechafinal, cursor, limite)
	}
}
