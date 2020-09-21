package formularios

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/formularios"
	types "rws/types/formularios"
)

func SolicitudAplazamientoPagoQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"solicitudaplazamientopago_ver": {
			Type:        types.SolicitudAplazamientoPagoType,
			Description: "Ver el de detalle de la solicitud de aplazamiento de pago",
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
					return resolvers.VerSolicitudAplazamientoPago(id), nil
				}
				return nil, nil
			},
		},
		"solicitudaplazamientopago_listar": {
			Type:        graphql.NewList(types.SolicitudAplazamientoPagoType),
			Description: "Listar las solicitudes de aplazamiento de pago",
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
		"solicitudaplazamientopago_modificar": &graphql.Field{
			Type:        types.SolicitudAplazamientoPagoType,
			Description: "Modificación de Solicitud de Aplazamiento de Pagos por la Crisis del COVID-19",
			Args: graphql.FieldConfigArgument{
				"gestion": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Manejo o Gestion del agente con el registro",
				},
				"estadoCuenta": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Estado de cuenta del cliente",
				},
				"acp": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "ACP del cliente",
				},
				"propuesta": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Propuesta que se aplica al cliente, gestionado por el agente",
				},
				"id": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "id del registro",
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

				if id, ok := params.Args["id"].(int); ok {
					gestion, _ := params.Args["gestion"].(string)
					estadoCuenta, _ := params.Args["estadoCuenta"].(string)
					acp, _ := params.Args["acp"].(string)
					propuesta, _ := params.Args["propuesta"].(string)

					return resolvers.ModificarSolicitudAplazamientoPago(
						gestion,
						estadoCuenta,
						acp,
						propuesta,
						id,
					), nil
				}
				return nil, nil
			},
		},
		"solicitudaplazamientopago_crear_comentario": &graphql.Field{
			Type:        types.SolicitudAplazamientoPagoComentarioType,
			Description: "Creación de comentario de la solicitud de aplazamiento de pago",
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

				return resolvers.CrearComentarioSolicitudAplazamientoPago(estado, comentario, formulario, usuario, correoUsuario, idFormulario), nil
			},
		},
	}
	return schemas
}
