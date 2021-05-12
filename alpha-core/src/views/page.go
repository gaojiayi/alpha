package views

type Page struct {
	PageIndex int64 `form:"page_index":"pageIndex" binding:"required"`
	PageSize  int64 `form:"page_size":"pageSize" binding:"required"`
}
