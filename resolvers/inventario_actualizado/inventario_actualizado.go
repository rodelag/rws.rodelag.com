package inventario_actualizado

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
)

type InventarioActualizado struct {
	RegistroNombreSucursal,
	RegistroSucBodegas,
	RegistroDepartamento,
	RegistroCategoria,
	RegistroCodigo,
	RegistroNombreVenta,
	RegistroDescripcion,
	RegistroCosto,
	RegistroPrecio,
	RegistroUnidadVenta,
	RegistroVentaTotal,
	RegistroMargen,
	RegistroOferta,
	RegistroInicioOferta,
	RegistroFinOferta,
	RegistroExistencia,
	RegistroExistenciaHerminia,
	RegistroUltimoRecibo,
	RegistroUltimaVenta,
	RegistroCantidadOrdenes,
	RegistroCantidadArticulosOrden,
	RegistroFechaUltimaOrden,
	RegistroProveedor,
	RegistroMarca string
}

type InventarioActualizadoTiendas struct {
	RegistroNombreSucursal string
}

type InventarioActualizadoDepartamentos struct {
	RegistroDepartamento string
}

type InventarioActualizadoCategorias struct {
	RegistroCategoria string
}

type InventarioActualizadoProducto struct {
	RegistroNombreSucursal,
	RegistroDepartamento,
	RegistroCategoria,
	RegistroCodigo,
	RegistroNombreVenta,
	RegistroDescripcion,
	RegistroCosto,
	RegistroPrecio,
	RegistroUnidadVenta,
	RegistroVentaTotal,
	RegistroMargen,
	RegistroOferta,
	RegistroInicioOferta,
	RegistroFinOferta,
	RegistroExistencia,
	RegistroExistenciaHerminia,
	RegistroUltimoRecibo,
	RegistroUltimaVenta,
	RegistroCantidadOrdenes,
	RegistroCantidadArticulosOrden,
	RegistroFechaUltimaOrden,
	RegistroProveedor,
	RegistroMarca string
}

func conexionInventarioActualizado() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.elconix.user"),
		viper.GetString("basedatos.mysql.elconix.password"),
		viper.GetString("basedatos.mysql.elconix.server"),
		viper.GetString("basedatos.mysql.elconix.database"),
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: (Inventario Actualizado) ", true, errMySQL)
	}
	return connMySQL
}

