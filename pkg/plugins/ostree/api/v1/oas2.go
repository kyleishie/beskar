// Code generated by kun; DO NOT EDIT.
// github.com/RussellLuo/kun

package apiv1

import (
	"reflect"

	"github.com/RussellLuo/kun/pkg/oas2"
)

var (
	base = `swagger: "2.0"
info:
  title: "OSTree Repository Management API"
  version: "1.0.0"
  description: "OSTree is used for managing ostree repositories.\nThis is the API documentation of OSTree.\n//"
  license:
    name: "MIT"
host: "example.com"
basePath: "/artifacts/ostree/api/v1"
schemes:
  - "https"
consumes:
  - "application/json"
produces:
  - "application/json"
`

	paths = `
paths:
  /repository/remote:
    post:
      description: "Add a new remote to the OSTree repository."
      operationId: "AddRemote"
      tags:
        - ostree
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/AddRemoteRequestBody"
      %s
  /repository:
    post:
      description: "Create an OSTree repository."
      operationId: "CreateRepository"
      tags:
        - ostree
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/CreateRepositoryRequestBody"
      %s
    delete:
      description: "Delete a OSTree repository."
      operationId: "DeleteRepository"
      tags:
        - ostree
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/DeleteRepositoryRequestBody"
      %s
  /repository/sync:
    get:
      description: "Get OSTree repository sync status."
      operationId: "GetRepositorySyncStatus"
      tags:
        - ostree
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/GetRepositorySyncStatusRequestBody"
      %s
    post:
      description: "Sync an ostree repository with one of the configured remotes."
      operationId: "SyncRepository"
      tags:
        - ostree
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/SyncRepositoryRequestBody"
      %s
`
)

func getResponses(schema oas2.Schema) []oas2.OASResponses {
	return []oas2.OASResponses{
		oas2.GetOASResponses(schema, "AddRemote", 200, &AddRemoteResponse{}),
		oas2.GetOASResponses(schema, "CreateRepository", 200, &CreateRepositoryResponse{}),
		oas2.GetOASResponses(schema, "DeleteRepository", 202, &DeleteRepositoryResponse{}),
		oas2.GetOASResponses(schema, "GetRepositorySyncStatus", 200, &GetRepositorySyncStatusResponse{}),
		oas2.GetOASResponses(schema, "SyncRepository", 202, &SyncRepositoryResponse{}),
	}
}

func getDefinitions(schema oas2.Schema) map[string]oas2.Definition {
	defs := make(map[string]oas2.Definition)

	oas2.AddDefinition(defs, "AddRemoteRequestBody", reflect.ValueOf(&struct {
		Repository string                  `json:"repository"`
		Properties *OSTreeRemoteProperties `json:"properties"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "AddRemote", 200, (&AddRemoteResponse{}).Body())

	oas2.AddDefinition(defs, "CreateRepositoryRequestBody", reflect.ValueOf(&struct {
		Repository string                      `json:"repository"`
		Properties *OSTreeRepositoryProperties `json:"properties"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "CreateRepository", 200, (&CreateRepositoryResponse{}).Body())

	oas2.AddDefinition(defs, "DeleteRepositoryRequestBody", reflect.ValueOf(&struct {
		Repository string `json:"repository"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "DeleteRepository", 202, (&DeleteRepositoryResponse{}).Body())

	oas2.AddDefinition(defs, "GetRepositorySyncStatusRequestBody", reflect.ValueOf(&struct {
		Repository string `json:"repository"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "GetRepositorySyncStatus", 200, (&GetRepositorySyncStatusResponse{}).Body())

	oas2.AddDefinition(defs, "SyncRepositoryRequestBody", reflect.ValueOf(&struct {
		Repository string                       `json:"repository"`
		Properties *OSTreeRepositorySyncRequest `json:"properties"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "SyncRepository", 202, (&SyncRepositoryResponse{}).Body())

	return defs
}

func OASv2APIDoc(schema oas2.Schema) string {
	resps := getResponses(schema)
	paths := oas2.GenPaths(resps, paths)

	defs := getDefinitions(schema)
	definitions := oas2.GenDefinitions(defs)

	return base + paths + definitions
}