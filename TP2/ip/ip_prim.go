package ip

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	_CANT_BLOQUES = 4
	_MAX_BLOQUE   = 255
	_SEPARADOR_IP = "."

)

type ip struct{
	ipPorBloques [CANT_BLOQUES]int
}

func ParsearIP(ipStr string) (IP, error) {
	bloquesIP := strings.Split(ipStr, SEPARADOR_IP)
	ip := IP{}
	if len(bloquesIP) != CANT_BLOQUES {
		return ip, errors.New("IP inválida: cantidad de bloques incorrecta")
	}
	for i, bloque := range bloquesIP {
		bloqueInt, _ := strconv.Atoi(bloque)
		ip.ipPorBloques[i] = bloqueInt
	}
	return ip, nil
}

func (ip IP) CompararIP(otraIP IP) int {
	for i := 0; i < CANT_BLOQUES; i++ {
		if ip.ipPorBloques[i] < otraIP.ipPorBloques[i] {
			return 0
		}
	}
	return 1
}

