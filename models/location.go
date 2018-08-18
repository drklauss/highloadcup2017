package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

type Location struct {
	Id       uint32 `json:"id"`
	Place    string `json:"place"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Distance uint32 `json:"distance"`
}

// Кэш данных достопримечательностей
type Locations struct {
	mx sync.RWMutex
	m  map[uint32]*Location
}

func (ls *Locations) Init() *Locations {
	ls.m = make(map[uint32]*Location)
	ls.readData()
	return ls
}

func (ls *Locations) Get(id uint32) *Location {
	ls.mx.RLock()
	defer ls.mx.RUnlock()
	return ls.m[id]
}

func (ls *Locations) Save(l *Location) {
	ls.mx.Lock()
	ls.m[l.Id] = l
	ls.mx.Unlock()
}

func (ls *Locations) Update(id uint32, l *Location) {
	ls.mx.Lock()
	l.Id = id
	ls.m[id] = l
	ls.mx.Unlock()
}

func (ls *Locations) readData() {
	t := time.Now()
	count := 1
	for {
		fName := fmt.Sprintf("data/locations_%d.json", count)
		fmt.Println(fName)
		b, err := ioutil.ReadFile(fName)
		if err != nil {
			println(err.Error())
			break
		}
		var locations struct {
			Locations []Location `json:"locations"`
		}
		err = json.Unmarshal(b, &locations)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		for _, l := range locations.Locations {
			ls.m[l.Id] = &l
		}
		count++
	}
	fmt.Printf("All Locs: %d\nLocations Time: %+v\n", len(ls.m), time.Since(t))
}
