package logs

import (
	IP "tp2/ip"
)

type Logs interface {
	AgregarArchivo(ruta string) ([]IP, error)

	VerVisitantes(desde, hasta IP) []IP

	//VerMasVisitados(n int) 
}