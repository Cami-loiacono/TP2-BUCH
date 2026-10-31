package ip

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	CANT_BLOQUES = 4
	MAX_BLOQUE   = 255
	SEPARADOR_IP = "."
)

type IP struct{
	combinacion [CANT_BLOQUES]int
	//cmp 
}

func ParsearIP(ipStr string) (IP) {}
func (ip IP) CompararIPs(ip2 IP) int {}
