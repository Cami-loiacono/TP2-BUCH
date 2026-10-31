package logs

type Recurso struct {
	Nombre   string
	Cantidad int
}

type Logs interface {
	AgregarArchivo(ruta string) ([]IP, error)

	VerVisitantes(desde, hasta IP) []IP

	VerMasVisitados(n int) []Recurso
}