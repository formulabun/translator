/*
 * Translator service between a srb2kart server and json
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "go.formulabun.club/translator/server/go"
)

func main() {
  validateEnvironment();
	log.Printf("Server started")

	DefaultApiService := openapi.NewDefaultApiService(TARGET)
	DefaultApiController := openapi.NewDefaultApiController(DefaultApiService)

	router := openapi.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":5092", router))
}
