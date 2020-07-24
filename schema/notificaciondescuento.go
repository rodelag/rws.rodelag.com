package schema

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	"rws/resolvers"
	"rws/types"
)

func NotificacionDescuentoQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"notificaciondescuento_listar": {
			Type: graphql.NewList(types.NotificacionDescuentoType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarNotificacionDescuento(), nil
			},
		},
	}
	return schemas
}

func NotificacionDescuentoMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"notificaciondescuento_crear": &graphql.Field{
			Type:        types.NotificacionDescuentoType,
			Description: "Creación de notificación de descuento",
			Args: graphql.FieldConfigArgument{
				"nombre": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre del cliente",
				},
				"apellido": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Apellido del cliente",
				},
				"sucursal": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Sucursal",
				},
				"fecha": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Fecha",
				},
				"nombreProducto": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre del producto",
				},
				"fotoPrecio": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Foto del precio",
				},
				"codigoParte": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Código de parte",
				},
				"cantidadVendidas": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Cantidad vendidas",
				},
				"nombreCompetidor": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre del competidor",
				},
				"precioCompetidor": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Precio del competidor",
				},
				"precioRodelag": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Precio Rodelag",
				},
				"fotoCotizacion": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Foto de la cotización",
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(params.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				nombre, _ := params.Args["nombre"].(string)
				apellido, _ := params.Args["apellido"].(string)
				sucursal, _ := params.Args["sucursal"].(string)
				fecha, _ := params.Args["fecha"].(string)
				nombreProducto, _ := params.Args["nombreProducto"].(string)
				fotoPrecio, _ := params.Args["fotoPrecio"].(string)
				codigoParte, _ := params.Args["codigoParte"].(string)
				cantidadVendidas, _ := params.Args["cantidadVendidas"].(string)
				nombreCompetidor, _ := params.Args["nombreCompetidor"].(string)
				precioCompetidor, _ := params.Args["precioCompetidor"].(string)
				precioRodelag, _ := params.Args["precioRodelag"].(string)
				fotoCotizacion, _ := params.Args["fotoCotizacion"].(string)

				return resolvers.CrearNotificacionDescuento(
					nombre,
					apellido,
					sucursal,
					fecha,
					nombreProducto,
					fotoPrecio,
					codigoParte,
					cantidadVendidas,
					nombreCompetidor,
					precioCompetidor,
					precioRodelag,
					fotoCotizacion,
				), nil
			},
		},
	}
	return schemas
}
