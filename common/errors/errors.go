package errors

import (
	"github.com/shiyan123/marvel.sy/common/rd"
)

var (
	RespErrorOK      = rd.RespWriteErrorWithCode(0, "ok")
	ErrInvalidParams = rd.RespWriteErrorWithCode(10000, "参数错误")
)