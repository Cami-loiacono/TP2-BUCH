package ip

type IP interface {
	ParsearIP(ipStr string) (IP, error)
	//VerMasVisitados(n int) 
}