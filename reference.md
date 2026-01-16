# Reference
<details><summary><code>client.PutPlanAudiencesPlanAudienceID(PlanAudienceID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.PutPlanAudiencesPlanAudienceID(
        context.TODO(),
        "plan_audience_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planAudienceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DeletePlanAudiencesPlanAudienceID(PlanAudienceID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.DeletePlanAudiencesPlanAudienceID(
        context.TODO(),
        "plan_audience_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planAudienceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## accounts
<details><summary><code>client.Accounts.ListAPIKeys() -> *schematichq.ListAPIKeysResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListAPIKeysRequest{
        EnvironmentID: schematichq.String(
            "environment_id",
        ),
        RequireEnvironment: true,
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Accounts.ListAPIKeys(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**environmentID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**requireEnvironment:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.CreateAPIKey(request) -> *schematichq.CreateAPIKeyResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateAPIKeyRequestBody{
        Name: "name",
    }
client.Accounts.CreateAPIKey(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**environmentID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.GetAPIKey(APIKeyID) -> *schematichq.GetAPIKeyResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounts.GetAPIKey(
        context.TODO(),
        "api_key_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKeyID:** `string` — api_key_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.UpdateAPIKey(APIKeyID, request) -> *schematichq.UpdateAPIKeyResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdateAPIKeyRequestBody{}
client.Accounts.UpdateAPIKey(
        context.TODO(),
        "api_key_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKeyID:** `string` — api_key_id
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.DeleteAPIKey(APIKeyID) -> *schematichq.DeleteAPIKeyResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounts.DeleteAPIKey(
        context.TODO(),
        "api_key_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKeyID:** `string` — api_key_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.CountAPIKeys() -> *schematichq.CountAPIKeysResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountAPIKeysRequest{
        EnvironmentID: schematichq.String(
            "environment_id",
        ),
        RequireEnvironment: true,
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Accounts.CountAPIKeys(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**environmentID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**requireEnvironment:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.ListAPIRequests() -> *schematichq.ListAPIRequestsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListAPIRequestsRequest{
        Q: schematichq.String(
            "q",
        ),
        RequestType: schematichq.String(
            "request_type",
        ),
        EnvironmentID: schematichq.String(
            "environment_id",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Accounts.ListAPIRequests(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**requestType:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**environmentID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.GetAPIRequest(APIRequestID) -> *schematichq.GetAPIRequestResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounts.GetAPIRequest(
        context.TODO(),
        "api_request_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiRequestID:** `string` — api_request_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.CountAPIRequests() -> *schematichq.CountAPIRequestsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountAPIRequestsRequest{
        Q: schematichq.String(
            "q",
        ),
        RequestType: schematichq.String(
            "request_type",
        ),
        EnvironmentID: schematichq.String(
            "environment_id",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Accounts.CountAPIRequests(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**requestType:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**environmentID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.ListEnvironments() -> *schematichq.ListEnvironmentsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListEnvironmentsRequest{
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Accounts.ListEnvironments(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.CreateEnvironment(request) -> *schematichq.CreateEnvironmentResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateEnvironmentRequestBody{
        EnvironmentType: schematichq.EnvironmentTypeDevelopment,
        Name: "name",
    }
client.Accounts.CreateEnvironment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**environmentType:** `*schematichq.EnvironmentType` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.GetEnvironment(EnvironmentID) -> *schematichq.GetEnvironmentResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounts.GetEnvironment(
        context.TODO(),
        "environment_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**environmentID:** `string` — environment_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.UpdateEnvironment(EnvironmentID, request) -> *schematichq.UpdateEnvironmentResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdateEnvironmentRequestBody{}
client.Accounts.UpdateEnvironment(
        context.TODO(),
        "environment_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**environmentID:** `string` — environment_id
    
</dd>
</dl>

<dl>
<dd>

**environmentType:** `*schematichq.EnvironmentType` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.DeleteEnvironment(EnvironmentID) -> *schematichq.DeleteEnvironmentResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounts.DeleteEnvironment(
        context.TODO(),
        "environment_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**environmentID:** `string` — environment_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.Quickstart() -> *schematichq.QuickstartResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounts.Quickstart(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## billing
<details><summary><code>client.Billing.ListCoupons() -> *schematichq.ListCouponsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListCouponsRequest{
        IsActive: schematichq.Bool(
            true,
        ),
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Billing.ListCoupons(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isActive:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.UpsertBillingCoupon(request) -> *schematichq.UpsertBillingCouponResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateCouponRequestBody{
        AmountOff: 1,
        Duration: "duration",
        DurationInMonths: 1,
        ExternalID: "external_id",
        MaxRedemptions: 1,
        Name: "name",
        PercentOff: 1.1,
        TimesRedeemed: 1,
    }
client.Billing.UpsertBillingCoupon(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**amountOff:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**duration:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**durationInMonths:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**maxRedemptions:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**percentOff:** `float64` 
    
</dd>
</dl>

<dl>
<dd>

**timesRedeemed:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.UpsertBillingCustomer(request) -> *schematichq.UpsertBillingCustomerResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateBillingCustomerRequestBody{
        Email: "email",
        ExternalID: "external_id",
        Meta: map[string]string{
            "key": "value",
        },
        Name: "name",
    }
client.Billing.UpsertBillingCustomer(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**defaultPaymentMethodID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**email:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `map[string]string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**providerType:** `*schematichq.BillingProviderType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.ListCustomersWithSubscriptions() -> *schematichq.ListCustomersWithSubscriptionsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListCustomersWithSubscriptionsRequest{
        Name: schematichq.String(
            "name",
        ),
        ProviderType: schematichq.BillingProviderTypeSchematic.Ptr(),
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Billing.ListCustomersWithSubscriptions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyIDs:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**providerType:** `*schematichq.BillingProviderType` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.CountCustomers() -> *schematichq.CountCustomersResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountCustomersRequest{
        Name: schematichq.String(
            "name",
        ),
        ProviderType: schematichq.BillingProviderTypeSchematic.Ptr(),
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Billing.CountCustomers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyIDs:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**providerType:** `*schematichq.BillingProviderType` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.ListInvoices() -> *schematichq.ListInvoicesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListInvoicesRequest{
        CompanyID: schematichq.String(
            "company_id",
        ),
        CustomerExternalID: "customer_external_id",
        SubscriptionExternalID: "subscription_external_id",
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Billing.ListInvoices(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**customerExternalID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**subscriptionExternalID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.UpsertInvoice(request) -> *schematichq.UpsertInvoiceResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateInvoiceRequestBody{
        AmountDue: 1,
        AmountPaid: 1,
        AmountRemaining: 1,
        CollectionMethod: "collection_method",
        Currency: "currency",
        CustomerExternalID: "customer_external_id",
        Subtotal: 1,
    }
client.Billing.UpsertInvoice(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**amountDue:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**amountPaid:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**amountRemaining:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**collectionMethod:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**currency:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**customerExternalID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**dueDate:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**paymentMethodExternalID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**subscriptionExternalID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**subtotal:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.ListMeters() -> *schematichq.ListMetersResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListMetersRequest{
        DisplayName: schematichq.String(
            "display_name",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Billing.ListMeters(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**displayName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.UpsertBillingMeter(request) -> *schematichq.UpsertBillingMeterResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateMeterRequestBody{
        DisplayName: "display_name",
        EventName: "event_name",
        EventPayloadKey: "event_payload_key",
        ExternalID: "external_id",
    }
client.Billing.UpsertBillingMeter(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**displayName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**eventName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**eventPayloadKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.ListPaymentMethods() -> *schematichq.ListPaymentMethodsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListPaymentMethodsRequest{
        CompanyID: schematichq.String(
            "company_id",
        ),
        CustomerExternalID: "customer_external_id",
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Billing.ListPaymentMethods(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**customerExternalID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.UpsertPaymentMethod(request) -> *schematichq.UpsertPaymentMethodResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreatePaymentMethodRequestBody{
        CustomerExternalID: "customer_external_id",
        ExternalID: "external_id",
        PaymentMethodType: "payment_method_type",
    }
client.Billing.UpsertPaymentMethod(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountLast4:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**accountName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**bankName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**billingEmail:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**billingName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**cardBrand:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**cardExpMonth:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**cardExpYear:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**cardLast4:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**customerExternalID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**paymentMethodType:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.ListBillingPrices() -> *schematichq.ListBillingPricesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListBillingPricesRequest{
        ForInitialPlan: schematichq.Bool(
            true,
        ),
        ForTrialExpiryPlan: schematichq.Bool(
            true,
        ),
        Interval: schematichq.String(
            "interval",
        ),
        IsActive: schematichq.Bool(
            true,
        ),
        Price: schematichq.Int(
            1,
        ),
        ProductID: schematichq.String(
            "product_id",
        ),
        ProviderType: schematichq.BillingProviderTypeSchematic.Ptr(),
        Q: schematichq.String(
            "q",
        ),
        TiersMode: schematichq.BillingTiersModeGraduated.Ptr(),
        UsageType: schematichq.BillingPriceUsageTypeLicensed.Ptr(),
        WithMeter: schematichq.Bool(
            true,
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Billing.ListBillingPrices(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**forInitialPlan:** `*bool` — Filter for prices valid for initial plans (free prices only)
    
</dd>
</dl>

<dl>
<dd>

**forTrialExpiryPlan:** `*bool` — Filter for prices valid for trial expiry plans (free prices only)
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**interval:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `*bool` — Filter for active prices on active products (defaults to true if not specified)
    
</dd>
</dl>

<dl>
<dd>

**price:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**productID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**productIDs:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**providerType:** `*schematichq.BillingProviderType` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**tiersMode:** `*schematichq.BillingTiersMode` 
    
</dd>
</dl>

<dl>
<dd>

**usageType:** `*schematichq.BillingPriceUsageType` 
    
</dd>
</dl>

<dl>
<dd>

**withMeter:** `*bool` — Filter for prices with a meter
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.UpsertBillingPrice(request) -> *schematichq.UpsertBillingPriceResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateBillingPriceRequestBody{
        BillingScheme: schematichq.BillingPriceSchemePerUnit,
        Currency: "currency",
        ExternalAccountID: "external_account_id",
        Interval: "interval",
        IsActive: true,
        Price: 1,
        PriceExternalID: "price_external_id",
        PriceTiers: []*schematichq.CreateBillingPriceTierRequestBody{
            &schematichq.CreateBillingPriceTierRequestBody{
                PriceExternalID: "price_external_id",
            },
        },
        ProductExternalID: "product_external_id",
        UsageType: schematichq.BillingPriceUsageTypeLicensed,
    }
client.Billing.UpsertBillingPrice(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**billingScheme:** `*schematichq.BillingPriceScheme` 
    
</dd>
</dl>

<dl>
<dd>

**currency:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**externalAccountID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**interval:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**meterID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**packageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**price:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**priceDecimal:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**priceExternalID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**priceTiers:** `[]*schematichq.CreateBillingPriceTierRequestBody` 
    
</dd>
</dl>

<dl>
<dd>

**productExternalID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**providerType:** `*schematichq.BillingProviderType` 
    
</dd>
</dl>

<dl>
<dd>

**tiersMode:** `*schematichq.BillingTiersMode` 
    
</dd>
</dl>

<dl>
<dd>

**usageType:** `*schematichq.BillingPriceUsageType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.DeleteBillingProduct(BillingID) -> *schematichq.DeleteBillingProductResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Billing.DeleteBillingProduct(
        context.TODO(),
        "billing_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**billingID:** `string` — billing_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.ListBillingProductPrices() -> *schematichq.ListBillingProductPricesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListBillingProductPricesRequest{
        ForInitialPlan: schematichq.Bool(
            true,
        ),
        ForTrialExpiryPlan: schematichq.Bool(
            true,
        ),
        Interval: schematichq.String(
            "interval",
        ),
        IsActive: schematichq.Bool(
            true,
        ),
        Price: schematichq.Int(
            1,
        ),
        ProductID: schematichq.String(
            "product_id",
        ),
        ProviderType: schematichq.BillingProviderTypeSchematic.Ptr(),
        Q: schematichq.String(
            "q",
        ),
        TiersMode: schematichq.BillingTiersModeGraduated.Ptr(),
        UsageType: schematichq.BillingPriceUsageTypeLicensed.Ptr(),
        WithMeter: schematichq.Bool(
            true,
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Billing.ListBillingProductPrices(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**forInitialPlan:** `*bool` — Filter for prices valid for initial plans (free prices only)
    
</dd>
</dl>

<dl>
<dd>

**forTrialExpiryPlan:** `*bool` — Filter for prices valid for trial expiry plans (free prices only)
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**interval:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `*bool` — Filter for active prices on active products (defaults to true if not specified)
    
</dd>
</dl>

<dl>
<dd>

**price:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**productID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**productIDs:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**providerType:** `*schematichq.BillingProviderType` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**tiersMode:** `*schematichq.BillingTiersMode` 
    
</dd>
</dl>

<dl>
<dd>

**usageType:** `*schematichq.BillingPriceUsageType` 
    
</dd>
</dl>

<dl>
<dd>

**withMeter:** `*bool` — Filter for prices with a meter
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.DeleteProductPrice(BillingID) -> *schematichq.DeleteProductPriceResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Billing.DeleteProductPrice(
        context.TODO(),
        "billing_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**billingID:** `string` — billing_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.UpsertBillingProduct(request) -> *schematichq.UpsertBillingProductResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateBillingProductRequestBody{
        ExternalID: "external_id",
        Name: "name",
        Price: 1.1,
    }
client.Billing.UpsertBillingProduct(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**externalID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**price:** `float64` 
    
</dd>
</dl>

<dl>
<dd>

**providerType:** `*schematichq.BillingProviderType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.ListBillingProducts() -> *schematichq.ListBillingProductsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListBillingProductsRequest{
        IsActive: schematichq.Bool(
            true,
        ),
        Name: schematichq.String(
            "name",
        ),
        PriceUsageType: schematichq.BillingPriceUsageTypeLicensed.Ptr(),
        ProviderType: schematichq.BillingProviderTypeSchematic.Ptr(),
        Q: schematichq.String(
            "q",
        ),
        WithOneTimeCharges: schematichq.Bool(
            true,
        ),
        WithPricesOnly: schematichq.Bool(
            true,
        ),
        WithZeroPrice: schematichq.Bool(
            true,
        ),
        WithoutLinkedToPlan: schematichq.Bool(
            true,
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Billing.ListBillingProducts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `*bool` — Filter products that are active. Defaults to true if not specified
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**priceUsageType:** `*schematichq.BillingPriceUsageType` 
    
</dd>
</dl>

<dl>
<dd>

**providerType:** `*schematichq.BillingProviderType` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**withOneTimeCharges:** `*bool` — Filter products that are one time charges
    
</dd>
</dl>

<dl>
<dd>

**withPricesOnly:** `*bool` — Filter products that have prices
    
</dd>
</dl>

<dl>
<dd>

**withZeroPrice:** `*bool` — Filter products that have zero price for free subscription type
    
</dd>
</dl>

<dl>
<dd>

**withoutLinkedToPlan:** `*bool` — Filter products that are not linked to any plan
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.CountBillingProducts() -> *schematichq.CountBillingProductsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountBillingProductsRequest{
        IsActive: schematichq.Bool(
            true,
        ),
        Name: schematichq.String(
            "name",
        ),
        PriceUsageType: schematichq.BillingPriceUsageTypeLicensed.Ptr(),
        ProviderType: schematichq.BillingProviderTypeSchematic.Ptr(),
        Q: schematichq.String(
            "q",
        ),
        WithOneTimeCharges: schematichq.Bool(
            true,
        ),
        WithPricesOnly: schematichq.Bool(
            true,
        ),
        WithZeroPrice: schematichq.Bool(
            true,
        ),
        WithoutLinkedToPlan: schematichq.Bool(
            true,
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Billing.CountBillingProducts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `*bool` — Filter products that are active. Defaults to true if not specified
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**priceUsageType:** `*schematichq.BillingPriceUsageType` 
    
</dd>
</dl>

<dl>
<dd>

**providerType:** `*schematichq.BillingProviderType` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**withOneTimeCharges:** `*bool` — Filter products that are one time charges
    
</dd>
</dl>

<dl>
<dd>

**withPricesOnly:** `*bool` — Filter products that have prices
    
</dd>
</dl>

<dl>
<dd>

**withZeroPrice:** `*bool` — Filter products that have zero price for free subscription type
    
</dd>
</dl>

<dl>
<dd>

**withoutLinkedToPlan:** `*bool` — Filter products that are not linked to any plan
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.UpsertBillingSubscription(request) -> *schematichq.UpsertBillingSubscriptionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateBillingSubscriptionRequestBody{
        CancelAtPeriodEnd: true,
        Currency: "currency",
        CustomerExternalID: "customer_external_id",
        Discounts: []*schematichq.BillingSubscriptionDiscount{
            &schematichq.BillingSubscriptionDiscount{
                CouponExternalID: "coupon_external_id",
                ExternalID: "external_id",
                IsActive: true,
                StartedAt: schematichq.MustParseDateTime(
                    "2024-01-15T09:30:00Z",
                ),
            },
        },
        ExpiredAt: schematichq.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
        ProductExternalIDs: []*schematichq.BillingProductPricing{
            &schematichq.BillingProductPricing{
                Currency: "currency",
                Interval: "interval",
                Price: 1,
                PriceExternalID: "price_external_id",
                ProductExternalID: "product_external_id",
                Quantity: 1,
                UsageType: schematichq.BillingPriceUsageTypeLicensed,
            },
        },
        SubscriptionExternalID: "subscription_external_id",
        TotalPrice: 1,
    }
client.Billing.UpsertBillingSubscription(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**applicationID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**cancelAt:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**cancelAtPeriodEnd:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**currency:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**customerExternalID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**defaultPaymentMethodExternalID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**defaultPaymentMethodID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**discounts:** `[]*schematichq.BillingSubscriptionDiscount` 
    
</dd>
</dl>

<dl>
<dd>

**expiredAt:** `time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**interval:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**periodEnd:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**periodStart:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**productExternalIDs:** `[]*schematichq.BillingProductPricing` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**subscriptionExternalID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**totalPrice:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**trialEnd:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**trialEndSetting:** `*schematichq.BillingSubscriptionTrialEndSetting` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## credits
<details><summary><code>client.Credits.ListBillingCredits() -> *schematichq.ListBillingCreditsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListBillingCreditsRequest{
        Name: schematichq.String(
            "name",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Credits.ListBillingCredits(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.CreateBillingCredit(request) -> *schematichq.CreateBillingCreditResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateBillingCreditRequestBody{
        Currency: "currency",
        Description: "description",
        Name: "name",
    }
client.Credits.CreateBillingCredit(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**burnStrategy:** `*schematichq.BillingCreditBurnStrategy` 
    
</dd>
</dl>

<dl>
<dd>

**currency:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**defaultExpiryUnit:** `*schematichq.BillingCreditExpiryUnit` 
    
</dd>
</dl>

<dl>
<dd>

**defaultExpiryUnitCount:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**defaultRolloverPolicy:** `*schematichq.BillingCreditRolloverPolicy` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**icon:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**perUnitPrice:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**perUnitPriceDecimal:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**pluralName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**singularName:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.GetSingleBillingCredit(CreditID) -> *schematichq.GetSingleBillingCreditResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Credits.GetSingleBillingCredit(
        context.TODO(),
        "credit_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**creditID:** `string` — credit_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.UpdateBillingCredit(CreditID, request) -> *schematichq.UpdateBillingCreditResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdateBillingCreditRequestBody{
        Description: "description",
        Name: "name",
    }
client.Credits.UpdateBillingCredit(
        context.TODO(),
        "credit_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**creditID:** `string` — credit_id
    
</dd>
</dl>

<dl>
<dd>

**burnStrategy:** `*schematichq.BillingCreditBurnStrategy` 
    
</dd>
</dl>

<dl>
<dd>

**defaultExpiryUnit:** `*schematichq.BillingCreditExpiryUnit` 
    
</dd>
</dl>

<dl>
<dd>

**defaultExpiryUnitCount:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**defaultRolloverPolicy:** `*schematichq.BillingCreditRolloverPolicy` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**icon:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**perUnitPrice:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**perUnitPriceDecimal:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**pluralName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**singularName:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.SoftDeleteBillingCredit(CreditID) -> *schematichq.SoftDeleteBillingCreditResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Credits.SoftDeleteBillingCredit(
        context.TODO(),
        "credit_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**creditID:** `string` — credit_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.ListCreditBundles() -> *schematichq.ListCreditBundlesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListCreditBundlesRequest{
        CreditID: schematichq.String(
            "credit_id",
        ),
        Status: schematichq.BillingCreditBundleStatusActive.Ptr(),
        BundleType: &schematichq.BillingCreditBundleType(
            "fixed",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Credits.ListCreditBundles(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**creditID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*schematichq.BillingCreditBundleStatus` 
    
</dd>
</dl>

<dl>
<dd>

**bundleType:** `*schematichq.BillingCreditBundleType` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.CreateCreditBundle(request) -> *schematichq.CreateCreditBundleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateCreditBundleRequestBody{
        BundleName: "bundle_name",
        CreditID: "credit_id",
        Currency: "currency",
        PricePerUnit: 1,
    }
client.Credits.CreateCreditBundle(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bundleName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**bundleType:** `*schematichq.BillingCreditBundleType` 
    
</dd>
</dl>

<dl>
<dd>

**creditID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**currency:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expiryType:** `*schematichq.BillingCreditExpiryType` 
    
</dd>
</dl>

<dl>
<dd>

**expiryUnit:** `*schematichq.BillingCreditExpiryUnit` 
    
</dd>
</dl>

<dl>
<dd>

**expiryUnitCount:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pricePerUnit:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**pricePerUnitDecimal:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**quantity:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*schematichq.BillingCreditBundleStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.GetCreditBundle(BundleID) -> *schematichq.GetCreditBundleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Credits.GetCreditBundle(
        context.TODO(),
        "bundle_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bundleID:** `string` — bundle_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.UpdateCreditBundleDetails(BundleID, request) -> *schematichq.UpdateCreditBundleDetailsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdateCreditBundleDetailsRequestBody{
        BundleName: "bundle_name",
        PricePerUnit: 1,
    }
client.Credits.UpdateCreditBundleDetails(
        context.TODO(),
        "bundle_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bundleID:** `string` — bundle_id
    
</dd>
</dl>

<dl>
<dd>

**bundleName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expiryType:** `*schematichq.BillingCreditExpiryType` 
    
</dd>
</dl>

<dl>
<dd>

**expiryUnit:** `*schematichq.BillingCreditExpiryUnit` 
    
</dd>
</dl>

<dl>
<dd>

**expiryUnitCount:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pricePerUnit:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**pricePerUnitDecimal:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**quantity:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*schematichq.BillingCreditBundleStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.DeleteCreditBundle(BundleID) -> *schematichq.DeleteCreditBundleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Credits.DeleteCreditBundle(
        context.TODO(),
        "bundle_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bundleID:** `string` — bundle_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.CountCreditBundles() -> *schematichq.CountCreditBundlesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountCreditBundlesRequest{
        CreditID: schematichq.String(
            "credit_id",
        ),
        Status: schematichq.BillingCreditBundleStatusActive.Ptr(),
        BundleType: &schematichq.BillingCreditBundleType(
            "fixed",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Credits.CountCreditBundles(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**creditID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*schematichq.BillingCreditBundleStatus` 
    
</dd>
</dl>

<dl>
<dd>

**bundleType:** `*schematichq.BillingCreditBundleType` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.CountBillingCredits() -> *schematichq.CountBillingCreditsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountBillingCreditsRequest{
        Name: schematichq.String(
            "name",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Credits.CountBillingCredits(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.ZeroOutGrant(GrantID, request) -> *schematichq.ZeroOutGrantResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ZeroOutGrantRequestBody{}
client.Credits.ZeroOutGrant(
        context.TODO(),
        "grant_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**grantID:** `string` — grant_id
    
</dd>
</dl>

<dl>
<dd>

**reason:** `*schematichq.BillingCreditGrantZeroedOutReason` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.GrantBillingCreditsToCompany(request) -> *schematichq.GrantBillingCreditsToCompanyResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateCompanyCreditGrant{
        CompanyID: "company_id",
        CreditID: "credit_id",
        Quantity: 1,
        Reason: schematichq.BillingCreditGrantReasonBillingCreditAutoTopup,
    }
client.Credits.GrantBillingCreditsToCompany(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**billingPeriodsCount:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**creditID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**expiryType:** `*schematichq.BillingCreditExpiryType` 
    
</dd>
</dl>

<dl>
<dd>

**expiryUnit:** `*schematichq.BillingCreditExpiryUnit` 
    
</dd>
</dl>

<dl>
<dd>

**expiryUnitCount:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**quantity:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**reason:** `*schematichq.BillingCreditGrantReason` 
    
</dd>
</dl>

<dl>
<dd>

**renewalEnabled:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**renewalPeriod:** `*schematichq.BillingPlanCreditGrantResetStart` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.CountCompanyGrants() -> *schematichq.CountCompanyGrantsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountCompanyGrantsRequest{
        CompanyID: schematichq.String(
            "company_id",
        ),
        Order: schematichq.CreditGrantSortOrderCreatedAt.Ptr(),
        Dir: schematichq.SortDirectionAsc.Ptr(),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Credits.CountCompanyGrants(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**order:** `*schematichq.CreditGrantSortOrder` 
    
</dd>
</dl>

<dl>
<dd>

**dir:** `*schematichq.SortDirection` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.ListCompanyGrants() -> *schematichq.ListCompanyGrantsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListCompanyGrantsRequest{
        CompanyID: schematichq.String(
            "company_id",
        ),
        Order: schematichq.CreditGrantSortOrderCreatedAt.Ptr(),
        Dir: schematichq.SortDirectionAsc.Ptr(),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Credits.ListCompanyGrants(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**order:** `*schematichq.CreditGrantSortOrder` 
    
</dd>
</dl>

<dl>
<dd>

**dir:** `*schematichq.SortDirection` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.CountBillingCreditsGrants() -> *schematichq.CountBillingCreditsGrantsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountBillingCreditsGrantsRequest{
        CreditID: schematichq.String(
            "credit_id",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Credits.CountBillingCreditsGrants(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**creditID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.ListGrantsForCredit() -> *schematichq.ListGrantsForCreditResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListGrantsForCreditRequest{
        CreditID: schematichq.String(
            "credit_id",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Credits.ListGrantsForCredit(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**creditID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.GetEnrichedCreditLedger() -> *schematichq.GetEnrichedCreditLedgerResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.GetEnrichedCreditLedgerRequest{
        CompanyID: "company_id",
        BillingCreditID: schematichq.String(
            "billing_credit_id",
        ),
        FeatureID: schematichq.String(
            "feature_id",
        ),
        Period: schematichq.CreditLedgerPeriodDaily,
        StartTime: schematichq.String(
            "start_time",
        ),
        EndTime: schematichq.String(
            "end_time",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Credits.GetEnrichedCreditLedger(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**billingCreditID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**featureID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**period:** `*schematichq.CreditLedgerPeriod` 
    
</dd>
</dl>

<dl>
<dd>

**startTime:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**endTime:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.CountCreditLedger() -> *schematichq.CountCreditLedgerResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountCreditLedgerRequest{
        CompanyID: "company_id",
        BillingCreditID: schematichq.String(
            "billing_credit_id",
        ),
        FeatureID: schematichq.String(
            "feature_id",
        ),
        Period: schematichq.CreditLedgerPeriodDaily,
        StartTime: schematichq.String(
            "start_time",
        ),
        EndTime: schematichq.String(
            "end_time",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Credits.CountCreditLedger(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**billingCreditID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**featureID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**period:** `*schematichq.CreditLedgerPeriod` 
    
</dd>
</dl>

<dl>
<dd>

**startTime:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**endTime:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.ListBillingPlanCreditGrants() -> *schematichq.ListBillingPlanCreditGrantsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListBillingPlanCreditGrantsRequest{
        CreditID: schematichq.String(
            "credit_id",
        ),
        PlanID: schematichq.String(
            "plan_id",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Credits.ListBillingPlanCreditGrants(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**creditID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**planID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**planIDs:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.CreateBillingPlanCreditGrant(request) -> *schematichq.CreateBillingPlanCreditGrantResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateBillingPlanCreditGrantRequestBody{
        CreditAmount: 1,
        CreditID: "credit_id",
        PlanID: "plan_id",
        ResetCadence: schematichq.BillingPlanCreditGrantResetCadenceDaily,
        ResetStart: schematichq.BillingPlanCreditGrantResetStartBillingPeriod,
    }
client.Credits.CreateBillingPlanCreditGrant(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.CreateBillingPlanCreditGrantRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.UpdateBillingPlanCreditGrant(PlanGrantID, request) -> *schematichq.UpdateBillingPlanCreditGrantResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdateBillingPlanCreditGrantRequestBody{
        ResetCadence: schematichq.BillingPlanCreditGrantResetCadenceDaily,
        ResetStart: schematichq.BillingPlanCreditGrantResetStartBillingPeriod,
    }
client.Credits.UpdateBillingPlanCreditGrant(
        context.TODO(),
        "plan_grant_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planGrantID:** `string` — plan_grant_id
    
</dd>
</dl>

<dl>
<dd>

**request:** `*schematichq.UpdateBillingPlanCreditGrantRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.DeleteBillingPlanCreditGrant(PlanGrantID) -> *schematichq.DeleteBillingPlanCreditGrantResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.DeleteBillingPlanCreditGrantRequest{
        ApplyToExisting: schematichq.Bool(
            true,
        ),
    }
client.Credits.DeleteBillingPlanCreditGrant(
        context.TODO(),
        "plan_grant_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planGrantID:** `string` — plan_grant_id
    
</dd>
</dl>

<dl>
<dd>

**applyToExisting:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credits.CountBillingPlanCreditGrants() -> *schematichq.CountBillingPlanCreditGrantsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountBillingPlanCreditGrantsRequest{
        CreditID: schematichq.String(
            "credit_id",
        ),
        PlanID: schematichq.String(
            "plan_id",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Credits.CountBillingPlanCreditGrants(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**creditID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**planID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**planIDs:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## checkout
<details><summary><code>client.Checkout.Internal(request) -> *schematichq.CheckoutInternalResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ChangeSubscriptionInternalRequestBody{
        AddOnIDs: []*schematichq.UpdateAddOnRequestBody{
            &schematichq.UpdateAddOnRequestBody{
                AddOnID: "add_on_id",
                PriceID: "price_id",
            },
        },
        CompanyID: "company_id",
        CreditBundles: []*schematichq.UpdateCreditBundleRequestBody{
            &schematichq.UpdateCreditBundleRequestBody{
                BundleID: "bundle_id",
                Quantity: 1,
            },
        },
        NewPlanID: "new_plan_id",
        NewPriceID: "new_price_id",
        PayInAdvance: []*schematichq.UpdatePayInAdvanceRequestBody{
            &schematichq.UpdatePayInAdvanceRequestBody{
                PriceID: "price_id",
                Quantity: 1,
            },
        },
        SkipTrial: true,
    }
client.Checkout.Internal(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.ChangeSubscriptionInternalRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Checkout.GetCheckoutData(request) -> *schematichq.GetCheckoutDataResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CheckoutDataRequestBody{
        CompanyID: "company_id",
    }
client.Checkout.GetCheckoutData(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**selectedPlanID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Checkout.PreviewCheckoutInternal(request) -> *schematichq.PreviewCheckoutInternalResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ChangeSubscriptionInternalRequestBody{
        AddOnIDs: []*schematichq.UpdateAddOnRequestBody{
            &schematichq.UpdateAddOnRequestBody{
                AddOnID: "add_on_id",
                PriceID: "price_id",
            },
        },
        CompanyID: "company_id",
        CreditBundles: []*schematichq.UpdateCreditBundleRequestBody{
            &schematichq.UpdateCreditBundleRequestBody{
                BundleID: "bundle_id",
                Quantity: 1,
            },
        },
        NewPlanID: "new_plan_id",
        NewPriceID: "new_price_id",
        PayInAdvance: []*schematichq.UpdatePayInAdvanceRequestBody{
            &schematichq.UpdatePayInAdvanceRequestBody{
                PriceID: "price_id",
                Quantity: 1,
            },
        },
        SkipTrial: true,
    }
client.Checkout.PreviewCheckoutInternal(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.ChangeSubscriptionInternalRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Checkout.ManagePlan(request) -> *schematichq.ManagePlanResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ManagePlanRequest{
        AddOnSelections: []*schematichq.PlanSelection{
            &schematichq.PlanSelection{
                PlanID: "plan_id",
            },
        },
        CompanyID: "company_id",
        CreditBundles: []*schematichq.UpdateCreditBundleRequestBody{
            &schematichq.UpdateCreditBundleRequestBody{
                BundleID: "bundle_id",
                Quantity: 1,
            },
        },
        PayInAdvanceEntitlements: []*schematichq.UpdatePayInAdvanceRequestBody{
            &schematichq.UpdatePayInAdvanceRequestBody{
                PriceID: "price_id",
                Quantity: 1,
            },
        },
    }
client.Checkout.ManagePlan(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.ManagePlanRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Checkout.PreviewManagePlan(request) -> *schematichq.PreviewManagePlanResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ManagePlanRequest{
        AddOnSelections: []*schematichq.PlanSelection{
            &schematichq.PlanSelection{
                PlanID: "plan_id",
            },
        },
        CompanyID: "company_id",
        CreditBundles: []*schematichq.UpdateCreditBundleRequestBody{
            &schematichq.UpdateCreditBundleRequestBody{
                BundleID: "bundle_id",
                Quantity: 1,
            },
        },
        PayInAdvanceEntitlements: []*schematichq.UpdatePayInAdvanceRequestBody{
            &schematichq.UpdatePayInAdvanceRequestBody{
                PriceID: "price_id",
                Quantity: 1,
            },
        },
    }
client.Checkout.PreviewManagePlan(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.ManagePlanRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Checkout.CancelSubscription(request) -> *schematichq.CancelSubscriptionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CancelSubscriptionRequest{
        CompanyID: "company_id",
    }
client.Checkout.CancelSubscription(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cancelImmediately:** `*bool` — If false, subscription cancels at period end. Defaults to true.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**prorate:** `*bool` — If true and cancel_immediately is true, issue prorated credit. Defaults to true.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Checkout.UpdateCustomerSubscriptionTrialEnd(SubscriptionID, request) -> *schematichq.UpdateCustomerSubscriptionTrialEndResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdateTrialEndRequestBody{}
client.Checkout.UpdateCustomerSubscriptionTrialEnd(
        context.TODO(),
        "subscription_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionID:** `string` — subscription_id
    
</dd>
</dl>

<dl>
<dd>

**trialEnd:** `*time.Time` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## companies
<details><summary><code>client.Companies.ListCompanies() -> *schematichq.ListCompaniesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListCompaniesRequest{
        PlanID: schematichq.String(
            "plan_id",
        ),
        Q: schematichq.String(
            "q",
        ),
        WithoutFeatureOverrideFor: schematichq.String(
            "without_feature_override_for",
        ),
        WithoutPlan: schematichq.Bool(
            true,
        ),
        WithSubscription: schematichq.Bool(
            true,
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Companies.ListCompanies(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` — Filter companies by multiple company IDs (starts with comp_)
    
</dd>
</dl>

<dl>
<dd>

**planID:** `*string` — Filter companies by plan ID (starts with plan_)
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Search for companies by name, keys or string traits
    
</dd>
</dl>

<dl>
<dd>

**withoutFeatureOverrideFor:** `*string` — Filter out companies that already have a company override for the specified feature ID
    
</dd>
</dl>

<dl>
<dd>

**withoutPlan:** `*bool` — Filter out companies that have a plan
    
</dd>
</dl>

<dl>
<dd>

**withSubscription:** `*bool` — Filter companies that have a subscription
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.UpsertCompany(request) -> *schematichq.UpsertCompanyResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpsertCompanyRequestBody{
        Keys: map[string]string{
            "key": "value",
        },
    }
client.Companies.UpsertCompany(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.UpsertCompanyRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.GetCompany(CompanyID) -> *schematichq.GetCompanyResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Companies.GetCompany(
        context.TODO(),
        "company_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `string` — company_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.DeleteCompany(CompanyID) -> *schematichq.DeleteCompanyResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.DeleteCompanyRequest{
        CancelSubscription: schematichq.Bool(
            true,
        ),
        Prorate: schematichq.Bool(
            true,
        ),
    }
client.Companies.DeleteCompany(
        context.TODO(),
        "company_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `string` — company_id
    
</dd>
</dl>

<dl>
<dd>

**cancelSubscription:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**prorate:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.CountCompanies() -> *schematichq.CountCompaniesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountCompaniesRequest{
        PlanID: schematichq.String(
            "plan_id",
        ),
        Q: schematichq.String(
            "q",
        ),
        WithoutFeatureOverrideFor: schematichq.String(
            "without_feature_override_for",
        ),
        WithoutPlan: schematichq.Bool(
            true,
        ),
        WithSubscription: schematichq.Bool(
            true,
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Companies.CountCompanies(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` — Filter companies by multiple company IDs (starts with comp_)
    
</dd>
</dl>

<dl>
<dd>

**planID:** `*string` — Filter companies by plan ID (starts with plan_)
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Search for companies by name, keys or string traits
    
</dd>
</dl>

<dl>
<dd>

**withoutFeatureOverrideFor:** `*string` — Filter out companies that already have a company override for the specified feature ID
    
</dd>
</dl>

<dl>
<dd>

**withoutPlan:** `*bool` — Filter out companies that have a plan
    
</dd>
</dl>

<dl>
<dd>

**withSubscription:** `*bool` — Filter companies that have a subscription
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.CountCompaniesForAdvancedFilter() -> *schematichq.CountCompaniesForAdvancedFilterResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountCompaniesForAdvancedFilterRequest{
        MonetizedSubscriptions: schematichq.Bool(
            true,
        ),
        Q: schematichq.String(
            "q",
        ),
        WithoutPlan: schematichq.Bool(
            true,
        ),
        WithoutSubscription: schematichq.Bool(
            true,
        ),
        SortOrderColumn: schematichq.String(
            "sort_order_column",
        ),
        SortOrderDirection: schematichq.SortDirectionAsc.Ptr(),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Companies.CountCompaniesForAdvancedFilter(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` — Filter companies by multiple company IDs (starts with comp_)
    
</dd>
</dl>

<dl>
<dd>

**planIDs:** `*string` — Filter companies by one or more plan IDs (each ID starts with plan_)
    
</dd>
</dl>

<dl>
<dd>

**featureIDs:** `*string` — Filter companies by one or more feature IDs (each ID starts with feat_)
    
</dd>
</dl>

<dl>
<dd>

**creditTypeIDs:** `*string` — Filter companies by one or more credit type IDs (each ID starts with bcrd_)
    
</dd>
</dl>

<dl>
<dd>

**subscriptionStatuses:** `*schematichq.SubscriptionStatus` — Filter companies by one or more subscription statuses (active, canceled, expired, incomplete, incomplete_expired, past_due, paused, trialing, unpaid)
    
</dd>
</dl>

<dl>
<dd>

**subscriptionTypes:** `*schematichq.SubscriptionType` — Filter companies by one or more subscription types (paid, free, trial)
    
</dd>
</dl>

<dl>
<dd>

**monetizedSubscriptions:** `*bool` — Filter companies that have monetized subscriptions
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Search for companies by name, keys or string traits
    
</dd>
</dl>

<dl>
<dd>

**withoutPlan:** `*bool` — Filter out companies that have a plan
    
</dd>
</dl>

<dl>
<dd>

**withoutSubscription:** `*bool` — Filter out companies that have a subscription
    
</dd>
</dl>

<dl>
<dd>

**sortOrderColumn:** `*string` — Column to sort by (e.g. name, created_at, last_seen_at)
    
</dd>
</dl>

<dl>
<dd>

**sortOrderDirection:** `*schematichq.SortDirection` — Direction to sort by (asc or desc)
    
</dd>
</dl>

<dl>
<dd>

**displayProperties:** `*string` — Select the display columns to return (e.g. plan, subscription, users, last_seen_at)
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.CreateCompany(request) -> *schematichq.CreateCompanyResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpsertCompanyRequestBody{
        Keys: map[string]string{
            "key": "value",
        },
    }
client.Companies.CreateCompany(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.UpsertCompanyRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.DeleteCompanyByKeys(request) -> *schematichq.DeleteCompanyByKeysResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.KeysRequestBody{
        Keys: map[string]string{
            "key": "value",
        },
    }
client.Companies.DeleteCompanyByKeys(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.KeysRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.ListCompaniesForAdvancedFilter() -> *schematichq.ListCompaniesForAdvancedFilterResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListCompaniesForAdvancedFilterRequest{
        MonetizedSubscriptions: schematichq.Bool(
            true,
        ),
        Q: schematichq.String(
            "q",
        ),
        WithoutPlan: schematichq.Bool(
            true,
        ),
        WithoutSubscription: schematichq.Bool(
            true,
        ),
        SortOrderColumn: schematichq.String(
            "sort_order_column",
        ),
        SortOrderDirection: schematichq.SortDirectionAsc.Ptr(),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Companies.ListCompaniesForAdvancedFilter(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` — Filter companies by multiple company IDs (starts with comp_)
    
</dd>
</dl>

<dl>
<dd>

**planIDs:** `*string` — Filter companies by one or more plan IDs (each ID starts with plan_)
    
</dd>
</dl>

<dl>
<dd>

**featureIDs:** `*string` — Filter companies by one or more feature IDs (each ID starts with feat_)
    
</dd>
</dl>

<dl>
<dd>

**creditTypeIDs:** `*string` — Filter companies by one or more credit type IDs (each ID starts with bcrd_)
    
</dd>
</dl>

<dl>
<dd>

**subscriptionStatuses:** `*schematichq.SubscriptionStatus` — Filter companies by one or more subscription statuses (active, canceled, expired, incomplete, incomplete_expired, past_due, paused, trialing, unpaid)
    
</dd>
</dl>

<dl>
<dd>

**subscriptionTypes:** `*schematichq.SubscriptionType` — Filter companies by one or more subscription types (paid, free, trial)
    
</dd>
</dl>

<dl>
<dd>

**monetizedSubscriptions:** `*bool` — Filter companies that have monetized subscriptions
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Search for companies by name, keys or string traits
    
</dd>
</dl>

<dl>
<dd>

**withoutPlan:** `*bool` — Filter out companies that have a plan
    
</dd>
</dl>

<dl>
<dd>

**withoutSubscription:** `*bool` — Filter out companies that have a subscription
    
</dd>
</dl>

<dl>
<dd>

**sortOrderColumn:** `*string` — Column to sort by (e.g. name, created_at, last_seen_at)
    
</dd>
</dl>

<dl>
<dd>

**sortOrderDirection:** `*schematichq.SortDirection` — Direction to sort by (asc or desc)
    
</dd>
</dl>

<dl>
<dd>

**displayProperties:** `*string` — Select the display columns to return (e.g. plan, subscription, users, last_seen_at)
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.LookupCompany() -> *schematichq.LookupCompanyResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.LookupCompanyRequest{
        Keys: map[string]string{
            "keys": "keys",
        },
    }
client.Companies.LookupCompany(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**keys:** `map[string]string` — Key/value pairs
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.ListCompanyMemberships() -> *schematichq.ListCompanyMembershipsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListCompanyMembershipsRequest{
        CompanyID: schematichq.String(
            "company_id",
        ),
        UserID: schematichq.String(
            "user_id",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Companies.ListCompanyMemberships(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.GetOrCreateCompanyMembership(request) -> *schematichq.GetOrCreateCompanyMembershipResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.GetOrCreateCompanyMembershipRequestBody{
        CompanyID: "company_id",
        UserID: "user_id",
    }
client.Companies.GetOrCreateCompanyMembership(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.DeleteCompanyMembership(CompanyMembershipID) -> *schematichq.DeleteCompanyMembershipResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Companies.DeleteCompanyMembership(
        context.TODO(),
        "company_membership_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyMembershipID:** `string` — company_membership_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.GetActiveCompanySubscription() -> *schematichq.GetActiveCompanySubscriptionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.GetActiveCompanySubscriptionRequest{
        CompanyID: schematichq.String(
            "company_id",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Companies.GetActiveCompanySubscription(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**companyIDs:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.UpsertCompanyTrait(request) -> *schematichq.UpsertCompanyTraitResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpsertTraitRequestBody{
        Keys: map[string]string{
            "key": "value",
        },
        Trait: "trait",
    }
client.Companies.UpsertCompanyTrait(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.UpsertTraitRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.ListEntityKeyDefinitions() -> *schematichq.ListEntityKeyDefinitionsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListEntityKeyDefinitionsRequest{
        EntityType: schematichq.EntityTypeCompany.Ptr(),
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Companies.ListEntityKeyDefinitions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**entityType:** `*schematichq.EntityType` 
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.CountEntityKeyDefinitions() -> *schematichq.CountEntityKeyDefinitionsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountEntityKeyDefinitionsRequest{
        EntityType: schematichq.EntityTypeCompany.Ptr(),
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Companies.CountEntityKeyDefinitions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**entityType:** `*schematichq.EntityType` 
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.ListEntityTraitDefinitions() -> *schematichq.ListEntityTraitDefinitionsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListEntityTraitDefinitionsRequest{
        EntityType: schematichq.EntityTypeCompany.Ptr(),
        Q: schematichq.String(
            "q",
        ),
        TraitType: schematichq.TraitTypeBoolean.Ptr(),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Companies.ListEntityTraitDefinitions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**entityType:** `*schematichq.EntityType` 
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**traitType:** `*schematichq.TraitType` 
    
</dd>
</dl>

<dl>
<dd>

**traitTypes:** `*schematichq.TraitType` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.GetOrCreateEntityTraitDefinition(request) -> *schematichq.GetOrCreateEntityTraitDefinitionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateEntityTraitDefinitionRequestBody{
        EntityType: schematichq.EntityTypeCompany,
        Hierarchy: []string{
            "hierarchy",
        },
        TraitType: schematichq.TraitTypeBoolean,
    }
client.Companies.GetOrCreateEntityTraitDefinition(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**displayName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**entityType:** `*schematichq.EntityType` 
    
</dd>
</dl>

<dl>
<dd>

**hierarchy:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**traitType:** `*schematichq.TraitType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.GetEntityTraitDefinition(EntityTraitDefinitionID) -> *schematichq.GetEntityTraitDefinitionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Companies.GetEntityTraitDefinition(
        context.TODO(),
        "entity_trait_definition_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**entityTraitDefinitionID:** `string` — entity_trait_definition_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.UpdateEntityTraitDefinition(EntityTraitDefinitionID, request) -> *schematichq.UpdateEntityTraitDefinitionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdateEntityTraitDefinitionRequestBody{
        TraitType: schematichq.TraitTypeBoolean,
    }
client.Companies.UpdateEntityTraitDefinition(
        context.TODO(),
        "entity_trait_definition_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**entityTraitDefinitionID:** `string` — entity_trait_definition_id
    
</dd>
</dl>

<dl>
<dd>

**displayName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**traitType:** `*schematichq.TraitType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.CountEntityTraitDefinitions() -> *schematichq.CountEntityTraitDefinitionsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountEntityTraitDefinitionsRequest{
        EntityType: schematichq.EntityTypeCompany.Ptr(),
        Q: schematichq.String(
            "q",
        ),
        TraitType: schematichq.TraitTypeBoolean.Ptr(),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Companies.CountEntityTraitDefinitions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**entityType:** `*schematichq.EntityType` 
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**traitType:** `*schematichq.TraitType` 
    
</dd>
</dl>

<dl>
<dd>

**traitTypes:** `*schematichq.TraitType` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.GetEntityTraitValues() -> *schematichq.GetEntityTraitValuesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.GetEntityTraitValuesRequest{
        DefinitionID: "definition_id",
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Companies.GetEntityTraitValues(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**definitionID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.ListPlanChanges() -> *schematichq.ListPlanChangesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListPlanChangesRequest{
        Action: schematichq.String(
            "action",
        ),
        BasePlanAction: schematichq.String(
            "base_plan_action",
        ),
        CompanyID: schematichq.String(
            "company_id",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Companies.ListPlanChanges(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**action:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**basePlanAction:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**companyIDs:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**planIDs:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.GetPlanChange(PlanChangeID) -> *schematichq.GetPlanChangeResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Companies.GetPlanChange(
        context.TODO(),
        "plan_change_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planChangeID:** `string` — plan_change_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.ListPlanTraits() -> *schematichq.ListPlanTraitsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListPlanTraitsRequest{
        PlanID: schematichq.String(
            "plan_id",
        ),
        TraitID: schematichq.String(
            "trait_id",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Companies.ListPlanTraits(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**planID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**traitID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**traitIDs:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.CreatePlanTrait(request) -> *schematichq.CreatePlanTraitResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreatePlanTraitRequestBody{
        PlanID: "plan_id",
        TraitID: "trait_id",
        TraitValue: "trait_value",
    }
client.Companies.CreatePlanTrait(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**traitID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**traitValue:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.GetPlanTrait(PlanTraitID) -> *schematichq.GetPlanTraitResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Companies.GetPlanTrait(
        context.TODO(),
        "plan_trait_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planTraitID:** `string` — plan_trait_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.UpdatePlanTrait(PlanTraitID, request) -> *schematichq.UpdatePlanTraitResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdatePlanTraitRequestBody{
        PlanID: "plan_id",
        TraitValue: "trait_value",
    }
client.Companies.UpdatePlanTrait(
        context.TODO(),
        "plan_trait_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planTraitID:** `string` — plan_trait_id
    
</dd>
</dl>

<dl>
<dd>

**planID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**traitValue:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.DeletePlanTrait(PlanTraitID) -> *schematichq.DeletePlanTraitResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Companies.DeletePlanTrait(
        context.TODO(),
        "plan_trait_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planTraitID:** `string` — plan_trait_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.UpdatePlanTraitsBulk(request) -> *schematichq.UpdatePlanTraitsBulkResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdatePlanTraitBulkRequestBody{
        ApplyToExistingCompanies: true,
        PlanID: "plan_id",
        Traits: []*schematichq.UpdatePlanTraitTraitRequestBody{
            &schematichq.UpdatePlanTraitTraitRequestBody{
                TraitID: "trait_id",
                TraitValue: "trait_value",
            },
        },
    }
client.Companies.UpdatePlanTraitsBulk(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**applyToExistingCompanies:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**planID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**traits:** `[]*schematichq.UpdatePlanTraitTraitRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.CountPlanTraits() -> *schematichq.CountPlanTraitsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountPlanTraitsRequest{
        PlanID: schematichq.String(
            "plan_id",
        ),
        TraitID: schematichq.String(
            "trait_id",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Companies.CountPlanTraits(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**planID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**traitID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**traitIDs:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.UpsertUserTrait(request) -> *schematichq.UpsertUserTraitResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpsertTraitRequestBody{
        Keys: map[string]string{
            "key": "value",
        },
        Trait: "trait",
    }
client.Companies.UpsertUserTrait(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.UpsertTraitRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.ListUsers() -> *schematichq.ListUsersResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListUsersRequest{
        CompanyID: schematichq.String(
            "company_id",
        ),
        PlanID: schematichq.String(
            "plan_id",
        ),
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Companies.ListUsers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` — Filter users by company ID (starts with comp_)
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — Filter users by multiple user IDs (starts with user_)
    
</dd>
</dl>

<dl>
<dd>

**planID:** `*string` — Filter users by plan ID (starts with plan_)
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Search for users by name, keys or string traits
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.UpsertUser(request) -> *schematichq.UpsertUserResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpsertUserRequestBody{
        Keys: map[string]string{
            "key": "value",
        },
    }
client.Companies.UpsertUser(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.UpsertUserRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.GetUser(UserID) -> *schematichq.GetUserResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Companies.GetUser(
        context.TODO(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — user_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.DeleteUser(UserID) -> *schematichq.DeleteUserResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Companies.DeleteUser(
        context.TODO(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — user_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.CountUsers() -> *schematichq.CountUsersResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountUsersRequest{
        CompanyID: schematichq.String(
            "company_id",
        ),
        PlanID: schematichq.String(
            "plan_id",
        ),
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Companies.CountUsers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` — Filter users by company ID (starts with comp_)
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — Filter users by multiple user IDs (starts with user_)
    
</dd>
</dl>

<dl>
<dd>

**planID:** `*string` — Filter users by plan ID (starts with plan_)
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Search for users by name, keys or string traits
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.CreateUser(request) -> *schematichq.CreateUserResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpsertUserRequestBody{
        Keys: map[string]string{
            "key": "value",
        },
    }
client.Companies.CreateUser(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.UpsertUserRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.DeleteUserByKeys(request) -> *schematichq.DeleteUserByKeysResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.KeysRequestBody{
        Keys: map[string]string{
            "key": "value",
        },
    }
client.Companies.DeleteUserByKeys(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.KeysRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.LookupUser() -> *schematichq.LookupUserResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.LookupUserRequest{
        Keys: map[string]string{
            "keys": "keys",
        },
    }
client.Companies.LookupUser(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**keys:** `map[string]string` — Key/value pairs
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## entitlements
<details><summary><code>client.Entitlements.ListCompanyOverrides() -> *schematichq.ListCompanyOverridesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListCompanyOverridesRequest{
        CompanyID: schematichq.String(
            "company_id",
        ),
        FeatureID: schematichq.String(
            "feature_id",
        ),
        WithoutExpired: schematichq.Bool(
            true,
        ),
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Entitlements.ListCompanyOverrides(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` — Filter company overrides by a single company ID (starting with comp_)
    
</dd>
</dl>

<dl>
<dd>

**companyIDs:** `*string` — Filter company overrides by multiple company IDs (starting with comp_)
    
</dd>
</dl>

<dl>
<dd>

**featureID:** `*string` — Filter company overrides by a single feature ID (starting with feat_)
    
</dd>
</dl>

<dl>
<dd>

**featureIDs:** `*string` — Filter company overrides by multiple feature IDs (starting with feat_)
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — Filter company overrides by multiple company override IDs (starting with cmov_)
    
</dd>
</dl>

<dl>
<dd>

**withoutExpired:** `*bool` — Filter company overrides by whether they have not expired
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Search for company overrides by feature or company name
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.CreateCompanyOverride(request) -> *schematichq.CreateCompanyOverrideResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateCompanyOverrideRequestBody{
        CompanyID: "company_id",
        FeatureID: "feature_id",
        ValueType: schematichq.EntitlementValueTypeBoolean,
    }
client.Entitlements.CreateCompanyOverride(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**creditConsumptionRate:** `*float64` 
    
</dd>
</dl>

<dl>
<dd>

**expirationDate:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**featureID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**metricPeriod:** `*schematichq.CreateCompanyOverrideRequestBodyMetricPeriod` 
    
</dd>
</dl>

<dl>
<dd>

**metricPeriodMonthReset:** `*schematichq.CreateCompanyOverrideRequestBodyMetricPeriodMonthReset` 
    
</dd>
</dl>

<dl>
<dd>

**note:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**valueBool:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**valueCreditID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**valueNumeric:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**valueTraitID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**valueType:** `*schematichq.EntitlementValueType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.GetCompanyOverride(CompanyOverrideID) -> *schematichq.GetCompanyOverrideResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Entitlements.GetCompanyOverride(
        context.TODO(),
        "company_override_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyOverrideID:** `string` — company_override_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.UpdateCompanyOverride(CompanyOverrideID, request) -> *schematichq.UpdateCompanyOverrideResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdateCompanyOverrideRequestBody{
        ValueType: schematichq.EntitlementValueTypeBoolean,
    }
client.Entitlements.UpdateCompanyOverride(
        context.TODO(),
        "company_override_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyOverrideID:** `string` — company_override_id
    
</dd>
</dl>

<dl>
<dd>

**creditConsumptionRate:** `*float64` 
    
</dd>
</dl>

<dl>
<dd>

**expirationDate:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**metricPeriod:** `*schematichq.UpdateCompanyOverrideRequestBodyMetricPeriod` 
    
</dd>
</dl>

<dl>
<dd>

**metricPeriodMonthReset:** `*schematichq.UpdateCompanyOverrideRequestBodyMetricPeriodMonthReset` 
    
</dd>
</dl>

<dl>
<dd>

**note:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**valueBool:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**valueCreditID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**valueNumeric:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**valueTraitID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**valueType:** `*schematichq.EntitlementValueType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.DeleteCompanyOverride(CompanyOverrideID) -> *schematichq.DeleteCompanyOverrideResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Entitlements.DeleteCompanyOverride(
        context.TODO(),
        "company_override_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyOverrideID:** `string` — company_override_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.CountCompanyOverrides() -> *schematichq.CountCompanyOverridesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountCompanyOverridesRequest{
        CompanyID: schematichq.String(
            "company_id",
        ),
        FeatureID: schematichq.String(
            "feature_id",
        ),
        WithoutExpired: schematichq.Bool(
            true,
        ),
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Entitlements.CountCompanyOverrides(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` — Filter company overrides by a single company ID (starting with comp_)
    
</dd>
</dl>

<dl>
<dd>

**companyIDs:** `*string` — Filter company overrides by multiple company IDs (starting with comp_)
    
</dd>
</dl>

<dl>
<dd>

**featureID:** `*string` — Filter company overrides by a single feature ID (starting with feat_)
    
</dd>
</dl>

<dl>
<dd>

**featureIDs:** `*string` — Filter company overrides by multiple feature IDs (starting with feat_)
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — Filter company overrides by multiple company override IDs (starting with cmov_)
    
</dd>
</dl>

<dl>
<dd>

**withoutExpired:** `*bool` — Filter company overrides by whether they have not expired
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Search for company overrides by feature or company name
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.ListFeatureCompanies() -> *schematichq.ListFeatureCompaniesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListFeatureCompaniesRequest{
        FeatureID: "feature_id",
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Entitlements.ListFeatureCompanies(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**featureID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.CountFeatureCompanies() -> *schematichq.CountFeatureCompaniesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountFeatureCompaniesRequest{
        FeatureID: "feature_id",
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Entitlements.CountFeatureCompanies(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**featureID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.ListFeatureUsage() -> *schematichq.ListFeatureUsageResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListFeatureUsageRequest{
        CompanyID: schematichq.String(
            "company_id",
        ),
        IncludeUsageAggregation: schematichq.Bool(
            true,
        ),
        Q: schematichq.String(
            "q",
        ),
        WithoutNegativeEntitlements: schematichq.Bool(
            true,
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Entitlements.ListFeatureUsage(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**companyKeys:** `map[string]string` 
    
</dd>
</dl>

<dl>
<dd>

**featureIDs:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**includeUsageAggregation:** `*bool` — Include time-bucketed usage aggregation (today, this week, this month, billing period) for credit-based entitlements. Defaults to false for performance.
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**withoutNegativeEntitlements:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.CountFeatureUsage() -> *schematichq.CountFeatureUsageResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountFeatureUsageRequest{
        CompanyID: schematichq.String(
            "company_id",
        ),
        IncludeUsageAggregation: schematichq.Bool(
            true,
        ),
        Q: schematichq.String(
            "q",
        ),
        WithoutNegativeEntitlements: schematichq.Bool(
            true,
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Entitlements.CountFeatureUsage(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**companyKeys:** `map[string]string` 
    
</dd>
</dl>

<dl>
<dd>

**featureIDs:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**includeUsageAggregation:** `*bool` — Include time-bucketed usage aggregation (today, this week, this month, billing period) for credit-based entitlements. Defaults to false for performance.
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**withoutNegativeEntitlements:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.ListFeatureUsers() -> *schematichq.ListFeatureUsersResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListFeatureUsersRequest{
        FeatureID: "feature_id",
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Entitlements.ListFeatureUsers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**featureID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.CountFeatureUsers() -> *schematichq.CountFeatureUsersResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountFeatureUsersRequest{
        FeatureID: "feature_id",
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Entitlements.CountFeatureUsers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**featureID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.ListPlanEntitlements() -> *schematichq.ListPlanEntitlementsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListPlanEntitlementsRequest{
        FeatureID: schematichq.String(
            "feature_id",
        ),
        PlanID: schematichq.String(
            "plan_id",
        ),
        Q: schematichq.String(
            "q",
        ),
        WithMeteredProducts: schematichq.Bool(
            true,
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Entitlements.ListPlanEntitlements(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**featureID:** `*string` — Filter plan entitlements by a single feature ID (starting with feat_)
    
</dd>
</dl>

<dl>
<dd>

**featureIDs:** `*string` — Filter plan entitlements by multiple feature IDs (starting with feat_)
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — Filter plan entitlements by multiple plan entitlement IDs (starting with pltl_)
    
</dd>
</dl>

<dl>
<dd>

**planID:** `*string` — Filter plan entitlements by a single plan ID (starting with plan_)
    
</dd>
</dl>

<dl>
<dd>

**planIDs:** `*string` — Filter plan entitlements by multiple plan IDs (starting with plan_)
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Search for plan entitlements by feature or company name
    
</dd>
</dl>

<dl>
<dd>

**withMeteredProducts:** `*bool` — Filter plan entitlements only with metered products
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.CreatePlanEntitlement(request) -> *schematichq.CreatePlanEntitlementResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreatePlanEntitlementRequestBody{
        FeatureID: "feature_id",
        PlanID: "plan_id",
        ValueType: schematichq.EntitlementValueTypeBoolean,
    }
client.Entitlements.CreatePlanEntitlement(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**billingProductID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**billingThreshold:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**creditConsumptionRate:** `*float64` 
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**featureID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**metricPeriod:** `*schematichq.CreatePlanEntitlementRequestBodyMetricPeriod` 
    
</dd>
</dl>

<dl>
<dd>

**metricPeriodMonthReset:** `*schematichq.CreatePlanEntitlementRequestBodyMetricPeriodMonthReset` 
    
</dd>
</dl>

<dl>
<dd>

**monthlyMeteredPriceID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**monthlyPriceTiers:** `[]*schematichq.CreatePriceTierRequestBody` 
    
</dd>
</dl>

<dl>
<dd>

**monthlyUnitPrice:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**monthlyUnitPriceDecimal:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**overageBillingProductID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**planID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**priceBehavior:** `*schematichq.EntitlementPriceBehavior` 
    
</dd>
</dl>

<dl>
<dd>

**priceTiers:** `[]*schematichq.CreatePriceTierRequestBody` — Use MonthlyPriceTiers or YearlyPriceTiers instead
    
</dd>
</dl>

<dl>
<dd>

**softLimit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**tierMode:** `*schematichq.BillingTiersMode` 
    
</dd>
</dl>

<dl>
<dd>

**valueBool:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**valueCreditID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**valueNumeric:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**valueTraitID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**valueType:** `*schematichq.EntitlementValueType` 
    
</dd>
</dl>

<dl>
<dd>

**yearlyMeteredPriceID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**yearlyPriceTiers:** `[]*schematichq.CreatePriceTierRequestBody` 
    
</dd>
</dl>

<dl>
<dd>

**yearlyUnitPrice:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**yearlyUnitPriceDecimal:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.GetPlanEntitlement(PlanEntitlementID) -> *schematichq.GetPlanEntitlementResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Entitlements.GetPlanEntitlement(
        context.TODO(),
        "plan_entitlement_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planEntitlementID:** `string` — plan_entitlement_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.UpdatePlanEntitlement(PlanEntitlementID, request) -> *schematichq.UpdatePlanEntitlementResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdatePlanEntitlementRequestBody{
        ValueType: schematichq.EntitlementValueTypeBoolean,
    }
client.Entitlements.UpdatePlanEntitlement(
        context.TODO(),
        "plan_entitlement_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planEntitlementID:** `string` — plan_entitlement_id
    
</dd>
</dl>

<dl>
<dd>

**billingProductID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**billingThreshold:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**creditConsumptionRate:** `*float64` 
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**metricPeriod:** `*schematichq.UpdatePlanEntitlementRequestBodyMetricPeriod` 
    
</dd>
</dl>

<dl>
<dd>

**metricPeriodMonthReset:** `*schematichq.UpdatePlanEntitlementRequestBodyMetricPeriodMonthReset` 
    
</dd>
</dl>

<dl>
<dd>

**monthlyMeteredPriceID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**monthlyPriceTiers:** `[]*schematichq.CreatePriceTierRequestBody` 
    
</dd>
</dl>

<dl>
<dd>

**monthlyUnitPrice:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**monthlyUnitPriceDecimal:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**overageBillingProductID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**priceBehavior:** `*schematichq.EntitlementPriceBehavior` 
    
</dd>
</dl>

<dl>
<dd>

**priceTiers:** `[]*schematichq.CreatePriceTierRequestBody` — Use MonthlyPriceTiers or YearlyPriceTiers instead
    
</dd>
</dl>

<dl>
<dd>

**softLimit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**tierMode:** `*schematichq.BillingTiersMode` 
    
</dd>
</dl>

<dl>
<dd>

**valueBool:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**valueCreditID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**valueNumeric:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**valueTraitID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**valueType:** `*schematichq.EntitlementValueType` 
    
</dd>
</dl>

<dl>
<dd>

**yearlyMeteredPriceID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**yearlyPriceTiers:** `[]*schematichq.CreatePriceTierRequestBody` 
    
</dd>
</dl>

<dl>
<dd>

**yearlyUnitPrice:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**yearlyUnitPriceDecimal:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.DeletePlanEntitlement(PlanEntitlementID) -> *schematichq.DeletePlanEntitlementResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Entitlements.DeletePlanEntitlement(
        context.TODO(),
        "plan_entitlement_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planEntitlementID:** `string` — plan_entitlement_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.CountPlanEntitlements() -> *schematichq.CountPlanEntitlementsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountPlanEntitlementsRequest{
        FeatureID: schematichq.String(
            "feature_id",
        ),
        PlanID: schematichq.String(
            "plan_id",
        ),
        Q: schematichq.String(
            "q",
        ),
        WithMeteredProducts: schematichq.Bool(
            true,
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Entitlements.CountPlanEntitlements(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**featureID:** `*string` — Filter plan entitlements by a single feature ID (starting with feat_)
    
</dd>
</dl>

<dl>
<dd>

**featureIDs:** `*string` — Filter plan entitlements by multiple feature IDs (starting with feat_)
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — Filter plan entitlements by multiple plan entitlement IDs (starting with pltl_)
    
</dd>
</dl>

<dl>
<dd>

**planID:** `*string` — Filter plan entitlements by a single plan ID (starting with plan_)
    
</dd>
</dl>

<dl>
<dd>

**planIDs:** `*string` — Filter plan entitlements by multiple plan IDs (starting with plan_)
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Search for plan entitlements by feature or company name
    
</dd>
</dl>

<dl>
<dd>

**withMeteredProducts:** `*bool` — Filter plan entitlements only with metered products
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.DuplicatePlanEntitlements(request) -> *schematichq.DuplicatePlanEntitlementsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.DuplicatePlanEntitlementsRequestBody{
        SourcePlanID: "source_plan_id",
        TargetPlanID: "target_plan_id",
    }
client.Entitlements.DuplicatePlanEntitlements(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sourcePlanID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**targetPlanID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entitlements.GetFeatureUsageByCompany() -> *schematichq.GetFeatureUsageByCompanyResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.GetFeatureUsageByCompanyRequest{
        Keys: map[string]string{
            "keys": "keys",
        },
    }
client.Entitlements.GetFeatureUsageByCompany(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**keys:** `map[string]string` — Key/value pairs
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## plans
<details><summary><code>client.Plans.UpdateCompanyPlans(CompanyPlanID, request) -> *schematichq.UpdateCompanyPlansResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdateCompanyPlansRequestBody{
        AddOnIDs: []string{
            "add_on_ids",
        },
    }
client.Plans.UpdateCompanyPlans(
        context.TODO(),
        "company_plan_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyPlanID:** `string` — company_plan_id
    
</dd>
</dl>

<dl>
<dd>

**addOnIDs:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**basePlanID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Plans.ListPlans() -> *schematichq.ListPlansResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListPlansRequest{
        CompanyID: schematichq.String(
            "company_id",
        ),
        ForFallbackPlan: schematichq.Bool(
            true,
        ),
        ForInitialPlan: schematichq.Bool(
            true,
        ),
        ForTrialExpiryPlan: schematichq.Bool(
            true,
        ),
        HasProductID: schematichq.Bool(
            true,
        ),
        PlanType: schematichq.PlanTypePlan.Ptr(),
        Q: schematichq.String(
            "q",
        ),
        WithoutEntitlementFor: schematichq.String(
            "without_entitlement_for",
        ),
        WithoutPaidProductID: schematichq.Bool(
            true,
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Plans.ListPlans(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**forFallbackPlan:** `*bool` — Filter for plans valid as fallback plans (not linked to billing)
    
</dd>
</dl>

<dl>
<dd>

**forInitialPlan:** `*bool` — Filter for plans valid as initial plans (not linked to billing, free, or auto-cancelling trial)
    
</dd>
</dl>

<dl>
<dd>

**forTrialExpiryPlan:** `*bool` — Filter for plans valid as trial expiry plans (not linked to billing or free)
    
</dd>
</dl>

<dl>
<dd>

**hasProductID:** `*bool` — Filter out plans that do not have a billing product ID
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**planType:** `*schematichq.PlanType` — Filter by plan type
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**withoutEntitlementFor:** `*string` — Filter out plans that already have a plan entitlement for the specified feature ID
    
</dd>
</dl>

<dl>
<dd>

**withoutPaidProductID:** `*bool` — Filter out plans that have a paid billing product ID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Plans.CreatePlan(request) -> *schematichq.CreatePlanResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreatePlanRequestBody{
        Description: "description",
        Name: "name",
        PlanType: schematichq.PlanTypePlan,
    }
client.Plans.CreatePlan(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.CreatePlanRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Plans.GetPlan(PlanID) -> *schematichq.GetPlanResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Plans.GetPlan(
        context.TODO(),
        "plan_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planID:** `string` — plan_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Plans.UpdatePlan(PlanID, request) -> *schematichq.UpdatePlanResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdatePlanRequestBody{
        Name: "name",
    }
client.Plans.UpdatePlan(
        context.TODO(),
        "plan_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planID:** `string` — plan_id
    
</dd>
</dl>

<dl>
<dd>

**request:** `*schematichq.UpdatePlanRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Plans.DeletePlan(PlanID) -> *schematichq.DeletePlanResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Plans.DeletePlan(
        context.TODO(),
        "plan_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planID:** `string` — plan_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Plans.UpsertBillingProductPlan(PlanID, request) -> *schematichq.UpsertBillingProductPlanResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpsertBillingProductRequestBody{
        ChargeType: schematichq.ChargeTypeFree,
        IsTrialable: true,
    }
client.Plans.UpsertBillingProductPlan(
        context.TODO(),
        "plan_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planID:** `string` — plan_id
    
</dd>
</dl>

<dl>
<dd>

**request:** `*schematichq.UpsertBillingProductRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Plans.CountPlans() -> *schematichq.CountPlansResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountPlansRequest{
        CompanyID: schematichq.String(
            "company_id",
        ),
        ForFallbackPlan: schematichq.Bool(
            true,
        ),
        ForInitialPlan: schematichq.Bool(
            true,
        ),
        ForTrialExpiryPlan: schematichq.Bool(
            true,
        ),
        HasProductID: schematichq.Bool(
            true,
        ),
        PlanType: schematichq.PlanTypePlan.Ptr(),
        Q: schematichq.String(
            "q",
        ),
        WithoutEntitlementFor: schematichq.String(
            "without_entitlement_for",
        ),
        WithoutPaidProductID: schematichq.Bool(
            true,
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Plans.CountPlans(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**forFallbackPlan:** `*bool` — Filter for plans valid as fallback plans (not linked to billing)
    
</dd>
</dl>

<dl>
<dd>

**forInitialPlan:** `*bool` — Filter for plans valid as initial plans (not linked to billing, free, or auto-cancelling trial)
    
</dd>
</dl>

<dl>
<dd>

**forTrialExpiryPlan:** `*bool` — Filter for plans valid as trial expiry plans (not linked to billing or free)
    
</dd>
</dl>

<dl>
<dd>

**hasProductID:** `*bool` — Filter out plans that do not have a billing product ID
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**planType:** `*schematichq.PlanType` — Filter by plan type
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**withoutEntitlementFor:** `*string` — Filter out plans that already have a plan entitlement for the specified feature ID
    
</dd>
</dl>

<dl>
<dd>

**withoutPaidProductID:** `*bool` — Filter out plans that have a paid billing product ID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Plans.ListPlanIssues() -> *schematichq.ListPlanIssuesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListPlanIssuesRequest{
        PlanID: "plan_id",
    }
client.Plans.ListPlanIssues(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## components
<details><summary><code>client.Components.ListComponents() -> *schematichq.ListComponentsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListComponentsRequest{
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Components.ListComponents(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Components.CreateComponent(request) -> *schematichq.CreateComponentResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateComponentRequestBody{
        EntityType: schematichq.ComponentEntityTypeBilling,
        Name: "name",
    }
client.Components.CreateComponent(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ast:** `map[string]float64` 
    
</dd>
</dl>

<dl>
<dd>

**entityType:** `*schematichq.ComponentEntityType` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Components.GetComponent(ComponentID) -> *schematichq.GetComponentResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Components.GetComponent(
        context.TODO(),
        "component_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**componentID:** `string` — component_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Components.UpdateComponent(ComponentID, request) -> *schematichq.UpdateComponentResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdateComponentRequestBody{}
client.Components.UpdateComponent(
        context.TODO(),
        "component_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**componentID:** `string` — component_id
    
</dd>
</dl>

<dl>
<dd>

**ast:** `map[string]float64` 
    
</dd>
</dl>

<dl>
<dd>

**entityType:** `*schematichq.ComponentEntityType` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**state:** `*schematichq.ComponentState` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Components.DeleteComponent(ComponentID) -> *schematichq.DeleteComponentResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Components.DeleteComponent(
        context.TODO(),
        "component_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**componentID:** `string` — component_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Components.CountComponents() -> *schematichq.CountComponentsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountComponentsRequest{
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Components.CountComponents(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Components.PreviewComponentData() -> *schematichq.PreviewComponentDataResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.PreviewComponentDataRequest{
        CompanyID: schematichq.String(
            "company_id",
        ),
        ComponentID: schematichq.String(
            "component_id",
        ),
    }
client.Components.PreviewComponentData(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**componentID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## dataexports
<details><summary><code>client.Dataexports.CreateDataExport(request) -> *schematichq.CreateDataExportResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateDataExportRequestBody{
        ExportType: schematichq.DataExportType(
            "company-feature-usage",
        ),
        Metadata: "metadata",
        OutputFileType: schematichq.DataExportOutputFileType(
            "csv",
        ),
    }
client.Dataexports.CreateDataExport(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**exportType:** `schematichq.DataExportType` 
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**outputFileType:** `schematichq.DataExportOutputFileType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Dataexports.GetDataExportArtifact(DataExportID) -> string</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Dataexports.GetDataExportArtifact(
        context.TODO(),
        "data_export_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**dataExportID:** `string` — data_export_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## events
<details><summary><code>client.Events.CreateEventBatch(request) -> *schematichq.CreateEventBatchResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateEventBatchRequestBody{
        Events: []*schematichq.CreateEventRequestBody{
            &schematichq.CreateEventRequestBody{
                EventType: schematichq.EventTypeFlagCheck,
            },
        },
    }
client.Events.CreateEventBatch(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**events:** `[]*schematichq.CreateEventRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Events.GetEventSummaries() -> *schematichq.GetEventSummariesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.GetEventSummariesRequest{
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Events.GetEventSummaries(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**eventSubtypes:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Events.ListEvents() -> *schematichq.ListEventsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListEventsRequest{
        CompanyID: schematichq.String(
            "company_id",
        ),
        EventSubtype: schematichq.String(
            "event_subtype",
        ),
        FlagID: schematichq.String(
            "flag_id",
        ),
        UserID: schematichq.String(
            "user_id",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Events.ListEvents(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**eventSubtype:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**eventTypes:** `*schematichq.EventType` 
    
</dd>
</dl>

<dl>
<dd>

**flagID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Events.CreateEvent(request) -> *schematichq.CreateEventResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateEventRequestBody{
        EventType: schematichq.EventTypeFlagCheck,
    }
client.Events.CreateEvent(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.CreateEventRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Events.GetEvent(EventID) -> *schematichq.GetEventResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Events.GetEvent(
        context.TODO(),
        "event_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**eventID:** `string` — event_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Events.GetSegmentIntegrationStatus() -> *schematichq.GetSegmentIntegrationStatusResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Events.GetSegmentIntegrationStatus(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## features
<details><summary><code>client.Features.ListFeatures() -> *schematichq.ListFeaturesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListFeaturesRequest{
        Q: schematichq.String(
            "q",
        ),
        WithoutCompanyOverrideFor: schematichq.String(
            "without_company_override_for",
        ),
        WithoutPlanEntitlementFor: schematichq.String(
            "without_plan_entitlement_for",
        ),
        BooleanRequireEvent: schematichq.Bool(
            true,
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Features.ListFeatures(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Search by feature name or ID
    
</dd>
</dl>

<dl>
<dd>

**withoutCompanyOverrideFor:** `*string` — Filter out features that already have a company override for the specified company ID
    
</dd>
</dl>

<dl>
<dd>

**withoutPlanEntitlementFor:** `*string` — Filter out features that already have a plan entitlement for the specified plan ID
    
</dd>
</dl>

<dl>
<dd>

**featureType:** `*schematichq.FeatureType` — Filter by one or more feature types (boolean, event, trait)
    
</dd>
</dl>

<dl>
<dd>

**booleanRequireEvent:** `*bool` — Only return boolean features if there is an associated event. Automatically includes boolean in the feature types filter.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Features.CreateFeature(request) -> *schematichq.CreateFeatureResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateFeatureRequestBody{
        Description: "description",
        FeatureType: schematichq.FeatureTypeBoolean,
        Name: "name",
    }
client.Features.CreateFeature(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**description:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**eventSubtype:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**featureType:** `*schematichq.FeatureType` 
    
</dd>
</dl>

<dl>
<dd>

**flag:** `*schematichq.CreateOrUpdateFlagRequestBody` 
    
</dd>
</dl>

<dl>
<dd>

**icon:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**lifecyclePhase:** `*schematichq.FeatureLifecyclePhase` 
    
</dd>
</dl>

<dl>
<dd>

**maintainerID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**pluralName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**singularName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**traitID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Features.GetFeature(FeatureID) -> *schematichq.GetFeatureResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Features.GetFeature(
        context.TODO(),
        "feature_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**featureID:** `string` — feature_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Features.UpdateFeature(FeatureID, request) -> *schematichq.UpdateFeatureResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdateFeatureRequestBody{}
client.Features.UpdateFeature(
        context.TODO(),
        "feature_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**featureID:** `string` — feature_id
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**eventSubtype:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**featureType:** `*schematichq.FeatureType` 
    
</dd>
</dl>

<dl>
<dd>

**flag:** `*schematichq.CreateOrUpdateFlagRequestBody` 
    
</dd>
</dl>

<dl>
<dd>

**icon:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**lifecyclePhase:** `*schematichq.FeatureLifecyclePhase` 
    
</dd>
</dl>

<dl>
<dd>

**maintainerID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**pluralName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**singularName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**traitID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Features.DeleteFeature(FeatureID) -> *schematichq.DeleteFeatureResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Features.DeleteFeature(
        context.TODO(),
        "feature_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**featureID:** `string` — feature_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Features.CountFeatures() -> *schematichq.CountFeaturesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountFeaturesRequest{
        Q: schematichq.String(
            "q",
        ),
        WithoutCompanyOverrideFor: schematichq.String(
            "without_company_override_for",
        ),
        WithoutPlanEntitlementFor: schematichq.String(
            "without_plan_entitlement_for",
        ),
        BooleanRequireEvent: schematichq.Bool(
            true,
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Features.CountFeatures(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Search by feature name or ID
    
</dd>
</dl>

<dl>
<dd>

**withoutCompanyOverrideFor:** `*string` — Filter out features that already have a company override for the specified company ID
    
</dd>
</dl>

<dl>
<dd>

**withoutPlanEntitlementFor:** `*string` — Filter out features that already have a plan entitlement for the specified plan ID
    
</dd>
</dl>

<dl>
<dd>

**featureType:** `*schematichq.FeatureType` — Filter by one or more feature types (boolean, event, trait)
    
</dd>
</dl>

<dl>
<dd>

**booleanRequireEvent:** `*bool` — Only return boolean features if there is an associated event. Automatically includes boolean in the feature types filter.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Features.ListFlags() -> *schematichq.ListFlagsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListFlagsRequest{
        FeatureID: schematichq.String(
            "feature_id",
        ),
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Features.ListFlags(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**featureID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Search by flag name, key, or ID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Features.CreateFlag(request) -> *schematichq.CreateFlagResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateFlagRequestBody{
        DefaultValue: true,
        Description: "description",
        FlagType: schematichq.FlagType(
            "boolean",
        ),
        Key: "key",
        Name: "name",
    }
client.Features.CreateFlag(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.CreateFlagRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Features.GetFlag(FlagID) -> *schematichq.GetFlagResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Features.GetFlag(
        context.TODO(),
        "flag_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**flagID:** `string` — flag_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Features.UpdateFlag(FlagID, request) -> *schematichq.UpdateFlagResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateFlagRequestBody{
        DefaultValue: true,
        Description: "description",
        FlagType: schematichq.FlagType(
            "boolean",
        ),
        Key: "key",
        Name: "name",
    }
client.Features.UpdateFlag(
        context.TODO(),
        "flag_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**flagID:** `string` — flag_id
    
</dd>
</dl>

<dl>
<dd>

**request:** `*schematichq.CreateFlagRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Features.DeleteFlag(FlagID) -> *schematichq.DeleteFlagResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Features.DeleteFlag(
        context.TODO(),
        "flag_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**flagID:** `string` — flag_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Features.UpdateFlagRules(FlagID, request) -> *schematichq.UpdateFlagRulesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdateFlagRulesRequestBody{
        Rules: []*schematichq.CreateOrUpdateRuleRequestBody{
            &schematichq.CreateOrUpdateRuleRequestBody{
                ConditionGroups: []*schematichq.CreateOrUpdateConditionGroupRequestBody{
                    &schematichq.CreateOrUpdateConditionGroupRequestBody{
                        Conditions: []*schematichq.CreateOrUpdateConditionRequestBody{
                            &schematichq.CreateOrUpdateConditionRequestBody{
                                ConditionType: schematichq.CreateOrUpdateConditionRequestBodyConditionTypeCompany,
                                Operator: schematichq.CreateOrUpdateConditionRequestBodyOperatorEq,
                                ResourceIDs: []string{
                                    "resource_ids",
                                },
                            },
                        },
                    },
                },
                Conditions: []*schematichq.CreateOrUpdateConditionRequestBody{
                    &schematichq.CreateOrUpdateConditionRequestBody{
                        ConditionType: schematichq.CreateOrUpdateConditionRequestBodyConditionTypeCompany,
                        Operator: schematichq.CreateOrUpdateConditionRequestBodyOperatorEq,
                        ResourceIDs: []string{
                            "resource_ids",
                        },
                    },
                },
                Name: "name",
                Priority: 1,
                Value: true,
            },
        },
    }
client.Features.UpdateFlagRules(
        context.TODO(),
        "flag_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**flagID:** `string` — flag_id
    
</dd>
</dl>

<dl>
<dd>

**rules:** `[]*schematichq.CreateOrUpdateRuleRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Features.CheckFlag(Key, request) -> *schematichq.CheckFlagResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CheckFlagRequestBody{}
client.Features.CheckFlag(
        context.TODO(),
        "key",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — key
    
</dd>
</dl>

<dl>
<dd>

**request:** `*schematichq.CheckFlagRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Features.CheckFlags(request) -> *schematichq.CheckFlagsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CheckFlagRequestBody{}
client.Features.CheckFlags(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*schematichq.CheckFlagRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Features.CheckFlagsBulk(request) -> *schematichq.CheckFlagsBulkResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CheckFlagsBulkRequestBody{
        Contexts: []*schematichq.CheckFlagRequestBody{
            &schematichq.CheckFlagRequestBody{},
        },
    }
client.Features.CheckFlagsBulk(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**contexts:** `[]*schematichq.CheckFlagRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Features.CountFlags() -> *schematichq.CountFlagsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountFlagsRequest{
        FeatureID: schematichq.String(
            "feature_id",
        ),
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Features.CountFlags(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**featureID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Search by flag name, key, or ID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## planbundle
<details><summary><code>client.Planbundle.CreatePlanBundle(request) -> *schematichq.CreatePlanBundleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreatePlanBundleRequestBody{
        Entitlements: []*schematichq.PlanBundleEntitlementRequestBody{
            &schematichq.PlanBundleEntitlementRequestBody{
                Action: schematichq.PlanBundleActionCreate,
            },
        },
    }
client.Planbundle.CreatePlanBundle(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**billingProduct:** `*schematichq.UpsertBillingProductRequestBody` 
    
</dd>
</dl>

<dl>
<dd>

**creditGrants:** `[]*schematichq.PlanBundleCreditGrantRequestBody` 
    
</dd>
</dl>

<dl>
<dd>

**entitlements:** `[]*schematichq.PlanBundleEntitlementRequestBody` 
    
</dd>
</dl>

<dl>
<dd>

**plan:** `*schematichq.CreatePlanRequestBody` 
    
</dd>
</dl>

<dl>
<dd>

**traits:** `[]*schematichq.UpdatePlanTraitTraitRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Planbundle.UpdatePlanBundle(PlanBundleID, request) -> *schematichq.UpdatePlanBundleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdatePlanBundleRequestBody{
        Entitlements: []*schematichq.PlanBundleEntitlementRequestBody{
            &schematichq.PlanBundleEntitlementRequestBody{
                Action: schematichq.PlanBundleActionCreate,
            },
        },
    }
client.Planbundle.UpdatePlanBundle(
        context.TODO(),
        "plan_bundle_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planBundleID:** `string` — plan_bundle_id
    
</dd>
</dl>

<dl>
<dd>

**billingProduct:** `*schematichq.UpsertBillingProductRequestBody` 
    
</dd>
</dl>

<dl>
<dd>

**creditGrants:** `[]*schematichq.PlanBundleCreditGrantRequestBody` 
    
</dd>
</dl>

<dl>
<dd>

**entitlements:** `[]*schematichq.PlanBundleEntitlementRequestBody` 
    
</dd>
</dl>

<dl>
<dd>

**plan:** `*schematichq.UpdatePlanRequestBody` 
    
</dd>
</dl>

<dl>
<dd>

**traits:** `[]*schematichq.UpdatePlanTraitTraitRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## plangroups
<details><summary><code>client.Plangroups.GetPlanGroup() -> *schematichq.GetPlanGroupResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Plangroups.GetPlanGroup(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Plangroups.CreatePlanGroup(request) -> *schematichq.CreatePlanGroupResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreatePlanGroupRequestBody{
        AddOnIDs: []string{
            "add_on_ids",
        },
        CheckoutCollectAddress: true,
        CheckoutCollectEmail: true,
        CheckoutCollectPhone: true,
        EnableTaxCollection: true,
        OrderedAddOns: []*schematichq.OrderedPlansInGroup{
            &schematichq.OrderedPlansInGroup{
                PlanID: "plan_id",
            },
        },
        OrderedBundleList: []*schematichq.PlanGroupBundleOrder{
            &schematichq.PlanGroupBundleOrder{
                BundleID: "bundleId",
            },
        },
        OrderedPlans: []*schematichq.OrderedPlansInGroup{
            &schematichq.OrderedPlansInGroup{
                PlanID: "plan_id",
            },
        },
        PreventDowngradesWhenOverLimit: true,
        PreventSelfServiceDowngrade: true,
        ProrationBehavior: schematichq.ProrationBehaviorCreateProrations,
        ShowAsMonthlyPrices: true,
        ShowCredits: true,
        ShowPeriodToggle: true,
        ShowZeroPriceAsFree: true,
        SyncCustomerBillingDetails: true,
    }
client.Plangroups.CreatePlanGroup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**addOnCompatibilities:** `[]*schematichq.CompatiblePlans` 
    
</dd>
</dl>

<dl>
<dd>

**addOnIDs:** `[]string` — Use OrderedAddOns instead
    
</dd>
</dl>

<dl>
<dd>

**checkoutCollectAddress:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**checkoutCollectEmail:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**checkoutCollectPhone:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**customPlanConfig:** `*schematichq.CustomPlanConfig` 
    
</dd>
</dl>

<dl>
<dd>

**customPlanID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**enableTaxCollection:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**fallbackPlanID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**initialPlanID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**initialPlanPriceID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**orderedAddOns:** `[]*schematichq.OrderedPlansInGroup` 
    
</dd>
</dl>

<dl>
<dd>

**orderedBundleList:** `[]*schematichq.PlanGroupBundleOrder` 
    
</dd>
</dl>

<dl>
<dd>

**orderedPlans:** `[]*schematichq.OrderedPlansInGroup` 
    
</dd>
</dl>

<dl>
<dd>

**preventDowngradesWhenOverLimit:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**preventSelfServiceDowngrade:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**preventSelfServiceDowngradeButtonText:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**preventSelfServiceDowngradeURL:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**prorationBehavior:** `*schematichq.ProrationBehavior` 
    
</dd>
</dl>

<dl>
<dd>

**showAsMonthlyPrices:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**showCredits:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**showPeriodToggle:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**showZeroPriceAsFree:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**syncCustomerBillingDetails:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**trialDays:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**trialExpiryPlanID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**trialExpiryPlanPriceID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**trialPaymentMethodRequired:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Plangroups.UpdatePlanGroup(PlanGroupID, request) -> *schematichq.UpdatePlanGroupResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdatePlanGroupRequestBody{
        AddOnIDs: []string{
            "add_on_ids",
        },
        CheckoutCollectAddress: true,
        CheckoutCollectEmail: true,
        CheckoutCollectPhone: true,
        EnableTaxCollection: true,
        OrderedAddOns: []*schematichq.OrderedPlansInGroup{
            &schematichq.OrderedPlansInGroup{
                PlanID: "plan_id",
            },
        },
        OrderedBundleList: []*schematichq.PlanGroupBundleOrder{
            &schematichq.PlanGroupBundleOrder{
                BundleID: "bundleId",
            },
        },
        OrderedPlans: []*schematichq.OrderedPlansInGroup{
            &schematichq.OrderedPlansInGroup{
                PlanID: "plan_id",
            },
        },
        PreventDowngradesWhenOverLimit: true,
        PreventSelfServiceDowngrade: true,
        ProrationBehavior: schematichq.ProrationBehaviorCreateProrations,
        ShowAsMonthlyPrices: true,
        ShowCredits: true,
        ShowPeriodToggle: true,
        ShowZeroPriceAsFree: true,
        SyncCustomerBillingDetails: true,
    }
client.Plangroups.UpdatePlanGroup(
        context.TODO(),
        "plan_group_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planGroupID:** `string` — plan_group_id
    
</dd>
</dl>

<dl>
<dd>

**addOnCompatibilities:** `[]*schematichq.CompatiblePlans` 
    
</dd>
</dl>

<dl>
<dd>

**addOnIDs:** `[]string` — Use OrderedAddOns instead
    
</dd>
</dl>

<dl>
<dd>

**checkoutCollectAddress:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**checkoutCollectEmail:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**checkoutCollectPhone:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**customPlanConfig:** `*schematichq.CustomPlanConfig` 
    
</dd>
</dl>

<dl>
<dd>

**customPlanID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**enableTaxCollection:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**fallbackPlanID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**initialPlanID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**initialPlanPriceID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**orderedAddOns:** `[]*schematichq.OrderedPlansInGroup` 
    
</dd>
</dl>

<dl>
<dd>

**orderedBundleList:** `[]*schematichq.PlanGroupBundleOrder` 
    
</dd>
</dl>

<dl>
<dd>

**orderedPlans:** `[]*schematichq.OrderedPlansInGroup` 
    
</dd>
</dl>

<dl>
<dd>

**preventDowngradesWhenOverLimit:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**preventSelfServiceDowngrade:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**preventSelfServiceDowngradeButtonText:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**preventSelfServiceDowngradeURL:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**prorationBehavior:** `*schematichq.ProrationBehavior` 
    
</dd>
</dl>

<dl>
<dd>

**showAsMonthlyPrices:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**showCredits:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**showPeriodToggle:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**showZeroPriceAsFree:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**syncCustomerBillingDetails:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**trialDays:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**trialExpiryPlanID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**trialExpiryPlanPriceID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**trialPaymentMethodRequired:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## accesstokens
<details><summary><code>client.Accesstokens.IssueTemporaryAccessToken(request) -> *schematichq.IssueTemporaryAccessTokenResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.IssueTemporaryAccessTokenRequestBody{
        Lookup: map[string]string{
            "key": "value",
        },
        ResourceType: schematichq.TemporaryAccessTokenResourceType(
            "company",
        ),
    }
client.Accesstokens.IssueTemporaryAccessToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**lookup:** `map[string]string` 
    
</dd>
</dl>

<dl>
<dd>

**resourceType:** `schematichq.TemporaryAccessTokenResourceType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## webhooks
<details><summary><code>client.Webhooks.ListWebhookEvents() -> *schematichq.ListWebhookEventsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListWebhookEventsRequest{
        Q: schematichq.String(
            "q",
        ),
        WebhookID: schematichq.String(
            "webhook_id",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Webhooks.ListWebhookEvents(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**webhookID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.GetWebhookEvent(WebhookEventID) -> *schematichq.GetWebhookEventResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Webhooks.GetWebhookEvent(
        context.TODO(),
        "webhook_event_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookEventID:** `string` — webhook_event_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.CountWebhookEvents() -> *schematichq.CountWebhookEventsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountWebhookEventsRequest{
        Q: schematichq.String(
            "q",
        ),
        WebhookID: schematichq.String(
            "webhook_id",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Webhooks.CountWebhookEvents(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**webhookID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.ListWebhooks() -> *schematichq.ListWebhooksResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.ListWebhooksRequest{
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Webhooks.ListWebhooks(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.CreateWebhook(request) -> *schematichq.CreateWebhookResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CreateWebhookRequestBody{
        Name: "name",
        RequestTypes: []schematichq.WebhookRequestType{
            schematichq.WebhookRequestTypeSubscriptionTrialEnded,
        },
        URL: "url",
    }
client.Webhooks.CreateWebhook(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**creditTriggerConfigs:** `[]*schematichq.CreditTriggerConfig` 
    
</dd>
</dl>

<dl>
<dd>

**entitlementTriggerConfigs:** `[]*schematichq.EntitlementTriggerConfig` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**requestTypes:** `[]*schematichq.WebhookRequestType` 
    
</dd>
</dl>

<dl>
<dd>

**url:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.GetWebhook(WebhookID) -> *schematichq.GetWebhookResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Webhooks.GetWebhook(
        context.TODO(),
        "webhook_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookID:** `string` — webhook_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.UpdateWebhook(WebhookID, request) -> *schematichq.UpdateWebhookResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.UpdateWebhookRequestBody{}
client.Webhooks.UpdateWebhook(
        context.TODO(),
        "webhook_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookID:** `string` — webhook_id
    
</dd>
</dl>

<dl>
<dd>

**creditTriggerConfigs:** `[]*schematichq.CreditTriggerConfig` 
    
</dd>
</dl>

<dl>
<dd>

**entitlementTriggerConfigs:** `[]*schematichq.EntitlementTriggerConfig` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**requestTypes:** `[]*schematichq.WebhookRequestType` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*schematichq.WebhookStatus` 
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.DeleteWebhook(WebhookID) -> *schematichq.DeleteWebhookResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Webhooks.DeleteWebhook(
        context.TODO(),
        "webhook_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookID:** `string` — webhook_id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.CountWebhooks() -> *schematichq.CountWebhooksResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &schematichq.CountWebhooksRequest{
        Q: schematichq.String(
            "q",
        ),
        Limit: schematichq.Int(
            1,
        ),
        Offset: schematichq.Int(
            1,
        ),
    }
client.Webhooks.CountWebhooks(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page limit (default 100)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Page offset (default 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>
