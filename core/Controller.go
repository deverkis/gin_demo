package core

type BaseAction struct {
	Code    int64    `json:"code"`
	Message string `json:"message"`
	Data    DataVo `json:"data"`
}

type DataVo struct {
	Total    int64         `json:"total"`
	PageSize int         `json:"page_size"`
	Page     int64         `json:"page"`
	List     interface{} `json:"list"`
}

func (con *BaseAction) Lists(pagination Pagination, list interface{}) {
	data := DataVo{
		Total:    pagination.Total,
		PageSize: pagination.PageSize,
		Page:     pagination.Page,
		List:     list,
	}
	con.Data = data
}
