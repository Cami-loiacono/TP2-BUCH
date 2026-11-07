package ip

type IP interface {
    // compara una IP entrante con la IP que queremos comparar. Devuelve 1 si la IP entrante es menor a la IP actual, -1 si es mayor y 0 si son iguales
    CompararIP(otra IP) int
    ObtenerOcteto(posicion int) int
}