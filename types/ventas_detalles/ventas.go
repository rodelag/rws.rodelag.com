package ventas_detalles

import "github.com/graphql-go/graphql"

var VentasType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Ventas",
	Description: "Detalle de las ventas",
	Fields: graphql.Fields{
		"registroFecha": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha de la venta.",
		},
		"registroSucursal": &graphql.Field{
			Type:        graphql.String,
			Description: "Sucursal de la venta.",
		},
		"registroFactura": &graphql.Field{
			Type:        graphql.String,
			Description: "Registro de la venta.",
		},
		"registroVendedor": &graphql.Field{
			Type:        graphql.String,
			Description: "Vendedor de la venta.",
		},
		"registroTipoVendedor": &graphql.Field{
			Type:        graphql.String,
			Description: "Tipo de vendedro de la venta.",
		},
		"registroEmpresa": &graphql.Field{
			Type:        graphql.String,
			Description: "Empresa que hizo la compra.",
		},
		"registroNombre": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre de la persona que hizo la compra.",
		},
		"registroApellido": &graphql.Field{
			Type:        graphql.String,
			Description: "Apellido de la persona que hizo la compra..",
		},
		"registroDepto": &graphql.Field{
			Type:        graphql.String,
			Description: "Departamento del producto.",
		},
		"registroCategory_L2": &graphql.Field{
			Type:        graphql.String,
			Description: "Categoria 1 del producto.",
		},
		"registroCategory_L3": &graphql.Field{
			Type:        graphql.String,
			Description: "Categoria 2 del producto.",
		},
		"registroCodigoID": &graphql.Field{
			Type:        graphql.String,
			Description: "Codigo ID del producto.",
		},
		"registroCodigo": &graphql.Field{
			Type:        graphql.String,
			Description: "Codigo del producto.",
		},
		"registroDescripcion": &graphql.Field{
			Type:        graphql.String,
			Description: "Descripcion corta del producto.",
		},
		"registroDescripLarga": &graphql.Field{
			Type:        graphql.String,
			Description: "Descripcion larga del producto.",
		},
		"registroMarca": &graphql.Field{
			Type:        graphql.String,
			Description: "Marca del producto.",
		},
		"registroParte": &graphql.Field{
			Type:        graphql.String,
			Description: "Numero de parte del producto.",
		},
		"registroUnidades": &graphql.Field{
			Type:        graphql.String,
			Description: "Unidades que tiene el producto.",
		},
		"registroCosto": &graphql.Field{
			Type:        graphql.String,
			Description: "Costo del producto.",
		},
		"registroVenta": &graphql.Field{
			Type:        graphql.String,
			Description: "Venta del prodcuto.",
		},
		"registroUtilidad": &graphql.Field{
			Type:        graphql.String,
			Description: "Utilidad del producto.",
		},
		"registroMargen": &graphql.Field{
			Type:        graphql.String,
			Description: "Margen dle producto.",
		},
		"registroProveedor": &graphql.Field{
			Type:        graphql.String,
			Description: "Proveedor del producto.",
		},
		"registroListaPrecio": &graphql.Field{
			Type:        graphql.String,
			Description: "Lista de precio del producto.",
		},
	},
})
