package college

type Result struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

func NewResult(total int, data interface{}) *Result {
	return &Result{Total: total, Data: data}
}
