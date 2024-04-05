package mocks

//go:generate mockgen -destination=api_mocks.go -package=mocks github.com/SchematicHQ/schematic-go/api HTTPClient,AccountsAPI,CompaniesAPI,EntitlementsAPI,EventsAPI,FeaturesAPI,PlansAPI
//go:generate mockgen -source=../schematic.go -destination=client_mock.go -package=mocks
