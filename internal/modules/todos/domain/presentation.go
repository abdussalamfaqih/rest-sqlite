package domain

// CreateRequest ...
type CreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateRequest ...
type UpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// TodoData ...
type TodoData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
