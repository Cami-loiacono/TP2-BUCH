package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACOLA "tdas/cola"
	"testing"
)

const (
	_MENSAJE_PANIC = "La cola esta vacia"
	_VOLUMEN       = 10000
)

func TestColaVacia(t *testing.T) {
	cola := TDACOLA.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { cola.VerPrimero() })
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { cola.Desencolar() })
}

func TestColaComoNueva(t *testing.T) {
	cola := TDACOLA.CrearColaEnlazada[string]()
	cola.Encolar("hola")
	require.False(t, cola.EstaVacia())
	require.Equal(t, "hola", cola.VerPrimero())
	require.Equal(t, "hola", cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { cola.VerPrimero() })
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { cola.Desencolar() })
}

func TestFIFO(t *testing.T) {
	cola := TDACOLA.CrearColaEnlazada[float64]()
	cola.Encolar(1.0)
	require.Equal(t, 1.0, cola.VerPrimero())
	cola.Encolar(2.0)
	require.Equal(t, 1.0, cola.VerPrimero())
	cola.Desencolar()
	require.False(t, cola.EstaVacia())
	require.Equal(t, 2.0, cola.VerPrimero())
	cola.Encolar(3.0)
	require.Equal(t, 2.0, cola.VerPrimero())
	cola.Desencolar()
	require.False(t, cola.EstaVacia())

	cola.Encolar(4.0)
	cola.Encolar(5.0)
	require.Equal(t, 3.0, cola.VerPrimero())
	cola.Desencolar()
	require.False(t, cola.EstaVacia())
	require.Equal(t, 4.0, cola.VerPrimero())
	cola.Desencolar()
	require.Equal(t, 5.0, cola.VerPrimero())
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { cola.VerPrimero() })
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { cola.Desencolar() })
}

func TestVolumen(t *testing.T) {
	cola := TDACOLA.CrearColaEnlazada[int]()
	volumen := _VOLUMEN
	for i := 0; i < volumen; i++ {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero())
		require.False(t, cola.EstaVacia())
	}
	for i := 0; i < volumen; i++ {
		require.Equal(t, i, cola.VerPrimero())
		require.False(t, cola.EstaVacia())
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { cola.VerPrimero() })
}

func TestDatosCambiantesPunteros(t *testing.T) {
	cola := TDACOLA.CrearColaEnlazada[*complex128]()
	a := complex(1, 2)
	b := complex(3, 4)
	c := complex(5, 6)
	cola.Encolar(&a)
	cola.Encolar(&b)
	cola.Encolar(&c)

	require.Equal(t, &a, cola.VerPrimero())
	require.Equal(t, &a, cola.Desencolar())
	require.Equal(t, &b, cola.VerPrimero())
	require.Equal(t, &b, cola.Desencolar())
	require.Equal(t, &c, cola.VerPrimero())
	require.Equal(t, &c, cola.Desencolar())
	require.True(t, cola.EstaVacia())

	cola.Encolar(&a)
	*cola.VerPrimero() = complex(7, 8)
	require.Equal(t, complex(7, 8), *cola.VerPrimero())
	require.Equal(t, complex(7, 8), *cola.Desencolar())
	require.True(t, cola.EstaVacia())
}
