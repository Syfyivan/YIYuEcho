package models

type Tag struct {
	Id  int    `json:"id"`
	Tag string `json:"tag"`
}

// ParamTagList 标签列表参数
type ParamTagList struct {
	Search  string `form:"search"`  // 搜索关键词
	Page    int    `form:"page"`    // 当前页码
	Size    int    `form:"size"`    // 每页数量
	Order   string `form:"order"`   // 排序方式
	OrderBy string `form:"orderBy"` // 排序字段
}
