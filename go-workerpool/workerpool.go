package go_workerpool

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"github.com/tungct/go-messqueue"
)

var MaxLenWorker int = 20
// Worker pool
var Worker chan(int)

func WriteToDisk(id int) bool{
	Worker <-id
	message := <-go_messqueue.Queue
	fmt.Println("Worker ", id, "execute Message")

	// check exits file output
	if _, err := os.Stat("output.json"); err == nil {
		f, _ := os.OpenFile("output.json", os.O_APPEND|os.O_WRONLY, 0600)
		rs, _ := json.Marshal(message)
		if _, err := f.Write(rs); err != nil {
			panic(err)
		}
		return true
	}else {
		jsonData, _  := json.Marshal(message)
		ioutil.WriteFile("output.json", jsonData, 0600)
		return true
	}
	return false
}