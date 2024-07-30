#!/bin/bash

# Clean existing generated code
git rm -rf api/
git rm mocks/api_mocks.go
git rm mocks/client_mock.go

# Run the OpenAPI generator
GO_POST_PROCESS_FILE="gofmt -w" openapi-generator generate -c config.yaml

# Run the MockGen commands
mockgen -destination=mocks/api_mocks.go -package=mocks github.com/SchematicHQ/schematic-go/api HTTPClient,AccountsAPI,WebhooksAPI,CrmAPI,BillingAPI,CompaniesAPI,EntitlementsAPI,EventsAPI,FeaturesAPI,PlansAPI
mockgen -source=schematic.go -destination=mocks/client_mock.go -package=mocks

# Add the generated code
git add api/
git add mocks/

# Unstage non-delete changes so that we can easily see what changed
git diff --cached --name-only --diff-filter=MA | xargs -r git restore --staged --
