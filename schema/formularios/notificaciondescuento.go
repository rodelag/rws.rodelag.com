package formularios

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/formularios"
	types "rws/types/formularios"
)

func NotificacionDescuentoQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"notificaciondescuento_ver": {
			Type:        types.NotificacionDescuentoType,
			Description: "Ver el de detalle de la solicitud de notificación de descuento",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "ID del registro",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				if id, ok := p.Args["id"].(int); ok {
					return resolvers.VerNotificacionDescuento(id), nil
				}
				return nil, nil
			},
		},
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
		"notificaciondescuento_crear_comentario": &graphql.Field{
			Type:        types.NotificacionDescuentoComentarioType,
			Description: "Creación de comentario de la notificación de descuento",
			Args: graphql.FieldConfigArgument{
				"estado": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Estado del registro",
				},
				"comentario": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Comentario del agente para con el registro",
				},
				"formulario": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Formulario al que pertenece el estado",
				},
				"usuario": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Usuario que gestiona el registro",
				},
				"correoUsuario": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Correo del usuario que gestiona el registro",
				},
				"idFormulario": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "ID del registro",
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
				estado, _ := params.Args["estado"].(string)
				comentario, _ := params.Args["comentario"].(string)
				formulario, _ := params.Args["formulario"].(string)
				usuario, _ := params.Args["usuario"].(string)
				correoUsuario, _ := params.Args["correoUsuario"].(string)
				idFormulario, _ := params.Args["idFormulario"].(int)

				return resolvers.CrearComentarioNotificacionDescuento(estado, comentario, formulario, usuario, correoUsuario, idFormulario), nil
			},
		},
	}
	return schemas
}
