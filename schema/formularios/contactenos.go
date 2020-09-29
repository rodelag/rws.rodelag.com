package formularios

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/formularios"
	types "rws/types/formularios"
)

func ContactenosQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"contactenos_ver": {
			Type:        types.ContactenosType,
			Description: "Ver el de detalle de ventas comerciales",
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
					return resolvers.VerContactenos(id), nil
				}
				return nil, nil
			},
		},
		"contactenos_listar": {
			Type: graphql.NewList(types.ContactenosType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarContactenos(), nil
			},
		},
	}
	return schemas
}

func ContactenosMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"contactenos_crear": &graphql.Field{
			Type:        types.ContactenosType,
			Description: "Contáctenos",
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
				"telefono": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Teléfono del cliente",
				},
				"mensaje": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Mensaje del cliente",
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
				telefono, _ := params.Args["telefono"].(string)
				mensaje, _ := params.Args["mensaje"].(string)

				return resolvers.CrearContactenos(nombre, apellido, correo, telefono, mensaje), nil
			},
		},
		"contactenos_crear_comentario": &graphql.Field{
			Type:        types.ContactenosComentarioType,
			Description: "Creación de comentario de contáctenos",
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

				return resolvers.CrearComentarioContactenos(estado, comentario, formulario, usuario, correoUsuario, idFormulario), nil
			},
		},
	}
	return schemas
}
