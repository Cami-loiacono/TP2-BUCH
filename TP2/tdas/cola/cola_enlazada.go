package cola

const (
	_MENSAJE_PANIC = "La cola esta vacia"
)

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

func crearNodo[T any](dato T) *nodo[T] {
	nodoNuevo := new(nodo[T])
	nodoNuevo.dato = dato
	return nodoNuevo
}

func CrearColaEnlazada[T any]() Cola[T] {
	return new(colaEnlazada[T])
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil && cola.ultimo == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic(_MENSAJE_PANIC)
	}
	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(elem T) {
	nuevoNodo := crearNodo(elem)
	if cola.EstaVacia() {
		cola.primero = nuevoNodo
	} else {
		cola.ultimo.siguiente = nuevoNodo
	}
	cola.ultimo = nuevoNodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic(_MENSAJE_PANIC)
	}
	dato := cola.primero.dato
	cola.primero = cola.primero.siguiente
	if cola.primero == nil {
		cola.ultimo = nil
	}
	return dato
}
