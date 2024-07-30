# schematic-go API

## API Endpoints

All URIs are relative to *https://api.schematichq.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccesstokensAPI* | [**IssueTemporaryAccessToken**](docs/AccesstokensAPI.md#issuetemporaryaccesstoken) | **Post** /temporary-access-tokens | Issue temporary access token
*AccountsAPI* | [**CountApiKeys**](docs/AccountsAPI.md#countapikeys) | **Get** /api-keys/count | Count api keys
*AccountsAPI* | [**CountApiRequests**](docs/AccountsAPI.md#countapirequests) | **Get** /api-requests/count | Count api requests
*AccountsAPI* | [**CreateApiKey**](docs/AccountsAPI.md#createapikey) | **Post** /api-keys | Create api key
*AccountsAPI* | [**CreateEnvironment**](docs/AccountsAPI.md#createenvironment) | **Post** /environments | Create environment
*AccountsAPI* | [**DeleteApiKey**](docs/AccountsAPI.md#deleteapikey) | **Delete** /api-keys/{api_key_id} | Delete api key
*AccountsAPI* | [**DeleteEnvironment**](docs/AccountsAPI.md#deleteenvironment) | **Delete** /environments/{environment_id} | Delete environment
*AccountsAPI* | [**GetApiKey**](docs/AccountsAPI.md#getapikey) | **Get** /api-keys/{api_key_id} | Get api key
*AccountsAPI* | [**GetApiRequest**](docs/AccountsAPI.md#getapirequest) | **Get** /api-requests/{api_request_id} | Get api request
*AccountsAPI* | [**GetEnvironment**](docs/AccountsAPI.md#getenvironment) | **Get** /environments/{environment_id} | Get environment
*AccountsAPI* | [**ListApiKeys**](docs/AccountsAPI.md#listapikeys) | **Get** /api-keys | List api keys
*AccountsAPI* | [**ListApiRequests**](docs/AccountsAPI.md#listapirequests) | **Get** /api-requests | List api requests
*AccountsAPI* | [**ListEnvironments**](docs/AccountsAPI.md#listenvironments) | **Get** /environments | List environments
*AccountsAPI* | [**UpdateApiKey**](docs/AccountsAPI.md#updateapikey) | **Put** /api-keys/{api_key_id} | Update api key
*AccountsAPI* | [**UpdateEnvironment**](docs/AccountsAPI.md#updateenvironment) | **Put** /environments/{environment_id} | Update environment
*BillingAPI* | [**CountCustomers**](docs/BillingAPI.md#countcustomers) | **Get** /billing/customers/count | Count customers
*BillingAPI* | [**ListBillingProducts**](docs/BillingAPI.md#listbillingproducts) | **Get** /billing/products | List billing products
*BillingAPI* | [**ListCustomers**](docs/BillingAPI.md#listcustomers) | **Get** /billing/customers | List customers
*BillingAPI* | [**ListProductPrices**](docs/BillingAPI.md#listproductprices) | **Get** /billing/product/prices | List product prices
*BillingAPI* | [**UpsertBillingCustomer**](docs/BillingAPI.md#upsertbillingcustomer) | **Post** /billing/customer/upsert | Upsert billing customer
*BillingAPI* | [**UpsertBillingPrice**](docs/BillingAPI.md#upsertbillingprice) | **Post** /billing/price/upsert | Upsert billing price
*BillingAPI* | [**UpsertBillingProduct**](docs/BillingAPI.md#upsertbillingproduct) | **Post** /billing/product/upsert | Upsert billing product
*BillingAPI* | [**UpsertBillingSubscription**](docs/BillingAPI.md#upsertbillingsubscription) | **Post** /billing/subscription/upsert | Upsert billing subscription
*CompaniesAPI* | [**CountCompanies**](docs/CompaniesAPI.md#countcompanies) | **Get** /companies/count | Count companies
*CompaniesAPI* | [**CountEntityKeyDefinitions**](docs/CompaniesAPI.md#countentitykeydefinitions) | **Get** /entity-key-definitions/count | Count entity key definitions
*CompaniesAPI* | [**CountEntityTraitDefinitions**](docs/CompaniesAPI.md#countentitytraitdefinitions) | **Get** /entity-trait-definitions/count | Count entity trait definitions
*CompaniesAPI* | [**CountUsers**](docs/CompaniesAPI.md#countusers) | **Get** /users/count | Count users
*CompaniesAPI* | [**CreateCompany**](docs/CompaniesAPI.md#createcompany) | **Post** /companies/create | Create company
*CompaniesAPI* | [**CreateUser**](docs/CompaniesAPI.md#createuser) | **Post** /users/create | Create user
*CompaniesAPI* | [**DeleteCompany**](docs/CompaniesAPI.md#deletecompany) | **Delete** /companies/{company_id} | Delete company
*CompaniesAPI* | [**DeleteCompanyByKeys**](docs/CompaniesAPI.md#deletecompanybykeys) | **Post** /companies/delete | Delete company by keys
*CompaniesAPI* | [**DeleteCompanyMembership**](docs/CompaniesAPI.md#deletecompanymembership) | **Delete** /company-memberships/{company_membership_id} | Delete company membership
*CompaniesAPI* | [**DeleteUser**](docs/CompaniesAPI.md#deleteuser) | **Delete** /users/{user_id} | Delete user
*CompaniesAPI* | [**DeleteUserByKeys**](docs/CompaniesAPI.md#deleteuserbykeys) | **Post** /users/delete | Delete user by keys
*CompaniesAPI* | [**GetActiveCompanySubscription**](docs/CompaniesAPI.md#getactivecompanysubscription) | **Get** /company-subscriptions | Get active company subscription
*CompaniesAPI* | [**GetActiveDeals**](docs/CompaniesAPI.md#getactivedeals) | **Get** /company-crm-deals | Get active deals
*CompaniesAPI* | [**GetCompany**](docs/CompaniesAPI.md#getcompany) | **Get** /companies/{company_id} | Get company
*CompaniesAPI* | [**GetEntityTraitDefinition**](docs/CompaniesAPI.md#getentitytraitdefinition) | **Get** /entity-trait-definitions/{entity_trait_definition_id} | Get entity trait definition
*CompaniesAPI* | [**GetEntityTraitValues**](docs/CompaniesAPI.md#getentitytraitvalues) | **Get** /entity-trait-values | Get entity trait values
*CompaniesAPI* | [**GetOrCreateCompanyMembership**](docs/CompaniesAPI.md#getorcreatecompanymembership) | **Post** /company-memberships | Get or create company membership
*CompaniesAPI* | [**GetOrCreateEntityTraitDefinition**](docs/CompaniesAPI.md#getorcreateentitytraitdefinition) | **Post** /entity-trait-definitions | Get or create entity trait definition
*CompaniesAPI* | [**GetUser**](docs/CompaniesAPI.md#getuser) | **Get** /users/{user_id} | Get user
*CompaniesAPI* | [**ListCompanies**](docs/CompaniesAPI.md#listcompanies) | **Get** /companies | List companies
*CompaniesAPI* | [**ListCompanyMemberships**](docs/CompaniesAPI.md#listcompanymemberships) | **Get** /company-memberships | List company memberships
*CompaniesAPI* | [**ListCompanyPlans**](docs/CompaniesAPI.md#listcompanyplans) | **Get** /company-plans | List company plans
*CompaniesAPI* | [**ListEntityKeyDefinitions**](docs/CompaniesAPI.md#listentitykeydefinitions) | **Get** /entity-key-definitions | List entity key definitions
*CompaniesAPI* | [**ListEntityTraitDefinitions**](docs/CompaniesAPI.md#listentitytraitdefinitions) | **Get** /entity-trait-definitions | List entity trait definitions
*CompaniesAPI* | [**ListUsers**](docs/CompaniesAPI.md#listusers) | **Get** /users | List users
*CompaniesAPI* | [**LookupCompany**](docs/CompaniesAPI.md#lookupcompany) | **Get** /companies/lookup | Lookup company
*CompaniesAPI* | [**LookupUser**](docs/CompaniesAPI.md#lookupuser) | **Get** /users/lookup | Lookup user
*CompaniesAPI* | [**UpdateEntityTraitDefinition**](docs/CompaniesAPI.md#updateentitytraitdefinition) | **Put** /entity-trait-definitions/{entity_trait_definition_id} | Update entity trait definition
*CompaniesAPI* | [**UpsertCompany**](docs/CompaniesAPI.md#upsertcompany) | **Post** /companies | Upsert company
*CompaniesAPI* | [**UpsertCompanyTrait**](docs/CompaniesAPI.md#upsertcompanytrait) | **Post** /company-traits | Upsert company trait
*CompaniesAPI* | [**UpsertUser**](docs/CompaniesAPI.md#upsertuser) | **Post** /users | Upsert user
*CompaniesAPI* | [**UpsertUserTrait**](docs/CompaniesAPI.md#upsertusertrait) | **Post** /user-traits | Upsert user trait
*ComponentsAPI* | [**CountComponents**](docs/ComponentsAPI.md#countcomponents) | **Get** /components/count | Count components
*ComponentsAPI* | [**CreateComponent**](docs/ComponentsAPI.md#createcomponent) | **Post** /components | Create component
*ComponentsAPI* | [**DeleteComponent**](docs/ComponentsAPI.md#deletecomponent) | **Delete** /components/{component_id} | Delete component
*ComponentsAPI* | [**GetComponent**](docs/ComponentsAPI.md#getcomponent) | **Get** /components/{component_id} | Get component
*ComponentsAPI* | [**HydrateComponent**](docs/ComponentsAPI.md#hydratecomponent) | **Get** /components/{component_id}/hydrate | Hydrate component
*ComponentsAPI* | [**ListComponents**](docs/ComponentsAPI.md#listcomponents) | **Get** /components | List components
*ComponentsAPI* | [**UpdateComponent**](docs/ComponentsAPI.md#updatecomponent) | **Put** /components/{component_id} | Update component
*CrmAPI* | [**ListCrmProducts**](docs/CrmAPI.md#listcrmproducts) | **Get** /crm/products | List crm products
*CrmAPI* | [**UpsertCrmDeal**](docs/CrmAPI.md#upsertcrmdeal) | **Post** /crm/deals/upsert | Upsert crm deal
*CrmAPI* | [**UpsertCrmProduct**](docs/CrmAPI.md#upsertcrmproduct) | **Post** /crm/products/upsert | Upsert crm product
*CrmAPI* | [**UpsertDealLineItemAssociation**](docs/CrmAPI.md#upsertdeallineitemassociation) | **Post** /crm/associations/deal-line-item | Upsert deal line item association
*CrmAPI* | [**UpsertLineItem**](docs/CrmAPI.md#upsertlineitem) | **Post** /crm/deal-line-item/upsert | Upsert line item
*EntitlementsAPI* | [**CountCompanyOverrides**](docs/EntitlementsAPI.md#countcompanyoverrides) | **Get** /company-overrides/count | Count company overrides
*EntitlementsAPI* | [**CountFeatureCompanies**](docs/EntitlementsAPI.md#countfeaturecompanies) | **Get** /feature-companies/count | Count feature companies
*EntitlementsAPI* | [**CountFeatureUsage**](docs/EntitlementsAPI.md#countfeatureusage) | **Get** /feature-usage/count | Count feature usage
*EntitlementsAPI* | [**CountFeatureUsers**](docs/EntitlementsAPI.md#countfeatureusers) | **Get** /feature-users/count | Count feature users
*EntitlementsAPI* | [**CountPlanEntitlements**](docs/EntitlementsAPI.md#countplanentitlements) | **Get** /plan-entitlements/count | Count plan entitlements
*EntitlementsAPI* | [**CreateCompanyOverride**](docs/EntitlementsAPI.md#createcompanyoverride) | **Post** /company-overrides | Create company override
*EntitlementsAPI* | [**CreatePlanEntitlement**](docs/EntitlementsAPI.md#createplanentitlement) | **Post** /plan-entitlements | Create plan entitlement
*EntitlementsAPI* | [**DeleteCompanyOverride**](docs/EntitlementsAPI.md#deletecompanyoverride) | **Delete** /company-overrides/{company_override_id} | Delete company override
*EntitlementsAPI* | [**DeletePlanEntitlement**](docs/EntitlementsAPI.md#deleteplanentitlement) | **Delete** /plan-entitlements/{plan_entitlement_id} | Delete plan entitlement
*EntitlementsAPI* | [**GetCompanyOverride**](docs/EntitlementsAPI.md#getcompanyoverride) | **Get** /company-overrides/{company_override_id} | Get company override
*EntitlementsAPI* | [**GetFeatureUsageByCompany**](docs/EntitlementsAPI.md#getfeatureusagebycompany) | **Get** /usage-by-company | Get feature usage by company
*EntitlementsAPI* | [**GetPlanEntitlement**](docs/EntitlementsAPI.md#getplanentitlement) | **Get** /plan-entitlements/{plan_entitlement_id} | Get plan entitlement
*EntitlementsAPI* | [**ListCompanyOverrides**](docs/EntitlementsAPI.md#listcompanyoverrides) | **Get** /company-overrides | List company overrides
*EntitlementsAPI* | [**ListFeatureCompanies**](docs/EntitlementsAPI.md#listfeaturecompanies) | **Get** /feature-companies | List feature companies
*EntitlementsAPI* | [**ListFeatureUsage**](docs/EntitlementsAPI.md#listfeatureusage) | **Get** /feature-usage | List feature usage
*EntitlementsAPI* | [**ListFeatureUsers**](docs/EntitlementsAPI.md#listfeatureusers) | **Get** /feature-users | List feature users
*EntitlementsAPI* | [**ListPlanEntitlements**](docs/EntitlementsAPI.md#listplanentitlements) | **Get** /plan-entitlements | List plan entitlements
*EntitlementsAPI* | [**UpdateCompanyOverride**](docs/EntitlementsAPI.md#updatecompanyoverride) | **Put** /company-overrides/{company_override_id} | Update company override
*EntitlementsAPI* | [**UpdatePlanEntitlement**](docs/EntitlementsAPI.md#updateplanentitlement) | **Put** /plan-entitlements/{plan_entitlement_id} | Update plan entitlement
*EventsAPI* | [**CreateEvent**](docs/EventsAPI.md#createevent) | **Post** /events | Create event
*EventsAPI* | [**CreateEventBatch**](docs/EventsAPI.md#createeventbatch) | **Post** /event-batch | Create event batch
*EventsAPI* | [**GetEvent**](docs/EventsAPI.md#getevent) | **Get** /events/{event_id} | Get event
*EventsAPI* | [**GetEventSummaries**](docs/EventsAPI.md#geteventsummaries) | **Get** /event-types | Get event summaries
*EventsAPI* | [**GetEventSummaryBySubtype**](docs/EventsAPI.md#geteventsummarybysubtype) | **Get** /event-types/{key} | Get event summary by subtype
*EventsAPI* | [**GetSegmentIntegrationStatus**](docs/EventsAPI.md#getsegmentintegrationstatus) | **Get** /segment-integration | Get segment integration status
*EventsAPI* | [**ListEvents**](docs/EventsAPI.md#listevents) | **Get** /events | List events
*EventsAPI* | [**ListMetricCounts**](docs/EventsAPI.md#listmetriccounts) | **Get** /metric-counts | List metric counts
*FeaturesAPI* | [**CheckFlag**](docs/FeaturesAPI.md#checkflag) | **Post** /flags/{key}/check | Check flag
*FeaturesAPI* | [**CheckFlags**](docs/FeaturesAPI.md#checkflags) | **Post** /flags/check | Check flags
*FeaturesAPI* | [**CountAudienceCompanies**](docs/FeaturesAPI.md#countaudiencecompanies) | **Post** /audience/count-companies | Count audience companies
*FeaturesAPI* | [**CountAudienceUsers**](docs/FeaturesAPI.md#countaudienceusers) | **Post** /audience/count-users | Count audience users
*FeaturesAPI* | [**CountFeatures**](docs/FeaturesAPI.md#countfeatures) | **Get** /features/count | Count features
*FeaturesAPI* | [**CountFlags**](docs/FeaturesAPI.md#countflags) | **Get** /flags/count | Count flags
*FeaturesAPI* | [**CreateFeature**](docs/FeaturesAPI.md#createfeature) | **Post** /features | Create feature
*FeaturesAPI* | [**CreateFlag**](docs/FeaturesAPI.md#createflag) | **Post** /flags | Create flag
*FeaturesAPI* | [**DeleteFeature**](docs/FeaturesAPI.md#deletefeature) | **Delete** /features/{feature_id} | Delete feature
*FeaturesAPI* | [**DeleteFlag**](docs/FeaturesAPI.md#deleteflag) | **Delete** /flags/{flag_id} | Delete flag
*FeaturesAPI* | [**GetFeature**](docs/FeaturesAPI.md#getfeature) | **Get** /features/{feature_id} | Get feature
*FeaturesAPI* | [**GetFlag**](docs/FeaturesAPI.md#getflag) | **Get** /flags/{flag_id} | Get flag
*FeaturesAPI* | [**ListAudienceCompanies**](docs/FeaturesAPI.md#listaudiencecompanies) | **Post** /audience/get-companies | List audience companies
*FeaturesAPI* | [**ListAudienceUsers**](docs/FeaturesAPI.md#listaudienceusers) | **Post** /audience/get-users | List audience users
*FeaturesAPI* | [**ListFeatures**](docs/FeaturesAPI.md#listfeatures) | **Get** /features | List features
*FeaturesAPI* | [**ListFlags**](docs/FeaturesAPI.md#listflags) | **Get** /flags | List flags
*FeaturesAPI* | [**UpdateFeature**](docs/FeaturesAPI.md#updatefeature) | **Put** /features/{feature_id} | Update feature
*FeaturesAPI* | [**UpdateFlag**](docs/FeaturesAPI.md#updateflag) | **Put** /flags/{flag_id} | Update flag
*FeaturesAPI* | [**UpdateFlagRules**](docs/FeaturesAPI.md#updateflagrules) | **Put** /flags/{flag_id}/rules | Update flag rules
*PlansAPI* | [**CountPlans**](docs/PlansAPI.md#countplans) | **Get** /plans/count | Count plans
*PlansAPI* | [**CreatePlan**](docs/PlansAPI.md#createplan) | **Post** /plans | Create plan
*PlansAPI* | [**DeleteAudience**](docs/PlansAPI.md#deleteaudience) | **Delete** /plan-audiences/{plan_audience_id} | Delete audience
*PlansAPI* | [**DeletePlan**](docs/PlansAPI.md#deleteplan) | **Delete** /plans/{plan_id} | Delete plan
*PlansAPI* | [**GetAudience**](docs/PlansAPI.md#getaudience) | **Get** /plan-audiences/{plan_audience_id} | Get audience
*PlansAPI* | [**GetPlan**](docs/PlansAPI.md#getplan) | **Get** /plans/{plan_id} | Get plan
*PlansAPI* | [**ListPlans**](docs/PlansAPI.md#listplans) | **Get** /plans | List plans
*PlansAPI* | [**UpdateAudience**](docs/PlansAPI.md#updateaudience) | **Put** /plan-audiences/{plan_audience_id} | Update audience
*PlansAPI* | [**UpdatePlan**](docs/PlansAPI.md#updateplan) | **Put** /plans/{plan_id} | Update plan
*PlansAPI* | [**UpsertBillingProductPlan**](docs/PlansAPI.md#upsertbillingproductplan) | **Put** /plans/{plan_id}/billing_products | Upsert billing product plan
*WebhooksAPI* | [**CountWebhookEvents**](docs/WebhooksAPI.md#countwebhookevents) | **Get** /webhook-events/count | Count webhook events
*WebhooksAPI* | [**CountWebhooks**](docs/WebhooksAPI.md#countwebhooks) | **Get** /webhooks/count | Count webhooks
*WebhooksAPI* | [**CreateWebhook**](docs/WebhooksAPI.md#createwebhook) | **Post** /webhooks | Create webhook
*WebhooksAPI* | [**DeleteWebhook**](docs/WebhooksAPI.md#deletewebhook) | **Delete** /webhooks/{webhook_id} | Delete webhook
*WebhooksAPI* | [**GetWebhook**](docs/WebhooksAPI.md#getwebhook) | **Get** /webhooks/{webhook_id} | Get webhook
*WebhooksAPI* | [**GetWebhookEvent**](docs/WebhooksAPI.md#getwebhookevent) | **Get** /webhook-events/{webhook_event_id} | Get webhook event
*WebhooksAPI* | [**ListWebhookEvents**](docs/WebhooksAPI.md#listwebhookevents) | **Get** /webhook-events | List webhook events
*WebhooksAPI* | [**ListWebhooks**](docs/WebhooksAPI.md#listwebhooks) | **Get** /webhooks | List webhooks
*WebhooksAPI* | [**UpdateWebhook**](docs/WebhooksAPI.md#updatewebhook) | **Put** /webhooks/{webhook_id} | Update webhook


## Models

 - [ApiError](docs/ApiError.md)
 - [ApiKeyCreateResponseData](docs/ApiKeyCreateResponseData.md)
 - [ApiKeyRequestListResponseData](docs/ApiKeyRequestListResponseData.md)
 - [ApiKeyRequestResponseData](docs/ApiKeyRequestResponseData.md)
 - [ApiKeyResponseData](docs/ApiKeyResponseData.md)
 - [AudienceRequestBody](docs/AudienceRequestBody.md)
 - [BillingCustomerResponseData](docs/BillingCustomerResponseData.md)
 - [BillingCustomerSubscription](docs/BillingCustomerSubscription.md)
 - [BillingCustomerWithSubscriptionsResponseData](docs/BillingCustomerWithSubscriptionsResponseData.md)
 - [BillingPriceResponseData](docs/BillingPriceResponseData.md)
 - [BillingProductDetailResponseData](docs/BillingProductDetailResponseData.md)
 - [BillingProductPlanResponseData](docs/BillingProductPlanResponseData.md)
 - [BillingProductPricing](docs/BillingProductPricing.md)
 - [BillingProductResponseData](docs/BillingProductResponseData.md)
 - [BillingSubscriptionResponseData](docs/BillingSubscriptionResponseData.md)
 - [CheckFlagOutputWithFlagKey](docs/CheckFlagOutputWithFlagKey.md)
 - [CheckFlagRequestBody](docs/CheckFlagRequestBody.md)
 - [CheckFlagResponse](docs/CheckFlagResponse.md)
 - [CheckFlagResponseData](docs/CheckFlagResponseData.md)
 - [CheckFlagsResponse](docs/CheckFlagsResponse.md)
 - [CheckFlagsResponseData](docs/CheckFlagsResponseData.md)
 - [CompanyCrmDealsResponseData](docs/CompanyCrmDealsResponseData.md)
 - [CompanyDetailResponseData](docs/CompanyDetailResponseData.md)
 - [CompanyMembershipDetailResponseData](docs/CompanyMembershipDetailResponseData.md)
 - [CompanyMembershipResponseData](docs/CompanyMembershipResponseData.md)
 - [CompanyOverrideResponseData](docs/CompanyOverrideResponseData.md)
 - [CompanyPlanResponseData](docs/CompanyPlanResponseData.md)
 - [CompanyResponseData](docs/CompanyResponseData.md)
 - [CompanySubscriptionResponseData](docs/CompanySubscriptionResponseData.md)
 - [ComponentHydrateResponseData](docs/ComponentHydrateResponseData.md)
 - [ComponentResponseData](docs/ComponentResponseData.md)
 - [CountApiKeysParams](docs/CountApiKeysParams.md)
 - [CountApiKeysResponse](docs/CountApiKeysResponse.md)
 - [CountApiRequestsParams](docs/CountApiRequestsParams.md)
 - [CountApiRequestsResponse](docs/CountApiRequestsResponse.md)
 - [CountAudienceCompaniesResponse](docs/CountAudienceCompaniesResponse.md)
 - [CountAudienceUsersResponse](docs/CountAudienceUsersResponse.md)
 - [CountCompaniesParams](docs/CountCompaniesParams.md)
 - [CountCompaniesResponse](docs/CountCompaniesResponse.md)
 - [CountCompanyOverridesParams](docs/CountCompanyOverridesParams.md)
 - [CountCompanyOverridesResponse](docs/CountCompanyOverridesResponse.md)
 - [CountComponentsParams](docs/CountComponentsParams.md)
 - [CountComponentsResponse](docs/CountComponentsResponse.md)
 - [CountCustomersParams](docs/CountCustomersParams.md)
 - [CountCustomersResponse](docs/CountCustomersResponse.md)
 - [CountEntityKeyDefinitionsParams](docs/CountEntityKeyDefinitionsParams.md)
 - [CountEntityKeyDefinitionsResponse](docs/CountEntityKeyDefinitionsResponse.md)
 - [CountEntityTraitDefinitionsParams](docs/CountEntityTraitDefinitionsParams.md)
 - [CountEntityTraitDefinitionsResponse](docs/CountEntityTraitDefinitionsResponse.md)
 - [CountFeatureCompaniesParams](docs/CountFeatureCompaniesParams.md)
 - [CountFeatureCompaniesResponse](docs/CountFeatureCompaniesResponse.md)
 - [CountFeatureUsageParams](docs/CountFeatureUsageParams.md)
 - [CountFeatureUsageResponse](docs/CountFeatureUsageResponse.md)
 - [CountFeatureUsersParams](docs/CountFeatureUsersParams.md)
 - [CountFeatureUsersResponse](docs/CountFeatureUsersResponse.md)
 - [CountFeaturesParams](docs/CountFeaturesParams.md)
 - [CountFeaturesResponse](docs/CountFeaturesResponse.md)
 - [CountFlagsParams](docs/CountFlagsParams.md)
 - [CountFlagsResponse](docs/CountFlagsResponse.md)
 - [CountPlanEntitlementsParams](docs/CountPlanEntitlementsParams.md)
 - [CountPlanEntitlementsResponse](docs/CountPlanEntitlementsResponse.md)
 - [CountPlansParams](docs/CountPlansParams.md)
 - [CountPlansResponse](docs/CountPlansResponse.md)
 - [CountResponse](docs/CountResponse.md)
 - [CountUsersParams](docs/CountUsersParams.md)
 - [CountUsersResponse](docs/CountUsersResponse.md)
 - [CountWebhookEventsParams](docs/CountWebhookEventsParams.md)
 - [CountWebhookEventsResponse](docs/CountWebhookEventsResponse.md)
 - [CountWebhooksParams](docs/CountWebhooksParams.md)
 - [CountWebhooksResponse](docs/CountWebhooksResponse.md)
 - [CreateApiKeyRequestBody](docs/CreateApiKeyRequestBody.md)
 - [CreateApiKeyResponse](docs/CreateApiKeyResponse.md)
 - [CreateBillingCustomerRequestBody](docs/CreateBillingCustomerRequestBody.md)
 - [CreateBillingPriceRequestBody](docs/CreateBillingPriceRequestBody.md)
 - [CreateBillingProductRequestBody](docs/CreateBillingProductRequestBody.md)
 - [CreateBillingSubscriptionsRequestBody](docs/CreateBillingSubscriptionsRequestBody.md)
 - [CreateCompanyOverrideRequestBody](docs/CreateCompanyOverrideRequestBody.md)
 - [CreateCompanyOverrideResponse](docs/CreateCompanyOverrideResponse.md)
 - [CreateCompanyResponse](docs/CreateCompanyResponse.md)
 - [CreateComponentRequestBody](docs/CreateComponentRequestBody.md)
 - [CreateComponentResponse](docs/CreateComponentResponse.md)
 - [CreateCrmDealLineItemAssociationRequestBody](docs/CreateCrmDealLineItemAssociationRequestBody.md)
 - [CreateCrmDealRequestBody](docs/CreateCrmDealRequestBody.md)
 - [CreateCrmLineItemRequestBody](docs/CreateCrmLineItemRequestBody.md)
 - [CreateCrmProductRequestBody](docs/CreateCrmProductRequestBody.md)
 - [CreateEntityTraitDefinitionRequestBody](docs/CreateEntityTraitDefinitionRequestBody.md)
 - [CreateEnvironmentRequestBody](docs/CreateEnvironmentRequestBody.md)
 - [CreateEnvironmentResponse](docs/CreateEnvironmentResponse.md)
 - [CreateEventBatchRequestBody](docs/CreateEventBatchRequestBody.md)
 - [CreateEventBatchResponse](docs/CreateEventBatchResponse.md)
 - [CreateEventRequestBody](docs/CreateEventRequestBody.md)
 - [CreateEventResponse](docs/CreateEventResponse.md)
 - [CreateFeatureRequestBody](docs/CreateFeatureRequestBody.md)
 - [CreateFeatureResponse](docs/CreateFeatureResponse.md)
 - [CreateFlagRequestBody](docs/CreateFlagRequestBody.md)
 - [CreateFlagResponse](docs/CreateFlagResponse.md)
 - [CreateOrUpdateConditionGroupRequestBody](docs/CreateOrUpdateConditionGroupRequestBody.md)
 - [CreateOrUpdateConditionRequestBody](docs/CreateOrUpdateConditionRequestBody.md)
 - [CreateOrUpdateFlagRequestBody](docs/CreateOrUpdateFlagRequestBody.md)
 - [CreateOrUpdateRuleRequestBody](docs/CreateOrUpdateRuleRequestBody.md)
 - [CreatePlanEntitlementRequestBody](docs/CreatePlanEntitlementRequestBody.md)
 - [CreatePlanEntitlementResponse](docs/CreatePlanEntitlementResponse.md)
 - [CreatePlanRequestBody](docs/CreatePlanRequestBody.md)
 - [CreatePlanResponse](docs/CreatePlanResponse.md)
 - [CreateReqCommon](docs/CreateReqCommon.md)
 - [CreateUserResponse](docs/CreateUserResponse.md)
 - [CreateWebhookRequestBody](docs/CreateWebhookRequestBody.md)
 - [CreateWebhookResponse](docs/CreateWebhookResponse.md)
 - [CrmDealLineItem](docs/CrmDealLineItem.md)
 - [CrmDealResponseData](docs/CrmDealResponseData.md)
 - [CrmLineItemResponseData](docs/CrmLineItemResponseData.md)
 - [CrmProductResponseData](docs/CrmProductResponseData.md)
 - [DeleteApiKeyResponse](docs/DeleteApiKeyResponse.md)
 - [DeleteAudienceResponse](docs/DeleteAudienceResponse.md)
 - [DeleteCompanyByKeysResponse](docs/DeleteCompanyByKeysResponse.md)
 - [DeleteCompanyMembershipResponse](docs/DeleteCompanyMembershipResponse.md)
 - [DeleteCompanyOverrideResponse](docs/DeleteCompanyOverrideResponse.md)
 - [DeleteCompanyResponse](docs/DeleteCompanyResponse.md)
 - [DeleteComponentResponse](docs/DeleteComponentResponse.md)
 - [DeleteEnvironmentResponse](docs/DeleteEnvironmentResponse.md)
 - [DeleteFeatureResponse](docs/DeleteFeatureResponse.md)
 - [DeleteFlagResponse](docs/DeleteFlagResponse.md)
 - [DeletePlanEntitlementResponse](docs/DeletePlanEntitlementResponse.md)
 - [DeletePlanResponse](docs/DeletePlanResponse.md)
 - [DeleteResponse](docs/DeleteResponse.md)
 - [DeleteUserByKeysResponse](docs/DeleteUserByKeysResponse.md)
 - [DeleteUserResponse](docs/DeleteUserResponse.md)
 - [DeleteWebhookResponse](docs/DeleteWebhookResponse.md)
 - [EntityKeyDefinitionResponseData](docs/EntityKeyDefinitionResponseData.md)
 - [EntityKeyDetailResponseData](docs/EntityKeyDetailResponseData.md)
 - [EntityKeyResponseData](docs/EntityKeyResponseData.md)
 - [EntityTraitDefinitionResponseData](docs/EntityTraitDefinitionResponseData.md)
 - [EntityTraitDetailResponseData](docs/EntityTraitDetailResponseData.md)
 - [EntityTraitResponseData](docs/EntityTraitResponseData.md)
 - [EntityTraitValue](docs/EntityTraitValue.md)
 - [EnvironmentDetailResponseData](docs/EnvironmentDetailResponseData.md)
 - [EnvironmentResponseData](docs/EnvironmentResponseData.md)
 - [EventBody](docs/EventBody.md)
 - [EventBodyFlagCheck](docs/EventBodyFlagCheck.md)
 - [EventBodyIdentify](docs/EventBodyIdentify.md)
 - [EventBodyIdentifyCompany](docs/EventBodyIdentifyCompany.md)
 - [EventBodyTrack](docs/EventBodyTrack.md)
 - [EventDetailResponseData](docs/EventDetailResponseData.md)
 - [EventResponseData](docs/EventResponseData.md)
 - [EventSummaryResponseData](docs/EventSummaryResponseData.md)
 - [FeatureCompanyResponseData](docs/FeatureCompanyResponseData.md)
 - [FeatureCompanyUserResponseData](docs/FeatureCompanyUserResponseData.md)
 - [FeatureDetailResponseData](docs/FeatureDetailResponseData.md)
 - [FeatureResponseData](docs/FeatureResponseData.md)
 - [FeatureUsageDetailResponseData](docs/FeatureUsageDetailResponseData.md)
 - [FeatureUsageResponseData](docs/FeatureUsageResponseData.md)
 - [FlagDetailResponseData](docs/FlagDetailResponseData.md)
 - [FlagResponseData](docs/FlagResponseData.md)
 - [GetActiveCompanySubscriptionParams](docs/GetActiveCompanySubscriptionParams.md)
 - [GetActiveCompanySubscriptionResponse](docs/GetActiveCompanySubscriptionResponse.md)
 - [GetActiveDealsParams](docs/GetActiveDealsParams.md)
 - [GetActiveDealsResponse](docs/GetActiveDealsResponse.md)
 - [GetApiKeyResponse](docs/GetApiKeyResponse.md)
 - [GetApiRequestResponse](docs/GetApiRequestResponse.md)
 - [GetAudienceResponse](docs/GetAudienceResponse.md)
 - [GetCompanyOverrideResponse](docs/GetCompanyOverrideResponse.md)
 - [GetCompanyResponse](docs/GetCompanyResponse.md)
 - [GetComponentResponse](docs/GetComponentResponse.md)
 - [GetEntityTraitDefinitionResponse](docs/GetEntityTraitDefinitionResponse.md)
 - [GetEntityTraitValuesParams](docs/GetEntityTraitValuesParams.md)
 - [GetEntityTraitValuesResponse](docs/GetEntityTraitValuesResponse.md)
 - [GetEnvironmentResponse](docs/GetEnvironmentResponse.md)
 - [GetEventResponse](docs/GetEventResponse.md)
 - [GetEventSummariesParams](docs/GetEventSummariesParams.md)
 - [GetEventSummariesResponse](docs/GetEventSummariesResponse.md)
 - [GetEventSummaryBySubtypeResponse](docs/GetEventSummaryBySubtypeResponse.md)
 - [GetFeatureResponse](docs/GetFeatureResponse.md)
 - [GetFeatureUsageByCompanyParams](docs/GetFeatureUsageByCompanyParams.md)
 - [GetFeatureUsageByCompanyResponse](docs/GetFeatureUsageByCompanyResponse.md)
 - [GetFlagResponse](docs/GetFlagResponse.md)
 - [GetOrCreateCompanyMembershipRequestBody](docs/GetOrCreateCompanyMembershipRequestBody.md)
 - [GetOrCreateCompanyMembershipResponse](docs/GetOrCreateCompanyMembershipResponse.md)
 - [GetOrCreateEntityTraitDefinitionResponse](docs/GetOrCreateEntityTraitDefinitionResponse.md)
 - [GetPlanEntitlementResponse](docs/GetPlanEntitlementResponse.md)
 - [GetPlanResponse](docs/GetPlanResponse.md)
 - [GetSegmentIntegrationStatusResponse](docs/GetSegmentIntegrationStatusResponse.md)
 - [GetUserResponse](docs/GetUserResponse.md)
 - [GetWebhookEventResponse](docs/GetWebhookEventResponse.md)
 - [GetWebhookResponse](docs/GetWebhookResponse.md)
 - [HydrateComponentResponse](docs/HydrateComponentResponse.md)
 - [IssueTemporaryAccessTokenRequestBody](docs/IssueTemporaryAccessTokenRequestBody.md)
 - [IssueTemporaryAccessTokenResponse](docs/IssueTemporaryAccessTokenResponse.md)
 - [IssueTemporaryAccessTokenResponseData](docs/IssueTemporaryAccessTokenResponseData.md)
 - [KeysRequestBody](docs/KeysRequestBody.md)
 - [ListApiKeysParams](docs/ListApiKeysParams.md)
 - [ListApiKeysResponse](docs/ListApiKeysResponse.md)
 - [ListApiRequestsParams](docs/ListApiRequestsParams.md)
 - [ListApiRequestsResponse](docs/ListApiRequestsResponse.md)
 - [ListAudienceCompaniesResponse](docs/ListAudienceCompaniesResponse.md)
 - [ListAudienceUsersResponse](docs/ListAudienceUsersResponse.md)
 - [ListBillingProductsParams](docs/ListBillingProductsParams.md)
 - [ListBillingProductsResponse](docs/ListBillingProductsResponse.md)
 - [ListCompaniesParams](docs/ListCompaniesParams.md)
 - [ListCompaniesResponse](docs/ListCompaniesResponse.md)
 - [ListCompanyMembershipsParams](docs/ListCompanyMembershipsParams.md)
 - [ListCompanyMembershipsResponse](docs/ListCompanyMembershipsResponse.md)
 - [ListCompanyOverridesParams](docs/ListCompanyOverridesParams.md)
 - [ListCompanyOverridesResponse](docs/ListCompanyOverridesResponse.md)
 - [ListCompanyPlansParams](docs/ListCompanyPlansParams.md)
 - [ListCompanyPlansResponse](docs/ListCompanyPlansResponse.md)
 - [ListComponentsParams](docs/ListComponentsParams.md)
 - [ListComponentsResponse](docs/ListComponentsResponse.md)
 - [ListCrmProductsParams](docs/ListCrmProductsParams.md)
 - [ListCrmProductsResponse](docs/ListCrmProductsResponse.md)
 - [ListCustomersParams](docs/ListCustomersParams.md)
 - [ListCustomersResponse](docs/ListCustomersResponse.md)
 - [ListEntityKeyDefinitionsParams](docs/ListEntityKeyDefinitionsParams.md)
 - [ListEntityKeyDefinitionsResponse](docs/ListEntityKeyDefinitionsResponse.md)
 - [ListEntityTraitDefinitionsParams](docs/ListEntityTraitDefinitionsParams.md)
 - [ListEntityTraitDefinitionsResponse](docs/ListEntityTraitDefinitionsResponse.md)
 - [ListEnvironmentsParams](docs/ListEnvironmentsParams.md)
 - [ListEnvironmentsResponse](docs/ListEnvironmentsResponse.md)
 - [ListEventsParams](docs/ListEventsParams.md)
 - [ListEventsResponse](docs/ListEventsResponse.md)
 - [ListFeatureCompaniesParams](docs/ListFeatureCompaniesParams.md)
 - [ListFeatureCompaniesResponse](docs/ListFeatureCompaniesResponse.md)
 - [ListFeatureUsageParams](docs/ListFeatureUsageParams.md)
 - [ListFeatureUsageResponse](docs/ListFeatureUsageResponse.md)
 - [ListFeatureUsersParams](docs/ListFeatureUsersParams.md)
 - [ListFeatureUsersResponse](docs/ListFeatureUsersResponse.md)
 - [ListFeaturesParams](docs/ListFeaturesParams.md)
 - [ListFeaturesResponse](docs/ListFeaturesResponse.md)
 - [ListFlagsParams](docs/ListFlagsParams.md)
 - [ListFlagsResponse](docs/ListFlagsResponse.md)
 - [ListMetricCountsParams](docs/ListMetricCountsParams.md)
 - [ListMetricCountsResponse](docs/ListMetricCountsResponse.md)
 - [ListPlanEntitlementsParams](docs/ListPlanEntitlementsParams.md)
 - [ListPlanEntitlementsResponse](docs/ListPlanEntitlementsResponse.md)
 - [ListPlansParams](docs/ListPlansParams.md)
 - [ListPlansResponse](docs/ListPlansResponse.md)
 - [ListProductPricesParams](docs/ListProductPricesParams.md)
 - [ListProductPricesResponse](docs/ListProductPricesResponse.md)
 - [ListUsersParams](docs/ListUsersParams.md)
 - [ListUsersResponse](docs/ListUsersResponse.md)
 - [ListWebhookEventsParams](docs/ListWebhookEventsParams.md)
 - [ListWebhookEventsResponse](docs/ListWebhookEventsResponse.md)
 - [ListWebhooksParams](docs/ListWebhooksParams.md)
 - [ListWebhooksResponse](docs/ListWebhooksResponse.md)
 - [LookupCompanyParams](docs/LookupCompanyParams.md)
 - [LookupCompanyResponse](docs/LookupCompanyResponse.md)
 - [LookupUserParams](docs/LookupUserParams.md)
 - [LookupUserResponse](docs/LookupUserResponse.md)
 - [MetricCountsHourlyResponseData](docs/MetricCountsHourlyResponseData.md)
 - [PaginationFilter](docs/PaginationFilter.md)
 - [PlanAudienceDetailResponseData](docs/PlanAudienceDetailResponseData.md)
 - [PlanAudienceResponseData](docs/PlanAudienceResponseData.md)
 - [PlanDetailResponseData](docs/PlanDetailResponseData.md)
 - [PlanEntitlementResponseData](docs/PlanEntitlementResponseData.md)
 - [PlanResponseData](docs/PlanResponseData.md)
 - [PreviewObject](docs/PreviewObject.md)
 - [RawEventBatchResponseData](docs/RawEventBatchResponseData.md)
 - [RawEventResponseData](docs/RawEventResponseData.md)
 - [RuleConditionDetailResponseData](docs/RuleConditionDetailResponseData.md)
 - [RuleConditionGroupDetailResponseData](docs/RuleConditionGroupDetailResponseData.md)
 - [RuleConditionGroupResponseData](docs/RuleConditionGroupResponseData.md)
 - [RuleConditionResourceResponseData](docs/RuleConditionResourceResponseData.md)
 - [RuleConditionResponseData](docs/RuleConditionResponseData.md)
 - [RuleDetailResponseData](docs/RuleDetailResponseData.md)
 - [RuleResponseData](docs/RuleResponseData.md)
 - [RulesDetailResponseData](docs/RulesDetailResponseData.md)
 - [SegmentStatusResp](docs/SegmentStatusResp.md)
 - [TemporaryAccessTokenResponseData](docs/TemporaryAccessTokenResponseData.md)
 - [UpdateApiKeyRequestBody](docs/UpdateApiKeyRequestBody.md)
 - [UpdateApiKeyResponse](docs/UpdateApiKeyResponse.md)
 - [UpdateAudienceRequestBody](docs/UpdateAudienceRequestBody.md)
 - [UpdateAudienceResponse](docs/UpdateAudienceResponse.md)
 - [UpdateCompanyOverrideRequestBody](docs/UpdateCompanyOverrideRequestBody.md)
 - [UpdateCompanyOverrideResponse](docs/UpdateCompanyOverrideResponse.md)
 - [UpdateComponentRequestBody](docs/UpdateComponentRequestBody.md)
 - [UpdateComponentResponse](docs/UpdateComponentResponse.md)
 - [UpdateEntityTraitDefinitionRequestBody](docs/UpdateEntityTraitDefinitionRequestBody.md)
 - [UpdateEntityTraitDefinitionResponse](docs/UpdateEntityTraitDefinitionResponse.md)
 - [UpdateEnvironmentRequestBody](docs/UpdateEnvironmentRequestBody.md)
 - [UpdateEnvironmentResponse](docs/UpdateEnvironmentResponse.md)
 - [UpdateFeatureRequestBody](docs/UpdateFeatureRequestBody.md)
 - [UpdateFeatureResponse](docs/UpdateFeatureResponse.md)
 - [UpdateFlagResponse](docs/UpdateFlagResponse.md)
 - [UpdateFlagRulesRequestBody](docs/UpdateFlagRulesRequestBody.md)
 - [UpdateFlagRulesResponse](docs/UpdateFlagRulesResponse.md)
 - [UpdatePlanEntitlementRequestBody](docs/UpdatePlanEntitlementRequestBody.md)
 - [UpdatePlanEntitlementResponse](docs/UpdatePlanEntitlementResponse.md)
 - [UpdatePlanRequestBody](docs/UpdatePlanRequestBody.md)
 - [UpdatePlanResponse](docs/UpdatePlanResponse.md)
 - [UpdateReqCommon](docs/UpdateReqCommon.md)
 - [UpdateRuleRequestBody](docs/UpdateRuleRequestBody.md)
 - [UpdateWebhookRequestBody](docs/UpdateWebhookRequestBody.md)
 - [UpdateWebhookResponse](docs/UpdateWebhookResponse.md)
 - [UpsertBillingCustomerResponse](docs/UpsertBillingCustomerResponse.md)
 - [UpsertBillingPriceResponse](docs/UpsertBillingPriceResponse.md)
 - [UpsertBillingProductPlanResponse](docs/UpsertBillingProductPlanResponse.md)
 - [UpsertBillingProductRequestBody](docs/UpsertBillingProductRequestBody.md)
 - [UpsertBillingProductResponse](docs/UpsertBillingProductResponse.md)
 - [UpsertBillingSubscriptionResponse](docs/UpsertBillingSubscriptionResponse.md)
 - [UpsertCompanyRequestBody](docs/UpsertCompanyRequestBody.md)
 - [UpsertCompanyResponse](docs/UpsertCompanyResponse.md)
 - [UpsertCompanyTraitResponse](docs/UpsertCompanyTraitResponse.md)
 - [UpsertCrmDealResponse](docs/UpsertCrmDealResponse.md)
 - [UpsertCrmProductResponse](docs/UpsertCrmProductResponse.md)
 - [UpsertDealLineItemAssociationResponse](docs/UpsertDealLineItemAssociationResponse.md)
 - [UpsertLineItemResponse](docs/UpsertLineItemResponse.md)
 - [UpsertTraitRequestBody](docs/UpsertTraitRequestBody.md)
 - [UpsertUserRequestBody](docs/UpsertUserRequestBody.md)
 - [UpsertUserResponse](docs/UpsertUserResponse.md)
 - [UpsertUserSubRequestBody](docs/UpsertUserSubRequestBody.md)
 - [UpsertUserTraitResponse](docs/UpsertUserTraitResponse.md)
 - [UserDetailResponseData](docs/UserDetailResponseData.md)
 - [UserResponseData](docs/UserResponseData.md)
 - [WebhookEventDetailResponseData](docs/WebhookEventDetailResponseData.md)
 - [WebhookEventResponseData](docs/WebhookEventResponseData.md)
 - [WebhookResponseData](docs/WebhookResponseData.md)