func ListarInventarioActualizado(sucursal, departamento, categoria string) []InventarioActualizado {
	connMySQL := conexionInventarioActualizado()
	defer connMySQL.Close()

	rows, err := connMySQL.Query(consultaDetalle(sucursal, departamento, categoria))
	utils.LogError("Problemas al listar los registros de la base de datos: (Inventario Actualizado) ", false, err)
	defer rows.Close()

	con, cons := InventarioActualizado{}, []InventarioActualizado{}

	for rows.Next() {
		err := rows.Scan(
			&con.RegistroNombreSucursal,
			&con.RegistroSucBodegas,
			&con.RegistroDepartamento,
			&con.RegistroCategoria,
			&con.RegistroCodigo,
			&con.RegistroNombreVenta,
			&con.RegistroDescripcion,
			&con.RegistroCosto,
			&con.RegistroPrecio,
			&con.RegistroUnidadVenta,
			&con.RegistroVentaTotal,
			&con.RegistroMargen,
			&con.RegistroOferta,
			&con.RegistroInicioOferta,
			&con.RegistroFinOferta,
			&con.RegistroExistencia,
			&con.RegistroExistenciaHerminia,
			&con.RegistroUltimoRecibo,
			&con.RegistroUltimaVenta,
			&con.RegistroCantidadOrdenes,
			&con.RegistroCantidadArticulosOrden,
			&con.RegistroFechaUltimaOrden,
			&con.RegistroProveedor,
			&con.RegistroMarca,
		)
		utils.LogError("Problemas leer los datos: (Inventario Actualizado) ", false, err)
		cons = append(cons, InventarioActualizado{
			RegistroNombreSucursal:         con.RegistroNombreSucursal,
			RegistroSucBodegas:             con.RegistroSucBodegas,
			RegistroDepartamento:           con.RegistroDepartamento,
			RegistroCategoria:              con.RegistroCategoria,
			RegistroCodigo:                 con.RegistroCodigo,
			RegistroNombreVenta:            con.RegistroNombreVenta,
			RegistroDescripcion:            con.RegistroDescripcion,
			RegistroCosto:                  con.RegistroCosto,
			RegistroPrecio:                 con.RegistroPrecio,
			RegistroUnidadVenta:            con.RegistroUnidadVenta,
			RegistroVentaTotal:             con.RegistroVentaTotal,
			RegistroMargen:                 con.RegistroMargen,
			RegistroOferta:                 con.RegistroOferta,
			RegistroInicioOferta:           con.RegistroInicioOferta,
			RegistroFinOferta:              con.RegistroFinOferta,
			RegistroExistencia:             con.RegistroExistencia,
			RegistroExistenciaHerminia:     con.RegistroExistenciaHerminia,
			RegistroUltimoRecibo:           con.RegistroUltimoRecibo,
			RegistroUltimaVenta:            con.RegistroUltimaVenta,
			RegistroCantidadOrdenes:        con.RegistroCantidadOrdenes,
			RegistroCantidadArticulosOrden: con.RegistroCantidadArticulosOrden,
			RegistroFechaUltimaOrden:       con.RegistroFechaUltimaOrden,
			RegistroProveedor:              con.RegistroProveedor,
			RegistroMarca:                  con.RegistroMarca,
		})
	}
	return cons
}

func ListarInventarioActualizadoTiendas() []InventarioActualizadoTiendas {
	connMySQL := conexionInventarioActualizado()
	defer connMySQL.Close()

	rows, err := connMySQL.Query(consultaTiendas())
	utils.LogError("Problemas al listar los registros de la base de datos: (Inventario Actualizado - Tiendas) ", false, err)
	defer rows.Close()

	con, cons := InventarioActualizadoTiendas{}, []InventarioActualizadoTiendas{}

	for rows.Next() {
		err := rows.Scan(
			&con.RegistroNombreSucursal,
		)
		utils.LogError("Problemas leer los datos: (Inventario Actualizado - Tiendas) ", false, err)
		cons = append(cons, InventarioActualizadoTiendas{
			RegistroNombreSucursal: con.RegistroNombreSucursal,
		})
	}
	return cons
}

func ListarInventarioActualizadoDepartamento(sucursal string) []InventarioActualizadoDepartamentos {
	connMySQL := conexionInventarioActualizado()
	defer connMySQL.Close()

	rows, err := connMySQL.Query(consultaDepartamentos(sucursal))
	utils.LogError("Problemas al listar los registros de la base de datos: (Inventario Actualizado - Departamento) ", false, err)
	defer rows.Close()

	con, cons := InventarioActualizadoDepartamentos{}, []InventarioActualizadoDepartamentos{}

	for rows.Next() {
		err := rows.Scan(
			&con.RegistroDepartamento,
		)
		utils.LogError("Problemas leer los datos: (Inventario Actualizado - Departamento) ", false, err)
		cons = append(cons, InventarioActualizadoDepartamentos{
			RegistroDepartamento: con.RegistroDepartamento,
		})
	}
	return cons
}

