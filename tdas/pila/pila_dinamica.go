package pila

const (
	_NUMERO_PARA_REDIMENSIONAR = 2
	_NUMERO_PARA_REDUCIR       = 4
	_CAPACIDAD_INICIAL         = 5
	_MENSAJE_PANIC             = "La pila esta vacia"
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, _CAPACIDAD_INICIAL)
	return pila
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic(_MENSAJE_PANIC)
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elem T) {
	if pila.cantidad == len(pila.datos) {
		nuevaCapacidad := len(pila.datos) * _NUMERO_PARA_REDIMENSIONAR
		pila.redimensionar(nuevaCapacidad)
	}
	pila.datos[pila.cantidad] = elem
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic(_MENSAJE_PANIC)
	}
	dato := pila.datos[pila.cantidad-1]
	pila.cantidad--
	if len(pila.datos) > _CAPACIDAD_INICIAL && pila.cantidad*_NUMERO_PARA_REDUCIR <= len(pila.datos) {
		pila.redimensionar(len(pila.datos) / _NUMERO_PARA_REDIMENSIONAR)
	}
	return dato
}

func (pila *pilaDinamica[T]) redimensionar(nuevaCapacidad int) {
	pilaNueva := make([]T, nuevaCapacidad)
	copy(pilaNueva, pila.datos[:pila.cantidad])
	pila.datos = pilaNueva
}
