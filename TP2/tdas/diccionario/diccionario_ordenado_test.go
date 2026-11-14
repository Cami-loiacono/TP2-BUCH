package diccionario_test

import (
	"math/rand"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_MENSAJE_CLAVE_NO_PERTENECE = "La clave no pertenece al diccionario"
	_MENSAJE_ITERADOR_TERMINADO = "El iterador termino de iterar"
	_VOLUMEN_ABB                = 100000
)

func compararEnteros(a, b int) int {
	return a - b
}

func TestABBVacio(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	require.Equal(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(0))
	require.PanicsWithValue(t, _MENSAJE_CLAVE_NO_PERTENECE, func() { abb.Obtener(0) })
	require.PanicsWithValue(t, _MENSAJE_CLAVE_NO_PERTENECE, func() { abb.Borrar(0) })
}

func TestABBUnElemento(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, float64](strings.Compare)
	abb.Guardar("pi", 3.14)
	require.Equal(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece("pi"))
	require.False(t, abb.Pertenece("e"))
	require.Equal(t, 3.14, abb.Obtener("pi"))
	require.PanicsWithValue(t, _MENSAJE_CLAVE_NO_PERTENECE, func() { abb.Obtener("e") })
}

func TestABBGuardarVarios(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](compararEnteros)
	abb.Guardar(5, "cinco")
	abb.Guardar(3, "tres")
	abb.Guardar(8, "ocho")
	abb.Guardar(1, "uno")
	abb.Guardar(4, "cuatro")

	require.Equal(t, 5, abb.Cantidad())
	require.Equal(t, "cinco", abb.Obtener(5))
	require.Equal(t, "tres", abb.Obtener(3))
	require.Equal(t, "ocho", abb.Obtener(8))
	require.Equal(t, "uno", abb.Obtener(1))
	require.Equal(t, "cuatro", abb.Obtener(4))
	require.False(t, abb.Pertenece(2))
}

func TestABBReemplazarDato(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	abb.Guardar("gato", 1)
	abb.Guardar("perro", 2)
	abb.Guardar("gato", 10)

	require.Equal(t, 2, abb.Cantidad())
	require.Equal(t, 10, abb.Obtener("gato"))
	require.Equal(t, 2, abb.Obtener("perro"))
}

func TestABBBorrarHoja(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(5, 50)
	abb.Guardar(3, 30)
	abb.Guardar(8, 80)

	require.Equal(t, 30, abb.Borrar(3))
	require.False(t, abb.Pertenece(3))
	require.True(t, abb.Pertenece(5))
	require.True(t, abb.Pertenece(8))
	require.Equal(t, 2, abb.Cantidad())
	require.PanicsWithValue(t, _MENSAJE_CLAVE_NO_PERTENECE, func() { abb.Borrar(3) })
}

func TestABBBorrarConUnHijo(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(5, 50)
	abb.Guardar(3, 30)
	abb.Guardar(1, 10)
	abb.Guardar(8, 80)
	abb.Guardar(9, 90)

	require.Equal(t, 30, abb.Borrar(3))
	require.False(t, abb.Pertenece(3))
	require.Equal(t, 10, abb.Obtener(1))

	require.Equal(t, 80, abb.Borrar(8))
	require.False(t, abb.Pertenece(8))
	require.Equal(t, 90, abb.Obtener(9))
	require.Equal(t, 3, abb.Cantidad())
}

func TestABBBorrarConDosHijos(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(5, 50)
	abb.Guardar(3, 30)
	abb.Guardar(8, 80)
	abb.Guardar(1, 10)
	abb.Guardar(4, 40)
	abb.Guardar(7, 70)
	abb.Guardar(9, 90)

	require.Equal(t, 30, abb.Borrar(3))
	require.False(t, abb.Pertenece(3))
	require.Equal(t, 10, abb.Obtener(1))
	require.Equal(t, 40, abb.Obtener(4))
	require.Equal(t, 6, abb.Cantidad())
}

func TestABBBorrarRaiz(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(5, 50)
	abb.Guardar(3, 30)
	abb.Guardar(8, 80)
	abb.Guardar(7, 70)

	require.Equal(t, 50, abb.Borrar(5))
	require.False(t, abb.Pertenece(5))
	require.Equal(t, 3, abb.Cantidad())

	var recorrido []int
	abb.Iterar(func(clave int, dato int) bool {
		recorrido = append(recorrido, clave)
		return true
	})
	require.Equal(t, []int{3, 7, 8}, recorrido)
}

func TestABBVaciarYVolverAUsar(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(1, 10)
	abb.Guardar(2, 20)
	abb.Borrar(1)
	abb.Borrar(2)
	require.Equal(t, 0, abb.Cantidad())
	require.False(t, abb.Iterador().HayAlgoMas())

	abb.Guardar(3, 30)
	require.Equal(t, 1, abb.Cantidad())
	require.Equal(t, 30, abb.Obtener(3))
}

