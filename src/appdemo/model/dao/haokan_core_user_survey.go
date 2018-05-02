package dao

import (
	"bdlib/mysql"
	"bdlib/util"
	"bytes"
	"fargo"
	"fmt"
	"io"
)

var (
	// TABLE_CORE_USER_SURVEY table name
	TABLE_CORE_USER_SURVEY = "haokan_core_user_survey"
)

// CoreUserSurvey table info
type CoreUserSurvey struct {
	ID      int64  `db:"id" json:"id,omitempty"`             // id
	CUID    string `db:"cuid" json:"cuid,omitempty"`         // cuid
	Content string `db:"content" json:"content,omitempty"`   // content
	AddTime int64  `db:"add_time" json:"add_time,omitempty"` // 添加时间
}

// Insert db为model对象
func Insert(db *mysql.DB, info *CoreUserSurvey) (err error) {
	return
}

// Update _
func Update(db *mysql.DB, info *CoreUserSurvey) (err error) {
	return
}

// Delete _
func Delete(db *mysql.DB, id int64) (err error) {
	return
}

// Select _
func Select(db *mysql.DB, fields, where string, orderByLimit string) (list []*CoreUserSurvey, err error) {
	sqlFind := bytes.Buffer{}
	sqlFind.WriteString(fmt.Sprintf("SELECT %s FROM %s WHERE 1=1 ", fields, TABLE_CORE_USER_SURVEY))
	sqlFind.WriteString(where)
	sqlFind.WriteString(" ")
	sqlFind.WriteString(orderByLimit)

	if err = db.Query(sqlFind.String()); err != nil {
		fargo.Error(err)
		return
	}

	rows, err := db.FetchAllMap()
	if err == io.EOF {
		err = nil
		return
	}
	if err != nil {
		fargo.Error(err)
		return
	}

	list = make([]*CoreUserSurvey, len(rows))
	for i, row := range rows {
		list[i] = &CoreUserSurvey{
			ID:      util.Int64(row["id"]),
			CUID:    row["id"],
			Content: row["content"],
			AddTime: util.Int64(row["id"]),
		}
	}

	// TODO orm 映射对象

	return
}

// Query _
func Query(db *mysql.DB, query string) (err error) {
	return
}
