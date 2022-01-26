package user

const (
	CreateUserQuery string = "INSERT INTO USER (name, pwd, age, add_info) VALUES (?,?,?,?)"
)
