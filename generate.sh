#!/bin/bash

# Run the MockGen commands
#mockgen -destination=mocks/mocks.go -package=mocks github.com/schematichq/schematic-go/client SchematicClient # Accounts,WebhooksAPI,CrmAPI,BillingAPI,CompaniesAPI,EntitlementsAPI,EventsAPI,FeaturesAPI,PlansAPI,PlangroupsAPI
# mockgen -source=client/schematic.go -destination=mocks/client_mock.go -package=mocks
mockgen -source=core/core.go -destination=mocks/http_mock.go -package=mocks HTTPClient