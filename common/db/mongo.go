package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/globalsign/mgo"
)

type MongoDB struct {
	Session *mgo.Session
	DBName  string
}

func NewMongoDB(conn string, dbName string, timeout time.Duration) (Mongo *MongoDB, err error) {
	var session *mgo.Session
	Mongo = &MongoDB{}
	session, err = mgo.DialWithTimeout(conn, timeout*time.Second)
	if err != nil {
		err = errors.New(fmt.Sprintf("Could not connect to %s: %s.", conn, err.Error()))
		return
	}
	session.SetMode(mgo.Strong, true)
	session.SetPoolLimit(100)
	Mongo.Session = session
	Mongo.DBName = dbName
	return
}

func (m *MongoDB) Close() {
	if m.Session != nil {
		m.Session.Close()
	}
}

func (m *MongoDB) C(collectionName string) *mgo.Collection {
	db := m.Session.DB(m.DBName)
	return db.C(collectionName)
}

func (m *MongoDB) DB() *mgo.Database {
	return m.Session.DB(m.DBName)
}

//能够明确是collection name
type CollName string

var (
	collIndexes = make(map[CollName][]mgo.Index)
)

func RegisterMgoIndex(coll CollName, index mgo.Index, rest ...mgo.Index) {
	indexes := collIndexes[coll]
	indexes = append(indexes, index)
	for _, idx := range rest {
		indexes = append(indexes, idx)
	}

	collIndexes[coll] = indexes
}

func EnsureIndex(mdb *mgo.Database) error {
	for coll, indexes := range collIndexes {
		for _, idx := range indexes {
			if err := mdb.C(string(coll)).EnsureIndex(idx); err != nil {
				return err
			}
		}
	}

	return nil
}
