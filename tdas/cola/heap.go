package cola

const (
	_CAPACIDAD_INICIAL = 5
	_PANIC_COLA_VACIA = "La cola esta vacia"
	_NUMERO_REDIMENSIONAR_DES = 4
	_NUMERO_REDIMENSIONAR_EN = 2
)

type cmpHeap[T any] func(T, T) int

type heap[T any] struct {
	datos []T
	cantidad int
	cmp cmpHeap[T]
}

func CrearHeap[T any](cmp cmpHeap[T]) ColaPrioridad[T] {
	return &heap[T]{
		datos: make([]T, _CAPACIDAD_INICIAL),
		cmp: cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, cmp cmpHeap[T]) ColaPrioridad[T] {
	datos := make([]T, max(len(arreglo), _CAPACIDAD_INICIAL))
	copy(datos, arreglo)
	heapNuevo := &heap[T]{datos: datos, cantidad: len(arreglo), cmp: cmp}
	heapify(heapNuevo.datos, heapNuevo.cantidad, cmp)
	return heapNuevo
}

func heapify[T any](datos []T, cantidad int, cmp cmpHeap[T]) {
	for i := cantidad/2 - 1; i >= 0; i-- {
		downheap(datos, cantidad, cmp, i)
	}
}

func downheap[T any](datos []T, cantidad int, cmp cmpHeap[T], padre int) {
	for {
		IndiceHijoIzq := 2*padre + 1
		IndiceHijoDer := 2*padre + 2
		mayor := padre

		if IndiceHijoIzq < cantidad && cmp(datos[IndiceHijoIzq], datos[mayor]) > 0 {
			mayor = IndiceHijoIzq
		}
		if IndiceHijoDer < cantidad && cmp(datos[IndiceHijoDer], datos[mayor]) > 0 {
			mayor = IndiceHijoDer
		}
		if mayor == padre {
			return
		}
		datos[padre], datos[mayor] = datos[mayor], datos[padre]
		padre = mayor
	}
}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic(_PANIC_COLA_VACIA)
	}
	return heap.datos[0]
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic(_PANIC_COLA_VACIA)
	}
	maximoPrimerElemento := heap.datos[0]
	heap.cantidad--
	heap.datos[0] = heap.datos[heap.cantidad]
	downheap(heap.datos, heap.cantidad, heap.cmp, 0)
	if len(heap.datos) > _CAPACIDAD_INICIAL && heap.cantidad*_NUMERO_REDIMENSIONAR_DES <= len(heap.datos) {
		redimensionar(heap, max(len(heap.datos)/_NUMERO_REDIMENSIONAR_EN, _CAPACIDAD_INICIAL))
	}
	return maximoPrimerElemento
}

func (heap *heap[T]) Encolar(elemento T){
	if heap.cantidad == len(heap.datos) {
		redimensionar(heap, len(heap.datos)*_NUMERO_REDIMENSIONAR_EN)
	}
	heap.datos[heap.cantidad] = elemento
	heap.cantidad++
	upheap(heap.datos, heap.cmp, heap.cantidad-1)
}

func upheap[T any](datos []T, cmp cmpHeap[T], indice int) {
	for indice > 0 {
		padre := (indice - 1) / 2
		if cmp(datos[indice], datos[padre]) <= 0 {
			return
		}
		datos[indice], datos[padre] = datos[padre], datos[indice]
		indice = padre
	}
}

func redimensionar[T any](heap *heap[T], nuevaCapacidad int) {
	nuevoDatos := make([]T, nuevaCapacidad)
	copy(nuevoDatos, heap.datos)
	heap.datos = nuevoDatos
}

func Heapsort[T any](arreglo []T, cmp cmpHeap[T]){
	heapify(arreglo, len(arreglo), cmp)
	for i := len(arreglo) - 1; i >= 0; i-- {
		arreglo[0], arreglo[i] = arreglo[i], arreglo[0]
		downheap(arreglo, i, cmp, 0)
	}
}
