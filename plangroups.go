// This file was auto-generated by Fern from our API Definition.

package schematichq

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/schematichq/schematic-go/core"
)

type CreatePlanGroupRequestBody struct {
	AddOnIDs      []string `json:"add_on_ids,omitempty" url:"-"`
	DefaultPlanID *string  `json:"default_plan_id,omitempty" url:"-"`
	PlanIDs       []string `json:"plan_ids,omitempty" url:"-"`
}

type CreatePlanGroupResponse struct {
	Data *PlanGroupResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreatePlanGroupResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreatePlanGroupResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreatePlanGroupResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreatePlanGroupResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreatePlanGroupResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type GetPlanGroupResponse struct {
	Data *PlanGroupDetailResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GetPlanGroupResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetPlanGroupResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetPlanGroupResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetPlanGroupResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetPlanGroupResponse) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type UpdatePlanGroupResponse struct {
	Data *PlanGroupResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UpdatePlanGroupResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdatePlanGroupResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdatePlanGroupResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdatePlanGroupResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdatePlanGroupResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UpdatePlanGroupRequestBody struct {
	AddOnIDs      []string `json:"add_on_ids,omitempty" url:"-"`
	DefaultPlanID *string  `json:"default_plan_id,omitempty" url:"-"`
	PlanIDs       []string `json:"plan_ids,omitempty" url:"-"`
}
