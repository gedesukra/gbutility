package gbutility

import (
	"net/http"
	"strconv"
)

func GetJsonObjRs(requestUrl string, httpStatus int, status bool, message string) string {
	result := "{\"requrl\":\"" + requestUrl + "\",\"stacod\":" + strconv.Itoa(httpStatus) + ",\"statxt\":\"" + http.StatusText(httpStatus) + "\",\"sta\":" + strconv.FormatBool(status) + ",\"type\":\"obj\",\"msg\":" + message + "}"
	switch httpStatus {
	case http.StatusContinue:
		return result
	case http.StatusSwitchingProtocols:
		return result
	case http.StatusProcessing:
		return result
	case http.StatusOK:
		return result
	case http.StatusCreated:
		return result
	case http.StatusAccepted:
		return result
	case http.StatusNonAuthoritativeInfo:
		return result
	case http.StatusNoContent:
		return result
	case http.StatusResetContent:
		return result
	case http.StatusPartialContent:
		return result
	case http.StatusMultiStatus:
		return result
	case http.StatusAlreadyReported:
		return result
	case http.StatusIMUsed:
		return result
	case http.StatusMultipleChoices:
		return result
	case http.StatusMovedPermanently:
		return result
	case http.StatusFound:
		return result
	case http.StatusSeeOther:
		return result
	case http.StatusNotModified:
		return result
	case http.StatusUseProxy:
		return result
	case http.StatusTemporaryRedirect:
		return result
	case http.StatusPermanentRedirect:
		return result
	case http.StatusBadRequest:
		return result
	case http.StatusUnauthorized:
		return result
	case http.StatusPaymentRequired:
		return result
	case http.StatusForbidden:
		return result
	case http.StatusNotFound:
		return result
	case http.StatusMethodNotAllowed:
		return result
	case http.StatusNotAcceptable:
		return result
	case http.StatusProxyAuthRequired:
		return result
	case http.StatusRequestTimeout:
		return result
	case http.StatusConflict:
		return result
	case http.StatusGone:
		return result
	case http.StatusLengthRequired:
		return result
	case http.StatusPreconditionFailed:
		return result
	case http.StatusRequestEntityTooLarge:
		return result
	case http.StatusRequestURITooLong:
		return result
	case http.StatusUnsupportedMediaType:
		return result
	case http.StatusRequestedRangeNotSatisfiable:
		return result
	case http.StatusExpectationFailed:
		return result
	case http.StatusTeapot:
		return result
	// case http.StatusMisdirectedRequest:
	// 	return result
	case http.StatusUnprocessableEntity:
		return result
	case http.StatusLocked:
		return result
	case http.StatusFailedDependency:
		return result
	case http.StatusUpgradeRequired:
		return result
	case http.StatusPreconditionRequired:
		return result
	case http.StatusTooManyRequests:
		return result
	case http.StatusRequestHeaderFieldsTooLarge:
		return result
	case http.StatusUnavailableForLegalReasons:
		return result
	case http.StatusInternalServerError:
		return result
	case http.StatusNotImplemented:
		return result
	case http.StatusBadGateway:
		return result
	case http.StatusServiceUnavailable:
		return result
	case http.StatusGatewayTimeout:
		return result
	case http.StatusHTTPVersionNotSupported:
		return result
	case http.StatusVariantAlsoNegotiates:
		return result
	case http.StatusInsufficientStorage:
		return result
	case http.StatusLoopDetected:
		return result
	case http.StatusNotExtended:
		return result
	case http.StatusNetworkAuthenticationRequired:
		return result
	default:
		return "{\"requrl\":\"" + requestUrl + "\",\"stacod\":520,\"statxt\":\"Unknown Error\",\"sta\":false,\"type\":\"obj\",\"msg\":\"" + message + "\"}"
	}
}

func GetJsonStrRs(requestUrl string, httpStatus int, status bool, message string) string {

	result := "{\"requrl\":\"" + requestUrl + "\",\"stacod\":" + strconv.Itoa(httpStatus) + ",\"statxt\":\"" + http.StatusText(httpStatus) + "\",\"sta\":" + strconv.FormatBool(status) + ",\"type\":\"str\",\"msg\":\"" + message + "\"}"
	switch httpStatus {
	case http.StatusContinue:
		return result
	case http.StatusSwitchingProtocols:
		return result
	case http.StatusProcessing:
		return result
	case http.StatusOK:
		return result
	case http.StatusCreated:
		return result
	case http.StatusAccepted:
		return result
	case http.StatusNonAuthoritativeInfo:
		return result
	case http.StatusNoContent:
		return result
	case http.StatusResetContent:
		return result
	case http.StatusPartialContent:
		return result
	case http.StatusMultiStatus:
		return result
	case http.StatusAlreadyReported:
		return result
	case http.StatusIMUsed:
		return result
	case http.StatusMultipleChoices:
		return result
	case http.StatusMovedPermanently:
		return result
	case http.StatusFound:
		return result
	case http.StatusSeeOther:
		return result
	case http.StatusNotModified:
		return result
	case http.StatusUseProxy:
		return result
	case http.StatusTemporaryRedirect:
		return result
	case http.StatusPermanentRedirect:
		return result
	case http.StatusBadRequest:
		return result
	case http.StatusUnauthorized:
		return result
	case http.StatusPaymentRequired:
		return result
	case http.StatusForbidden:
		return result
	case http.StatusNotFound:
		return result
	case http.StatusMethodNotAllowed:
		return result
	case http.StatusNotAcceptable:
		return result
	case http.StatusProxyAuthRequired:
		return result
	case http.StatusRequestTimeout:
		return result
	case http.StatusConflict:
		return result
	case http.StatusGone:
		return result
	case http.StatusLengthRequired:
		return result
	case http.StatusPreconditionFailed:
		return result
	case http.StatusRequestEntityTooLarge:
		return result
	case http.StatusRequestURITooLong:
		return result
	case http.StatusUnsupportedMediaType:
		return result
	case http.StatusRequestedRangeNotSatisfiable:
		return result
	case http.StatusExpectationFailed:
		return result
	case http.StatusTeapot:
		return result
	// case http.StatusMisdirectedRequest:
	// 	return result
	case http.StatusUnprocessableEntity:
		return result
	case http.StatusLocked:
		return result
	case http.StatusFailedDependency:
		return result
	case http.StatusUpgradeRequired:
		return result
	case http.StatusPreconditionRequired:
		return result
	case http.StatusTooManyRequests:
		return result
	case http.StatusRequestHeaderFieldsTooLarge:
		return result
	case http.StatusUnavailableForLegalReasons:
		return result
	case http.StatusInternalServerError:
		return result
	case http.StatusNotImplemented:
		return result
	case http.StatusBadGateway:
		return result
	case http.StatusServiceUnavailable:
		return result
	case http.StatusGatewayTimeout:
		return result
	case http.StatusHTTPVersionNotSupported:
		return result
	case http.StatusVariantAlsoNegotiates:
		return result
	case http.StatusInsufficientStorage:
		return result
	case http.StatusLoopDetected:
		return result
	case http.StatusNotExtended:
		return result
	case http.StatusNetworkAuthenticationRequired:
		return result
	default:
		return "{\"requrl\":\"" + requestUrl + "\",\"stacod\":520,\"statxt\":\"Unknown Error\",\"sta\":false,\"type\":\"obj\",\"msg\":\"" + message + "\"}"
	}

}
