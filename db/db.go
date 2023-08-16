package db

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")

var db *bolt.DB

type Task struct {
	Key   int
	Value Entry
}

type Entry struct {
	Task        string
	Done        bool
	CompletedOn time.Time
}

func Init(dbPath string) error {
	var err error
	// We define this as it is so the db goes to the global package variable
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	// Creates our bucket (table) called as taskBucket
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})

}

func ClearDone() error {
	allTasks, err := GetTasks()
	if err != nil {
		return err
	}
	doneTasks := FilterDone(allTasks)
	for _, t := range doneTasks {
		if t.Value.CompletedOn.Add(24 * time.Hour).Before(time.Now()) {
			err := RmTask(t.Key)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func CreateTask(task string) (int, error) {
	var w bytes.Buffer
	encoder := gob.NewEncoder(&w)
	err := encoder.Encode(Entry{Task: task, Done: false, CompletedOn: time.Now()})
	if err != nil {
		return 0, err
	}
	var id int
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		key, err := b.NextSequence()
		if err != nil {
			id = 0
			return err
		}
		id = int(key)
		keyb := itob(id)
		return b.Put(keyb, w.Bytes())
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetTasks() ([]Task, error) {
	tasks := []Task{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		b.ForEach(func(k, v []byte) error {
			vbuffer := bytes.NewReader(v)
			decoder := gob.NewDecoder(vbuffer)
			var e Entry
			err := decoder.Decode(&e)
			if err != nil {
				log.Fatalln(err)
			}
			tasks = append(tasks, Task{Key: btoi(k), Value: e})
			return nil
		})
		return nil
	})
	return tasks, err
}

func DoTask(key int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		u_task := b.Get(itob(key))
		r := bytes.NewReader(u_task)
		decoder := gob.NewDecoder(r)
		var e Entry
		err := decoder.Decode(&e)
		if err != nil {
			return err
		}
		d_task := Entry{Task: e.Task, Done: true, CompletedOn: time.Now()}

		var w bytes.Buffer
		encoder := gob.NewEncoder(&w)
		err = encoder.Encode(d_task)
		if err != nil {
			return err
		}
		return b.Put(itob(key), w.Bytes())
	})
	return err
}

func RmTask(key int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
	return err
}

func itob(v int) []byte {
	byteSlice := make([]byte, 8)
	binary.BigEndian.PutUint64(byteSlice, uint64(v))
	return byteSlice
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
