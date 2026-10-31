package logs

import (
	"bufio"
	"os"
	"slices"
	"time"

	TDACola "tdas/cola"
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	IP "tp2/ip"
)

type logsServidor struct {
	visitantes TDADiccionario.DiccionarioOrdenado[IP, bool] //LISTA DE IPS
	visitas    TDADiccionario.Diccionario[string, int]
}

func CrearLogs() Logs {
	return &logsServidor{
		visitantes: TDADiccionario.CrearABB[IP, bool](),
		visitas:    TDADiccionario.CrearHash[string, int](),
	}
}

func (l *logsServidor) AgregarArchivo(ruta string) ([]IP, error){}
func (l *logsServidor) VerVisitantes(desde, hasta IP) []IP{}
func (l *logsServidor) VerMasVisitados(n int) []Recurso{}
func (l *logsServidor) registrarVisita(e entrada){}

func buscarIPsSospechosas(ips []IP) []IP {}