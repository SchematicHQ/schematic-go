/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// EventBody - struct for EventBody
type EventBody struct {
	EventBodyFlagCheck *EventBodyFlagCheck
	EventBodyIdentify  *EventBodyIdentify
	EventBodyTrack     *EventBodyTrack
}

// EventBodyFlagCheckAsEventBody is a convenience function that returns EventBodyFlagCheck wrapped in EventBody
func EventBodyFlagCheckAsEventBody(v *EventBodyFlagCheck) EventBody {
	return EventBody{
		EventBodyFlagCheck: v,
	}
}

// EventBodyIdentifyAsEventBody is a convenience function that returns EventBodyIdentify wrapped in EventBody
func EventBodyIdentifyAsEventBody(v *EventBodyIdentify) EventBody {
	return EventBody{
		EventBodyIdentify: v,
	}
}

// EventBodyTrackAsEventBody is a convenience function that returns EventBodyTrack wrapped in EventBody
func EventBodyTrackAsEventBody(v *EventBodyTrack) EventBody {
	return EventBody{
		EventBodyTrack: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EventBody) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EventBodyFlagCheck
	err = newStrictDecoder(data).Decode(&dst.EventBodyFlagCheck)
	if err == nil {
		jsonEventBodyFlagCheck, _ := json.Marshal(dst.EventBodyFlagCheck)
		if string(jsonEventBodyFlagCheck) == "{}" { // empty struct
			dst.EventBodyFlagCheck = nil
		} else {
			match++
		}
	} else {
		dst.EventBodyFlagCheck = nil
	}

	// try to unmarshal data into EventBodyIdentify
	err = newStrictDecoder(data).Decode(&dst.EventBodyIdentify)
	if err == nil {
		jsonEventBodyIdentify, _ := json.Marshal(dst.EventBodyIdentify)
		if string(jsonEventBodyIdentify) == "{}" { // empty struct
			dst.EventBodyIdentify = nil
		} else {
			match++
		}
	} else {
		dst.EventBodyIdentify = nil
	}

	// try to unmarshal data into EventBodyTrack
	err = newStrictDecoder(data).Decode(&dst.EventBodyTrack)
	if err == nil {
		jsonEventBodyTrack, _ := json.Marshal(dst.EventBodyTrack)
		if string(jsonEventBodyTrack) == "{}" { // empty struct
			dst.EventBodyTrack = nil
		} else {
			match++
		}
	} else {
		dst.EventBodyTrack = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.EventBodyFlagCheck = nil
		dst.EventBodyIdentify = nil
		dst.EventBodyTrack = nil

		return fmt.Errorf("data matches more than one schema in oneOf(EventBody)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(EventBody)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EventBody) MarshalJSON() ([]byte, error) {
	if src.EventBodyFlagCheck != nil {
		return json.Marshal(&src.EventBodyFlagCheck)
	}

	if src.EventBodyIdentify != nil {
		return json.Marshal(&src.EventBodyIdentify)
	}

	if src.EventBodyTrack != nil {
		return json.Marshal(&src.EventBodyTrack)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EventBody) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EventBodyFlagCheck != nil {
		return obj.EventBodyFlagCheck
	}

	if obj.EventBodyIdentify != nil {
		return obj.EventBodyIdentify
	}

	if obj.EventBodyTrack != nil {
		return obj.EventBodyTrack
	}

	// all schemas are nil
	return nil
}

type NullableEventBody struct {
	value *EventBody
	isSet bool
}

func (v NullableEventBody) Get() *EventBody {
	return v.value
}

func (v *NullableEventBody) Set(val *EventBody) {
	v.value = val
	v.isSet = true
}

func (v NullableEventBody) IsSet() bool {
	return v.isSet
}

func (v *NullableEventBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventBody(val *EventBody) *NullableEventBody {
	return &NullableEventBody{value: val, isSet: true}
}

func (v NullableEventBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}