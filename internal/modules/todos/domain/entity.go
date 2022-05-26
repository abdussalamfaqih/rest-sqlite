package domain

import (
	"encoding/base64"
	"fmt"
)

// Todo ...
type Todo struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

func (d *Todo) TransformToPresentation() TodoData {
	return TodoData{
		ID:          base64.RawStdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", d.ID))),
		Name:        d.Name,
		Description: d.Description,
	}
}

// CreateData ...
type CreateData struct {
	UserID      int64  `db:"user_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

// CreateData ...
type UpdateData struct {
	UserID      int64  `db:"user_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
