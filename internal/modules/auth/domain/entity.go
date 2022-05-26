package domain

// Todo ...
type User struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

// CreateData ...
type CreateData struct {
	LoginData
}

// CreateData ...
type LoginData struct {
	Name     string `db:"name"`
	Password string `db:"password"`
}

// CreateData ...
type UpdateData struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
}
