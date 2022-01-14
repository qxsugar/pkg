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
		Info:      err.Error(),
		Desc:      "客户端指定了无效的参数。 检查错误消息和错误详细信息以获取更多信息。",
	}
}

func NewFailedPreconditionError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewFailedPreconditionError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      FailedPrecondition,
		Info:      err.Error(),
		Desc:      "请求不能在当前系统状态下执行，例如删除非空目录。",
	}
}

func NewOutOfRangeError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewOutOfRangeError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      OutOfRange,
		Info:      err.Error(),
		Desc:      "客户端指定了无效的范围。",
	}
}

func NewUnauthenticatedError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewUnauthenticatedError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      Unauthenticated,
		Info:      err.Error(),
		Desc:      "由于遗失，无效或过期的OAuth令牌而导致请求未通过身份验证。",
	}
}

func NewPermissionDeniedError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewPermissionDeniedError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      PermissionDenied,
		Info:      err.Error(),
		Desc:      "客户端没有足够的权限。这可能是因为OAuth令牌没有正确的范围，客户端没有权限，或者客户端项目尚未启用API。",
	}
}

func NewNotFoundError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewNotFoundError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      NotFound,
		Info:      err.Error(),
		Desc:      "找不到指定的资源，或者该请求被未公开的原因（例如白名单）拒绝。",
	}
}

func NewAbortedError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewAbortedError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      Aborted,
		Info:      err.Error(),
		Desc:      "并发冲突，例如读-修改-写冲突。",
	}
}

func NewAlreadyExistsError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewAlreadyExistsError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      AlreadyExists,
		Info:      err.Error(),
		Desc:      "客户端尝试创建的资源已存在。",
	}
}

func NewResourceExhaustedError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewResourceExhaustedError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      ResourceExhausted,
		Info:      err.Error(),
		Desc:      "资源配额达到速率限制。 客户端应该查找google.rpc.QuotaFailure错误详细信息以获取更多信息。",
	}
}

func NewCancelledError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewCancelledError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      Cancelled,
		Info:      err.Error(),
		Desc:      "客户端取消请求。",
	}
}

func NewDataLossError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewDataLossError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      DataLoss,
		Info:      err.Error(),
		Desc:      "不可恢复的数据丢失或数据损坏。 客户端应该向用户报告错误。",
	}
}

func NewUnknownError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewUnknownError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      Unknown,
		Info:      err.Error(),
		Desc:      "未知的服务器错误。 通常是服务器错误。",
	}
}

func NewInternalError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewInternalError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      Internal,
		Info:      err.Error(),
		Desc:      "内部服务错误。 通常是服务器错误。",
	}
}

func NewNotImplementedError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewNotImplementedError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      NotImplemented,
		Info:      err.Error(),
		Desc:      "服务器未实现该API方法。",
	}
}

func NewUnavailableError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewUnavailableError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      Unavailable,
		Info:      err.Error(),
		Desc:      "暂停服务。通常是服务器已经关闭。",
	}
}

func NewDeadlineExceededError(err error) FailedRespBody {
	zapx.GetLogger().Warnw(TAG, "fn", "NewDeadlineExceededError", "err", err)
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      DeadlineExceeded,
		Info:      err.Error(),
		Desc:      "已超过请求期限。如果重复发生，请考虑降低请求的复杂性。",
	}
}
