package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

type Visit struct {
	Id        uint32 `json:"id"`
	Location  uint32 `json:"location"`
	User      uint32 `json:"user"`
	VisitedAt uint32 `json:"visited_at"`
	Mark      uint8  `json:"mark"`
}

// Кэш данных посещений
type Visits struct {
	mx sync.RWMutex
	m  map[uint32]Visit
}

func (vs *Visits) Init() *Visits {
	vs.m = make(map[uint32]Visit)
	vs.readData()
	return vs
}

func (vs *Visits) Get(id uint32) *Visit {
	vs.mx.RLock()
	defer vs.mx.RUnlock()
	if v, ok := vs.m[id]; ok {
		return &v
	}
	return nil
}

func (vs *Visits) Save(v *Visit) {
	vs.mx.Lock()
	vs.m[v.Id] = *v
	UvlCache.Save(v.User, v.Id)     // проверить
	LvlCache.Save(v.Location, v.Id) // проверить
	vs.mx.Unlock()
}

func (vs *Visits) Update(v *Visit) {
	vs.mx.Lock()
	vs.m[v.Id] = *v
	vs.mx.Unlock()
}

func (vs *Visits) readData() {
	t := time.Now()
	count := 1
	for {
		fName := fmt.Sprintf("data/visits_%d.json", count)
		fmt.Println(fName)
		b, err := ioutil.ReadFile(fName)
		if err != nil {
			println(err.Error())
			break
		}
		var visits struct {
			Visits []Visit `json:"visits"`
		}
		err = json.Unmarshal(b, &visits)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		for _, v := range visits.Visits {
			LvlCache.Save(v.Location, v.Id)
			UvlCache.Save(v.User, v.Id)
			vs.m[v.Id] = v
		}
		count++
	}
	fmt.Printf("All visits: %d\nVisits time: %+v\n", len(vs.m), time.Since(t))
}
