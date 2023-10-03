package account

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

const codeAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Account interface {
	ID() string
	Name() string
	Role() string
	Code() string
}

type account struct {
	id   string
	name string
	role string
	code string
}

func (a account) Code() string {
	return a.code
}

func (a account) ID() string {
	return a.id
}

func (a account) Name() string {
	return a.name
}

func (a account) Role() string {
	return a.role
}

func NewTeacher(name string) Account {
	return account{
		name: name,
		role: "teacher",
		id:   gonanoid.Must(),
		code: newAccountCode(),
	}
}

func newAccountCode() string {
	return gonanoid.MustGenerate(codeAlphabet, 4)
}

func NewStudent(name string) Account {
	return account{
		name: name,
		role: "student",
		id:   gonanoid.Must(),
		code: newAccountCode(),
	}
}

type Store interface {
	SaveAccount(account Account) error
	Login(code string) (Account, error)
	AccountExists(id string) bool
}
