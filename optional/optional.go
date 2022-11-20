package optional

import (
	"encoding/json"
	"errors"
)

var ErrValueNotSet = errors.New("value is not set")

var marshalledNil, _ = json.Marshal(nil)

type Optional[T any] struct {
	data  T
	isSet bool
}

func New[T any](data T) Optional[T] {
	return Optional[T]{
		data:  data,
		isSet: true,
	}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{
		data:  *new(T),
		isSet: false,
	}
}

func (o *Optional[T]) Get() T {
	if !o.isSet {
		panic(ErrValueNotSet)
	}

	return o.data
}

func (o *Optional[T]) GetOrFail() (T, error) {
	if !o.isSet {
		return *new(T), ErrValueNotSet
	}

	return o.data, nil
}

func (o *Optional[T]) GetOrElse(t T) T {
	if !o.isSet {
		return t
	}

	return o.data
}

func (o *Optional[T]) IsSet() bool {
	return o.isSet
}

func (o *Optional[T]) Set(data T) {
	o.data = data
	o.isSet = true
}

func (o *Optional[T]) Clear(data T) {
	o.data = *new(T)
	o.isSet = false
}

func (o *Optional[T]) MarshalJSON() ([]byte, error) {
	if o.isSet {
		return json.Marshal(o.data)
	}

	return marshalledNil, nil
}

func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	if string(data) == string(marshalledNil) {
		o.data = *new(T)
		o.isSet = false
		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	o.data = value
	o.isSet = true
	return nil
}
