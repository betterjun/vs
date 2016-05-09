package bdb

import (
	"encoding/binary"
	"fmt"
	"github.com/boltdb/bolt"
	"os"
)

/*
db对象
*/
type BoltDB interface {
	Open(dbname string, mode os.FileMode) error // 打开
	Close()                                     // 关闭
	CreateTable(tn string) error                // 创建一张表
	DeleteTable(tn string) error                // 删除一张表
	GetDBName() string                          // 获取数据库名

	Set(tn string, key, value interface{}) // 设置键值,key,value只支持int64,string,[]byte
	Get(tn string, key interface{}) []byte // 获取键值
	Delete(tn string, key interface{})     // 删除键
}

// 实现BoltDB接口
type dbConnection struct {
	name string   // 数据库名字
	bdb  *bolt.DB // 数据库连接对象
}

// 打开一个数据库对象
func Open(db string, mode os.FileMode) BoltDB {
	bdb := &dbConnection{name: db}
	bdb.Open(db, mode)
	return bdb
}

func (b *dbConnection) Open(dbname string, mode os.FileMode) error {
	db, err := bolt.Open(dbname, mode, nil)
	if err != nil {
		return err
	}
	b.bdb = db
	return nil
}

func (b *dbConnection) Close() {
	if b.bdb != nil {
		b.bdb.Close()
	}
}

func (b *dbConnection) CreateTable(tn string) error {
	if b.bdb != nil {
		return b.bdb.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(tn))
			if err != nil {
				return fmt.Errorf("create bucket (%v) failed: %s", tn, err)
			}
			return nil
		})
	}

	return fmt.Errorf("invalid boltdb connection")
}

func (b *dbConnection) DeleteTable(tn string) error {
	if b.bdb != nil {
		return b.bdb.Update(func(tx *bolt.Tx) error {
			err := tx.DeleteBucket([]byte(tn))
			if err != nil {
				return fmt.Errorf("delete bucket (%v) failed: %s", tn, err)
			}
			return nil
		})
	}

	return fmt.Errorf("invalid boltdb connection")
}

func (b *dbConnection) GetDBName() string {
	return b.name
}

func (b *dbConnection) Set(tn string, key, value interface{}) {
	b.bdb.Update(func(tx *bolt.Tx) error {
		k, err := processData(key)
		if err != nil {
			panic(fmt.Errorf("invalid key:%v", err))
		}
		v, err := processData(value)
		if err != nil {
			panic(fmt.Errorf("invalid value:%v", err))
		}

		bucket := tx.Bucket([]byte(tn))
		err = bucket.Put(k, v)
		if err != nil {
			fmt.Printf("set %v.%v failed: %v\n", tn, k, err)
		}
		return err
	})
}

func (b *dbConnection) Get(tn string, key interface{}) (ret []byte) {
	b.bdb.Update(func(tx *bolt.Tx) error {
		k, err := processData(key)
		if err != nil {
			panic(fmt.Errorf("invalid key:%v", err))
		}

		bucket := tx.Bucket([]byte(tn))
		v := bucket.Get(k)
		// do make space before copy
		ret = make([]byte, len(v))
		copy(ret, v)
		return nil
	})
	return ret
}

func (b *dbConnection) Delete(tn string, key interface{}) {
	b.bdb.Update(func(tx *bolt.Tx) error {
		k, err := processData(key)
		if err != nil {
			panic(fmt.Errorf("invalid key:%v", err))
		}

		bucket := tx.Bucket([]byte(tn))
		bucket.Delete(k)
		return nil
	})
}

// 处理支持的key，value类型
func processData(value interface{}) (v []byte, err error) {
	switch val := value.(type) {
	case string:
		v = []byte(val)
	case []byte:
		v = val
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		v = []byte(fmt.Sprintf("%d", val))
	case float64, float32:
		v = []byte(fmt.Sprintf("%f", val))
	default:
		err = fmt.Errorf("non integer, string, []byte types")
	}
	return
}

// btoi returns an 8-byte big endian representation of v.
func btoi(b []byte) (int64, error) {
	i, n := binary.Uvarint(b)
	if n < 0 {
		return 0, fmt.Errorf("overflow: value larger than 64 bits")
	} else if n == 0 {
		return 0, fmt.Errorf("buf too small")
	}

	return int64(i), nil
}

// itob returns an 8-byte big endian representation of v.
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
