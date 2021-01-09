// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package wf_business_checker

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table wf_business_checker.
type Entity struct {
	Id           uint64 `orm:"id,primary"    json:"id"`            // 主键
	FromTable    string `orm:"from_table"    json:"from_table"`    // 业务表名
	FromId       uint64 `orm:"from_id"       json:"from_id"`       // 业务id
	UserId       string `orm:"user_id"       json:"user_id"`       // 用户id
	DepartmentId string `orm:"department_id" json:"department_id"` // 部门id
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}