package entities

type BaseResponse[T interface{}] struct {
	CommonResult
	Message string `json:"message"`
	Data    *T     `json:"data"`
}

type BaseResponseArray[T interface{}] struct {
	CommonResult
	Message string `json:"message"`
	Data    []*T   `json:"data"`
}
