package schema

import (
	"github.com/graphql-go/graphql"
)

var Schema graphql.Schema

func init() {
	Query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			//Formulario de Pago por ACH
			"pagoach_ver":    PagoACHQuery()["pagoach_ver"],
			"pagoach_listar": PagoACHQuery()["pagoach_listar"],
			//Formulario de Solicitud de Estado de Cuenta
			"estadocuenta_listar": EstadoCuentaQuery()["estadocuenta_listar"],
			//Formulario de Comprobante de Pago
			"comprobantepago_listar": ComprobantePagoQuery()["comprobantepago_listar"],
			//Formulario de Solicitud de Aplazamiento de Pagos por la Crisis del COVID-19
			"solicitudaplazamientopago_listar": SolicitudAplazamientoPagoQuery()["solicitudaplazamientopago_listar"],
			//Formulario de Prevención de Fraude
			"prevencionfraude_listar": PrevencionFraudeQuery()["prevencionfraude_listar"],
			//Formulario de Solicitud Tarjeta Rodelag
			"solicitudtarjetarodelag_listar": SolicitudTarjetaRodelagQuery()["solicitudtarjetarodelag_listar"],
			//Formulario de Reclamo
			"reclamo_listar": ReclamoQuery()["reclamo_listar"],
			//Formulario de Notificación Descuento
			"notificaciondescuento_listar": NotificacionDescuentoQuery()["notificaciondescuento_listar"],
			//Formulario de Contáctenos
			"contactenos_listar": ContactenosQuery()["contactenos_listar"],
			//Formulario de Contáctenos de ventas comerciales
			"contactenosvc_listar": ContactenosVCQuery()["contactenosvc_listar"],
			//Encuesta de Satisfacción Instalación AA
			"esiaa_listar": EsiaaQuery()["esiaa_listar"],
		},
	})

	Mutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			//Formulario de Pago por ACH
			"pagoach_crear":            PagoACHMutation()["pagoach_crear"],
			"pagoach_crear_comentario": PagoACHMutation()["pagoach_crear_comentario"],
			//Formulario de Solicitud de Estado de Cuenta
			"estadocuenta_crear": EstadoCuentaMutation()["estadocuenta_crear"],
			//Formulario de Comprobante de Pago
			"comprobantepago_crear": ComprobantePagoMutation()["comprobantepago_crear"],
			//Formulario de Solicitud de Aplazamiento de Pagos por la Crisis del COVID-19
			"solicitudaplazamientopago_crear": SolicitudAplazamientoPagoMutation()["solicitudaplazamientopago_crear"],
			//Formulario de Prevención de Fraude
			"prevencionfraude_crear": PrevencionFraudeMutation()["prevencionfraude_crear"],
			//Formulario de Solicitud Tarjeta Rodelag
			"solicitudtarjetarodelag_crear": SolicitudTarjetaRodelagMutation()["solicitudtarjetarodelag_crear"],
			//Formulario de Reclamo
			"reclamo_crear": ReclamoMutation()["reclamo_crear"],
			//Formulario de Notificación Descuento
			"notificaciondescuento_crear": NotificacionDescuentoMutation()["notificaciondescuento_crear"],
			//Formulario de Contáctenos
			"contactenos_crear": ContactenosMutation()["contactenos_crear"],
			//Formulario de Contáctenos de ventas comerciales
			"contactenosvc_crear": ContactenosVCMutation()["contactenosvc_crear"],
			//Encuesta de Satisfacción Instalación AA
			"esiaa_crear": EsiaaMutation()["esiaa_crear"],
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    Query,
		Mutation: Mutation,
	})
	if err != nil {
		panic(err)
	}
	Schema = schema
}
