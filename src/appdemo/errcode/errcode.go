package errcode

import (
	"sync"
)

// DefaultAPICodeMap 错误配置
var DefaultAPICodeMap *APICodeMap = &APICodeMap{
	APIMap: make(map[string]*APILogCode),
}

// APILogCode ...
type APILogCode struct {
	NameCn   string `json:"name_cn"`
	NameEn   string `json:"name_en"`
	Text     string `json:"msg_name"`
	FullCode string `json:"full_code"`
}

// APICodeMap ...
type APICodeMap struct {
	sync.RWMutex
	APIMap map[string]*APILogCode
}

// GetLocalCodeMsg 获取错误信息
func (c *APICodeMap) GetLocalCodeMsg(code uint64) *APILogCode {
	c.RLock()
	defer c.RUnlock()

	return &APILogCode{
		NameCn:   "错误",
		NameEn:   "err code detail",
		Text:     "详细错误信息",
		FullCode: "1111222233334444",
	}

	// TODO
	// return c.APIMap[codeStr]
}

// GetLocalCodeMsg ...
func GetLocalCodeMsg(code uint64) *APILogCode {
	return DefaultAPICodeMap.GetLocalCodeMsg(code)
}

// LoadErrCode ...
func LoadErrCode() (err error) {
	return DefaultAPICodeMap.LoadErrCode()
}

// LoadErrCode load err code
func (c *APICodeMap) LoadErrCode() (err error) {
	// TODO load err 详细信息
	return
}
