# README
API usando GraphQL, más detalles en: [GraphQL](https://graphql.org/) Libreria usada: [graphql-go](https://github.com/graphql-go/graphql)

Es necesario que junto a la aplicación esté al mismo nivel el archivo de configuración:
`configuracion.yml`

Estructura que debe tener el archivo de configuración yaml:

```console
basedatos:
    mysql:
        rodelag:
            server: tcp(mysql.rodelag.com)
            user: ***
            password: ***
            database: ***
        elconix:
            server: tcp(rodelag.app.enxcloud.com)
            user: ***
            password: ***
            database: ***
correo:
    host: email-smtp.us-east-2.amazonaws.com
    port: 587
    mail: no-responder-no-supervisado@portal.rodelag.com
    user: ***
    password: ***
    to: murray.greer@rodelag.com
    asunto: Notificacion de alerta - rws.rodelag.com
frasesecreta: ***
puerto: 8888
```