package option_test

import (
	"testing"
	"time"

	"github.com/iizudev/option"
)

func TestOption(t *testing.T) {
	t.Run("IsNone", func(t *testing.T) {
		t.Run("returns true for empty option", func(t *testing.T) {
			s := option.None[any]()
			if v := s.IsNone(); v != true {
				t.Error("IsNone should return true for an empty option, but got false")
			}
		})

		t.Run("returns false for option with value", func(t *testing.T) {
			s := option.Some(0)
			if v := s.IsNone(); v != false {
				t.Error("IsNone should return false for an option with a value, but got true")
			}
		})
	})

	t.Run("IsSome", func(t *testing.T) {
		t.Run("returns false for empty option", func(t *testing.T) {
			s := option.None[any]()
			if v := s.IsSome(); v != false {
				t.Error("IsSome should return false for an empty option, but got true")
			}
		})

		t.Run("returns true for option with value", func(t *testing.T) {
			s := option.Some(0)
			if v := s.IsSome(); v != true {
				t.Error("IsSome should return true for an option with a value, but got false")
			}
		})
	})

	t.Run("Value", func(t *testing.T) {
		t.Run("returns nil and !ok for empty option", func(t *testing.T) {
			s := option.None[any]()
			v, ok := s.Value()
			if v != nil {
				t.Error("Value should return a nil pointer for an empty option, but got a non-nil pointer")
			}
			if ok != false {
				t.Error("Value should return a false ok for an empty option, but got true")
			}
		})

		t.Run("returns actual value and ok for option with value", func(t *testing.T) {
			var (
				e     = time.Now().Unix()
				s     = option.Some(e)
				v, ok = s.Value()
			)
			if v == nil {
				t.Error("Value should return a non-nil pointer for an option with value, but got a nil")
			} else {
				if *v != e {
					t.Errorf(
						"Value should have returned a pointer to the inner value %v, but got a pointer to %v",
						e, *v,
					)
				}
			}
			if ok == false {
				t.Error("Value should return a true ok for an option with a value, but got false")
			}
		})
	})

	t.Run("Or", func(t *testing.T) {
		t.Run("returns func result for empty option", func(t *testing.T) {
			var (
				e = time.Now().Unix()
				a = option.None[int64]().Or(func() int64 { return e })
			)
			if a != e {
				t.Errorf(
					"Or should have returned %v, which is the value that the fallback func returns, but got %v",
					e, a,
				)
			}
		})

		t.Run("returns actual value for option with value", func(t *testing.T) {
			var (
				e = time.Now().Unix()
				a = option.Some(e).Or(func() int64 { return 0 })
			)
			if a != e {
				t.Errorf(
					"Or should have returned %v, which is the inner value of option, but got %v",
					e, a,
				)
			}
		})
	})

	t.Run("OrDefault", func(t *testing.T) {
		t.Run("returns fallback for empty option", func(t *testing.T) {
			var (
				e = time.Now().Unix()
				a = option.None[int64]().OrDefault(e)
			)
			if a != e {
				t.Errorf(
					"Or should have returned %v, which is the fallback value, but got %v",
					e, a,
				)
			}
		})

		t.Run("returns actual value for option with value", func(t *testing.T) {
			var (
				e = time.Now().Unix()
				a = option.Some(e).OrDefault(0)
			)
			if a != e {
				t.Errorf(
					"OrDefault should have returned %v, which is the inner value of option, but got %v",
					e, a,
				)
			}
		})
	})
}
