package logs

import (
	IP "tp2/ip"
)

type Logs interface {
	AgregarArchivo(ruta string) ([]IP.IP)
	VerVisitantes(desde, hasta IP.IP) []IP.IP
	VerMasVisitados(n int) 
}