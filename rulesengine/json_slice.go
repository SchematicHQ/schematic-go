package rulesengine

import "encoding/json"

// JSONSlice serializes as an empty array rather than null when nil.
//
// Two consumers depend on that. Strictly-typed SDK clients reject `null` for a
// list-typed field, and the WebAssembly engine's Rust types deserialize a list
// with serde's #[serde(default)], which supplies a value for an absent key but
// not for one explicitly set to null -- `"rules": null` fails to deserialize
// into Vec<Rule>. Declaring wire slices as JSONSlice keeps both happy without
// each caller having to sanitize its payload.
type JSONSlice[T any] []T

func NewJSONSlice[T any](slice []T) JSONSlice[T] {
	if slice == nil {
		return JSONSlice[T]{}
	}
	return JSONSlice[T](slice)
}

func (s JSONSlice[T]) MarshalJSON() ([]byte, error) {
	if s == nil {
		return json.Marshal([]T{})
	}
	return json.Marshal([]T(s))
}

func (s *JSONSlice[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*s = JSONSlice[T]{}
		return nil
	}

	var slice []T
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}

	*s = JSONSlice[T](slice)
	return nil
}

func (s JSONSlice[T]) Slice() []T {
	if s == nil {
		return []T{}
	}
	return []T(s)
}
