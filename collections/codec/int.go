package codec

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
)

// NewInt64Key returns a NameableKeyCodec instance for encoding and decoding int64 keys.
func NewInt64Key[T ~int64]() NameableKeyCodec[T] { return int64Key[T]{} }

type int64Key[T ~int64] struct{}

// EncodeJSON encodes an int64 value into JSON format.
func (i int64Key[T]) EncodeJSON(value T) ([]byte, error) {
	return []byte(`"` + strconv.FormatInt((int64)(value), 10) + `"`), nil
}

// DecodeJSON decodes an int64 value from JSON format.
func (i int64Key[T]) DecodeJSON(b []byte) (T, error) {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return 0, err
	}
	k, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return (T)(k), nil
}

// Stringify converts an int64 value to its string representation.
func (i int64Key[T]) Stringify(key T) string {
	return strconv.FormatInt((int64)(key), 10)
}

// KeyType returns the type identifier for an int64 key.
func (i int64Key[T]) KeyType() string {
	return "int64"
}

// WithName returns a KeyCodec instance with a specified name for int64 keys.
func (i int64Key[T]) WithName(name string) KeyCodec[T] {
	return NamedKeyCodec[T]{KeyCodec: i, Name: name}
}

// NewInt32Key returns a NameableKeyCodec instance for encoding and decoding int32 keys.
func NewInt32Key[T ~int32]() NameableKeyCodec[T] {
	return int32Key[T]{}
}

type int32Key[T ~int32] struct{}

// EncodeJSON encodes an int32 value into JSON format.
func (i int32Key[T]) EncodeJSON(value T) ([]byte, error) {
	return []byte(`"` + strconv.FormatInt((int64)(value), 10) + `"`), nil
}

// DecodeJSON decodes an int32 value from JSON format.
func (i int32Key[T]) DecodeJSON(b []byte) (T, error) {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return 0, err
	}
	k, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return (T)(k), nil
}

// Stringify converts an int32 value to its string representation.
func (i int32Key[T]) Stringify(key T) string {
	return strconv.FormatInt((int64)(key), 10)
}

// KeyType returns the type identifier for an int32 key.
func (i int32Key[T]) KeyType() string {
	return "int32"
}

// WithName returns a KeyCodec instance with a specified name for int32 keys.
func (i int32Key[T]) WithName(name string) KeyCodec[T] {
	return NamedKeyCodec[T]{KeyCodec: i, Name: name}
}
