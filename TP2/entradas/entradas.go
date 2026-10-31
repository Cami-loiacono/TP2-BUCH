package entradas

import (
	"errors"
	"strings"
	"time"
	"tp2/ip"
)

const (
	LAYOUT_FECHA = "2006-01-02T15:04:05-07:00"
	SEPARADOR    = "\t"
	CANT_CAMPOS  = 4
)

type entrada struct {
	ip      ip.IP
	fecha   time.Time
	metodo  string
	recurso string
}

func parsearEntrada(linea string) (entrada, error) {
	campos := strings.Fields(linea)
	if len(campos) != CANT_CAMPOS {
		return entrada{}, errors.New("línea de log inválida")
	}
	ip := ip.ParsearIP(campos[0])
	fecha, err := time.Parse(LAYOUT_FECHA, campos[1])
	if err != nil {
		return entrada{}, err
	}
	return entrada{ip: ip, fecha: fecha, metodo: campos[2], recurso: campos[3]}, nil
}