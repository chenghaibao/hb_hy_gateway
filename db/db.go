package db

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
)

var instance *Db

func GetDbInstance() *Db {
	if instance == nil {
		db, err := leveldb.OpenFile("storage/db", nil)
		if err != nil {
			panic(err)
		}
		instance = &Db{
			db: db,
		}
	}
	return instance
}

type Db struct {
	db *leveldb.DB
}

func (m *Db) Close() (err error) {
	err = m.db.Close()
	if err != nil {
		return err
	}

	m.db = nil
	return nil
}

func (m *Db) PutString(key, value string) (err error) {
	return m.db.Put([]byte(key), []byte(value), nil)
}

func (m *Db) PutStringByteValue(key string, value []byte) (err error) {
	return m.db.Put([]byte(key), value, nil)
}

func (m *Db) HasStringKey(key string) (has bool, err error) {
	return m.db.Has([]byte(key), nil)
}

func (m *Db) GetStringKey(key string) (value []byte, err error) {
	return m.db.Get([]byte(key), nil)
}

func (m *Db) GetStringKeyStringValue(key string) (value string, err error) {
	byteValue, err := m.GetStringKey(key)
	if err != nil {
		return "", err
	}

	return string(byteValue), nil
}

func (m *Db) DeleteStringKey(key string) (err error) {
	return m.db.Delete([]byte(key), nil)
}

func (m *Db) DeleteBytesKey(key []byte) (err error) {
	return m.db.Delete(key, nil)
}

func (m *Db) DeleteStringPrefixKey(prefix string) (err error) {
	iter := m.NewStringPrefixIterator(prefix)
	defer iter.Release()
	for iter.Next() {
		err := m.DeleteBytesKey(iter.Key())
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Db) NewStringPrefixIterator(prefixKey string) iterator.Iterator {
	return m.db.NewIterator(util.BytesPrefix([]byte(prefixKey)), nil)
}

func (m *Db) DebugDump() {
	iter := m.db.NewIterator(nil, nil)
	for iter.Next() {
		log.Printf("%s => %s", iter.Key(), iter.Value())
	}
}
