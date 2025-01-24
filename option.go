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
