package formularios

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/formularios"
	types "rws/types/formularios"
)

func PrevencionFraudeQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"prevencionfraude_ver": {
			Type:        types.PrevencionFraudeType,
			Description: "Ver el de detalle de Prevención de fraude",
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
					return resolvers.VerPrevencionFraude(id), nil
				}
				return nil, nil
			},
		},
		"prevencionfraude_listar": {
			Type: graphql.NewList(types.PrevencionFraudeType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarPrevencionFraude(), nil
			},
		},
	}
	return schemas
}

func PrevencionFraudeMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"prevencionfraude_crear": &graphql.Field{
			Type:        types.PrevencionFraudeType,
			Description: "Creación Prevención de Fraude",
			Args: graphql.FieldConfigArgument{
				"nombre": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre del cliente",
				},
				"apellido": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Apellido del cliente",
				},
				"fechaNacimiento": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Fecha de nacimiento del cliente",
				},
				"lugarResidencia": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Lugar de residencia",
				},
				"celular": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Celular del cliente",
				},
				"fotoCedula": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Foto de la cédula o pasaporte",
				},
				"fotoTarjeta": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Foto frontal de la tarjeta",
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
				fechaNacimiento, _ := params.Args["fechaNacimiento"].(string)
				lugarResidencia, _ := params.Args["lugarResidencia"].(string)
				celular, _ := params.Args["celular"].(string)
				fotoCedula, _ := params.Args["fotoCedula"].(string)
				fotoTarjeta, _ := params.Args["fotoTarjeta"].(string)

				return resolvers.CrearPrevencionFraude(nombre, apellido, fechaNacimiento, lugarResidencia, celular, fotoCedula, fotoTarjeta), nil
			},
		},
		"prevencionfraude_crear_comentario": &graphql.Field{
			Type:        types.PrevencionFraudeComentarioType,
			Description: "Creación de comentario de Prevención de Fraude",
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

				return resolvers.CrearComentarioPrevencionFraude(estado, comentario, formulario, usuario, correoUsuario, idFormulario), nil
			},
		},
	}
	return schemas
}
