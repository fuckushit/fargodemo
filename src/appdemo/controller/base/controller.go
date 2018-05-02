package base

import (
	"fargo"
	"fmt"
	"math/rand"
	"net/url"
	"runtime"
	"strings"
)

// Controller base controller
type Controller struct {
	fargo.Controller
	Model        string
	Action       string
	LogID        int64
	UserID       int64
	ActionResult uint32
	isReply      bool              // 是否有返回
	LogMap       map[string]string // 打本地日志的额外数据
	ActionCode   uint64
}

// Prepare 对请求做预处理
func (c *Controller) Prepare() {

	// 获取version, model, action
	c.parseURI()

	// 记录logid，便于查询
	c.LogID = rand.Int63()
	c.LogMap = make(map[string]string)

	// TODO 打点

	// TODO 区分渠道

	return
}

// Filter 过滤请求
func (c *Controller) Filter() bool {

	// 过滤非法请求
	if c.Model == "" || c.Action == "" {
		return false
	}

	// TODO 校验登陆

	return true
}

// Error 自带user_id的Error
func (c *Controller) Error(err error) {
	nerr := fmt.Errorf("ERROR: %d %s controller %v", c.LogID, c.Ctx.Input.Uri(), err)
	fargo.Log.PrintN(3, nerr)
	return
}

// Errorf 自带user_id的Errorf
func (c *Controller) Errorf(format string, args ...interface{}) {
	format = fmt.Sprintf("ERROR: %d %s controller %v", c.LogID, c.Ctx.Input.Uri(), format)
	fargo.Log.PrintfN(3, format, args...)
	return
}

// Info 自带user_id的Info
func (c *Controller) Info(err error) {
	nerr := fmt.Errorf("INFO: %d %s controller %v", c.LogID, c.Ctx.Input.Uri(), err)
	fargo.Log.PrintN(3, nerr)
	return
}

// Infof 自带user_id的Infof
func (c *Controller) Infof(format string, args ...interface{}) {
	format = fmt.Sprintf("INFO: %d %s controller %v", c.LogID, c.Ctx.Input.Uri(), format)
	fargo.Log.PrintfN(3, format, args...)
	return
}

// DumpStack panic打印堆栈(带user_id)
func (c *Controller) DumpStack() {
	cnt := 1
	fargo.Log.Printf("------- DumpStack --------")
	for {
		_, file, line, ok := runtime.Caller(cnt)
		if !ok {
			break
		}
		fargo.Log.Printf("%s:%d", file, line)
		cnt++
	}
}

// 提取version，model，action
func (c *Controller) parseURI() {
	uri := c.Ctx.Input.Uri()
	purl, err := url.Parse(uri)
	if err != nil {
		c.Error(err)
		return
	}
	fields := strings.Split(purl.Path, "/")
	if len(fields) < 3 {
		return
	}
	c.Model, c.Action = fields[1], fields[2]
}
