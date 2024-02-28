package response

type CategoryDTO struct {
	ID   uint
	Name string
	Icon string
}

type CategoryCreateResp struct {
	Category *CategoryDTO `json:"category"`
}

type CategoryListResp struct {
	Categories *[]CategoryDTO `json:"categories"`
}
