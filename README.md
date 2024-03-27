# schematic-go

## Installation and Setup

1. Install the Go library:

```bash
go get github.com/SchematicHQ/schematic-go
```

2. [Issue an API key](https://docs.schematichq.com/quickstart#create-an-api-key) for the appropriate environment using the [Schematic app](https://app.schematichq.com/settings/api-keys). Be sure to capture the secret key when you issue the API key; you'll only see this key once, and this is what you'll use with the Schematic Go library.

3. Using this secret key, initialize a client in your Go application:

```go
import (
  "os"

  "github.com/SchematicHQ/schematic-go"
)

func main() {
  apiKey := os.Getenv("SCHEMATIC_API_KEY")
  client := schematic.NewClient(apiKey)
  defer client.Close()
}
```

## Usage examples

### Creating and updating companies

Because you use your own identifiers to identify companies, rather than a Schematic company ID, creating and updating companies are both done via the same upsert operation:

```go
func main() {
  client := schematic.NewClient(os.Getenv("SCHEMATIC_API_KEY"))
  defer client.Close()

  body := schematic.NewUpsertCompanyRequestBody(map[string]any{
    "id": "your-company-id",
  })
  body.SetName("Acme Widgets, Inc.")
  body.SetTraits(map[string]any{
    "city":       "Atlanta",
    "high_score": 25,
    "is_active":  true,
  })

  resp, r, err := client.API().CompaniesAPI.CreateCompany(context.Background()).UpsertCompanyRequestBody(*body).Execute()
}
```

You can define any number of company keys; these are used to address the company in the future, for example by updating the company's traits or checking a flag for the company.
You can also define any number of company traits; these can then be used as targeting parameters.


### Creating and updating users

Because you use your own identifiers to identify users, rather than a Schematic user ID, creating and updating users are both done via the same upsert operation:

```go
func main() {
  client := schematic.NewClient(os.Getenv("SCHEMATIC_API_KEY"))
  defer client.Close()

  companyKeys := map[string]any{
    "id": "your-company-id",
  }
  userKeys := map[string]any{
    "email":   "wcoyote@acme.net",
    "user-id": "your-user-id",
  }
  body := schematic.NewUpsertUserRequestBody(companyKeys, userKeys)
  body.SetName("Wile E. Coyote")
  body.SetTraits(map[string]any{
    "city":        "Atlanta",
    "login_count": 24,
    "is_staff":    false,
  })

  resp, r, err := client.API().CompaniesAPI.CreateUser(context.Background()).UpsertUserRequestBody(*body).Execute()
}
```

You can define any number of user keys; these are used to address the user in the future, for example by updating the user's traits or checking a flag for the user.
You can also define any number of user traits; these can then be used as targeting parameters.

### Checking flags

When checking a flag, you'll provide keys for a company and/or keys for a user. You can also provide no keys at all, in which case you'll get the default value for the flag.

You can check all flags at once:

```go
func main() {
  client := schematic.NewClient(os.Getenv("SCHEMATIC_API_KEY"))
  defer client.Close()

  flagContext := schematic.CheckFlagRequestBody{
    Company: map[string]any{
      "id": "your-company-id",
    },
    User: map[string]any{
      "email":   "wcoyote@acme.net",
      "user-id": "your-user-id",
    },
  }

  resp, r, err := client.API().FeaturesAPI.CheckFlags(ctx).CheckFlagRequestBody(flagContext).Execute()
}
```

Or, you can check a single flag using its key:

```go
func main() {
  client := schematic.NewClient(os.Getenv("SCHEMATIC_API_KEY"))
  defer client.Close()

  flagContext := schematic.CheckFlagRequestBody{
    Company: map[string]any{
      "id": "your-company-id",
    },
    User: map[string]any{
      "email":   "wcoyote@acme.net",
      "user-id": "your-user-id",
    },
  }

  resp, r, err := client.API().FeaturesAPI.CheckFlag(ctx, "some-flag-key").CheckFlagRequestBody(flagContext).Execute()
}
```

### Sending identify events

Create or update users and companies using identify events.

```go
func main() {
  client := schematic.NewClient(os.Getenv("SCHEMATIC_API_KEY"))
  defer client.Close()

  eventBody := NewEventBodyIdentify(map[string]any{
    "email":   "wcoyote@acme.net",
    "user-id": "your-user-id",
  })
  eventBody.SetCompany(map[string]any{
    "id": "your-company-id",
  })
  eventBody.SetName("Wile E. Coyote")
  eventBody.SetTraits(map[string]any{
    "city":        "Atlanta",
    "login_count": 24,
    "is_staff":    false,
  })

  err := client.Identify(context.Background(), eventBody)
}
```

### Sending track events

Track activity in your application using track events; these events can later be used to produce metrics for targeting.

```go
func main() {
  client := schematic.NewClient(os.Getenv("SCHEMATIC_API_KEY"))
  defer client.Close()

  eventBody := NewEventBodyTrack("some-action")
  eventBody.SetUser(map[string]any{
    "email":   "wcoyote@acme.net",
    "user-id": "your-user-id",
  })
  eventBody.SetCompany(map[string]any{
    "id": "your-company-id",
  })

  err := client.Track(context.Background(), eventBody)
}
```

### Other operations

The Schematic API supports many operations beyond these, accessible via `client.API()`. See the [API submodule readme](https://github.com/SchematicHQ/schematic-go/tree/main/api#readme) for a full list and documentation of supported operations.
