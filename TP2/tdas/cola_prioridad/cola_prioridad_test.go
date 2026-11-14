package cola_prioridad_test

import (
	"github.com/stretchr/testify/require"
	TDACOLA_PRIORIDAD "tdas/cola_prioridad"
	"testing"
)

const (
	_MENSAJE_PANIC_HEAP = "La cola esta vacia"
	_VOLUMEN_HEAP       = 10000
)

func cmpInt(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func cmpString(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func TestHeapify(t *testing.T) {
	arreglo := []int{5, 3, 8, 1, 2, 7}
	heap := TDACOLA_PRIORIDAD.CrearHeapArr(arreglo, cmpInt)

	require.Equal(t, len(arreglo), heap.Cantidad())
	require.Equal(t, 8, heap.VerMax())

	arregloOrdenadoEsperado := []int{8, 7, 5, 3, 2, 1}
	for _, elementoEsperado := range arregloOrdenadoEsperado {
		require.Equal(t, elementoEsperado, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestHeapVacio(t *testing.T) {
	heap := TDACOLA_PRIORIDAD.CrearHeap[int](cmpInt)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_HEAP, func() { heap.VerMax() })
	require.PanicsWithValue(t, _MENSAJE_PANIC_HEAP, func() { heap.Desencolar() })
}

func TestHeapComoNuevo(t *testing.T) {
	heap := TDACOLA_PRIORIDAD.CrearHeap[string](cmpString)
	heap.Encolar("hola")
	require.False(t, heap.EstaVacia())
	require.Equal(t, "hola", heap.VerMax())
	require.Equal(t, "hola", heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_HEAP, func() { heap.VerMax() })
	require.PanicsWithValue(t, _MENSAJE_PANIC_HEAP, func() { heap.Desencolar() })
}

func TestHeapFIFO(t *testing.T) {
	heap := TDACOLA_PRIORIDAD.CrearHeap[int](cmpInt)
	heap.Encolar(1)
	require.Equal(t, 1, heap.VerMax())
	heap.Encolar(2)
	require.Equal(t, 2, heap.VerMax())
	heap.Desencolar()
	require.False(t, heap.EstaVacia())
	require.Equal(t, 1, heap.VerMax())
	heap.Encolar(3)
	require.Equal(t, 3, heap.VerMax())
	heap.Desencolar()
	require.False(t, heap.EstaVacia())

	heap.Encolar(4)
	heap.Encolar(5)
	require.Equal(t, 5, heap.VerMax())
	heap.Desencolar()
	require.False(t, heap.EstaVacia())
	require.Equal(t, 4, heap.VerMax())
	heap.Desencolar()
	require.Equal(t, 1, heap.VerMax())
	heap.Desencolar()
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_HEAP, func() { heap.VerMax() })
	require.PanicsWithValue(t, _MENSAJE_PANIC_HEAP, func() { heap.Desencolar() })
}

func TestHeapVolumen(t *testing.T) {
	heap := TDACOLA_PRIORIDAD.CrearHeap[int](cmpInt)
	for i := 0; i < _VOLUMEN_HEAP; i++ {
		heap.Encolar(i)
		require.Equal(t, i, heap.VerMax())
	}
	require.Equal(t, _VOLUMEN_HEAP-1, heap.VerMax())
	for i := _VOLUMEN_HEAP - 1; i >= 0; i-- {
		require.Equal(t, i, heap.VerMax())
		require.Equal(t, i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_HEAP, func() { heap.VerMax() })
	require.PanicsWithValue(t, _MENSAJE_PANIC_HEAP, func() { heap.Desencolar() })
}

func TestHeapsortVacio(t *testing.T) {
	arreglo := []int{}
	TDACOLA_PRIORIDAD.HeapSort(arreglo, cmpInt)
	require.Equal(t, []int{}, arreglo)
}

func TestHeapsortUnElemento(t *testing.T) {
	arreglo := []int{7}
	TDACOLA_PRIORIDAD.HeapSort(arreglo, cmpInt)
	require.Equal(t, []int{7}, arreglo)
}

func TestHeapsortYaOrdenado(t *testing.T) {
	arreglo := []int{1, 2, 3, 4, 5, 6}
	TDACOLA_PRIORIDAD.HeapSort(arreglo, cmpInt)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6}, arreglo)
}

func TestHeapsortRepetidos(t *testing.T) {
	arreglo := []int{4, 1, 4, 2, 1, 4, 3}
	TDACOLA_PRIORIDAD.HeapSort(arreglo, cmpInt)
	require.Equal(t, []int{1, 1, 2, 3, 4, 4, 4}, arreglo)
}

func TestHeapsortDesordenado(t *testing.T) {
	arreglo := []int{9, 3, 7, 1, 8, 2, 6, 0, 5, 4}
	TDACOLA_PRIORIDAD.HeapSort(arreglo, cmpInt)
	require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arreglo)
}

func TestHeapsortStrings(t *testing.T) {
	arreglo := []string{"hola", "algo", "docentes", "buenas noches"}
	TDACOLA_PRIORIDAD.HeapSort(arreglo, cmpString)
	require.Equal(t, []string{"algo", "buenas noches", "docentes", "hola"}, arreglo)
}

func TestHeapsortVolumen(t *testing.T) {
	arreglo := make([]int, _VOLUMEN_HEAP)
	for i := range arreglo {
		arreglo[i] = _VOLUMEN_HEAP - 1 - i
	}
	TDACOLA_PRIORIDAD.HeapSort(arreglo, cmpInt)
	for i := range arreglo {
		require.Equal(t, i, arreglo[i])
	}
}
