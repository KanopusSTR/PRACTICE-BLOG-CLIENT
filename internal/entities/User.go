package entities

type User struct {
	Name     string
	Mail     string // unique value
	Password []byte
}
