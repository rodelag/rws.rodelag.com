package proveedores

import "github.com/graphql-go/graphql"

var ProveedorVentasType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ProveedorVentas",
	Description: "Ventas de proveedores",
	Fields: graphql.Fields{
		"Consecutivo": &graphql.Field{
			Type:        graphql.String,
			Description: "Cursor para identificar cada registro, id único y secuencial",
		},
		"NombreSucursal": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre de la sucursal",
		},
		"Departamento": &graphql.Field{
			Type:        graphql.String,
			Description: "Departamento",
		},
		"Categoria": &graphql.Field{
			Type:        graphql.String,
			Description: "Categoría",
		},
		"Parte": &graphql.Field{
			Type:        graphql.String,
			Description: "Número de parte",
		},
		"Codigo": &graphql.Field{
			Type:        graphql.String,
			Description: "Código del producto",
		},
		"Descripcion": &graphql.Field{
			Type:        graphql.String,
			Description: "Descripcion del producto",
		},
		"PrecioRegular": &graphql.Field{
			Type:        graphql.String,
			Description: "Precio regular del producto",
		},
		"PrecioOferta": &graphql.Field{
			Type:        graphql.String,
			Description: "Oferta del producto",
		},
		"Existencia": &graphql.Field{
			Type:        graphql.String,
			Description: "Existencia del producto",
		},
		"Cantidad": &graphql.Field{
			Type:        graphql.String,
			Description: "Cantidad del producto",
		},
		"VentaTotal": &graphql.Field{
			Type:        graphql.String,
			Description: "Venta total del producto",
		},
		"Fecha": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha",
		},
		"Marca": &graphql.Field{
			Type:        graphql.String,
			Description: "Marca del producto",
		},
		"DescripcionLarga": &graphql.Field{
			Type:        graphql.String,
			Description: "Descripción del producto",
		},
		"ProveedorID": &graphql.Field{
			Type:        graphql.String,
			Description: "ID del proveedor en RMS",
		},
		"NombreProveedor": &graphql.Field{
			Type:        graphql.String,
			Description: "Nobre del proveedor",
		},
		"OfertaInicial": &graphql.Field{
			Type:        graphql.String,
			Description: "Oferta inicial",
		},
		"OfertaFinal": &graphql.Field{
			Type:        graphql.String,
			Description: "Oferta final",
		},
		"CategoriaID": &graphql.Field{
			Type:        graphql.String,
			Description: "ID Categoría (Campo en desuso)",
		},
	},
})
