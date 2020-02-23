// WARNING: THIS FILE IS GENERATED USING gen/main.go. Edit that script and regenerate this file.

package erkhttp

// StatusContinue = 100
type StatusContinue struct{}

func (StatusContinue) HTTPStatus() int { return 100 }

var _ HTTPStatusable = StatusContinue{}

// StatusSwitchingProtocols = 101
type StatusSwitchingProtocols struct{}

func (StatusSwitchingProtocols) HTTPStatus() int { return 101 }

var _ HTTPStatusable = StatusSwitchingProtocols{}

// StatusProcessing = 102
type StatusProcessing struct{}

func (StatusProcessing) HTTPStatus() int { return 102 }

var _ HTTPStatusable = StatusProcessing{}

// StatusEarlyHints = 103
type StatusEarlyHints struct{}

func (StatusEarlyHints) HTTPStatus() int { return 103 }

var _ HTTPStatusable = StatusEarlyHints{}

// StatusOK = 200
type StatusOK struct{}

func (StatusOK) HTTPStatus() int { return 200 }

var _ HTTPStatusable = StatusOK{}

// StatusCreated = 201
type StatusCreated struct{}

func (StatusCreated) HTTPStatus() int { return 201 }

var _ HTTPStatusable = StatusCreated{}

// StatusAccepted = 202
type StatusAccepted struct{}

func (StatusAccepted) HTTPStatus() int { return 202 }

var _ HTTPStatusable = StatusAccepted{}

// StatusNonAuthoritativeInfo = 203
type StatusNonAuthoritativeInfo struct{}

func (StatusNonAuthoritativeInfo) HTTPStatus() int { return 203 }

var _ HTTPStatusable = StatusNonAuthoritativeInfo{}

// StatusNoContent = 204
type StatusNoContent struct{}

func (StatusNoContent) HTTPStatus() int { return 204 }

var _ HTTPStatusable = StatusNoContent{}

// StatusResetContent = 205
type StatusResetContent struct{}

func (StatusResetContent) HTTPStatus() int { return 205 }

var _ HTTPStatusable = StatusResetContent{}

// StatusPartialContent = 206
type StatusPartialContent struct{}

func (StatusPartialContent) HTTPStatus() int { return 206 }

var _ HTTPStatusable = StatusPartialContent{}

// StatusMultiStatus = 207
type StatusMultiStatus struct{}

func (StatusMultiStatus) HTTPStatus() int { return 207 }

var _ HTTPStatusable = StatusMultiStatus{}

// StatusAlreadyReported = 208
type StatusAlreadyReported struct{}

func (StatusAlreadyReported) HTTPStatus() int { return 208 }

var _ HTTPStatusable = StatusAlreadyReported{}

// StatusIMUsed = 226
type StatusIMUsed struct{}

func (StatusIMUsed) HTTPStatus() int { return 226 }

var _ HTTPStatusable = StatusIMUsed{}

// StatusMultipleChoices = 300
type StatusMultipleChoices struct{}

func (StatusMultipleChoices) HTTPStatus() int { return 300 }

var _ HTTPStatusable = StatusMultipleChoices{}

// StatusMovedPermanently = 301
type StatusMovedPermanently struct{}

func (StatusMovedPermanently) HTTPStatus() int { return 301 }

var _ HTTPStatusable = StatusMovedPermanently{}

// StatusFound = 302
type StatusFound struct{}

func (StatusFound) HTTPStatus() int { return 302 }

var _ HTTPStatusable = StatusFound{}

// StatusSeeOther = 303
type StatusSeeOther struct{}

func (StatusSeeOther) HTTPStatus() int { return 303 }

var _ HTTPStatusable = StatusSeeOther{}

// StatusNotModified = 304
type StatusNotModified struct{}

func (StatusNotModified) HTTPStatus() int { return 304 }

var _ HTTPStatusable = StatusNotModified{}

// StatusUseProxy = 305
type StatusUseProxy struct{}

func (StatusUseProxy) HTTPStatus() int { return 305 }

var _ HTTPStatusable = StatusUseProxy{}

// StatusTemporaryRedirect = 307
type StatusTemporaryRedirect struct{}

func (StatusTemporaryRedirect) HTTPStatus() int { return 307 }

var _ HTTPStatusable = StatusTemporaryRedirect{}

// StatusPermanentRedirect = 308
type StatusPermanentRedirect struct{}

func (StatusPermanentRedirect) HTTPStatus() int { return 308 }

