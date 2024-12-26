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
