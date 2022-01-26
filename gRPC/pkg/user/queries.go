package user

const (
	CreateUserQuery string = "INSERT INTO user (name, pwd_hash, age, add_info) VALUES (?,?,?,?)"
)
