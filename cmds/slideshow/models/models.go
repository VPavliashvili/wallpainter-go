package models

type Operation interface {
	Execute() error
}