func TestABBIteradorExternoVacio(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	iterador := abb.Iterador()
	require.False(t, iterador.HayAlgoMas())
	require.PanicsWithValue(t, _MENSAJE_ITERADOR_TERMINADO, func() { iterador.VerActual() })
	require.PanicsWithValue(t, _MENSAJE_ITERADOR_TERMINADO, func() { iterador.Avanzar() })
}

func TestABBIteradorExternoEnOrden(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](compararEnteros)
	abb.Guardar(5, "cinco")
	abb.Guardar(3, "tres")
	abb.Guardar(8, "ocho")
	abb.Guardar(1, "uno")
	abb.Guardar(4, "cuatro")

	iterador := abb.Iterador()
	clave, dato := iterador.VerActual()
	require.Equal(t, 1, clave)
	require.Equal(t, "uno", dato)
	iterador.Avanzar()
	clave, dato = iterador.VerActual()
	require.Equal(t, 3, clave)
	require.Equal(t, "tres", dato)
	iterador.Avanzar()
	clave, dato = iterador.VerActual()
	require.Equal(t, 4, clave)
	require.Equal(t, "cuatro", dato)
	iterador.Avanzar()
	clave, dato = iterador.VerActual()
	require.Equal(t, 5, clave)
	require.Equal(t, "cinco", dato)
	iterador.Avanzar()
	clave, dato = iterador.VerActual()
	require.Equal(t, 8, clave)
	require.Equal(t, "ocho", dato)
	iterador.Avanzar()

	require.False(t, iterador.HayAlgoMas())
	require.PanicsWithValue(t, _MENSAJE_ITERADOR_TERMINADO, func() { iterador.VerActual() })
	require.PanicsWithValue(t, _MENSAJE_ITERADOR_TERMINADO, func() { iterador.Avanzar() })
}

func TestABBIteradorExternoTrasBorrar(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(5, 50)
	abb.Guardar(3, 30)
	abb.Guardar(8, 80)
	abb.Guardar(1, 10)
	abb.Borrar(5)

	var recorrido []int
	for iterador := abb.Iterador(); iterador.HayAlgoMas(); iterador.Avanzar() {
		clave, _ := iterador.VerActual()
		recorrido = append(recorrido, clave)
	}
	require.Equal(t, []int{1, 3, 8}, recorrido)
}

func TestABBRecorrerOrdenIteradorInterno(t *testing.T) {
	var recorrido []string
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	abb.Guardar("hola", 1)
	abb.Guardar("que", 2)
	abb.Guardar("tal", 3)

	abb.Iterar(func(clave string, dato int) bool {
		recorrido = append(recorrido, clave)
		return true
	})
	require.Equal(t, []string{"hola", "que", "tal"}, recorrido)
}

