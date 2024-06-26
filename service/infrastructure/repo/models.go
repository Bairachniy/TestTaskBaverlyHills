package repo

type SignUpInfo struct {
	ID       int    `json:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}
