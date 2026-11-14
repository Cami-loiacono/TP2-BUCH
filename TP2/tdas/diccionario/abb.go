package diccionario

import (
	TDAPila "tdas/pila"
)

const (
	_PANIC_CLAVE_NO_PERTENECE_DICCIONARIO = "La clave no pertenece al diccionario"
	_PANIC_ITERADOR_TERMINADO             = "El iterador termino de iterar"
)

type iterABB[K comparable, V any] struct {
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
	cmp   func(K, K) int
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cmp      func(K, K) int
	cantidad int
}

type nodoAbb[K comparable, V any] struct {
	clave     K
	dato      V
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
}

func buscarNodo[K comparable, V any](raiz **nodoAbb[K, V], clave K, cmp func(K, K) int) **nodoAbb[K, V] {
	if *raiz == nil {
		return raiz
	}
	comp := cmp(clave, (*raiz).clave)
	if comp == 0 {
		return raiz
	} else if comp < 0 {
		return buscarNodo(&(*raiz).izquierdo, clave, cmp)
	}
	return buscarNodo(&(*raiz).derecho, clave, cmp)
}

func CrearABB[K comparable, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	arbol := new(abb[K, V])
	arbol.cmp = funcionCmp
	return arbol
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{
		clave:     clave,
		dato:      dato,
		izquierdo: nil,
		derecho:   nil,
	}
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	nodo := buscarNodo(&abb.raiz, clave, abb.cmp)
	if *nodo != nil {
		(*nodo).dato = dato
	} else {
		nuevoNodo := crearNodo(clave, dato)
		*nodo = nuevoNodo
		abb.cantidad++
	}
}

func buscarMinimoDerecho[K comparable, V any](nodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	if (*nodo).izquierdo == nil {
		return nodo
	}
	return buscarMinimoDerecho(&(*nodo).izquierdo)
}

func verHijo[K comparable, V any](nodo *nodoAbb[K, V]) (hijo *nodoAbb[K, V]) {
	if nodo.derecho != nil {
		return nodo.derecho
	}
	return nodo.izquierdo
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodo := buscarNodo(&abb.raiz, clave, abb.cmp)
	if *nodo != nil {
		abb.cantidad--
		dato := (*nodo).dato
		if (*nodo).izquierdo != nil && (*nodo).derecho != nil {
			menorDerecho := buscarMinimoDerecho(&(*nodo).derecho)
			(*nodo).clave = (*menorDerecho).clave
			(*nodo).dato = (*menorDerecho).dato
			*menorDerecho = (*menorDerecho).derecho
		} else {
			*nodo = verHijo(*nodo)
		}
		return dato
	}
	panic(_PANIC_CLAVE_NO_PERTENECE_DICCIONARIO)
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	nodo := buscarNodo(&abb.raiz, clave, abb.cmp)
	return *nodo != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo := buscarNodo(&abb.raiz, clave, abb.cmp)
	if *nodo != nil {
		return (*nodo).dato
	}
	panic(_PANIC_CLAVE_NO_PERTENECE_DICCIONARIO)
}

func (abb *abb[K, V]) Iterar(visitar func(K, V) bool) {
	iterarInOrden(abb.raiz, visitar)
}

func iterarInOrden[K comparable, V any](nodo *nodoAbb[K, V], visitar func(K, V) bool) bool {
	if nodo == nil {
		return true
	}
	if !iterarInOrden(nodo.izquierdo, visitar) {
		return false
	}
	if !visitar(nodo.clave, nodo.dato) {
		return false
	}
	if !iterarInOrden(nodo.derecho, visitar) {
		return false
	}
	return true
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(K, V) bool) {
	iterarRango(abb.raiz, desde, hasta, abb.cmp, visitar)
}

func iterarRango[K comparable, V any](nodo *nodoAbb[K, V], desde *K, hasta *K, cmp func(K, K) int, visitar func(K, V) bool) bool {
	if nodo == nil {
		return true
	}
	if desde == nil || cmp(*desde, nodo.clave) < 0 {
		if !iterarRango(nodo.izquierdo, desde, hasta, cmp, visitar) {
			return false
		}
	}
	if (desde == nil || cmp(nodo.clave, *desde) >= 0) && (hasta == nil || cmp(nodo.clave, *hasta) <= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}
	if hasta == nil || cmp(nodo.clave, *hasta) < 0 {
		if !iterarRango(nodo.derecho, desde, hasta, cmp, visitar) {
			return false
		}
	}
	return true
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iterador := new(iterABB[K, V])
	iterador.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iterador.desde = desde
	iterador.hasta = hasta
	iterador.cmp = abb.cmp
	iterador.apilarIzquierdos(abb.raiz)
	return iterador
}

func (iter *iterABB[K, V]) apilarIzquierdos(nodo *nodoAbb[K, V]) {
	for nodo != nil {
		if iter.desde != nil && iter.cmp(nodo.clave, *iter.desde) < 0 {
			nodo = nodo.derecho
		} else if iter.hasta != nil && iter.cmp(nodo.clave, *iter.hasta) > 0 {
			nodo = nodo.izquierdo
		} else {
			iter.pila.Apilar(nodo)
			nodo = nodo.izquierdo
		}
	}
}

func (iter *iterABB[K, V]) HayAlgoMas() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iterABB[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic(_PANIC_ITERADOR_TERMINADO)
	}
	tope := iter.pila.VerTope()
	return tope.clave, tope.dato
}

func (iter *iterABB[K, V]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic(_PANIC_ITERADOR_TERMINADO)
	}
	desapilado := iter.pila.Desapilar()
	iter.apilarIzquierdos(desapilado.derecho)
}
