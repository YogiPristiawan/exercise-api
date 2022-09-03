package entities

import "log"

type CommonResult struct {
	ResCode    int    `json:"-"`
	ResMessage string `json:"-"`
}

func (c *CommonResult) SetCode(code int, err error) {
	switch true {
	case code >= 400 && code < 500:
		c.ResCode = code
		c.ResMessage = err.Error()
	case code >= 500:
		log.Print(err)
		c.ResCode = code
		c.ResMessage = "internal server error"
	default:
		c.ResCode = code
	}
}
