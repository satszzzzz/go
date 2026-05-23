package main1

import (
	"time"
	"sync"
	"fmt"
)

var dbData = []string{"db1", "db2", "db3", "db4", "db5"}
var results = []string{}
func main() {
	var wg = sync.WaitGroup{}

	for i:=0; i<len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Println("Results: ", results)

}

func dbCall(i int) {
	delay:=2000
	time.Sleep(time.Duration(delay*time.Millisecond))
	fmt.Println("Database called is ", dbData[i])
	results = append(results, dbData[i])
}