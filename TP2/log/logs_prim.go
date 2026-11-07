package logs

import (
	TDADiccionario "tdas/diccionario"
	IP "tp2/ip"
)

type logsServidor struct {
	visitantes         TDADiccionario.DiccionarioOrdenado[IP.IP, bool]
	visitasArchivosIPs TDADiccionario.Diccionario[string, int]
}

func CrearLogs() Logs {
	return &logsServidor{
		visitantes:         TDADiccionario.CrearABB[IP.IP, bool](),
		visitasArchivosIPs: TDADiccionario.CrearHash[string, int](),
	}
}

func (l *logsServidor) AgregarArchivo(ruta string) ([]IP.IP, error) {}
func (l *logsServidor) VerVisitantes(desde, hasta IP.IP) []IP.IP    {}
func (l *logsServidor) VerMasVisitados(n int) []Recurso             {}

func buscarIPsSospechosas(ips []IP.IP) []IP.IP {}