var _ HTTPStatusable = StatusPermanentRedirect{}

// StatusBadRequest = 400
type StatusBadRequest struct{}

func (StatusBadRequest) HTTPStatus() int { return 400 }

var _ HTTPStatusable = StatusBadRequest{}

// StatusUnauthorized = 401
type StatusUnauthorized struct{}

func (StatusUnauthorized) HTTPStatus() int { return 401 }

var _ HTTPStatusable = StatusUnauthorized{}

// StatusPaymentRequired = 402
type StatusPaymentRequired struct{}

func (StatusPaymentRequired) HTTPStatus() int { return 402 }

var _ HTTPStatusable = StatusPaymentRequired{}

// StatusForbidden = 403
type StatusForbidden struct{}

func (StatusForbidden) HTTPStatus() int { return 403 }

var _ HTTPStatusable = StatusForbidden{}

// StatusNotFound = 404
type StatusNotFound struct{}

func (StatusNotFound) HTTPStatus() int { return 404 }

var _ HTTPStatusable = StatusNotFound{}

// StatusMethodNotAllowed = 405
type StatusMethodNotAllowed struct{}

func (StatusMethodNotAllowed) HTTPStatus() int { return 405 }

var _ HTTPStatusable = StatusMethodNotAllowed{}

// StatusNotAcceptable = 406
type StatusNotAcceptable struct{}

func (StatusNotAcceptable) HTTPStatus() int { return 406 }

var _ HTTPStatusable = StatusNotAcceptable{}

// StatusProxyAuthRequired = 407
type StatusProxyAuthRequired struct{}

func (StatusProxyAuthRequired) HTTPStatus() int { return 407 }

var _ HTTPStatusable = StatusProxyAuthRequired{}

// StatusRequestTimeout = 408
type StatusRequestTimeout struct{}

func (StatusRequestTimeout) HTTPStatus() int { return 408 }

var _ HTTPStatusable = StatusRequestTimeout{}

// StatusConflict = 409
type StatusConflict struct{}

func (StatusConflict) HTTPStatus() int { return 409 }

var _ HTTPStatusable = StatusConflict{}

// StatusGone = 410
type StatusGone struct{}

func (StatusGone) HTTPStatus() int { return 410 }

var _ HTTPStatusable = StatusGone{}

// StatusLengthRequired = 411
type StatusLengthRequired struct{}

func (StatusLengthRequired) HTTPStatus() int { return 411 }

var _ HTTPStatusable = StatusLengthRequired{}

// StatusPreconditionFailed = 412
type StatusPreconditionFailed struct{}

func (StatusPreconditionFailed) HTTPStatus() int { return 412 }

var _ HTTPStatusable = StatusPreconditionFailed{}

// StatusRequestEntityTooLarge = 413
type StatusRequestEntityTooLarge struct{}

func (StatusRequestEntityTooLarge) HTTPStatus() int { return 413 }

var _ HTTPStatusable = StatusRequestEntityTooLarge{}

// StatusRequestURITooLong = 414
type StatusRequestURITooLong struct{}

func (StatusRequestURITooLong) HTTPStatus() int { return 414 }

var _ HTTPStatusable = StatusRequestURITooLong{}

// StatusUnsupportedMediaType = 415
type StatusUnsupportedMediaType struct{}

func (StatusUnsupportedMediaType) HTTPStatus() int { return 415 }

var _ HTTPStatusable = StatusUnsupportedMediaType{}

// StatusRequestedRangeNotSatisfiable = 416
type StatusRequestedRangeNotSatisfiable struct{}

func (StatusRequestedRangeNotSatisfiable) HTTPStatus() int { return 416 }

var _ HTTPStatusable = StatusRequestedRangeNotSatisfiable{}

// StatusExpectationFailed = 417
type StatusExpectationFailed struct{}

func (StatusExpectationFailed) HTTPStatus() int { return 417 }

var _ HTTPStatusable = StatusExpectationFailed{}

// StatusTeapot = 418
type StatusTeapot struct{}

func (StatusTeapot) HTTPStatus() int { return 418 }

var _ HTTPStatusable = StatusTeapot{}

// StatusMisdirectedRequest = 421
type StatusMisdirectedRequest struct{}

func (StatusMisdirectedRequest) HTTPStatus() int { return 421 }

var _ HTTPStatusable = StatusMisdirectedRequest{}

// StatusUnprocessableEntity = 422
type StatusUnprocessableEntity struct{}

