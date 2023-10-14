package container

type Container struct {
	RestClient RestClient
}

func NewContainer() Container {
	return Container{CreateRestClient()}
}
