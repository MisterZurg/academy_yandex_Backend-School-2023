#!/bin/bash

# Generate Contracts
oapi-codegen -generate server -package api ./internal/api/openapi.json > ./internal/api/server_generated.go
# Generate Models
oapi-codegen -generate types -package api ./internal/api/openapi.json > ./internal/api/types_generated.go