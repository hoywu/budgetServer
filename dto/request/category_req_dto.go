package request

type CategoryCreateRequest struct {
	Name string
	Icon string
}

type CategoryUpdateRequest struct {
	ID uint
	CategoryCreateRequest
}

type CategoryRemoveRequest struct {
	Name string `json:"name"`
}
