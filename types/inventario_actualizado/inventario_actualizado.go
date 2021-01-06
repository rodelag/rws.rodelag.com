package inventario_actualizado

import "github.com/graphql-go/graphql"

var InventarioActualizadoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "InventarioActualizado",
	Description: "Inventario Actualizado",
	Fields: graphql.Fields{
		"RegistroNombreSucursal": &graphql.Field{
			Type:        graphql.String,
			Description: "Sucursal",
		},
		"RegistroSucBodegas": &graphql.Field{
			Type:        graphql.String,
			Description: "Bodega",
		},
		"RegistroDepartamento": &graphql.Field{
			Type:        graphql.String,
			Description: "Departamento",
		},
		"RegistroCategoria": &graphql.Field{
			Type:        graphql.String,
			Description: "Categoria",
		},
		"RegistroCodigo": &graphql.Field{
			Type:        graphql.String,
			Description: "Codigo",
		},
		"RegistroNombreVenta": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre de venta",
		},
		"RegistroDescripcion": &graphql.Field{
			Type:        graphql.String,
			Description: "Descripcion",
		},
		"RegistroCosto": &graphql.Field{
			Type:        graphql.String,
			Description: "Costo",
		},
		"RegistroPrecio": &graphql.Field{
			Type:        graphql.String,
			Description: "Precio",
		},
		"RegistroUnidadVenta": &graphql.Field{
			Type:        graphql.String,
			Description: "Unidad de Venta",
		},
		"RegistroVentaTotal": &graphql.Field{
			Type:        graphql.String,
			Description: "Venta Total",
		},
		"RegistroMargen": &graphql.Field{
			Type:        graphql.String,
			Description: "Margen",
		},
		"RegistroOferta": &graphql.Field{
			Type:        graphql.String,
			Description: "Oferta",
		},
		"RegistroInicioOferta": &graphql.Field{
			Type:        graphql.String,
			Description: "Inicio de Oferta",
		},
		"RegistroFinOferta": &graphql.Field{
			Type:        graphql.String,
			Description: "Fin de Oferta",
		},
		"RegistroExistencia": &graphql.Field{
			Type:        graphql.String,
			Description: "Existencia",
		},
		"RegistroExistenciaHerminia": &graphql.Field{
			Type:        graphql.String,
			Description: "Existencia en Herminia",
		},
		"RegistroUltimoRecibo": &graphql.Field{
			Type:        graphql.String,
			Description: "Ultimo Recibo",
		},
		"RegistroUltimaVenta": &graphql.Field{
			Type:        graphql.String,
			Description: "Ultima Venta",
		},
		"RegistroCantidadOrdenes": &graphql.Field{
			Type:        graphql.String,
			Description: "Cantidad de Ordenes",
		},
		"RegistroCantidadArticulosOrden": &graphql.Field{
			Type:        graphql.String,
			Description: "Cantidad de Articulos en la Orden",
		},
		"RegistroFechaUltimaOrden": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha de la Ultima Orden",
		},
		"RegistroProveedor": &graphql.Field{
			Type:        graphql.String,
			Description: "Proveedor",
		},
		"RegistroMarca": &graphql.Field{
			Type:        graphql.String,
			Description: "Marca",
		},
	},
})

var InventarioProductosType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "InventarioProductos",
	Description: "Productos con Inventario Actualizado",
	Fields: graphql.Fields{
		"RegistroNombreSucursal": &graphql.Field{
			Type:        graphql.String,
			Description: "Sucursal",
		},
		"RegistroDepartamento": &graphql.Field{
			Type:        graphql.String,
			Description: "Departamento",
		},
		"RegistroCategoria": &graphql.Field{
			Type:        graphql.String,
			Description: "Categoria",
		},
		"RegistroCodigo": &graphql.Field{
			Type:        graphql.String,
			Description: "Codigo",
		},
		"RegistroNombreVenta": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre de venta",
		},
		"RegistroDescripcion": &graphql.Field{
			Type:        graphql.String,
			Description: "Descripcion",
		},
		"RegistroCosto": &graphql.Field{
			Type:        graphql.String,
			Description: "Costo",
		},
		"RegistroPrecio": &graphql.Field{
			Type:        graphql.String,
			Description: "Precio",
		},
		"RegistroUnidadVenta": &graphql.Field{
			Type:        graphql.String,
			Description: "Unidad de Venta",
		},
		"RegistroVentaTotal": &graphql.Field{
			Type:        graphql.String,
			Description: "Venta Total",
		},
		"RegistroMargen": &graphql.Field{
			Type:        graphql.String,
			Description: "Margen",
		},
		"RegistroOferta": &graphql.Field{
			Type:        graphql.String,
			Description: "Oferta",
		},
		"RegistroInicioOferta": &graphql.Field{
			Type:        graphql.String,
			Description: "Inicio de Oferta",
		},
		"RegistroFinOferta": &graphql.Field{
			Type:        graphql.String,
			Description: "Fin de Oferta",
		},
		"RegistroExistencia": &graphql.Field{
			Type:        graphql.String,
			Description: "Existencia",
		},
		"RegistroExistenciaHerminia": &graphql.Field{
			Type:        graphql.String,
			Description: "Existencia en Herminia",
		},
		"RegistroUltimoRecibo": &graphql.Field{
			Type:        graphql.String,
			Description: "Ultimo Recibo",
		},
		"RegistroUltimaVenta": &graphql.Field{
			Type:        graphql.String,
			Description: "Ultima Venta",
		},
		"RegistroCantidadOrdenes": &graphql.Field{
			Type:        graphql.String,
			Description: "Cantidad de Ordenes",
		},
		"RegistroCantidadArticulosOrden": &graphql.Field{
			Type:        graphql.String,
			Description: "Cantidad de Articulos en la Orden",
		},
		"RegistroFechaUltimaOrden": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha de la Ultima Orden",
		},
		"RegistroProveedor": &graphql.Field{
			Type:        graphql.String,
			Description: "Proveedor",
		},
		"RegistroMarca": &graphql.Field{
			Type:        graphql.String,
			Description: "Marca",
		},
	},
})

var InventarioTiendasType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "InventarioTiendas",
	Description: "Tiendas con inventario actualizado",
	Fields: graphql.Fields{
		"RegistroNombreSucursal": &graphql.Field{
			Type:        graphql.String,
			Description: "Sucursal",
		},
	},
})

var InventarioDepartamentosType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "InventarioDepartamentos",
	Description: "Departamentos con inventario actualizado",
	Fields: graphql.Fields{
		"RegistroDepartamento": &graphql.Field{
			Type:        graphql.String,
			Description: "Departamento",
		},
	},
})

var InventarioCategoriasType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "InventarioCategorias",
	Description: "Categorias con inventario actualizado",
	Fields: graphql.Fields{
		"RegistroCategoria": &graphql.Field{
			Type:        graphql.String,
			Description: "Categoria",
		},
	},
})