func ListarInventarioActualizadoCategorias(sucursal, departamento string) []InventarioActualizadoCategorias {
	connMySQL := conexionInventarioActualizado()
	defer connMySQL.Close()

	rows, err := connMySQL.Query(consultaCategorias(sucursal, departamento))
	utils.LogError("Problemas al listar los registros de la base de datos: (Inventario Actualizado - Categorias) ", false, err)
	defer rows.Close()

	con, cons := InventarioActualizadoCategorias{}, []InventarioActualizadoCategorias{}

	for rows.Next() {
		err := rows.Scan(
			&con.RegistroCategoria,
		)
		utils.LogError("Problemas leer los datos: (Inventario Actualizado - Categorias) ", false, err)
		cons = append(cons, InventarioActualizadoCategorias{
			RegistroCategoria: con.RegistroCategoria,
		})
	}
	return cons
}

func ListarInventarioActualizadoProducto(codigo, sucursal string) InventarioActualizadoProducto {
	connMySQL := conexionInventarioActualizado()
	defer connMySQL.Close()

	con := InventarioActualizadoProducto{}

	err := connMySQL.QueryRow(consultaProductos(codigo, sucursal)).Scan(
		&con.RegistroNombreSucursal,
		&con.RegistroDepartamento,
		&con.RegistroCategoria,
		&con.RegistroCodigo,
		&con.RegistroNombreVenta,
		&con.RegistroDescripcion,
		&con.RegistroCosto,
		&con.RegistroPrecio,
		&con.RegistroUnidadVenta,
		&con.RegistroVentaTotal,
		&con.RegistroMargen,
		&con.RegistroOferta,
		&con.RegistroInicioOferta,
		&con.RegistroFinOferta,
		&con.RegistroExistencia,
		&con.RegistroExistenciaHerminia,
		&con.RegistroUltimoRecibo,
		&con.RegistroUltimaVenta,
		&con.RegistroCantidadOrdenes,
		&con.RegistroCantidadArticulosOrden,
		&con.RegistroFechaUltimaOrden,
		&con.RegistroProveedor,
		&con.RegistroMarca,
	)
	utils.LogError("Problemas al listar los registros de la base de datos: (Inventario Actualizado - Producto) ", false, err)

	return con
}

