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
*AccountsAPI* | [**UpdateApiKey**](docs/AccountsAPI.md#updateapikey) | **Put** /api-keys/{api_key_id} | Update api key
*AccountsAPI* | [**UpdateEnvironment**](docs/AccountsAPI.md#updateenvironment) | **Put** /environments/{environment_id} | Update environment
*CompaniesAPI* | [**CreateCompany**](docs/CompaniesAPI.md#createcompany) | **Post** /companies/create | Create company
*CompaniesAPI* | [**CreateUser**](docs/CompaniesAPI.md#createuser) | **Post** /users/create | Create user
*CompaniesAPI* | [**DeleteCompany**](docs/CompaniesAPI.md#deletecompany) | **Delete** /companies/{company_id} | Delete company
*CompaniesAPI* | [**DeleteCompanyByKeys**](docs/CompaniesAPI.md#deletecompanybykeys) | **Post** /companies/delete | Delete company by keys
*CompaniesAPI* | [**DeleteCompanyMembership**](docs/CompaniesAPI.md#deletecompanymembership) | **Delete** /company-memberships/{company_membership_id} | Delete company membership
*CompaniesAPI* | [**DeleteUser**](docs/CompaniesAPI.md#deleteuser) | **Delete** /users/{user_id} | Delete user
*CompaniesAPI* | [**DeleteUserByKeys**](docs/CompaniesAPI.md#deleteuserbykeys) | **Post** /users/delete | Delete user by keys
*CompaniesAPI* | [**GetCompany**](docs/CompaniesAPI.md#getcompany) | **Get** /companies/{company_id} | Get company
*CompaniesAPI* | [**GetOrCreateCompanyMembership**](docs/CompaniesAPI.md#getorcreatecompanymembership) | **Post** /company-memberships | Get or create company membership
*CompaniesAPI* | [**GetOrCreateEntityTraitDefinition**](docs/CompaniesAPI.md#getorcreateentitytraitdefinition) | **Post** /entity-trait-definitions | Get or create entity trait definition
*CompaniesAPI* | [**GetUser**](docs/CompaniesAPI.md#getuser) | **Get** /users/{user_id} | Get user
*CompaniesAPI* | [**ListCompanies**](docs/CompaniesAPI.md#listcompanies) | **Get** /companies | List companies
*CompaniesAPI* | [**ListCompanyMemberships**](docs/CompaniesAPI.md#listcompanymemberships) | **Get** /company-memberships | List company memberships
*CompaniesAPI* | [**ListCompanyPlans**](docs/CompaniesAPI.md#listcompanyplans) | **Get** /company-plans | List company plans
*CompaniesAPI* | [**ListUsers**](docs/CompaniesAPI.md#listusers) | **Get** /users | List users
*CompaniesAPI* | [**LookupCompany**](docs/CompaniesAPI.md#lookupcompany) | **Get** /companies/lookup | Lookup company
*CompaniesAPI* | [**LookupUser**](docs/CompaniesAPI.md#lookupuser) | **Get** /users/lookup | Lookup user
*CompaniesAPI* | [**UpdateEntityTraitDefinition**](docs/CompaniesAPI.md#updateentitytraitdefinition) | **Put** /entity-trait-definitions/{entity_trait_definition_id} | Update entity trait definition
*CompaniesAPI* | [**UpsertCompany**](docs/CompaniesAPI.md#upsertcompany) | **Post** /companies | Upsert company
*CompaniesAPI* | [**UpsertCompanyTrait**](docs/CompaniesAPI.md#upsertcompanytrait) | **Post** /company-traits | Upsert company trait
*CompaniesAPI* | [**UpsertUser**](docs/CompaniesAPI.md#upsertuser) | **Post** /users | Upsert user
*CompaniesAPI* | [**UpsertUserTrait**](docs/CompaniesAPI.md#upsertusertrait) | **Post** /user-traits | Upsert user trait
*EntitlementsAPI* | [**CreateCompanyOverride**](docs/EntitlementsAPI.md#createcompanyoverride) | **Post** /company-overrides | Create company override
*EntitlementsAPI* | [**CreatePlanEntitlement**](docs/EntitlementsAPI.md#createplanentitlement) | **Post** /plan-entitlements | Create plan entitlement
*EntitlementsAPI* | [**DeleteCompanyOverride**](docs/EntitlementsAPI.md#deletecompanyoverride) | **Delete** /company-overrides/{company_override_id} | Delete company override
*EntitlementsAPI* | [**DeletePlanEntitlement**](docs/EntitlementsAPI.md#deleteplanentitlement) | **Delete** /plan-entitlements/{plan_entitlement_id} | Delete plan entitlement
*EntitlementsAPI* | [**GetCompanyOverride**](docs/EntitlementsAPI.md#getcompanyoverride) | **Get** /company-overrides/{company_override_id} | Get company override
*EntitlementsAPI* | [**GetPlanEntitlement**](docs/EntitlementsAPI.md#getplanentitlement) | **Get** /plan-entitlements/{plan_entitlement_id} | Get plan entitlement
*EntitlementsAPI* | [**ListCompanyOverrides**](docs/EntitlementsAPI.md#listcompanyoverrides) | **Get** /company-overrides | List company overrides
*EntitlementsAPI* | [**ListPlanEntitlements**](docs/EntitlementsAPI.md#listplanentitlements) | **Get** /plan-entitlements | List plan entitlements
*EntitlementsAPI* | [**UpdateCompanyOverride**](docs/EntitlementsAPI.md#updatecompanyoverride) | **Put** /company-overrides/{company_override_id} | Update company override
*EntitlementsAPI* | [**UpdatePlanEntitlement**](docs/EntitlementsAPI.md#updateplanentitlement) | **Put** /plan-entitlements/{plan_entitlement_id} | Update plan entitlement
*EventsAPI* | [**CountEventSummaries**](docs/EventsAPI.md#counteventsummaries) | **Get** /event-types/count | Count event summaries
*EventsAPI* | [**CountEvents**](docs/EventsAPI.md#countevents) | **Get** /events/count | Count events
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
*FeaturesAPI* | [**CountFlagChecks**](docs/FeaturesAPI.md#countflagchecks) | **Get** /flag-checks/count | Count flag checks
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
*PlansAPI* | [**CreatePlan**](docs/PlansAPI.md#createplan) | **Post** /plans | Create plan
*PlansAPI* | [**DeleteAudience**](docs/PlansAPI.md#deleteaudience) | **Delete** /plan-audiences/{plan_audience_id} | Delete audience
*PlansAPI* | [**DeletePlan**](docs/PlansAPI.md#deleteplan) | **Delete** /plans/{plan_id} | Delete plan
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
 - [CountApiKeysParams](docs/CountApiKeysParams.md)
 - [CountApiKeysResponse](docs/CountApiKeysResponse.md)
 - [CountApiRequestsParams](docs/CountApiRequestsParams.md)
 - [CountApiRequestsResponse](docs/CountApiRequestsResponse.md)
 - [CountAudienceCompaniesResponse](docs/CountAudienceCompaniesResponse.md)
 - [CountAudienceUsersResponse](docs/CountAudienceUsersResponse.md)
 - [CountEventSummariesParams](docs/CountEventSummariesParams.md)
 - [CountEventSummariesResponse](docs/CountEventSummariesResponse.md)
 - [CountEventsParams](docs/CountEventsParams.md)
 - [CountEventsResponse](docs/CountEventsResponse.md)
 - [CountFlagChecksParams](docs/CountFlagChecksParams.md)
 - [CountFlagChecksResponse](docs/CountFlagChecksResponse.md)
 - [CountResponse](docs/CountResponse.md)
 - [CreateApiKeyRequestBody](docs/CreateApiKeyRequestBody.md)
 - [CreateApiKeyResponse](docs/CreateApiKeyResponse.md)
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
 - [EntityKeyResponseData](docs/EntityKeyResponseData.md)
 - [EntityTraitDefinitionResponseData](docs/EntityTraitDefinitionResponseData.md)
 - [EnvironmentDetailResponseData](docs/EnvironmentDetailResponseData.md)
 - [EnvironmentResponseData](docs/EnvironmentResponseData.md)
 - [EventBody](docs/EventBody.md)
 - [EventBodyIdentify](docs/EventBodyIdentify.md)
 - [EventBodyIdentifyCompany](docs/EventBodyIdentifyCompany.md)
 - [EventBodyTrack](docs/EventBodyTrack.md)
 - [EventListResponseData](docs/EventListResponseData.md)
 - [EventResponseData](docs/EventResponseData.md)
 - [EventSummaryResponseData](docs/EventSummaryResponseData.md)
 - [FeatureDetailResponseData](docs/FeatureDetailResponseData.md)
 - [FeatureResponseData](docs/FeatureResponseData.md)
 - [FlagCheckLogDetailResponseData](docs/FlagCheckLogDetailResponseData.md)
 - [FlagCheckLogResponseData](docs/FlagCheckLogResponseData.md)
 - [FlagDetailResponseData](docs/FlagDetailResponseData.md)
 - [FlagResponseData](docs/FlagResponseData.md)
 - [GetApiKeyResponse](docs/GetApiKeyResponse.md)
 - [GetApiRequestResponse](docs/GetApiRequestResponse.md)
 - [GetCompanyOverrideResponse](docs/GetCompanyOverrideResponse.md)
 - [GetCompanyResponse](docs/GetCompanyResponse.md)
 - [GetEnvironmentResponse](docs/GetEnvironmentResponse.md)
 - [GetEventResponse](docs/GetEventResponse.md)
 - [GetEventSummariesParams](docs/GetEventSummariesParams.md)
 - [GetEventSummariesResponse](docs/GetEventSummariesResponse.md)
 - [GetEventSummaryBySubtypeResponse](docs/GetEventSummaryBySubtypeResponse.md)
 - [GetFeatureResponse](docs/GetFeatureResponse.md)
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
 - [ListCompaniesParams](docs/ListCompaniesParams.md)
 - [ListCompaniesResponse](docs/ListCompaniesResponse.md)
 - [ListCompanyMembershipsParams](docs/ListCompanyMembershipsParams.md)
 - [ListCompanyMembershipsResponse](docs/ListCompanyMembershipsResponse.md)
 - [ListCompanyOverridesParams](docs/ListCompanyOverridesParams.md)
 - [ListCompanyOverridesResponse](docs/ListCompanyOverridesResponse.md)
 - [ListCompanyPlansParams](docs/ListCompanyPlansParams.md)
 - [ListCompanyPlansResponse](docs/ListCompanyPlansResponse.md)
 - [ListEventsParams](docs/ListEventsParams.md)
 - [ListEventsResponse](docs/ListEventsResponse.md)
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
 - [PlanEntitlementResponseData](docs/PlanEntitlementResponseData.md)
 - [PlanResponseData](docs/PlanResponseData.md)
 - [RawEventBatchResponseData](docs/RawEventBatchResponseData.md)
 - [RawEventResponseData](docs/RawEventResponseData.md)
 - [RuleConditionGroupDetailResponseData](docs/RuleConditionGroupDetailResponseData.md)
 - [RuleConditionGroupResponseData](docs/RuleConditionGroupResponseData.md)
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

