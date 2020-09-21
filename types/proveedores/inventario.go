package proveedores

import "github.com/graphql-go/graphql"

var ProveedorInventarioType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ProveedorInventario",
	Description: "Inventario de proveedores",
	Fields: graphql.Fields{
		"consecutivo": &graphql.Field{
			Type:        graphql.String,
			Description: "Consecutivo para cada registro",
		},
		"nomSuc": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre de la sucursal",
		},
		"depto": &graphql.Field{
			Type:        graphql.String,
			Description: "Departamento",
		},
		"categ": &graphql.Field{
			Type:        graphql.String,
			Description: "Categoría",
		},
		"parte": &graphql.Field{
			Type:        graphql.String,
			Description: "Número de parte",
		},
		"codigo": &graphql.Field{
			Type:        graphql.String,
			Description: "Código del producto",
		},
		"descripcion": &graphql.Field{
			Type:        graphql.String,
			Description: "Descripcion del producto",
		},
		"preReg": &graphql.Field{
			Type:        graphql.String,
			Description: "Precio regular del producto",
		},
		"oferta": &graphql.Field{
			Type:        graphql.String,
			Description: "Oferta del producto",
		},
		"exist": &graphql.Field{
			Type:        graphql.String,
			Description: "Existencia del producto",
		},
		"marca": &graphql.Field{
			Type:        graphql.String,
			Description: "Marca del producto",
		},
		"ofeIni": &graphql.Field{
			Type:        graphql.String,
			Description: "Oferta inicial del producto",
		},
		"ofeFin": &graphql.Field{
			Type:        graphql.String,
			Description: "Oferta final del producto",
		},
		"provIDElconix": &graphql.Field{
			Type:        graphql.String,
			Description: "ID del proveedor en ElConix",
		},
		"nomProv": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del proveedor",
		},
		"provID": &graphql.Field{
			Type:        graphql.String,
			Description: "ID del proveedor en RMS",
		},
		"categID": &graphql.Field{
			Type:        graphql.String,
			Description: "ID Categoría (Campo en desuso)",
		},
	},
})
