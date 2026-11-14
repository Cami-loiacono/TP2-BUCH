package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	ip "tp2/ip"
	log "tp2/log"
)

const (
	_AGREGAR_ARCHIVO                       = "agregar_archivo"
	_VER_VISITANTES                        = "ver_visitantes"
	_VER_MAS_VISITADOS                     = "ver_mas_visitados"
	_REGISTRO_CORRECTO                     = "OK"
	_CANTIDAD_PARAMETROS_AGREGAR           = 1
	_CANTIDAD_PARAMETROS_VER_VISITANTES    = 2
	_CANTIDAD_PARAMETROS_VER_TOP_VISITADOS = 1
	_MENSAJE_VER_MAS_VISITADOS             = "Sitios más visitados:"
	_MENSAJE_VER_VISITANTES                = "Visitantes:"
)

func verificarValidezRango(desdeIP, hastaIP ip.IP) bool {
	if desdeIP.CompararIP(hastaIP) == 1 {
		return false
	}
	return true
}

func parsearEntrada(linea string) (string, []string) {
	campos := strings.Fields(linea)
	if len(campos) == 0 {
		return "", []string{}
	}
	return campos[0], campos[1:]
}

func comunicarIpsSospechosas(ipsSospechosas []ip.IP) {
	for _, ip := range ipsSospechosas {
		if ip != nil {
			fmt.Printf("DoS: %d.%d.%d.%d\n",
				ip.ObtenerOcteto(0),
				ip.ObtenerOcteto(1),
				ip.ObtenerOcteto(2),
				ip.ObtenerOcteto(3),
			)
		}
	}
}

func procesarRegistro(linea string, registro log.Logs) bool {
	accion, restante := parsearEntrada(linea)
	if accion != _AGREGAR_ARCHIVO && accion != _VER_VISITANTES && accion != _VER_MAS_VISITADOS {
		return false
	}
	if accion == _AGREGAR_ARCHIVO {
		return agregarArchivo(restante, registro)
	} else if accion == _VER_VISITANTES {
		return verVisitantes(restante, registro)
	} else if accion == _VER_MAS_VISITADOS {
		return verMasVisitados(restante, registro)
	}
	return false
}

func agregarArchivo(restante []string, registro log.Logs) bool {
	if len(restante) != _CANTIDAD_PARAMETROS_AGREGAR {
		return false
	}
	archivo, err := os.Open(restante[0])
	if err != nil {
		return false
	}
	archivo.Close()
	ipsSospechosas := registro.AgregarArchivo(restante[0])
	comunicarIpsSospechosas(ipsSospechosas)
	return true
}

func verVisitantes(restante []string, registro log.Logs) bool {
	if len(restante) != _CANTIDAD_PARAMETROS_VER_VISITANTES {
		return false
	}
	desdeIP, errDesde := ip.ParsearIP(restante[0])
	hastaIP, errHasta := ip.ParsearIP(restante[1])
	if !verificarValidezRango(desdeIP, hastaIP) || errDesde != nil || errHasta != nil {
		return false
	}
	fmt.Println(_MENSAJE_VER_VISITANTES)
	registro.VerVisitantes(desdeIP, hastaIP)
	return true
}

func verMasVisitados(restante []string, registro log.Logs) bool {
	if len(restante) != _CANTIDAD_PARAMETROS_VER_TOP_VISITADOS {
		return false
	}
	numeroTop, _ := strconv.Atoi(restante[0])
	fmt.Println(_MENSAJE_VER_MAS_VISITADOS)
	registro.VerMasVisitados(numeroTop)
	return true
}

func main() {
	registro := log.CrearLogs() //nomrbe registro es raro, no agarra un solo registro, sino que es un conjunto de registros, pero bueno
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		if !procesarRegistro(linea, registro) {
			accion, _ := parsearEntrada(linea)
			fmt.Fprintf(os.Stderr, "Error en comando %s\n", accion)
			return
		}
		fmt.Println(_REGISTRO_CORRECTO)
	}
}
