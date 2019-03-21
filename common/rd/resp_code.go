package rd

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func Error(err *ErrorType, c *gin.Context) {
	c.AbortWithStatusJSON(err.StatusCode, NewRespModelWithErrType(err))
}

func JSON(data interface{}, c *gin.Context) {
	res := NewRespModel(data)
	c.JSON(http.StatusOK, res)
}

func JSONPaging(data interface{}, currentPage, pageSize, total int, c *gin.Context) {
	res := NewRespModelWithPaging(data, currentPage, pageSize, total)
	c.JSON(http.StatusOK, res)
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

func NewRespModelWithErrType(err *ErrorType) *RespModel {
	return &RespModel{
		Meta: &RespMeta{ErrCode: err.ErrorCode, ErrMsg: err.ErrorMsg},
	}
}
