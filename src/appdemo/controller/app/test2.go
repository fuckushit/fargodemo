package app

import (
	"fmt"
)

func (c *Controller) test2() {

	if err := fmt.Errorf("this is a fail demo"); err != nil {
		c.Error(err)
		c.ReplyErrCode(1)
		return
	}

	c.ReplySucc(true)
	return
}