func consultaDetalle(sucursal, departamento, categoria string) string {
	consulta := `
		SELECT
			a.Sucursal AS NombreSucursal,
			a.warehouse AS SucBodegas,
			Category AS Departamento,
			Category_L2 AS Categoria,
			Item_number AS Codigo,
			Nombre_Venta AS NombreVenta,
			nombre AS Descripcion,
			Item_AvgCost_Calc(b.id, 'NACIONALIZADA', SYSDATE()) AS Costo,
			Get_Price(b.id, '') AS Precio,
			ifnull((select sum(unidades) from mviews_sales_details msd where msd.Sucursal =a.sucursal and codigo =b.id and date between date_add(date_add(LAST_DAY(now()),interval 1 DAY),interval -6 MONTH) and LAST_DAY(sysdate())),0) AS UnidadVenta,
			ifnull((select sum(venta) from mviews_sales_details msd where msd.Sucursal =a.sucursal and codigo =b.id and date between date_add(date_add(LAST_DAY(now()),interval 1 DAY),interval -6 MONTH) and LAST_DAY(sysdate())),0) AS VentaTotal,
			round((1-(Item_AvgCost_Calc(b.id, 'NACIONALIZADA', SYSDATE())/Get_Price(b.id, ''))) *100,2) AS Margen,
			Ifnull((select NewPrice from promotions p join promotions_products pp on p.id =pp.Promotion where CodeType =b.id and pricelist ='PRECIO REGULAR' and expira >CURDATE() order by expira desc limit 1)  ,0) AS Oferta,
			Ifnull((select Fecha_Inicio from promotions p join promotions_products pp on p.id =pp.Promotion where CodeType =b.id and pricelist ='PRECIO REGULAR' and expira >CURDATE() order by expira desc limit 1)  ,'') AS InicioOferta,
			Ifnull((select Expira from promotions p join promotions_products pp on p.id =pp.Promotion where CodeType =b.id and pricelist ='PRECIO REGULAR' and expira >CURDATE() order by expira desc limit 1)  ,'') AS FinOferta,
			a.instock AS Existencia,
			Get_InStock(b.id,'herminia') AS ExistenciaHerminia,
			IFNULL((select IFNULL((MAX(t1.Actualizado)),'') from purchases_inventory_items t1 join purchases t2 on t1.Purchase = t2.id where t1.Item = b.id and t2.Status <> 'ABORTED' and t2.Status <> 'IN-TRANSIT' and t2.Type ='PURCHASE-INV' and t2.Sucursal =a.sucursal), '') AS UltimoRecibo,
			IFNULL((select max(date) from mviews_sales_details msd where msd.Sucursal =a.sucursal and msd.Codigo =b.ID), '') AS UltimaVenta,
			(select IFNULL(count(distinct t2.id),'') from purchases_inventory_items t1 join purchases t2 on t1.Purchase = t2.id where t2.Sucursal =a.sucursal and t1.Item = b.id and t2.Status <> 'ABORTED' and t2.Status in('IN-TRANSIT','PURCHASED')) AS CantidadOrdenes,
			(select IFNULL(sum(t1.units),'') from purchases_inventory_items t1 join purchases t2 on t1.Purchase = t2.id where t2.Sucursal =a.sucursal and t1.Item = b.id and t2.Status <> 'ABORTED' and t2.Status in('IN-TRANSIT','PURCHASED')) AS CantidadArticulosOrden,
			(select IFNULL(MAX(t1.Actualizado),'') from purchases_inventory_items t1 join purchases t2 on t1.Purchase = t2.id where t2.Sucursal =a.sucursal and t1.Item = b.id and t2.Status <> 'ABORTED' and t2.Status <> 'IN-TRANSIT') AS FechaUltimaOrden,
			(select empresa from providers p2 where id =b.Proveedor_Principal) AS Proveedor,
			b.Marca AS Marca
		FROM
			 enx_rodelag.products_mview_instock_actualizado AS a
				 INNER JOIN products AS b ON a.Item = b.ID
		WHERE
			a.sucursal <> 'Ventas Comerciales'
			AND a.Sucursal = '%s' AND Category = '%s' AND Category_L2 = '%s';
	`
	return fmt.Sprintf(consulta, sucursal, departamento, categoria)
}

func consultaTiendas() string {
	consulta := `
		SELECT
			a.Sucursal AS NombreSucursal
		FROM
			enx_rodelag.products_mview_instock_actualizado AS a
				INNER JOIN products AS b ON a.Item = b.ID
		WHERE
				a.sucursal <> 'Ventas Comerciales'
				AND a.Sucursal != '' GROUP BY a.Sucursal;
	`
	return consulta
}

func consultaDepartamentos(sucursal string) string {
	consulta := `
		SELECT
			Category AS Departamento
		FROM
			enx_rodelag.products_mview_instock_actualizado AS a
				INNER JOIN products AS b ON a.Item = b.ID
		WHERE
				a.sucursal <> 'Ventas Comerciales'
		  AND a.Sucursal = '%s' AND Category != '' GROUP BY Category;
	`
	return fmt.Sprintf(consulta, sucursal)
}

func consultaCategorias(sucursal, departamento string) string {
	consulta := `
		SELECT
			Category_L2 AS Categoria
		FROM
			enx_rodelag.products_mview_instock_actualizado AS a
				INNER JOIN products AS b ON a.Item = b.ID
		WHERE
				a.sucursal <> 'Ventas Comerciales'
		  AND a.Sucursal = '%s' AND Category = '%s' AND Category_L2 != '' GROUP BY Category_L2;
	`
	return fmt.Sprintf(consulta, sucursal, departamento)
}

