package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_MENSAJE_LISTA_VACIA        = "La lista esta vacia"
	_MENSAJE_ITERADOR_TERMINADO = "El iterador termino de iterar"
	_VOLUMEN                    = 100000
)

func TestListaVaciaEsVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, _MENSAJE_LISTA_VACIA, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, _MENSAJE_LISTA_VACIA, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, _MENSAJE_LISTA_VACIA, func() { lista.BorrarPrimero() })
}
func TestInsertarPrimeroUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[complex128]()
	lista.InsertarPrimero(complex(1, 2))
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, complex(1, 2), lista.VerPrimero())
	require.Equal(t, complex(1, 2), lista.VerUltimo())
}
func TestInsertarUltimoUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()
	lista.InsertarUltimo(3.14)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 3.14, lista.VerPrimero())
	require.Equal(t, 3.14, lista.VerUltimo())
}
func TestInsertarPrimeroVariosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)

	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
}
func TestInsertarUltimoVariosStrings(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
}
func TestInsertarPrimeroYUltimo(t *testing.T) {

	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("hola")
	lista.InsertarPrimero("que")
	lista.InsertarUltimo("tal")

	require.Equal(t, "que", lista.VerPrimero())
	require.Equal(t, "tal", lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())
}
func TestBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())

	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 3, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}
func TestVaciarListaYVolverAUsar(t *testing.T) {

	lista := TDALista.CrearListaEnlazada[float64]()
	lista.InsertarUltimo(1.2)
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())

	lista.InsertarUltimo(2.3)
	require.Equal(t, 2.3, lista.VerPrimero())
	require.Equal(t, 2.3, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
}
func TestVolumenInsertarUltimoYBorrarPrimero(t *testing.T) {

	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < _VOLUMEN; i++ {
		lista.InsertarUltimo(i)
	}
	require.Equal(t, _VOLUMEN, lista.Largo())
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, _VOLUMEN-1, lista.VerUltimo())

	for i := 0; i < _VOLUMEN; i++ {
		require.Equal(t, i, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}
func TestVolumenInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < _VOLUMEN; i++ {
		lista.InsertarPrimero(i)
	}
	require.Equal(t, _VOLUMEN, lista.Largo())
	require.Equal(t, _VOLUMEN-1, lista.VerPrimero())
	require.Equal(t, 0, lista.VerUltimo())

	for i := _VOLUMEN - 1; i >= 0; i-- {
		require.Equal(t, i, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}
func TestInicioIteradorExterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iterador := lista.Iterador()
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, iterador.VerActual())
}
func TestInsertarConIteradorExterno(t *testing.T) {

	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	require.Equal(t, 2, lista.Largo())

	iterador := lista.Iterador()
	for iterador.HayAlgoMas() {
		iterador.Avanzar()
	}
	iterador.Insertar(3)
	require.Equal(t, 3, iterador.VerActual())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 3, lista.VerUltimo())
	require.True(t, iterador.HayAlgoMas())

}
func TestInsertarMedioIteradorExterno(t *testing.T) {

	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)

	iter := lista.Iterador()
	iter.Avanzar()
	iter.Insertar(3)

	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 3, iter.VerActual())
	iter.Avanzar()
	require.Equal(t, 2, iter.VerActual())
	iter.Avanzar()
	require.False(t, iter.HayAlgoMas())
}
func TestBorrarPrimeroIteradorExterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iterador := lista.Iterador()
	require.Equal(t, 1, iterador.Borrar())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 2, iterador.VerActual())
}
func TestBorrarUltimoIteradorExterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iterador := lista.Iterador()
	iterador.Avanzar()
	iterador.Avanzar()
	require.Equal(t, 3, iterador.VerActual())

	require.Equal(t, 3, iterador.Borrar())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 2, lista.VerUltimo())
	require.False(t, iterador.HayAlgoMas())
}
func TestBorrarMedioIteradorExterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iterador := lista.Iterador()
	iterador.Avanzar()
	require.Equal(t, 2, iterador.VerActual())

	require.Equal(t, 2, iterador.Borrar())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 3, iterador.VerActual())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
}
func TestVaciarListaIteradorExterno(t *testing.T) {

	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iterador := lista.Iterador()
	iterador.Borrar()
	require.True(t, iterador.HayAlgoMas())
	iterador.Borrar()
	require.True(t, iterador.HayAlgoMas())
	iterador.Borrar()
	require.True(t, lista.EstaVacia())
	require.False(t, iterador.HayAlgoMas())
}
func TestInsertarMismaPosicionIteradorExterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(3)

	iterador := lista.Iterador()
	iterador.Insertar(2)
	require.Equal(t, 2, lista.Largo())
	require.True(t, iterador.HayAlgoMas())
	require.Equal(t, 2, iterador.VerActual())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	iterador.Insertar(1)
	require.True(t, iterador.HayAlgoMas())
	require.Equal(t, 1, iterador.VerActual())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())
}
func TestRecorrerOrdenIteradorInterno(t *testing.T) {
	var recorridoLista []int
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	lista.Iterar(func(elemento int) bool {
		recorridoLista = append(recorridoLista, elemento)
		return true
	})
	require.Equal(t, []int{1, 2, 3}, recorridoLista)
}
func TestCorteAnticipadoIteradorInterno(t *testing.T) {
	var recorridoLista []int
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(0)

	lista.Iterar(func(elemento int) bool {
		recorridoLista = append(recorridoLista, elemento)
		if elemento%2 == 0 {
			return false
		}
		return true
	})

	require.Equal(t, []int{1, 2}, recorridoLista)
}
func TestSumarElementosIteradorInterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(0)
	suma := 0
	lista.Iterar(func(v int) bool {
		suma += v
		return true
	})
	require.Equal(t, 5, suma)
}
