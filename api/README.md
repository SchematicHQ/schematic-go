# schematic-go API

## API Endpoints

All URIs are relative to *https://api.schematichq.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
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
*BillingAPI* | [**ListProducts**](docs/BillingAPI.md#listproducts) | **Get** /billing/products | List products
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
*CrmAPI* | [**ListCRMProducts**](docs/CrmAPI.md#listcrmproducts) | **Get** /crm/products | List c r m products
*CrmAPI* | [**UpsertCRMProduct**](docs/CrmAPI.md#upsertcrmproduct) | **Post** /crm/products/upsert | Upsert c r m product
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
*EventsAPI* | [**ListEvents**](docs/EventsAPI.md#listevents) | **Get** /events | List events
*EventsAPI* | [**ListMetricCounts**](docs/EventsAPI.md#listmetriccounts) | **Get** /metric-counts | List metric counts
*FeaturesAPI* | [**CheckFlag**](docs/FeaturesAPI.md#checkflag) | **Post** /flags/{key}/check | Check flag
*FeaturesAPI* | [**CheckFlags**](docs/FeaturesAPI.md#checkflags) | **Post** /flags/check | Check flags
*FeaturesAPI* | [**CountAudienceCompanies**](docs/FeaturesAPI.md#countaudiencecompanies) | **Post** /audience/count-companies | Count audience companies
*FeaturesAPI* | [**CountAudienceUsers**](docs/FeaturesAPI.md#countaudienceusers) | **Post** /audience/count-users | Count audience users
*FeaturesAPI* | [**CountFeatures**](docs/FeaturesAPI.md#countfeatures) | **Get** /features/count | Count features
*FeaturesAPI* | [**CountFlagChecks**](docs/FeaturesAPI.md#countflagchecks) | **Get** /flag-checks/count | Count flag checks
*FeaturesAPI* | [**CountFlags**](docs/FeaturesAPI.md#countflags) | **Get** /flags/count | Count flags
*FeaturesAPI* | [**CreateFeature**](docs/FeaturesAPI.md#createfeature) | **Post** /features | Create feature
*FeaturesAPI* | [**CreateFlag**](docs/FeaturesAPI.md#createflag) | **Post** /flags | Create flag
*FeaturesAPI* | [**DeleteFeature**](docs/FeaturesAPI.md#deletefeature) | **Delete** /features/{feature_id} | Delete feature
*FeaturesAPI* | [**DeleteFlag**](docs/FeaturesAPI.md#deleteflag) | **Delete** /flags/{flag_id} | Delete flag
*FeaturesAPI* | [**GetFeature**](docs/FeaturesAPI.md#getfeature) | **Get** /features/{feature_id} | Get feature
*FeaturesAPI* | [**GetFlag**](docs/FeaturesAPI.md#getflag) | **Get** /flags/{flag_id} | Get flag
*FeaturesAPI* | [**GetFlagCheck**](docs/FeaturesAPI.md#getflagcheck) | **Get** /flag-checks/{flag_check_id} | Get flag check
*FeaturesAPI* | [**GetLatestFlagChecks**](docs/FeaturesAPI.md#getlatestflagchecks) | **Get** /flag-checks/latest | Get latest flag checks
*FeaturesAPI* | [**ListAudienceCompanies**](docs/FeaturesAPI.md#listaudiencecompanies) | **Post** /audience/get-companies | List audience companies
*FeaturesAPI* | [**ListAudienceUsers**](docs/FeaturesAPI.md#listaudienceusers) | **Post** /audience/get-users | List audience users
*FeaturesAPI* | [**ListFeatures**](docs/FeaturesAPI.md#listfeatures) | **Get** /features | List features
*FeaturesAPI* | [**ListFlagChecks**](docs/FeaturesAPI.md#listflagchecks) | **Get** /flag-checks | List flag checks
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


## Models

 - [ApiError](docs/ApiError.md)
 - [ApiKeyCreateResponseData](docs/ApiKeyCreateResponseData.md)
 - [ApiKeyRequestListResponseData](docs/ApiKeyRequestListResponseData.md)
 - [ApiKeyRequestResponseData](docs/ApiKeyRequestResponseData.md)
 - [ApiKeyResponseData](docs/ApiKeyResponseData.md)
 - [AudienceRequestBody](docs/AudienceRequestBody.md)
 - [BillingProductResponseData](docs/BillingProductResponseData.md)
 - [BillingSubscriptionResponseData](docs/BillingSubscriptionResponseData.md)
 - [CRMProductResponseData](docs/CRMProductResponseData.md)
 - [CheckFlagOutputWithFlagKey](docs/CheckFlagOutputWithFlagKey.md)
 - [CheckFlagRequestBody](docs/CheckFlagRequestBody.md)
 - [CheckFlagResponse](docs/CheckFlagResponse.md)
 - [CheckFlagResponseData](docs/CheckFlagResponseData.md)
 - [CheckFlagsResponse](docs/CheckFlagsResponse.md)
 - [CheckFlagsResponseData](docs/CheckFlagsResponseData.md)
 - [CompanyDetailResponseData](docs/CompanyDetailResponseData.md)
 - [CompanyMembershipDetailResponseData](docs/CompanyMembershipDetailResponseData.md)
 - [CompanyMembershipResponseData](docs/CompanyMembershipResponseData.md)
 - [CompanyOverrideResponseData](docs/CompanyOverrideResponseData.md)
 - [CompanyPlanResponseData](docs/CompanyPlanResponseData.md)
 - [CompanyResponseData](docs/CompanyResponseData.md)
 - [CompanySubscriptionResponseData](docs/CompanySubscriptionResponseData.md)
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
 - [CountFlagChecksParams](docs/CountFlagChecksParams.md)
 - [CountFlagChecksResponse](docs/CountFlagChecksResponse.md)
 - [CountFlagsParams](docs/CountFlagsParams.md)
 - [CountFlagsResponse](docs/CountFlagsResponse.md)
 - [CountPlanEntitlementsParams](docs/CountPlanEntitlementsParams.md)
 - [CountPlanEntitlementsResponse](docs/CountPlanEntitlementsResponse.md)
 - [CountPlansParams](docs/CountPlansParams.md)
 - [CountPlansResponse](docs/CountPlansResponse.md)
 - [CountResponse](docs/CountResponse.md)
 - [CountUsersParams](docs/CountUsersParams.md)
 - [CountUsersResponse](docs/CountUsersResponse.md)
 - [CreateApiKeyRequestBody](docs/CreateApiKeyRequestBody.md)
 - [CreateApiKeyResponse](docs/CreateApiKeyResponse.md)
 - [CreateBillingProductRequestBody](docs/CreateBillingProductRequestBody.md)
 - [CreateBillingSubscriptionsRequestBody](docs/CreateBillingSubscriptionsRequestBody.md)
 - [CreateCRMProductRequestBody](docs/CreateCRMProductRequestBody.md)
 - [CreateCompanyOverrideRequestBody](docs/CreateCompanyOverrideRequestBody.md)
 - [CreateCompanyOverrideResponse](docs/CreateCompanyOverrideResponse.md)
 - [CreateCompanyResponse](docs/CreateCompanyResponse.md)
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
 - [DeleteApiKeyResponse](docs/DeleteApiKeyResponse.md)
 - [DeleteAudienceResponse](docs/DeleteAudienceResponse.md)
 - [DeleteCompanyByKeysResponse](docs/DeleteCompanyByKeysResponse.md)
 - [DeleteCompanyMembershipResponse](docs/DeleteCompanyMembershipResponse.md)
 - [DeleteCompanyOverrideResponse](docs/DeleteCompanyOverrideResponse.md)
 - [DeleteCompanyResponse](docs/DeleteCompanyResponse.md)
 - [DeleteEnvironmentResponse](docs/DeleteEnvironmentResponse.md)
 - [DeleteFeatureResponse](docs/DeleteFeatureResponse.md)
 - [DeleteFlagResponse](docs/DeleteFlagResponse.md)
 - [DeletePlanEntitlementResponse](docs/DeletePlanEntitlementResponse.md)
 - [DeletePlanResponse](docs/DeletePlanResponse.md)
 - [DeleteResponse](docs/DeleteResponse.md)
 - [DeleteUserByKeysResponse](docs/DeleteUserByKeysResponse.md)
 - [DeleteUserResponse](docs/DeleteUserResponse.md)
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
 - [FlagCheckLogDetailResponseData](docs/FlagCheckLogDetailResponseData.md)
 - [FlagCheckLogResponseData](docs/FlagCheckLogResponseData.md)
 - [FlagDetailResponseData](docs/FlagDetailResponseData.md)
 - [FlagResponseData](docs/FlagResponseData.md)
 - [GetActiveCompanySubscriptionParams](docs/GetActiveCompanySubscriptionParams.md)
 - [GetActiveCompanySubscriptionResponse](docs/GetActiveCompanySubscriptionResponse.md)
 - [GetApiKeyResponse](docs/GetApiKeyResponse.md)
 - [GetApiRequestResponse](docs/GetApiRequestResponse.md)
 - [GetAudienceResponse](docs/GetAudienceResponse.md)
 - [GetCompanyOverrideResponse](docs/GetCompanyOverrideResponse.md)
 - [GetCompanyResponse](docs/GetCompanyResponse.md)
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
 - [GetFlagCheckResponse](docs/GetFlagCheckResponse.md)
 - [GetFlagResponse](docs/GetFlagResponse.md)
 - [GetLatestFlagChecksParams](docs/GetLatestFlagChecksParams.md)
 - [GetLatestFlagChecksResponse](docs/GetLatestFlagChecksResponse.md)
 - [GetOrCreateCompanyMembershipRequestBody](docs/GetOrCreateCompanyMembershipRequestBody.md)
 - [GetOrCreateCompanyMembershipResponse](docs/GetOrCreateCompanyMembershipResponse.md)
 - [GetOrCreateEntityTraitDefinitionResponse](docs/GetOrCreateEntityTraitDefinitionResponse.md)
 - [GetPlanEntitlementResponse](docs/GetPlanEntitlementResponse.md)
 - [GetPlanResponse](docs/GetPlanResponse.md)
 - [GetUserResponse](docs/GetUserResponse.md)
 - [KeysRequestBody](docs/KeysRequestBody.md)
 - [ListApiKeysParams](docs/ListApiKeysParams.md)
 - [ListApiKeysResponse](docs/ListApiKeysResponse.md)
 - [ListApiRequestsParams](docs/ListApiRequestsParams.md)
 - [ListApiRequestsResponse](docs/ListApiRequestsResponse.md)
 - [ListAudienceCompaniesResponse](docs/ListAudienceCompaniesResponse.md)
 - [ListAudienceUsersResponse](docs/ListAudienceUsersResponse.md)
 - [ListCRMProductsParams](docs/ListCRMProductsParams.md)
 - [ListCRMProductsResponse](docs/ListCRMProductsResponse.md)
 - [ListCompaniesParams](docs/ListCompaniesParams.md)
 - [ListCompaniesResponse](docs/ListCompaniesResponse.md)
 - [ListCompanyMembershipsParams](docs/ListCompanyMembershipsParams.md)
 - [ListCompanyMembershipsResponse](docs/ListCompanyMembershipsResponse.md)
 - [ListCompanyOverridesParams](docs/ListCompanyOverridesParams.md)
 - [ListCompanyOverridesResponse](docs/ListCompanyOverridesResponse.md)
 - [ListCompanyPlansParams](docs/ListCompanyPlansParams.md)
 - [ListCompanyPlansResponse](docs/ListCompanyPlansResponse.md)
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
 - [ListFlagChecksParams](docs/ListFlagChecksParams.md)
 - [ListFlagChecksResponse](docs/ListFlagChecksResponse.md)
 - [ListFlagsParams](docs/ListFlagsParams.md)
 - [ListFlagsResponse](docs/ListFlagsResponse.md)
 - [ListMetricCountsParams](docs/ListMetricCountsParams.md)
 - [ListMetricCountsResponse](docs/ListMetricCountsResponse.md)
 - [ListPlanEntitlementsParams](docs/ListPlanEntitlementsParams.md)
 - [ListPlanEntitlementsResponse](docs/ListPlanEntitlementsResponse.md)
 - [ListPlansParams](docs/ListPlansParams.md)
 - [ListPlansResponse](docs/ListPlansResponse.md)
 - [ListProductsParams](docs/ListProductsParams.md)
 - [ListProductsResponse](docs/ListProductsResponse.md)
 - [ListUsersParams](docs/ListUsersParams.md)
 - [ListUsersResponse](docs/ListUsersResponse.md)
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
 - [UpdateApiKeyRequestBody](docs/UpdateApiKeyRequestBody.md)
 - [UpdateApiKeyResponse](docs/UpdateApiKeyResponse.md)
 - [UpdateAudienceRequestBody](docs/UpdateAudienceRequestBody.md)
 - [UpdateAudienceResponse](docs/UpdateAudienceResponse.md)
 - [UpdateCompanyOverrideRequestBody](docs/UpdateCompanyOverrideRequestBody.md)
 - [UpdateCompanyOverrideResponse](docs/UpdateCompanyOverrideResponse.md)
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
 - [UpsertBillingProductResponse](docs/UpsertBillingProductResponse.md)
 - [UpsertBillingSubscriptionResponse](docs/UpsertBillingSubscriptionResponse.md)
 - [UpsertCRMProductResponse](docs/UpsertCRMProductResponse.md)
 - [UpsertCompanyRequestBody](docs/UpsertCompanyRequestBody.md)
 - [UpsertCompanyResponse](docs/UpsertCompanyResponse.md)
 - [UpsertCompanyTraitResponse](docs/UpsertCompanyTraitResponse.md)
 - [UpsertTraitRequestBody](docs/UpsertTraitRequestBody.md)
 - [UpsertUserRequestBody](docs/UpsertUserRequestBody.md)
 - [UpsertUserResponse](docs/UpsertUserResponse.md)
 - [UpsertUserSubRequestBody](docs/UpsertUserSubRequestBody.md)
 - [UpsertUserTraitResponse](docs/UpsertUserTraitResponse.md)
 - [UserDetailResponseData](docs/UserDetailResponseData.md)
 - [UserResponseData](docs/UserResponseData.md)

