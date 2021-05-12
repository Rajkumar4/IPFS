package store

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"os"

	badger "github.com/dgraph-io/badger/v3"
	logger "github.com/ipfs/go-log/v2"
)

type File struct {
	Name string
}

type ContentBased struct {
	Id string
}

func (cb *ContentBased) GetID() string {
	return ""
}

func (f *File) GetID() string {
	return fmt.Sprintf("badgeripfs_%s", f.Name)
}

type StoreConf struct {
	DB *badger.DB
}

type store interface {
	Create(Item) error
	Read(Item) Item
	Delete(Item) error
	Update(Item) error
}

type Item interface {
	GetID() string
}

type exchange struct {
	name string
}

func marhsal(data *exchange) ([]byte, error) {
	return []byte(data.name), nil
}

func unMarshal(data []byte, ex *exchange) error {
	ex.name = string(data)
	return nil
}

var log = logger.Logger("badger")

func genCid(val []byte) (string, error) {
	hash := sha256.New()
	hash.Write(val)
	en := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return en, nil
}

func DBIntialization() (*badger.DB, error) {
	defer os.RemoveAll("./db/ipfsdb")
	db, err := badger.Open(badger.DefaultOptions("./db/ipfsdb"))
	if err != nil {
		log.Errorf("Failed to open db %s", err.Error())
		return nil, err
	}
	return db, nil
}

func (ss *StoreConf) Create(item Item) (string, error) {
	var id string
	err := ss.DB.Update(func(tx *badger.Txn) error {
		file, ok := item.(*File)
		if !ok {
			log.Errorf("File is not item type")
			return errors.New("File is not item type")
		}
		var ex = &exchange{
			name: file.Name,
		}
		value, err := marhsal(ex)
		if err != nil {
			log.Errorf("Failed to marshal value %s", err.Error())
			return err
		}
		hash, err := genCid(value)
		if err != nil {
			log.Errorf("Failed to get cid %s", err.Error())
			return err
		}
		entry := badger.NewEntry([]byte(hash), value)
		err = tx.SetEntry(entry)
		if err != nil {
			log.Errorf("Failed to set entry in database %s", err.Error())
			return err
		}
		id = hash
		return nil
	})
	if err != nil {
		log.Errorf("Failed to create entry in db %s", err.Error())
		return id, err
	}
	return id, nil
}

func (ss *StoreConf) Read(item Item) (Item, error) {
	var ex = &exchange{}
	err := ss.DB.View(func(tx *badger.Txn) error {
		f, ok := item.(*ContentBased)
		if !ok {
			return errors.New("item is not type of file")
		}

		value, err := tx.Get([]byte(f.Id))
		if err != nil {
			log.Errorf("Failed to get item from db %s", err.Error())
			return err
		}
		var fileinfo []byte
		err = value.Value(func(val []byte) error {
			fileinfo = val
			return nil
		})
		err = unMarshal(fileinfo, ex)
		if err != nil {
			log.Errorf("Failedto unmarshal data %s", err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		log.Errorf("Failed to read item %s", err.Error())
		return nil, err
	}
	return &File{Name: ex.name}, nil
}

func (ss *StoreConf) Delete(item Item) error {
	return ss.DB.Update(func(tx *badger.Txn) error {
		key, err := marhsal(&exchange{
			name: item.GetID(),
		})
		if err != nil {
			log.Errorf("Failed to marshal key %s", err.Error())
			return err
		}
		if err = tx.Delete(key); err != nil {
			log.Errorf("Failed to delete item %s", err.Error())
			return err
		}
		return nil
	})
}

func (ss *StoreConf) Update(item Item) error {
	return ss.DB.Update(func(tx *badger.Txn) error {
		key, err := marhsal(&exchange{
			name: item.GetID(),
		})
		if err != nil {
			log.Errorf("Failed to marshal key %s", err.Error())
			return err
		}
		file, ok := item.(*File)
		if !ok {
			log.Errorf("File is not item type")
			return errors.New("File is not item type")
		}
		value, err := marhsal(&exchange{name: file.Name})
		if err != nil {
			log.Errorf("Failed to marshal value %s", err.Error())
			return err
		}
		err = tx.Set(key, value)
		if err != nil {
			log.Errorf("Failed to set update item %s", err.Error())
			return err
		}
		return nil
	})
}
