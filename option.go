// Package option provides Option struct that acts as a container
// of an optional value, meaning that value may exist or be omitted.
//
// The Option struct is designed to provide a convenient way to handle
// optional fields in structs without resorting to pointers or magic values.
//
// Unlike pointers, Option avoids unnecessary heap allocations in certain
// scenarios.
//
// Unlike magic values (e.g., empty strings), Option explicitly represents
// the absence of a value.
//
// Example use case:
//
//	type User struct {
//	  Name  string
//	  Email option.Option[string]
//	}
package option

// Option is an optional T, that may not contain an actual
// value.
type Option[T any] struct {
	inner T
	isset bool
}

// Some returns a new Option with v.
func Some[T any](v T) Option[T] {
	return Option[T]{inner: v, isset: true}
}

// None returns a new Option without a value.
func None[T any]() Option[T] {
	return Option[T]{isset: false}
}

// Value returns a pointer to the inner value of this Option
// and a bool indicating that the value is present or not.
//
// If the bool is equal to false, the pointer is nil.
func (o *Option[T]) Value() (*T, bool) {
	if o.isset {
		return &o.inner, true
	} else {
		return nil, false
	}
}

// IsSome returns true if this Option has a value.
func (o *Option[T]) IsSome() bool {
	_, r := o.Value()
	return r
}

// IsNone returns true if this Option is considered empty.
func (o *Option[T]) IsNone() bool {
	return !o.IsSome()
}

// Or returns the inner value of this Option if such
// exists or the result of f is used.
func (o Option[T]) Or(f func() T) T {
	if v, ok := o.Value(); ok {
		return *v
	} else {
		return f()
	}
}

// OrDefault returns the inner value of this Option
// or the fallback value if it's considered empty.
func (o Option[T]) OrDefault(fallback T) T {
	return o.Or(func() T { return fallback })
}

// From returns an option based on input v and ok.
func From[T any](v T, ok bool) Option[T] {
	if ok {
		return Some(v)
	} else {
		return None[T]()
	}
}

// FromFunc returns an option based on the result of the func f.
func FromFunc[T any](f func() (T, bool)) Option[T] {
	return From(f())
}

// FromMap returns an option with the value inside m by key or an empty
// option if the map m does not contain such entry.
func FromMap[M ~map[K]V, K comparable, V any](m M, key K) Option[V] {
	return FromFunc(func() (v V, ok bool) { v, ok = m[key]; return })
}
