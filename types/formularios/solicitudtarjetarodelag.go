package formularios

import "github.com/graphql-go/graphql"

var SolicitudTarjetaRodelagComentarioType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "SolicitudTarjetaRodelagComentarios",
	Description: "Comentario de registro para la solicitud de tarjeta Rodelag",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID del estado",
		},
		"estado": &graphql.Field{
			Type:        graphql.String,
			Description: "Estado del registro",
		},
		"comentario": &graphql.Field{
			Type:        graphql.String,
			Description: "Comentario del agente para con el registro",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha de creación del estado",
		},
		"formulario": &graphql.Field{
			Type:        graphql.String,
			Description: "Formulario al que pertenece el estado",
		},
		"usuario": &graphql.Field{
			Type:        graphql.String,
			Description: "Usuario que gestiona el registro",
		},
		"correoUsuario": &graphql.Field{
			Type:        graphql.String,
			Description: "Correo del usuario que gestiona el registro",
		},
		"idFormulario": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID del registro",
		},
	},
})

var SolicitudTarjetaRodelagType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "SolicitudTarjetaRodelag",
	Description: "Solicitud de Tarjeta Rodelag",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "id del registro",
		},
		"sucursal": &graphql.Field{
			Type:        graphql.String,
			Description: "Sucursal más cercana",
		},
		"nombre": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del cliente",
		},
		"apellido": &graphql.Field{
			Type:        graphql.String,
			Description: "Apellido del cliente",
		},
		"fechaNacimiento": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha de nacimiento",
		},
		"lugarNacimiento": &graphql.Field{
			Type:        graphql.String,
			Description: "Lugar de nacimiento",
		},
		"nacionalidad": &graphql.Field{
			Type:        graphql.String,
			Description: "Nacionalidad del cliente",
		},
		"cedula": &graphql.Field{
			Type:        graphql.String,
			Description: "Cédula del cliente",
		},
		"fotoCedula": &graphql.Field{
			Type:        graphql.String,
			Description: "Foto de la cédula",
		},
		"fotoFicha": &graphql.Field{
			Type:        graphql.String,
			Description: "Foto de la ficha del seguro social",
		},
		"fotoRecibo": &graphql.Field{
			Type:        graphql.String,
			Description: "Foto de recibo de servicios",
		},
		"estadoCivil": &graphql.Field{
			Type:        graphql.String,
			Description: "Estado civil",
		},
		"correo": &graphql.Field{
			Type:        graphql.String,
			Description: "Correo del cliente",
		},
		"direccionResidencia": &graphql.Field{
			Type:        graphql.String,
			Description: "",
		},
		"barrio": &graphql.Field{
			Type:        graphql.String,
			Description: "Barriada de residencia",
		},
		"provincia": &graphql.Field{
			Type:        graphql.String,
			Description: "Provincia de residencia",
		},
		"telefonoResidencia": &graphql.Field{
			Type:        graphql.String,
			Description: "Teléfono de la residencia",
		},
		"celular": &graphql.Field{
			Type:        graphql.String,
			Description: "Celular del cliente",
		},
		"credito": &graphql.Field{
			Type:        graphql.String,
			Description: "Monto de Línea de Crédito",
		},
		"educacion": &graphql.Field{
			Type:        graphql.String,
			Description: "Nivel educativo del cliente",
		},
		"nombreEmpresa": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre de la Empresa para la cual trabaja el cliente",
		},
		"tipoNegocio": &graphql.Field{
			Type:        graphql.String,
			Description: "Tipo de negocio en donde labora el cliente",
		},
		"cargo": &graphql.Field{
			Type:        graphql.String,
			Description: "Cargo o posición en la empresa donde trabaja el cliente",
		},
		"tiempoLaboral": &graphql.Field{
			Type:        graphql.String,
			Description: "Tiempo laborado en la empresa del cliente",
		},
		"direccionTrabajo": &graphql.Field{
			Type:        graphql.String,
			Description: "Dirección de la empresa donde labora el cliente",
		},
		"telefonoTrabajo": &graphql.Field{
			Type:        graphql.String,
			Description: "Teléfono de la empresa donde labora el cliente",
		},
		"extension": &graphql.Field{
			Type:        graphql.String,
			Description: "Extensión teléfono de la empresa donde labora el cliente",
		},
		"salarioMensual": &graphql.Field{
			Type:        graphql.String,
			Description: "Salario mensual del cliente",
		},
		"fuentesIngreso": &graphql.Field{
			Type:        graphql.String,
			Description: "Otras fuentes de ingreso del cliente",
		},
		"montoFuentesIngreso": &graphql.Field{
			Type:        graphql.String,
			Description: "Monto provenientes de otra fuente de ingreso",
		},
		"detalleFuentesIngreso": &graphql.Field{
			Type:        graphql.String,
			Description: "Detalle de la otra fuentes de ingreso",
		},
		"nombreReferenciaUno": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre de referencia número uno",
		},
		"telefonoReferenciaUno": &graphql.Field{
			Type:        graphql.String,
			Description: "Teléfono de referencia número uno",
		},
		"nombreReferenciaDos": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre de referencia número dos",
		},
		"telefonoReferenciaDos": &graphql.Field{
			Type:        graphql.String,
			Description: "Teléfono de referencia número dos",
		},
		"nombreReferenciaTres": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre de referencia número tres",
		},
		"telefonoReferenciaTres": &graphql.Field{
			Type:        graphql.String,
			Description: "Teléfono de referencia número tres",
		},
		"estado": &graphql.Field{
			Type:        graphql.String,
			Description: "Estado del registro",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha del registro",
		},
		"comentarios": &graphql.Field{
			Type:        graphql.NewList(SolicitudTarjetaRodelagComentarioType),
			Description: "Comentarios del registro",
		},
	},
})
