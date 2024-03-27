package mocks

import (
	schematic "github.com/SchematicHQ/schematic-go"
	"github.com/SchematicHQ/schematic-go/api"
	gomock "go.uber.org/mock/gomock"
)

// Type containing all mock objects; for use in tests to set up expectations
type mockObjects struct {
	ctrl            *gomock.Controller
	AccountsAPI     *MockAccountsAPI
	CompaniesAPI    *MockCompaniesAPI
	EntitlementsAPI *MockEntitlementsAPI
	EventsAPI       *MockEventsAPI
	FeaturesAPI     *MockFeaturesAPI
	PlansAPI        *MockPlansAPI
}

// Create an API client instance with mock instances for all API services
// Return both the API client instance and the mock objects for use in tests
func NewMockAPIClient(ctrl *gomock.Controller) (*api.APIClient, *mockObjects) {
	// Set up mocks
	mocks := &mockObjects{
		ctrl:            ctrl,
		AccountsAPI:     NewMockAccountsAPI(ctrl),
		CompaniesAPI:    NewMockCompaniesAPI(ctrl),
		EntitlementsAPI: NewMockEntitlementsAPI(ctrl),
		EventsAPI:       NewMockEventsAPI(ctrl),
		FeaturesAPI:     NewMockFeaturesAPI(ctrl),
		PlansAPI:        NewMockPlansAPI(ctrl),
	}

	// Set up API client
	apiClient := api.NewAPIClient(api.NewConfiguration())
	apiClient.AccountsAPI = mocks.AccountsAPI
	apiClient.CompaniesAPI = mocks.CompaniesAPI
	apiClient.EntitlementsAPI = mocks.EntitlementsAPI
	apiClient.EventsAPI = mocks.EventsAPI
	apiClient.FeaturesAPI = mocks.FeaturesAPI
	apiClient.PlansAPI = mocks.PlansAPI

	// Provide API client instance and mocks
	return apiClient, mocks
}

// Create a schematic client instance with a mock API client
func NewMockClientWithAPI(ctrl *gomock.Controller) (schematic.Client, *mockObjects) {
	client := NewMockClient(ctrl)

	mockAPI, mocks := NewMockAPIClient(ctrl)

	client.EXPECT().API().Return(mockAPI).AnyTimes()

	return client, mocks
}
