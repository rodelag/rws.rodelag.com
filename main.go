package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"rws/schema"
)

func configuracion() {
	viper.SetConfigName("configuracion")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func main() {
	configuracion()
	fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/graphql", graphqlHandler)
	http.Handle("/", fs)
	http.ListenAndServe(":"+viper.GetString("puerto"), nil)
}

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
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

		token := r.Header.Get("token")

		result := graphql.Do(graphql.Params{
			Schema:         schema.Schema,
			RequestString:  graphQLPostBody.Query,
			VariableValues: graphQLPostBody.Variables,
			OperationName:  graphQLPostBody.OperationName,
			Context:        context.WithValue(context.Background(), "token", token),
		})
		json.NewEncoder(w).Encode(result)

	default:
		fmt.Fprintf(w, "Solo está soportado el metodo POST.")
	}
}
