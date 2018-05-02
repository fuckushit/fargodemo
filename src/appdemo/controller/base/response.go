package base

import (
	"appdemo/errcode"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
	"sync"
)

// RespHead shopapi返回结果头
type RespHead struct {
	Code int         `json:"code"`
	Info string      `json:"info"`
	Desc string      `json:"desc"`
	Ext  interface{} `json:"ext,omitempty"`
}

// Resp 返回resp
type Resp struct {
	Code int         `json:"code"`
	Info string      `json:"info"`
	Desc string      `json:"desc"`
	Data interface{} `json:"data"`
}

// Reply304 返回304
func (c *Controller) Reply304() {
	c.LogMap["304"] = "true"
	c.Ctx.ResponseWriter.WriteHeader(304)
	return

}

func (c *Controller) replyContent(content []byte) {
	c.isReply = true
	sum := md5.Sum(content)
	etag := hex.EncodeToString(sum[:])
	reqETag := c.Ctx.Input.Header("If-None-Match")

	// TODO 判断页面内容相同，直接返回304
	c.LogMap["_page_size"] = strconv.FormatInt(int64(len(content)), 10)
	if etag == reqETag {
		c.Ctx.ResponseWriter.WriteHeader(304)
		c.Reply304()
		return
	}

	header := c.Ctx.ResponseWriter.Header()
	header["ETag"] = []string{etag}
	if len(content) < 1024 {
		c.Ctx.Output.EnableGzip = false
	}

	c.gzipReply(etag, content)
	return
}

// ReplyErrCode ...
func (c *Controller) ReplyErrCode(code uint64) {

	// TODO 自定义code字典

	// err 内部对应错误
	// ext_err 外部错误head
	// ext_source 错误来源

	ins := errcode.GetLocalCodeMsg(code)
	if ins == nil {
		ins = &errcode.APILogCode{
			NameCn:   "错误code 未定义",
			NameEn:   "error code undefined",
			Text:     "服务器异常，请稍后重试",
			FullCode: strconv.FormatUint(code, 10),
		}
	}
	cont, _ := json.Marshal(ins)
	c.LogMap["err"] = string(cont)
	// TODO 回收cont

	code, _ = strconv.ParseUint(ins.FullCode, 10, 64)
	c.ActionCode = code

	// TODO 错误信息写到observer

	content, _ := json.Marshal(map[string]interface{}{
		"err": ins,
		//"ext_err":    c.LogMap["ext_err"],
		"ext_source": c.LogMap["ext_source"],
	})

	// TODO 回收[]byte

	head := RespHead{
		Code: int(code),
		Info: ins.NameEn,
		Desc: ins.Text,
	}
	content, _ = json.Marshal(head)
	c.replyContent(content)
}

// ReplySucc 返回成功
func (c *Controller) ReplySucc(data interface{}) {

	c.ActionResult = 1

	head := Resp{
		Code: 0,
		Info: "ok",
		Desc: "成功",
		Data: data,
	}
	content, _ := json.Marshal(&head)

	// TODO json pool
	c.replyContent(content)
}

var gzipPool, flatePool sync.Pool

// gzipReply TODO 将内容按json格式输出，gzip时做缓存优化，降低CPU消耗
func (c *Controller) gzipReply(etag string, content []byte) (err error) {

	m := c.Ctx.Output
	m.Header("Content-Type", "application/json; charset=utf-8")

	acceptEncoding := m.Context.Input.Header("Accept-Encoding")
	if m.EnableGzip && acceptEncoding != "" {
		splitted := strings.SplitN(acceptEncoding, ",", -1)
		encodings := make([]string, len(splitted))

		for i, val := range splitted {
			encodings[i] = strings.TrimSpace(val)
		}
		for _, val := range encodings {
			if val == "gzip" {
				m.Header("Content-Encoding", "gzip")
				break
			} else if val == "deflate" {
				m.Header("Content-Encoding", "deflate")
				break
			}
		}
	} else {
		m.Header("Content-Length", strconv.Itoa(len(content)))
	}

	m.Context.ResponseWriter.Write(content)

	// TODO 先从cache里寻找压缩以后的内容，节省CPU

	// TODO  区分用户个人信息与公共信息，个人信息放redis，公共信息放localcache

	// TODO 非压缩的没有存到localstore，可以回收

	return
}
