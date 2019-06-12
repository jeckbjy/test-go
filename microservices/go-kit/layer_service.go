package main

type IService interface {
	Greeting(name string) string
}

type GreeterSrv struct {
}

func (s *GreeterSrv) Greeting(name string) string {
	return "GO-KIT Hello " + name
}
