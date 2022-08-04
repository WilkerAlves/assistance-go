package service

type IGeneratedIds interface {
	Create() (string, error)
}
