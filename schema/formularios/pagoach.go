package formularios

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/formularios"
	types "rws/types/formularios"
)

func PagoACHQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"pagoach_ver": {
			Type:        types.PagoACHType,
			Description: "Ver pago ACH",
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
					return resolvers.VerPagoACH(id), nil
				}
				return nil, nil
			},
		},
		"pagoach_listar": {
			Type:        graphql.NewList(types.PagoACHType),
			Description: "Listar pagos ACH",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarPagoACH(), nil
			},
		},
	}
	return schemas
}

func PagoACHMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"pagoach_crear": &graphql.Field{
			Type:        types.PagoACHType,
			Description: "Creación de pago ACH",
			Args: graphql.FieldConfigArgument{
				"nombre": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre del cliente",
				},
				"apellido": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Apellido del cliente",
				},
				"titularCuenta": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Titular de la cuenta",
				},
				"cedula": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Cédula del cliente",
				},
				"correo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Correo del cliente",
				},
				"telefono": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Teléfono del cliente",
				},
				"compraOrigen": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Lugar donde hizo la compra",
				},
				"numeroOrden": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Pedido o Cotización",
				},
				"fotoComprobante": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Foto del comprobante de pago",
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
				titularCuenta, _ := params.Args["titularCuenta"].(string)
				cedula, _ := params.Args["cedula"].(string)
				correo, _ := params.Args["correo"].(string)
				telefono, _ := params.Args["telefono"].(string)
				compraOrigen, _ := params.Args["compraOrigen"].(string)
				numeroOrden, _ := params.Args["numeroOrden"].(string)
				fotoComprobante, _ := params.Args["fotoComprobante"].(string)

				return resolvers.CrearPagoACH(nombre, apellido, titularCuenta, cedula, correo, telefono, compraOrigen, numeroOrden, fotoComprobante), nil
			},
		},
		"pagoach_crear_comentario": &graphql.Field{
			Type:        types.PagoACHComentarioType,
			Description: "Creación de los comentarios del pago por ACH",
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

				return resolvers.CrearComentarioPagoACH(estado, comentario, formulario, usuario, correoUsuario, idFormulario), nil
			},
		},
	}
	return schemas
}
