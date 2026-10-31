package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	log "tp2/log"
	ip "tp2/ip"
)

const (
	_AGREGAR_ARCHIVO = "agregar_archivo"
	_VER_VISITANTES = "ver_visitantes"
	_VER_MAS_VISITADOS = "ver_mas_visitados"
	_REGISTRO_CORRECTO = "OK"
	_CANTIDAD_PARAMETROS_AGREGAR = 1
)

func main() {
	registro := log.CrearLogs()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		if !procesarRegistro(linea, registro) {
			fmt.Fprintf(os.Stderr, "Error en comando %s\n", linea)
			os.Exit(1)
		}
		fmt.Println(_REGISTRO_CORRECTO)
	}
}

func procesarRegistro(linea string, registro log.Logs) bool {
	campos := strings.Fields(linea)
	if len(campos) == 0 {
		return false
	}
	accion, restante := campos[0], campos[1:]
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
	registro.AgregarArchivo(restante[0])
	return true
}

func verVisitantes(restante []string, registro log.Logs) bool {
	if len(restante) != 2 {
		return false
	}
	desde, hasta := restante[0], restante[1]
	desdeIP := ip.ParsearIP(desde)
	hastaIP := ip.ParsearIP(hasta)
   	registro.VerVisitantes(desdeIP, hastaIP)
   	return true
}

func verMasVisitados(restante []string, registro log.Logs) bool {}