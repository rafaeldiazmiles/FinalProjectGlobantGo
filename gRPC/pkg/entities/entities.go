package entities

// User - should contain the definition of our user
type User struct {
	Id      int32
	Pwd     string
	Name    string
	Age     uint32
	AddInfo string
	// parents []Parent   --> Para implementar cuando haya parents
}
