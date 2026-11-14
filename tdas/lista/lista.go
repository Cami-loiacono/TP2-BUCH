package lista

type IteradorLista[T any] interface {
	// Muestra el elemento actual del iterador.
	// Si el iterador está al final de la lista, entra en pánico con un mensaje diciendo: "El iterador termino de iterar"
	VerActual() T

	// Devuelve verdadero si el iterador tiene un elemento actual, falso en caso contrario.
	HayAlgoMas() bool

	// Avanza el iterador al siguiente elemento de la lista.
	// entra en panic si está en el final de la lista con un mensaje diciendo: "El iterador termino de iterar"
	Avanzar()

	// Inserta un elemento en la posición actual del iterador.
	// El elemento insertado pasa a ser el actual y aumenta el contador de elementos de la lista.
	Insertar(T)

	// Elimina el elemento actual del iterador y lo devuelve.
	// entra en panico si el actual es nil, con el mensaje: "El iterador termino de iterar"
	Borrar() T
}

type Lista[T any] interface {
	// EstaVacia devuelve verdadero si la lista no tiene elementos enlazados, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un elemento al principio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo agrega un elemento al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer elemento de la lista y lo devuelve.
	// Si está vacía, entra en pánico con un mensaje diciendo: "La lista está vacía"
	BorrarPrimero() T

	// VerPrimero devuelve el primer elemento de la lista.
	// Si está vacía, entra en pánico con un mensaje diciendo: "La lista está vacía"
	VerPrimero() T

	// VerUltimo devuelve el último elemento de la lista.
	// Si está vacía, entra en pánico con un mensaje diciendo: "La lista está vacía"
	VerUltimo() T

	// Largo devuelve la cantidad de elementos en la lista.
	Largo() int

	// Iterar recorre la lista y aplica la función visitar a cada elemento.
	// Si la función visitar devuelve false, se detiene la iteración.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un iterador externo de la lista.
	Iterador() IteradorLista[T]
}
