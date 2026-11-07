package logs

import (
	"bufio"
	"os"
	"strings"
	TDADiccionario "tdas/diccionario"
	"time"
	IP "tp2/ip"
	"slices"
	TDALista "tdas/lista"
)

const (
	_CANTIDAD_PARAMETROS_LOGS = 4
	_LAYOUT_FECHA = "2006-01-02T15:04:05-07:00"
	_CANTIDAD_MINIMA_VISITAS_DOS = 5
	_SEGUNDOS_MINIMOS_DOS = 2
)

type logsServidor struct {
	visitantes         TDADiccionario.DiccionarioOrdenado[IP.IP, bool]
	visitasArchivosIPs TDADiccionario.Diccionario[string, int]
}

func CrearLogs() Logs {
	return &logsServidor{
		visitantes:         TDADiccionario.CrearABB[IP.IP, bool](), //le falta un comparar, para que el abb sepa como acomodarlos. COMPARAR DEBERIA DE DEJAR SER PRIMITIVTIVA Y PASEAR PODRIA SER LA PRMITIVA.
		visitasArchivosIPs: TDADiccionario.CrearHash[string, int](),
	}
}

func verificarValidezLogs(campos []string) bool {
	if len(campos) != _CANTIDAD_PARAMETROS_LOGS {
		return false
	}
	return true
}

// func ordenarIPs(sospechosas []IP.IP) { SEGURAMENTE RADIX

func esDoS(horas TDALista.Lista[time.Time]) bool {
	if horas.Largo() < _CANTIDAD_MINIMA_VISITAS_DOS {
		return false
	}
	hayDoS := false
	iteradorPrimero := horas.Iterador()
	iteradorSegundo := horas.Iterador()
	for i := 0; i < _CANTIDAD_MINIMA_VISITAS_DOS; i++ {
		iteradorSegundo.Siguiente()
	}
	
	for iteradorSegundo.HayAlgoMas() {
		tiempoPrimero := iteradorPrimero.VerActual()
		tiempoSegundo := iteradorSegundo.VerActual()
		diferencia := tiempoSegundo.Sub(tiempoPrimero)
		if diferencia.Seconds() <= _SEGUNDOS_MINIMOS_DOS {
			hayDoS = true
			break
		}
		tiempoPrimero = iteradorPrimero.Siguiente()
		tiempoSegundo = iteradorSegundo.Siguiente()
	}
	return hayDoS
}

func identificarIPsSospechosasDoS(horasIPs TDADiccionario.Diccionario[IP.IP, time.Time]) []IP.IP {
	var sospechosas []IP.IP
	horasIPs.Iterar(func(dir IP.IP, horas TDALista.Lista[time.Time]) bool {
		if esDoS(horas) {
			sospechosas = append(sospechosas, dir)
		} else {
			sospechosas = append(sospechosas, nil)
		}
		return true
	})
	//ordenarIPs(sospechosas)
	return sospechosas
}

func (l *logsServidor) AgregarArchivo(ruta string) ([]IP.IP){
	archivoAbierto, _ := os.Open(ruta)
	defer archivoAbierto.Close()
	scanner := bufio.NewScanner(archivoAbierto)
	var ipsSospechosas []IP.IP	
	horasIPs := TDADiccionario.CrearHash[IP.IP, time.Time]()

	for scanner.Scan() {
		linea := scanner.Text()
		campos := strings.Fields(linea)
		if !verificarValidezLogs(campos) {
			continue
		}
		ip, _ := IP.ParsearIP(campos[0])
		l.visitantes.Guardar(ip, true)
		l.visitasArchivosIPs.Guardar(campos[3], l.visitasArchivosIPs.Obtener(campos[3]) + 1)
		horasIPs.Guardar(ip, campos[1])
	}
	ipsSospechosas = identificarIPsSospechosasDoS(horasIPs)
	return ipsSospechosas
}
func (l *logsServidor) VerVisitantes(desde, hasta IP.IP) []IP.IP    {}
func (l *logsServidor) VerMasVisitados(n int) []Recurso             {}

func buscarIPsSospechosas(ips []IP.IP) []IP.IP {}
