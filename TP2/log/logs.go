package logs

import (
	IP "tp2/ip"
)

type Logs interface {
	// Agregar archivo informa cuales ips del log entrante es sospecha de DoS. Devuelve las ips que tienen 5 o más consultas con menos de 2 segundos de diferencia. Valida si tiene la cantidad de campos correspondiete para realizar la validez, en caso negativo devuelve false y sino devuelve true después de informar las ips sospechosas.
	AgregarArchivo(ruta string) ([]IP.IP)
	VerVisitantes(desde, hasta IP.IP)
	VerMasVisitados(n int) 
}