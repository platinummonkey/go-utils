package atomic

import (
	goAtomic "sync/atomic"
)

// Atomic is a go generics accessible interface to the sync/atomic.
type Atomic[T any] struct {
	v *goAtomic.Value
}

// NewAtomic will allow you to create a go generic Atomic.
func NewAtomic[T any](value T) *Atomic[T] {
	v := &goAtomic.Value{}
	v.Store(value)
	return &Atomic[T]{
		v: v,
	}
}

// Load returns the value set by the most recent Store.
// It returns nil if there has been no call to Store for this Value.
func (a *Atomic[T]) Load() T {
	iv := a.v.Load()
	return iv.(T)
}

// Store sets the value of the Value to x.
// All calls to Store for a given Value must use values of the same concrete type.
// Store of an inconsistent type panics, as does Store(nil).
func (a *Atomic[T]) Store(value T) {
	a.v.Store(value)
}

// Swap stores new into Value and returns the previous value. It returns nil if
// the Value is empty.
//
// All calls to Swap for a given Value must use values of the same concrete
// type. Swap of an inconsistent type panics, as does Swap(nil).
func (a *Atomic[T]) Swap(newValue T) (oldValue T) {
	oldUntypedValue := a.v.Swap(newValue)
	return oldUntypedValue.(T)
}

// CompareAndSwap executes the compare-and-swap operation for the Value.
//
// All calls to CompareAndSwap for a given Value must use values of the same
// concrete type. CompareAndSwap of an inconsistent type panics, as does
// CompareAndSwap(old, nil).
func (a *Atomic[T]) CompareAndSwap(old, new T) (swapped bool) {
	return a.v.CompareAndSwap(old, new)
}
