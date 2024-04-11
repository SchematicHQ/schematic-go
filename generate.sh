#!/bin/bash

# Run the OpenAPI generator
GO_POST_PROCESS_FILE="gofmt -w" openapi-generator generate -c config.yaml

# Run the MockGen commands
mockgen -destination=mocks/api_mocks.go -package=mocks github.com/SchematicHQ/schematic-go/api HTTPClient,AccountsAPI,CompaniesAPI,EntitlementsAPI,EventsAPI,FeaturesAPI,PlansAPI
mockgen -source=schematic.go -destination=mocks/client_mock.go -package=mocks
