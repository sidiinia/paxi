package state

import (
//		"sync"
	//"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	//"encoding/binary"
	//"fmt"
	"fmt"
	//"encoding/binary"
	"strconv"
	"dlog"
)

type Operation uint8

const (
	NONE Operation = iota
	PUT
	GET
	DELETE
	RLOCK
	WLOCK
)

type Value int64

const NIL Value = 0

type Key int64

type Command struct {
	Op Operation
	K  Key
	V  Value
}

type State struct {
	//mutex *sync.Mutex
	//Store map[Key]Value
	DB *leveldb.DB
}

func InitState(id int) *State {
	   rs := strconv.Itoa(id)
	   d, err := leveldb.OpenFile("/Users/shannon/Desktop/dbtest" + rs, nil)

	   if err != nil {
	       fmt.Printf("Leveldb open failed: %v\n", err)
	   }

	   return &State{d}


	//return &State{new(sync.Mutex), make(map[Key]Value)}
}

func Conflict(gamma *Command, delta *Command) bool {
	if gamma.K == delta.K {
		if gamma.Op == PUT || delta.Op == PUT {
			return true
		}
	}
	return false
}

func ConflictBatch(batch1 []Command, batch2 []Command) bool {
	for i := 0; i < len(batch1); i++ {
		for j := 0; j < len(batch2); j++ {
			if Conflict(&batch1[i], &batch2[j]) {
				return true
			}
		}
	}
	return false
}

func IsRead(command *Command) bool {
	return command.Op == GET
}

func (c *Command) Execute(st *State) Value {
	fmt.Printf("Executing (%d, %d)\n", c.K, c.V)

	//var key, value [8]byte

	//    st.mutex.Lock()
	//    defer st.mutex.Unlock()

	switch c.Op {
	case PUT:

		   /*binary.BigEndian.PutUint64(key[:], uint64(c.K))
		   binary.BigEndian.PutUint64(value[:], uint64(c.V))
		   st.DB.Put(key[:], value[:], nil)*/
		//st.DB.Put(uint64(c.K), uint64(c.V), nil)
		//
		//dlog.Println("PUT command executed: key: ", c.K, " value: ", c.V)
		//defer st.DB.Close()
		return c.V
		/*st.Store[c.K] = c.V
		dlog.Println(fmt.Sprintf("state after PUT: %d\n", st.Store[c.K]))
		return c.V*/

	case GET:
		dlog.Println("GET command executing...")
		/*if val, present := st.Store[c.K]; present {
			dlog.Println(fmt.Sprintf("state after GET: %d\n", val))
			return val
		}*/
		/*binary.BigEndian.PutUint64(key[:], uint64(c.K))
		fmt.Println("key[:] is ", key[:])
		val, err := st.DB.Get(key[:], nil)

		if err == nil {
			val1 := binary.BigEndian.Uint64(val)
			fmt.Println("in state.go get v is ", Value(val1))
			fmt.Println("finished GET")
			//defer st.DB.Close()
			return Value(val1)  // returns the correct number here
		}*/
		//val, err := st.DB.Get(uint64(c.K), nil)
		//if err == nil {
		//	fmt.Println("in state.go get v is ", Value(val))
		//	fmt.Println("finished GET")
		//	return Value(val)  // returns the correct number here
		//}
	}

	defer st.DB.Close()
	return NIL
}
