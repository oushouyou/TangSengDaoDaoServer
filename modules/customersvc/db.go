package customersvc

import (
	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/pkg/db"
	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/pkg/util"
	"github.com/gocraft/dbr/v2"
)

// DB DB
type DB struct {
	session *dbr.Session
}

func newDB(session *dbr.Session) *DB {
	return &DB{
		session: session,
	}
}

func (d *DB) queryWithCustomerExternalNo(customerExternalNo string) (*model, error) {
	var m *model
	_, err := d.session.Select("*").From("customersvc_group").Where("customer_external_no=?", customerExternalNo).Load(&m)
	return m, err
}

func (d *DB) insert(m *model) error {
	_, err := d.session.InsertInto("customersvc_group").Columns(util.AttrToUnderscore(m)...).Record(m).Exec()
	return err
}

type model struct {
	CustomerExternalNo string
	UserId             string
	GroupId            string
	db.BaseModel
}
