package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	log "tp2/log"
	ip "tp2/ip"
)

const (
	_AGREGAR_ARCHIVO = "agregar_archivo"
	_VER_VISITANTES = "ver_visitantes"
	_VER_MAS_VISITADOS = "ver_mas_visitados"
	_REGISTRO_CORRECTO = "OK"
	_CANTIDAD_PARAMETROS_AGREGAR = 1
	_CANTIDAD_PARAMETROS_VER_VISITANTES = 2
	_CANTIDAD_PARAMETROS_VER_TOP_VISITADOS = 1
)

func verificarValidezRango(desdeIP, hastaIP ip.IP) bool {
	for i := 0; i < ip.CANT_BLOQUES; i++ {
		if desdeIP.ipPorBloques[i] > hastaIP.ipPorBloques[i] {
			return false
		}
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
	registro.AgregarArchivo(restante[0])
	return true
}

func verVisitantes(restante []string, registro log.Logs) bool {
	if len(restante) != _CANTIDAD_PARAMETROS_VER_VISITANTES {
		return false
	}
	desdeIP, errDesde := ip.ParsearIP(restante[0])
	hastaIP, errHasta := ip.ParsearIP(restante[1])
	if !verificarValidezRango(desdeIP, hastaIP) && errDesde != nil && errHasta != nil {
		return false
	}
   	registro.VerVisitantes(desdeIP, hastaIP)
   	return true
}

func verMasVisitados(restante []string, registro log.Logs) bool {
	if len(restante) != _CANTIDAD_PARAMETROS_VER_TOP_VISITADOS{
		return false
	}
	numeroTop, _ := strconv.Atoi(restante[0])
	if numeroTop <= 0 {
		return false
	}
	fmt.Println("Sitios más visitados:")
	registro.VerMasVisitados()
	return true
}



func main() {
	registro := log.CrearLogs() //nomrbe registro es raro, no agarra un solo registro, sino que es un conjunto de registros, pero bueno
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		if !procesarRegistro(linea, registro) {
			fmt.Fprintln(os.Stderr, "Error en comando %s\n", linea)
			os.Exit(1)
		}
		fmt.Println(_REGISTRO_CORRECTO)
	}
}
