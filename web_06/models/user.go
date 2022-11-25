package models

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username" form:"username" `
	Password string `db:"password" form:"password" `
}

type Question struct {
	Food     string `db:"food"`
	Favor    string `db:"favor"`
	School   string `db:"school"`
	Username string `db:"username"`
}

type Info struct {
	Message string `db:"message"`
}
