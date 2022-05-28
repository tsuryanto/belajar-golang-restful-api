package web

// Modelnya Category , Aksinya Create, Bentuk nya Request
type CategoryCreateRequest struct {
	Name string `validate:"required, min=1, max=100"`
}
