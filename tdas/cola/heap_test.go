package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACOLA "tdas/cola"
	"testing"
)

const (
	_MENSAJE_PANIC_HEAP = "La cola esta vacia"
	_VOLUMEN_HEAP       = 10000
)

func TestHeapify(t *testing.T) {
	arreglo := []int{5, 3, 8, 1, 2, 7}
	heap := TDACOLA.CrearHeapArr(arreglo, func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	require.Equal(t, len(arreglo), heap.Cantidad())
	require.Equal(t, 8, heap.VerMax())

	arregloOrdenadoEsperado := []int{8, 7, 5, 3, 2, 1}
	for _, elementoEsperado := range arregloOrdenadoEsperado {
		require.Equal(t, elementoEsperado, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestHeapVacio(t *testing.T) {
	heap := TDACOLA.CrearHeap[int](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_HEAP, func() { heap.VerMax() })
	require.PanicsWithValue(t, _MENSAJE_PANIC_HEAP, func() { heap.Desencolar() })
}

func TestHeapComoNuevo(t *testing.T) {
	heap := TDACOLA.CrearHeap[string](func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	heap.Encolar("hola")
	require.False(t, heap.EstaVacia())
	require.Equal(t, "hola", heap.VerMax())
	require.Equal(t, "hola", heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_HEAP, func() { heap.VerMax() })
	require.PanicsWithValue(t, _MENSAJE_PANIC_HEAP, func() { heap.Desencolar() })
}

func TestHeapFIFO(t *testing.T) {
	heap := TDACOLA.CrearHeap[int](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	heap.Encolar(1)
	require.Equal(t, 1, heap.VerMax())
	heap.Encolar(2)
	require.Equal(t, 2, heap.VerMax())
	heap.Desencolar()
	require.False(t, heap.EstaVacia())
	require.Equal(t, 2, heap.VerMax())
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
	require.Equal(t, 4, heap.VerMax())
	heap.Desencolar()
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_HEAP, func() { heap.VerMax() })
	require.PanicsWithValue(t, _MENSAJE_PANIC_HEAP, func() { heap.Desencolar() })
}

func TestHeapVolumen(t *testing.T) {
	heap := TDACOLA.CrearHeap[int](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	for i := 0; i < _VOLUMEN_HEAP; i++ {
		heap.Encolar(i)
		require.Equal(t, _VOLUMEN_HEAP-1, heap.VerMax())
	}
	for i := _VOLUMEN_HEAP - 1; i >= 0; i-- {
		require.Equal(t, i, heap.VerMax())
		require.Equal(t, i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_HEAP, func() { heap.VerMax() })
	require.PanicsWithValue(t, _MENSAJE_PANIC_HEAP, func() { heap.Desencolar() })
}


