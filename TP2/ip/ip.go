package ip

type IP interface {
    // compara una IP entrante con la IP que queremos comparar. Devuelve 1 si la IP entrante es menor a la IP actual y 0 si es mayor.
    CompararIP(otra IP) int
}