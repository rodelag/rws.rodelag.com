package schema

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	"rws/resolvers"
	"rws/types"
)

func SolicitudAplazamientoPagoQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"solicitudaplazamientopago_listar": {
			Type: graphql.NewList(types.SolicitudAplazamientoPagoType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarSolicitudAplazamientoPago(), nil
			},
		},
	}
	return schemas
}

func SolicitudAplazamientoPagoMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"solicitudaplazamientopago_crear": &graphql.Field{
			Type:        types.SolicitudAplazamientoPagoType,
			Description: "Creación de Solicitud de Aplazamiento de Pagos por la Crisis del COVID-19",
			Args: graphql.FieldConfigArgument{
				"nombre": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre del cliente",
				},
				"apellido": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Apellido del cliente",
				},
				"correo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Correo del cliente",
				},
				"telefonoCasa": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Teléfono de la casa del cliente",
				},
				"cedula": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Cédula del cliente",
				},
				"celular": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Celular del cliente",
				},
				"tipoProducto": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Tipo de producto",
				},
				"tipoCliente": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Tipo de cliente",
				},
				"tipoActividadEconomica": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Tipo de actividad del cliente",
				},
				"lugarTrabajo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Lugar de trabajo del cliente",
				},
				"motivoSolicitud": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Motivo de la solicitud",
				},
				"detalleMotivo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Detalle del motivo",
				},
				"talonario": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Talonario",
				},
				"cartaMotivo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Carta del motivo",
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
				correo, _ := params.Args["correo"].(string)
				telefonoCasa, _ := params.Args["telefonoCasa"].(string)
				cedula, _ := params.Args["cedula"].(string)
				celular, _ := params.Args["celular"].(string)
				tipoProducto, _ := params.Args["tipoProducto"].(string)
				tipoCliente, _ := params.Args["tipoCliente"].(string)
				tipoActividadEconomica, _ := params.Args["tipoActividadEconomica"].(string)
				lugarTrabajo, _ := params.Args["lugarTrabajo"].(string)
				motivoSolicitud, _ := params.Args["motivoSolicitud"].(string)
				detalleMotivo, _ := params.Args["detalleMotivo"].(string)
				talonario, _ := params.Args["talonario"].(string)
				cartaMotivo, _ := params.Args["cartaMotivo"].(string)

				return resolvers.CrearSolicitudAplazamientoPago(
					nombre,
					apellido,
					correo,
					telefonoCasa,
					cedula,
					celular,
					tipoProducto,
					tipoCliente,
					tipoActividadEconomica,
					lugarTrabajo,
					motivoSolicitud,
					detalleMotivo,
					talonario,
					cartaMotivo,
				), nil
			},
		},
	}
	return schemas
}