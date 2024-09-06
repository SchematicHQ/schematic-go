// This file was auto-generated by Fern from our API Definition.

package schematichq

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/schematichq/schematic-go/core"
)

type IssueTemporaryAccessTokenRequestBody struct {
	Lookup       map[string]string `json:"lookup,omitempty" url:"-"`
	ResourceType string            `json:"resource_type" url:"-"`
}

type IssueTemporaryAccessTokenResponse struct {
	Data *IssueTemporaryAccessTokenResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (i *IssueTemporaryAccessTokenResponse) GetExtraProperties() map[string]interface{} {
	return i.extraProperties
}

func (i *IssueTemporaryAccessTokenResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler IssueTemporaryAccessTokenResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = IssueTemporaryAccessTokenResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties

	i._rawJSON = json.RawMessage(data)
	return nil
}

func (i *IssueTemporaryAccessTokenResponse) String() string {
	if len(i._rawJSON) > 0 {
		if value, err := core.StringifyJSON(i._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}