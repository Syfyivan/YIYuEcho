package models

type Diary struct {
	DiaryID   uint32 `json:"id"`
	Tag       string `json:"tag"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	AuthorID  int64  `json:"author_id"`
	User      string `json:"user"`
}

//todo ： UnmarshalJSON
//UnmarshalJSON 为diary类型实现自定义的UnmarshalJSON方法 ，可以在解析JSON数据时进行一些额外的验证和转换
//还是先用默认的json解析方法吧

// ApiDiaryDetail 日记返回的详情结构体
type ApiDiaryDetail struct {
	*Diary            //嵌入一个指向 Diary 结构体的指针
	AuthorPhoneNumber string
}

// Page 分页结构体,表示分页信息
type Page struct {
	Total int64 `json:"total"` //表示总记录数
	Page  int64 `json:"page"`  //表示当前页码
	Size  int64 `json:"size"`  //表示每页的记录数
}

// ApiDiaryDetailRes 表示详细的日记分页结果
type ApiDiaryDetailRes struct {
	Page Page             `json:"page"` //嵌入一个 Page 结构体，表示分页信息。
	List []ApiDiaryDetail `json:"list"` //List：指向 ApiPostDetail 结构体指针的切片，表示当前页的帖子详情列表。
}

// ParamDiaryList 表示查询日记列表的参数结构体
type ParamDiaryList struct {
	Search    string `json:"search" form:"search"`                    // 关键字搜索
	AuthorID  uint64 `json:"author_id" form:"author_id"`              // 可以为空
	Page      int    `json:"page" form:"page"`                        // 页码
	Size      int    `json:"size" form:"size"`                        // 每页数量
	Order     string `json:"order" form:"order" example:"created_at"` // 排序依据(按创建时间/最近编辑时间排序)
	CreatedAt string `json:"created_at" from:"created_at"`
}
