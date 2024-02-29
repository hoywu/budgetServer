package response

type CategoryDTO struct {
	ID   uint
	Name string
	Icon string
}

type CategoryCreateResp struct {
	Category *CategoryDTO
}

type CategoryListResp struct {
	Categories *[]CategoryDTO
}
