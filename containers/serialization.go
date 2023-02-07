package containers

type JSONSerializer interface {
	// ToJSON output the JSON representation of container's element
	ToJSON() ([]byte, error)
	// MarshalJSON @implements json.Marshaler
	MarshalJSON() ([]byte, error)
}

// JSONDeserializer provides JSON deserialization
type JSONDeserializer interface {
	// FromJSON populates containers's elements from the input JSON representation.
	FromJSON([]byte) error
	// UnmarshalJSON @implements json.Unmarshaler
	UnmarshalJSON([]byte) error
}
