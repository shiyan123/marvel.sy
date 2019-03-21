package rd

const (
	ErrCode_Success = 0
	ErrMsg_Success  = "success"
)

type RespModel struct {
	Meta *RespMeta   `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

type RespMeta struct {
	ErrCode int     `json:"errCode"`
	ErrMsg  string  `json:"errMsg"`
	Paging  *Paging `json:"paging,omitempty"`
}

type Paging struct {
	Total       int `json:"total"`
	TotalPage   int `json:"totalPage"`
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
}

type ErrorType struct {
	StatusCode int
	ErrorCode  int
	ErrorMsg   string
}

func Data(data interface{}) interface{}{
	return NewRespModel(data)
}

func DataPaging(data interface{}, currentPage, pageSize, total int)interface{} {
	return NewRespModelWithPaging(data, currentPage, pageSize, total)
}

func NewRespModel(data interface{}) *RespModel {
	return &RespModel{
		Meta: NewMeta(),
		Data: data,
	}
}

func NewRespModelWithPaging(data interface{}, page, size, total int) *RespModel {
	return &RespModel{
		Meta: NewMetaWithPaging(page, size, total),
		Data: data,
	}
}

func NewMeta() *RespMeta {
	return &RespMeta{
		ErrCode: ErrCode_Success,
		ErrMsg:  ErrMsg_Success,
	}
}

func NewMetaWithPaging(page, size, total int) *RespMeta {
	return &RespMeta{
		ErrCode: ErrCode_Success,
		ErrMsg:  ErrMsg_Success,
		Paging:  NewPaging(page, size, total),
	}
}

func NewPaging(currentPage, pageSize, total int) *Paging {
	totalPage := total / pageSize
	if total%pageSize != 0 {
		totalPage += 1
	}

	return &Paging{
		Total:       total,
		TotalPage:   totalPage,
		CurrentPage: currentPage,
		PageSize:    pageSize,
	}
}

func RespWriteErrorWithCode(errCode int, msg string) *RespModel {
	return &RespModel{
		Meta: &RespMeta{ErrCode: errCode, ErrMsg: msg},
	}
}