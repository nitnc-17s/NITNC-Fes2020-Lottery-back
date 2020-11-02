package lottery

import (
	"fmt"
	"sync"
)

// JSONData is JSON of data (count etc)
type JSONData struct {
	Count int `json:"count"`
}

// GetJSON is convert data to JSON
func (data *Data) GetJSON() JSONData {
	data.RLock()
	defer data.RUnlock()
	return JSONData{
		Count: data.count,
	}
}

// Data is clicker count data :)
type Data struct {
	count int
	sync.RWMutex
}

func (data *Data) AddCount(cnt int) {
	data.Lock()
	defer data.Unlock()
	fmt.Println(data.count)

	data.count += cnt
}
