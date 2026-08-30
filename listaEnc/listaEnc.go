package listaEnc

import "fmt"

type No[T comparable] struct {
	Id    int
	Value T
	Next  *No[T]
}

type ListaEnc[T comparable] struct {
	Head      *No[T]
	IdCounter int
}

func InitListaEnc[T comparable]() *ListaEnc[T] {
	return &ListaEnc[T]{
		Head:      nil,
		IdCounter: 0,
	}
}

func (l *ListaEnc[T]) ShowAll() {
	current := l.Head
	for current != nil {
		fmt.Println(current.Id, " : ", current.Value)
		current = current.Next
	}
}

func (l *ListaEnc[T]) Size() int {

	if l.Head == nil {
		return 0
	}

	current := l.Head
	counter := 1
	for current.Next != nil {
		current = current.Next
		counter++
	}
	return counter
}

func (l *ListaEnc[T]) IsEmpty() bool {
	return l.Head == nil
}

func (l *ListaEnc[T]) GetNext(id int) *No[T] {

	if l.Head == nil {
		return nil
	}

	current := l.Head
	for current != nil && current.Id != id {
		current = current.Next
	}
	return current.Next
}

func (l *ListaEnc[T]) GetValue(id int) *T {
	if l.Head == nil {
		return nil
	}

	current := l.Head
	for current.Id != id {
		if current.Next == nil {
			return nil
		}

		current = current.Next
	}

	return &current.Value
}

func (l *ListaEnc[T]) Insert(indice int, value T) {
	l.IdCounter++

	if l.Head == nil {
		new := &No[T]{Id: l.IdCounter, Value: value, Next: nil}
		l.Head = new
		return
	}

	current := l.Head
	for i := 0; i < indice-1 && current.Next != nil; i++ {
		current = current.Next
	}

	new := &No[T]{Id: l.IdCounter, Value: value, Next: current.Next}
	current.Next = new

}

func (l *ListaEnc[T]) InsertHead(value T) {

	l.IdCounter++

	if l.Head == nil {
		new := &No[T]{Id: l.IdCounter, Value: value, Next: nil}
		l.Head = new
		return
	}

	new := &No[T]{Id: l.IdCounter, Value: value, Next: l.Head}
	l.Head = new

}

func (l *ListaEnc[T]) InsertTail(value T) {

	l.IdCounter++

	if l.Head == nil {
		new := &No[T]{Id: l.IdCounter, Value: value, Next: nil}
		l.Head = new
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	new := &No[T]{Id: l.IdCounter, Value: value, Next: nil}
	current.Next = new
}

func (l *ListaEnc[T]) Remove(id int) {
	if l.Head == nil {
		return
	}
	if id == l.Head.Id {
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Id == id {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}

}

func (l *ListaEnc[T]) RemoveHead() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
}

func (l *ListaEnc[T]) RemoveTail() {
	if l.Head == nil {
		return
	}
	if l.Head.Next == nil {
		l.Head = nil
		return
	}

	current := l.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
}

func (l *ListaEnc[T]) Update(id int, value T) {

	current := l.Head

	for current != nil {
		if current.Id == id {
			current.Value = value
			return
		}
		current = current.Next
	}

}

func (l *ListaEnc[T]) UpdateHead(value T) {
	if l.Head == nil {
		return
	}
	l.Head.Value = value

}

func (l *ListaEnc[T]) UpdateTail(value T) {

	if l.Head == nil {
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Value = value

}

func (l *ListaEnc[T]) Exist(value T) bool {

	current := l.Head

	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}

	return false
}

func (l *ListaEnc[T]) Search(value T) *No[T] {

	if value == l.Head.Value {
		return l.Head
	}

	current := l.Head

	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}

	return nil
}

func (l *ListaEnc[T]) Destroy(id int) {

	fmt.Println("Go não precisa de destroy, ele tem o 'Garbage Collector' que limpa da memória qualquer variavel ou espaço utilizado que já não tenha mais referências")

}
