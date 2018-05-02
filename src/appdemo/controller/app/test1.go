package app

import (
	appmodel "appdemo/model/data"
)

func (c *Controller) test1() {

	param1 := c.GetString("param1")
	param2, _ := c.GetInt("param2")
	c.Infof("get param1: %s, param2: %d ", param1, param2)

	list, err := appmodel.GetList()
	if err != nil {
		c.Error(err)
		c.ReplyErrCode(1)
		return
	}

	c.ReplySucc(list)
	return
}
