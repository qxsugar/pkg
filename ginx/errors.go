package ginx

import (
	"github.com/qxsugar/pkg/zapx"
)

const TAG = "ExceptionX"

func NewInvalidArgumentError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewInvalidArgumentError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      InvalidArgument,
		Msg:       "参数异常",
		Desc:      err.Error(),
	}
}

func NewFailedPreconditionError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewFailedPreconditionError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      FailedPrecondition,
		Msg:       "执行条件异常",
		Desc:      err.Error(),
	}
}

func NewOutOfRangeError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewOutOfRangeError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      OutOfRange,
		Msg:       "无效范围",
		Desc:      err.Error(),
	}
}

func NewUnauthenticatedError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewUnauthenticatedError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      Unauthenticated,
		Msg:       "无效身份",
		Desc:      err.Error(),
	}
}

func NewPermissionDeniedError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewPermissionDeniedError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      PermissionDenied,
		Msg:       "权限不足",
		Desc:      err.Error(),
	}
}

func NewNotFoundError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewNotFoundError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      NotFound,
		Msg:       "资源不存在",
		Desc:      err.Error(),
	}
}

func NewAbortedError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewAbortedError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      Aborted,
		Msg:       "重复操作",
		Desc:      err.Error(),
	}
}

func NewAlreadyExistsError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewAlreadyExistsError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      AlreadyExists,
		Msg:       "资源已存在",
		Desc:      err.Error(),
	}
}

func NewResourceExhaustedError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewResourceExhaustedError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      ResourceExhausted,
		Msg:       "系统繁忙",
		Desc:      err.Error(),
	}
}

func NewCancelledError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewCancelledError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      Cancelled,
		Msg:       "客户端取消请求",
		Desc:      err.Error(),
	}
}

func NewDataLossError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewDataLossError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      DataLoss,
		Msg:       "数据已损坏",
		Desc:      err.Error(),
	}
}

func NewUnknownError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewUnknownError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      Unknown,
		Msg:       "未知错误",
		Desc:      err.Error(),
	}
}

func NewInternalError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewInternalError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      Internal,
		Msg:       "内部错误",
		Desc:      err.Error(),
	}
}

func NewNotImplementedError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewNotImplementedError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      NotImplemented,
		Msg:       "方法未实现",
		Desc:      err.Error(),
	}
}

func NewUnavailableError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewUnavailableError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      Unavailable,
		Msg:       "暂停服务",
		Desc:      err.Error(),
	}
}

func NewDeadlineExceededError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewDeadlineExceededError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      DeadlineExceeded,
		Msg:       "系统无法执行",
		Desc:      err.Error(),
	}
}
