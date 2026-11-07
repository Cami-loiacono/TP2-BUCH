package ip

import (
	"errors"
	"strconv"
	"strings"
)

const (
	_CANT_BLOQUES = 4
	_MAX_BLOQUE   = 255
	_SEPARADOR_IP = "."
)

type ip struct {
	ipPorBloques [_CANT_BLOQUES]int
}

func ParsearIP(ipStr string) (IP, error) {
	bloquesIP := strings.Split(ipStr, _SEPARADOR_IP)
	nuevaIP := ip{}

	if len(bloquesIP) != _CANT_BLOQUES {
		return nil, errors.New("IP inválida: cantidad de bloques incorrecta")
	}

	for i, bloque := range bloquesIP {
		bloqueInt, err := strconv.Atoi(bloque)
		if err != nil || bloqueInt < 0 || bloqueInt > _MAX_BLOQUE {
			return nil, errors.New("IP inválida")
		}

		nuevaIP.ipPorBloques[i] = bloqueInt
	}

	return nuevaIP, nil
}

func (direccion ip) CompararIP(otraIp IP) int {
	otraDireccion := otraIp.(ip)
	for i := 0; i < _CANT_BLOQUES; i++ {
		if direccion.ipPorBloques[i] > otraDireccion.ipPorBloques[i] {
			return 1
		} else if direccion.ipPorBloques[i] < otraDireccion.ipPorBloques[i] {
			return -1
		}
	}
	return 0
}

func (direccion ip) ObtenerOcteto(posicion int) int {
	if posicion < 0 || posicion >= _CANT_BLOQUES {
		panic("posición de octeto inválida")
	}
	return direccion.ipPorBloques[posicion]
}