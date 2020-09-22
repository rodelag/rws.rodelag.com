package main

import (
	"context"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"rws/schema"
	"rws/utils"
	"strings"
)

func main() {
	utils.Configuracion()
	http.HandleFunc("/", graphqlHandler)
	http.ListenAndServe(":"+viper.GetString("puerto"), nil)
}

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		type GraphQLPostBody struct {
			Query         string                 `json:"query"`
			Variables     map[string]interface{} `json:"variables"`
			OperationName string                 `json:"operationName"`
		}

		var graphQLPostBody GraphQLPostBody
		err = json.Unmarshal(body, &graphQLPostBody)
		if err != nil {
			panic(err)
		}

		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader != "" {
			authHeader := strings.Split(authorizationHeader, "Bearer ")
			if len(authHeader) == 2 {
				var jwtToken string

				if authHeader[1] != "" {
					jwtToken = authHeader[1]
				}

				result := graphql.Do(graphql.Params{
					Schema:         schema.Schema,
					RequestString:  graphQLPostBody.Query,
					VariableValues: graphQLPostBody.Variables,
					OperationName:  graphQLPostBody.OperationName,
					Context:        context.WithValue(context.Background(), "token", jwtToken),
				})
				json.NewEncoder(w).Encode(result)
			} else {
				json.NewEncoder(w).Encode("Token de autorización inválido")
			}
		} else {
			json.NewEncoder(w).Encode("Se requiere un encabezado de autorización")
		}
	}
}