func (StatusUnprocessableEntity) HTTPStatus() int { return 422 }

var _ HTTPStatusable = StatusUnprocessableEntity{}

// StatusLocked = 423
type StatusLocked struct{}

func (StatusLocked) HTTPStatus() int { return 423 }

var _ HTTPStatusable = StatusLocked{}

// StatusFailedDependency = 424
type StatusFailedDependency struct{}

func (StatusFailedDependency) HTTPStatus() int { return 424 }

var _ HTTPStatusable = StatusFailedDependency{}

// StatusTooEarly = 425
type StatusTooEarly struct{}

func (StatusTooEarly) HTTPStatus() int { return 425 }

var _ HTTPStatusable = StatusTooEarly{}

// StatusUpgradeRequired = 426
type StatusUpgradeRequired struct{}

func (StatusUpgradeRequired) HTTPStatus() int { return 426 }

var _ HTTPStatusable = StatusUpgradeRequired{}

// StatusPreconditionRequired = 428
type StatusPreconditionRequired struct{}

func (StatusPreconditionRequired) HTTPStatus() int { return 428 }

var _ HTTPStatusable = StatusPreconditionRequired{}

// StatusTooManyRequests = 429
type StatusTooManyRequests struct{}

func (StatusTooManyRequests) HTTPStatus() int { return 429 }

var _ HTTPStatusable = StatusTooManyRequests{}

// StatusRequestHeaderFieldsTooLarge = 431
type StatusRequestHeaderFieldsTooLarge struct{}

func (StatusRequestHeaderFieldsTooLarge) HTTPStatus() int { return 431 }

var _ HTTPStatusable = StatusRequestHeaderFieldsTooLarge{}

// StatusUnavailableForLegalReasons = 451
type StatusUnavailableForLegalReasons struct{}

func (StatusUnavailableForLegalReasons) HTTPStatus() int { return 451 }

var _ HTTPStatusable = StatusUnavailableForLegalReasons{}

// StatusInternalServerError = 500
type StatusInternalServerError struct{}

func (StatusInternalServerError) HTTPStatus() int { return 500 }

var _ HTTPStatusable = StatusInternalServerError{}

// StatusNotImplemented = 501
type StatusNotImplemented struct{}

func (StatusNotImplemented) HTTPStatus() int { return 501 }

var _ HTTPStatusable = StatusNotImplemented{}

// StatusBadGateway = 502
type StatusBadGateway struct{}

func (StatusBadGateway) HTTPStatus() int { return 502 }

var _ HTTPStatusable = StatusBadGateway{}

// StatusServiceUnavailable = 503
type StatusServiceUnavailable struct{}

func (StatusServiceUnavailable) HTTPStatus() int { return 503 }

var _ HTTPStatusable = StatusServiceUnavailable{}

// StatusGatewayTimeout = 504
type StatusGatewayTimeout struct{}

func (StatusGatewayTimeout) HTTPStatus() int { return 504 }

var _ HTTPStatusable = StatusGatewayTimeout{}

// StatusHTTPVersionNotSupported = 505
type StatusHTTPVersionNotSupported struct{}

func (StatusHTTPVersionNotSupported) HTTPStatus() int { return 505 }

var _ HTTPStatusable = StatusHTTPVersionNotSupported{}

// StatusVariantAlsoNegotiates = 506
type StatusVariantAlsoNegotiates struct{}

func (StatusVariantAlsoNegotiates) HTTPStatus() int { return 506 }

var _ HTTPStatusable = StatusVariantAlsoNegotiates{}

// StatusInsufficientStorage = 507
type StatusInsufficientStorage struct{}

func (StatusInsufficientStorage) HTTPStatus() int { return 507 }

var _ HTTPStatusable = StatusInsufficientStorage{}

// StatusLoopDetected = 508
type StatusLoopDetected struct{}

func (StatusLoopDetected) HTTPStatus() int { return 508 }

var _ HTTPStatusable = StatusLoopDetected{}

// StatusNotExtended = 510
type StatusNotExtended struct{}

func (StatusNotExtended) HTTPStatus() int { return 510 }

var _ HTTPStatusable = StatusNotExtended{}

// StatusNetworkAuthenticationRequired = 511
type StatusNetworkAuthenticationRequired struct{}

func (StatusNetworkAuthenticationRequired) HTTPStatus() int { return 511 }

var _ HTTPStatusable = StatusNetworkAuthenticationRequired{}
