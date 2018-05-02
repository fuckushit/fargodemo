package data

import (
	"appdemo/model"
	"appdemo/model/dao"
	"bytes"
	"fargo"
)

// GetList _
func GetList() (rs []*dao.CoreUserSurvey, err error) {
	db := model.GetDefaultDB()
	defer model.PutDefaultDB(db)

	where := bytes.Buffer{}
	where.WriteString("AND `cuid` <> ''")
	if rs, err = dao.Select(db, "*", where.String(), "order by id desc limit 10"); err != nil {
		fargo.Error(err)
		return
	}

	return
}
