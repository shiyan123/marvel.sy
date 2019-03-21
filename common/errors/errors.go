package errors

import "custom-switch/common/dto"

var (
	RespErrorOK      = dto.RespWriteErrorWithCode(0, "ok")
	ErrInvalidParams = dto.RespWriteErrorWithCode(10000, "参数错误")
)