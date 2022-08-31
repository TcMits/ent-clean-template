package lazy

import "sync"

type LazyObject[T any] interface {
	Value() T
}

type lazyObject[T any] struct {
	once      sync.Once
	value     T
	setupFunc func() T
}

func (l *lazyObject[T]) lazyInit() {
	l.once.Do(func() {
		l.value = l.setupFunc()
	})
}

func (l *lazyObject[T]) Value() T {
	l.lazyInit()
	return l.value
}

func NewLazyObject[T any](setupFunc func() T) LazyObject[T] {
	return &lazyObject[T]{setupFunc: setupFunc}
}
