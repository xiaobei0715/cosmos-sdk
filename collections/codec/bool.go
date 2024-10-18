package codec

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// NewBoolKey returns a NameableKeyCodec instance for encoding and decoding bool keys.
func NewBoolKey[T ~bool]() NameableKeyCodec[T] {
	return boolKey[T]{}
}

// EncodeJSON encodes a boolean value into JSON format.
func (b boolKey[T]) EncodeJSON(value T) ([]byte, error) {
	return json.Marshal(value)
}

// DecodeJSON decodes a boolean value from JSON format.
func (b boolKey[T]) DecodeJSON(buffer []byte) (T, error) {
	var t T
	err := json.Unmarshal(buffer, &t)
	return t, err
}

// Stringify converts a boolean value to its string representation.
func (b boolKey[T]) Stringify(key T) string {
	return strconv.FormatBool((bool)(key))
}

// KeyType returns the type identifier for a bool key.
func (b boolKey[T]) KeyType() string {
	return "bool"
}

// WithName returns a KeyCodec instance with a specified name for bool keys.
func (b boolKey[T]) WithName(name string) KeyCodec[T] {
	return NamedKeyCodec[T]{KeyCodec: b, Name: name}
}
