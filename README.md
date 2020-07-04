# README
API usando GraphQL, más detalles en: [GraphQL](https://graphql.org/) Libreria usada: [graphql-go](https://github.com/graphql-go/graphql)

Es necesario que junto a la aplicación esté al mismo nivel el archivo de configuración:
`config.yml`

Estructura que debe tener el archivo de configuración yaml:

```console
basedatos:
    mysql:
        server: tcp(localhost)
        user: user
        password: *****
        database: ovnicom_formularios
correo:
    host: smtp.gmail.com
    port: 587
    mail: web.rodelag@gmail.com
    password: *****
    to: murray.greer@rodelag.com
    asunto: Notificacion de alerta - rws.rodelag.com
frasesecreta: *****
puerto: 8080
```