package core

type BaseAction struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    DataVo `json:"data"`
}

type DataVo struct {
	Total    int         `json:"total"`
	PageSize int         `json:"page_size"`
	Page     int         `json:"page"`
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
