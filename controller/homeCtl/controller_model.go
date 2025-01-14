package homeCtl

type SearchActivityParams struct {
	ActivityDateEnd   string `form:"end"`
	ActivityDateStart string `form:"start"`
	ActivityNature    int    `form:"nature"`
	ActivityStatus    int    `form:"status"`
	ActivityTypeID    uint   `form:"type_id"`
	Keyword           string `form:"keyword"`
	Page              int    `form:"page"`
}

type EditType struct {
	Id       int    `json:"id"`
	TypeName string `json:"typeName"`
}

type AddType struct {
	TypeName string `json:"typeName"`
	ImageId  int    `json:"imageId"`
}

type EditImage struct {
	Id       int    `json:"id"`
	FileName string `json:"fileName"`
}

type AddImage struct {
	FileName string `json:"fileName"`
	Url      string `json:"url"`
}