func TestABBSumarDatosIteradorInterno(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(2, 20)
	abb.Guardar(1, 10)
	abb.Guardar(3, 30)
	suma := 0
	abb.Iterar(func(clave int, dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 60, suma)
}

func TestABBIterarRangoIncluyeExtremos(t *testing.T) {
	var recorrido []int
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(5, 50)
	abb.Guardar(3, 30)
	abb.Guardar(8, 80)
	abb.Guardar(1, 10)
	abb.Guardar(4, 40)
	abb.Guardar(9, 90)

	desde, hasta := 3, 8
	abb.IterarRango(&desde, &hasta, func(clave int, dato int) bool {
		recorrido = append(recorrido, clave)
		return true
	})
	require.Equal(t, []int{3, 4, 5, 8}, recorrido)
}

func TestABBIterarRangoExtremosQueNoEstan(t *testing.T) {
	var recorrido []int
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(10, 100)
	abb.Guardar(20, 200)
	abb.Guardar(30, 300)
	abb.Guardar(40, 400)

	desde, hasta := 15, 35
	abb.IterarRango(&desde, &hasta, func(clave int, dato int) bool {
		recorrido = append(recorrido, clave)
		return true
	})
	require.Equal(t, []int{20, 30}, recorrido)
}

func TestABBIterarRangoSinDesdeOSinHasta(t *testing.T) {
	var recorridoHasta, recorridoDesde []int
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(5, 50)
	abb.Guardar(3, 30)
	abb.Guardar(8, 80)
	abb.Guardar(1, 10)

	limite := 4
	abb.IterarRango(nil, &limite, func(clave int, dato int) bool {
		recorridoHasta = append(recorridoHasta, clave)
		return true
	})
	abb.IterarRango(&limite, nil, func(clave int, dato int) bool {
		recorridoDesde = append(recorridoDesde, clave)
		return true
	})
	require.Equal(t, []int{1, 3}, recorridoHasta)
	require.Equal(t, []int{5, 8}, recorridoDesde)
}

func TestABBIterarRangoVacio(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(5, 50)
	abb.Guardar(3, 30)
	abb.Guardar(8, 80)

	visitados := 0
	desde, hasta := 6, 7
	abb.IterarRango(&desde, &hasta, func(clave int, dato int) bool {
		visitados++
		return true
	})
	require.Equal(t, 0, visitados)
}

func TestABBIteradorRangoIncluyeExtremos(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(5, 50)
	abb.Guardar(3, 30)
	abb.Guardar(8, 80)
	abb.Guardar(1, 10)
	abb.Guardar(4, 40)
	abb.Guardar(9, 90)

	desde, hasta := 3, 8
	iterador := abb.IteradorRango(&desde, &hasta)
	clave, dato := iterador.VerActual()
	require.Equal(t, 3, clave)
	require.Equal(t, 30, dato)
	iterador.Avanzar()
	clave, _ = iterador.VerActual()
	require.Equal(t, 4, clave)
	iterador.Avanzar()
	clave, _ = iterador.VerActual()
	require.Equal(t, 5, clave)
	iterador.Avanzar()
	clave, dato = iterador.VerActual()
	require.Equal(t, 8, clave)
	require.Equal(t, 80, dato)
	iterador.Avanzar()

	require.False(t, iterador.HayAlgoMas())
	require.PanicsWithValue(t, _MENSAJE_ITERADOR_TERMINADO, func() { iterador.VerActual() })
	require.PanicsWithValue(t, _MENSAJE_ITERADOR_TERMINADO, func() { iterador.Avanzar() })
}

func TestABBIteradorRangoExtremosQueNoEstan(t *testing.T) {
	var recorrido []int
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(10, 100)
	abb.Guardar(20, 200)
	abb.Guardar(30, 300)
	abb.Guardar(40, 400)

	desde, hasta := 15, 35
	for iterador := abb.IteradorRango(&desde, &hasta); iterador.HayAlgoMas(); iterador.Avanzar() {
		clave, _ := iterador.VerActual()
		recorrido = append(recorrido, clave)
	}
	require.Equal(t, []int{20, 30}, recorrido)
}

func TestABBIteradorRangoSinDesdeOSinHasta(t *testing.T) {
	var recorridoHasta, recorridoDesde []int
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(5, 50)
	abb.Guardar(3, 30)
	abb.Guardar(8, 80)
	abb.Guardar(1, 10)

	limite := 4
	for iterador := abb.IteradorRango(nil, &limite); iterador.HayAlgoMas(); iterador.Avanzar() {
		clave, _ := iterador.VerActual()
		recorridoHasta = append(recorridoHasta, clave)
	}
	for iterador := abb.IteradorRango(&limite, nil); iterador.HayAlgoMas(); iterador.Avanzar() {
		clave, _ := iterador.VerActual()
		recorridoDesde = append(recorridoDesde, clave)
	}
	require.Equal(t, []int{1, 3}, recorridoHasta)
	require.Equal(t, []int{5, 8}, recorridoDesde)
}

func TestABBIteradorRangoVacio(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(5, 50)
	abb.Guardar(3, 30)
	abb.Guardar(8, 80)

	desde, hasta := 6, 7
	iterador := abb.IteradorRango(&desde, &hasta)
	require.False(t, iterador.HayAlgoMas())
	require.PanicsWithValue(t, _MENSAJE_ITERADOR_TERMINADO, func() { iterador.VerActual() })
	require.PanicsWithValue(t, _MENSAJE_ITERADOR_TERMINADO, func() { iterador.Avanzar() })
}

// Volumen

func clavesDesordenadas(n int) []int {
	return rand.New(rand.NewSource(42)).Perm(n)
}

func TestABBVolumenGuardarObtenerBorrar(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	claves := clavesDesordenadas(_VOLUMEN_ABB)

	for _, clave := range claves {
		abb.Guardar(clave, clave*2)
	}
	require.Equal(t, _VOLUMEN_ABB, abb.Cantidad())

	for _, clave := range claves {
		require.True(t, abb.Pertenece(clave))
		require.Equal(t, clave*2, abb.Obtener(clave))
	}

	for _, clave := range claves {
		require.Equal(t, clave*2, abb.Borrar(clave))
		require.False(t, abb.Pertenece(clave))
	}
	require.Equal(t, 0, abb.Cantidad())
}

func TestABBVolumenIteradoresEnOrden(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	for _, clave := range clavesDesordenadas(_VOLUMEN_ABB) {
		abb.Guardar(clave, clave)
	}

	esperada := 0
	for iterador := abb.Iterador(); iterador.HayAlgoMas(); iterador.Avanzar() {
		clave, dato := iterador.VerActual()
		require.Equal(t, esperada, clave)
		require.Equal(t, esperada, dato)
		esperada++
	}
	require.Equal(t, _VOLUMEN_ABB, esperada)

	esperada = 0
	abb.Iterar(func(clave int, dato int) bool {
		require.Equal(t, esperada, clave)
		esperada++
		return true
	})
	require.Equal(t, _VOLUMEN_ABB, esperada)

	desde, hasta := _VOLUMEN_ABB/4, _VOLUMEN_ABB/2
	esperada = desde
	for iterador := abb.IteradorRango(&desde, &hasta); iterador.HayAlgoMas(); iterador.Avanzar() {
		clave, _ := iterador.VerActual()
		require.Equal(t, esperada, clave)
		esperada++
	}
	require.Equal(t, hasta+1, esperada)
}