func consultaProductos(codigo, sucursal string) string {
	consulta := `
		SELECT
			a.Sucursal AS NombreSucursal,
			Category AS Departamento,
			Category_L2 AS Categoria,
			Item_number AS Codigo,
			Nombre_Venta AS NombreVenta,
			nombre AS Descripcion,
			Item_AvgCost_Calc(b.id, 'NACIONALIZADA', SYSDATE()) AS Costo,
			Get_Price(b.id, '') AS Precio,
			ifnull((select sum(unidades) from mviews_sales_details msd where msd.Sucursal =a.sucursal and codigo =b.id and date between date_add(date_add(LAST_DAY(now()),interval 1 DAY),interval -6 MONTH) and LAST_DAY(sysdate())),0) AS UnidadVenta,
			ifnull((select sum(venta) from mviews_sales_details msd where msd.Sucursal =a.sucursal and codigo =b.id and date between date_add(date_add(LAST_DAY(now()),interval 1 DAY),interval -6 MONTH) and LAST_DAY(sysdate())),0) AS VentaTotal,
			round((1-(Item_AvgCost_Calc(b.id, 'NACIONALIZADA', SYSDATE())/Get_Price(b.id, ''))) *100,2) AS Margen,
			Ifnull((select NewPrice from promotions p join promotions_products pp on p.id =pp.Promotion where CodeType =b.id and pricelist ='PRECIO REGULAR' and expira >CURDATE() order by expira desc limit 1)  ,0) AS Oferta,
			Ifnull((select Fecha_Inicio from promotions p join promotions_products pp on p.id =pp.Promotion where CodeType =b.id and pricelist ='PRECIO REGULAR' and expira >CURDATE() order by expira desc limit 1)  ,'') AS InicioOferta,
			Ifnull((select Expira from promotions p join promotions_products pp on p.id =pp.Promotion where CodeType =b.id and pricelist ='PRECIO REGULAR' and expira >CURDATE() order by expira desc limit 1)  ,'') AS FinOferta,
			a.instock AS Existencia,
			Get_InStock(b.id,'herminia') AS ExistenciaHerminia,
			IFNULL((select IFNULL((MAX(t1.Actualizado)),'') from purchases_inventory_items t1 join purchases t2 on t1.Purchase = t2.id where t1.Item = b.id and t2.Status <> 'ABORTED' and t2.Status <> 'IN-TRANSIT' and t2.Type ='PURCHASE-INV' and t2.Sucursal =a.sucursal), '') AS UltimoRecibo,
			IFNULL((select max(date) from mviews_sales_details msd where msd.Sucursal =a.sucursal and msd.Codigo =b.ID), '') AS UltimaVenta,
			(select IFNULL(count(distinct t2.id),'') from purchases_inventory_items t1 join purchases t2 on t1.Purchase = t2.id where t2.Sucursal =a.sucursal and t1.Item = b.id and t2.Status <> 'ABORTED' and t2.Status in('IN-TRANSIT','PURCHASED')) AS CantidadOrdenes,
			(select IFNULL(sum(t1.units),'') from purchases_inventory_items t1 join purchases t2 on t1.Purchase = t2.id where t2.Sucursal =a.sucursal and t1.Item = b.id and t2.Status <> 'ABORTED' and t2.Status in('IN-TRANSIT','PURCHASED')) AS CantidadArticulosOrden,
			(select IFNULL(MAX(t1.Actualizado),'') from purchases_inventory_items t1 join purchases t2 on t1.Purchase = t2.id where t2.Sucursal =a.sucursal and t1.Item = b.id and t2.Status <> 'ABORTED' and t2.Status <> 'IN-TRANSIT') AS FechaUltimaOrden,
			(select empresa from providers p2 where id =b.Proveedor_Principal) AS Proveedor,
			b.Marca AS Marca
		FROM
			 enx_rodelag.products_mview_instock_actualizado AS a
				 INNER JOIN products AS b ON a.Item = b.ID
		WHERE
			a.sucursal <> 'Ventas Comerciales'
			AND Item_number = '%s' AND a.Sucursal = '%s' LIMIT 1;
	`
	return fmt.Sprintf(consulta, codigo, sucursal)
}
