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
  title: "Yum Repository Management API"
  version: "1.0.0"
  description: "YUM is used for managing YUM repositories.\nThis is the API documentation of YUM.\n//"
  license:
    name: "MIT"
host: "example.com"
basePath: "/artifacts/yum/api/v1"
schemes:
  - "https"
consumes:
  - "application/json"
produces:
  - "application/json"
`

	paths = `
paths:
  /repository:
    post:
      description: "Create a YUM repository."
      operationId: "CreateRepository"
      tags:
        - yum
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/CreateRepositoryRequestBody"
      %s
    delete:
      description: "Delete a YUM repository."
      operationId: "DeleteRepository"
      tags:
        - yum
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/DeleteRepositoryRequestBody"
      %s
    get:
      description: "Get YUM repository properties."
      operationId: "GetRepository"
      tags:
        - yum
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/GetRepositoryRequestBody"
      %s
    put:
      description: "Update YUM repository properties."
      operationId: "UpdateRepository"
      tags:
        - yum
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/UpdateRepositoryRequestBody"
      %s
  /repository/package:
    get:
      description: "Get RPM package from YUM repository."
      operationId: "GetRepositoryPackage"
      tags:
        - yum
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/GetRepositoryPackageRequestBody"
      %s
    delete:
      description: "Remove RPM package from YUM repository."
      operationId: "RemoveRepositoryPackage"
      tags:
        - yum
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/RemoveRepositoryPackageRequestBody"
      %s
  /repository/sync:status:
    get:
      description: "Get YUM repository sync status."
      operationId: "GetRepositorySyncStatus"
      tags:
        - yum
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/GetRepositorySyncStatusRequestBody"
      %s
  /repository/logs:
    get:
      description: "List YUM repository logs."
      operationId: "ListRepositoryLogs"
      tags:
        - yum
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/ListRepositoryLogsRequestBody"
      %s
  /repository/package:list:
    get:
      description: "List RPM packages for a YUM repository."
      operationId: "ListRepositoryPackages"
      tags:
        - yum
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/ListRepositoryPackagesRequestBody"
      %s
  /repository/sync:
    get:
      description: "Sync YUM repository with an upstream repository."
      operationId: "SyncRepository"
      tags:
        - yum
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
		oas2.GetOASResponses(schema, "CreateRepository", 200, &CreateRepositoryResponse{}),
		oas2.GetOASResponses(schema, "DeleteRepository", 200, &DeleteRepositoryResponse{}),
		oas2.GetOASResponses(schema, "GetRepository", 200, &GetRepositoryResponse{}),
		oas2.GetOASResponses(schema, "UpdateRepository", 200, &UpdateRepositoryResponse{}),
		oas2.GetOASResponses(schema, "GetRepositoryPackage", 200, &GetRepositoryPackageResponse{}),
		oas2.GetOASResponses(schema, "RemoveRepositoryPackage", 200, &RemoveRepositoryPackageResponse{}),
		oas2.GetOASResponses(schema, "GetRepositorySyncStatus", 200, &GetRepositorySyncStatusResponse{}),
		oas2.GetOASResponses(schema, "ListRepositoryLogs", 200, &ListRepositoryLogsResponse{}),
		oas2.GetOASResponses(schema, "ListRepositoryPackages", 200, &ListRepositoryPackagesResponse{}),
		oas2.GetOASResponses(schema, "SyncRepository", 200, &SyncRepositoryResponse{}),
	}
}

func getDefinitions(schema oas2.Schema) map[string]oas2.Definition {
	defs := make(map[string]oas2.Definition)

	oas2.AddDefinition(defs, "CreateRepositoryRequestBody", reflect.ValueOf(&struct {
		Repository string                `json:"repository"`
		Properties *RepositoryProperties `json:"properties"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "CreateRepository", 200, (&CreateRepositoryResponse{}).Body())

	oas2.AddDefinition(defs, "DeleteRepositoryRequestBody", reflect.ValueOf(&struct {
		Repository string `json:"repository"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "DeleteRepository", 200, (&DeleteRepositoryResponse{}).Body())

	oas2.AddDefinition(defs, "GetRepositoryRequestBody", reflect.ValueOf(&struct {
		Repository string `json:"repository"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "GetRepository", 200, (&GetRepositoryResponse{}).Body())

	oas2.AddDefinition(defs, "GetRepositoryPackageRequestBody", reflect.ValueOf(&struct {
		Repository string `json:"repository"`
		Id         string `json:"id"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "GetRepositoryPackage", 200, (&GetRepositoryPackageResponse{}).Body())

	oas2.AddDefinition(defs, "GetRepositorySyncStatusRequestBody", reflect.ValueOf(&struct {
		Repository string `json:"repository"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "GetRepositorySyncStatus", 200, (&GetRepositorySyncStatusResponse{}).Body())

	oas2.AddDefinition(defs, "ListRepositoryLogsRequestBody", reflect.ValueOf(&struct {
		Repository string `json:"repository"`
		Page       *Page  `json:"page"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "ListRepositoryLogs", 200, (&ListRepositoryLogsResponse{}).Body())

	oas2.AddDefinition(defs, "ListRepositoryPackagesRequestBody", reflect.ValueOf(&struct {
		Repository string `json:"repository"`
		Page       *Page  `json:"page"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "ListRepositoryPackages", 200, (&ListRepositoryPackagesResponse{}).Body())

	oas2.AddDefinition(defs, "RemoveRepositoryPackageRequestBody", reflect.ValueOf(&struct {
		Repository string `json:"repository"`
		Id         string `json:"id"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "RemoveRepositoryPackage", 200, (&RemoveRepositoryPackageResponse{}).Body())

	oas2.AddDefinition(defs, "SyncRepositoryRequestBody", reflect.ValueOf(&struct {
		Repository string `json:"repository"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "SyncRepository", 200, (&SyncRepositoryResponse{}).Body())

	oas2.AddDefinition(defs, "UpdateRepositoryRequestBody", reflect.ValueOf(&struct {
		Repository string                `json:"repository"`
		Properties *RepositoryProperties `json:"properties"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "UpdateRepository", 200, (&UpdateRepositoryResponse{}).Body())

	return defs
}

func OASv2APIDoc(schema oas2.Schema) string {
	resps := getResponses(schema)
	paths := oas2.GenPaths(resps, paths)

	defs := getDefinitions(schema)
	definitions := oas2.GenDefinitions(defs)

	return base + paths + definitions
}