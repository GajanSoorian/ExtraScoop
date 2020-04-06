package persistence

//To do: WHy the array of bytes.
type DatabaseHandler interface {
	AddService(Service) ([]byte, error)
	FindService([]byte) (Service, error)
	FindServiceByType(string) (Service, error)
	FindAllAvailableServices() ([]Service, error)
}
