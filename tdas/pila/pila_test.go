package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

const (
	_MENSAJE_PANIC = "La pila esta vacia"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { pila.VerTope() })
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { pila.Desapilar() })
}

func TestPilaComoNueva(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("Hola")
	require.False(t, pila.EstaVacia())
	require.Equal(t, "Hola", pila.VerTope())
	require.Equal(t, "Hola", pila.Desapilar())
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { pila.VerTope() })
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { pila.Desapilar() })
}

func TestLIFO(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[complex128]()
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { pila.VerTope() })
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { pila.Desapilar() })

	pila.Apilar(complex(1, 2))
	pila.Apilar(complex(2, 3))
	pila.Apilar(complex(3, 4))
	require.False(t, pila.EstaVacia())
	require.Equal(t, complex(3, 4), pila.VerTope())

	pila.Apilar(complex(4, 5))
	require.Equal(t, complex(4, 5), pila.VerTope())
	require.Equal(t, complex(4, 5), pila.Desapilar())
	require.False(t, pila.EstaVacia())
	require.Equal(t, complex(3, 4), pila.VerTope())
	require.Equal(t, complex(3, 4), pila.Desapilar())
	require.False(t, pila.EstaVacia())
	require.Equal(t, complex(2, 3), pila.VerTope())
	require.Equal(t, complex(2, 3), pila.Desapilar())
	require.False(t, pila.EstaVacia())
	require.Equal(t, complex(1, 2), pila.VerTope())
	require.Equal(t, complex(1, 2), pila.Desapilar())
	require.True(t, pila.EstaVacia())

	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { pila.VerTope() })
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { pila.Desapilar() })
}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	volumen := 1000.0
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { pila.VerTope() })
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { pila.Desapilar() })
	for i := 0.0; i < volumen; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
		require.False(t, pila.EstaVacia())
	}
	for i := volumen - 1; i >= 0; i-- {
		require.Equal(t, i, pila.VerTope())
		require.Equal(t, i, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { pila.VerTope() })
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { pila.Desapilar() })
}

func TestFalsoVacio(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[*int]()
	pila.Apilar(nil)
	require.False(t, pila.EstaVacia(), "La pila no debería estar vacía después de apilar un elemento nil")
}

func TestDatosPunterosCambiantes(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[*int]()
	elemento1, elemento2, elemento3 := 1, 2, 3
	pila.Apilar(&elemento1)
	pila.Apilar(&elemento2)
	pila.Apilar(&elemento3)
	require.Equal(t, &elemento3, pila.VerTope())
	require.Equal(t, &elemento3, pila.Desapilar())
	require.Equal(t, &elemento2, pila.VerTope())
	require.Equal(t, &elemento2, pila.Desapilar())
	require.Equal(t, &elemento1, pila.VerTope())
	require.Equal(t, &elemento1, pila.Desapilar())
	require.True(t, pila.EstaVacia())

	pila.Apilar(&elemento1)
	*pila.VerTope() = 1302
	require.Equal(t, 1302, *pila.VerTope())
}
