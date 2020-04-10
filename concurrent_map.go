package cmap

import (
    "sync"
    "hash/fnv"
)

type ConcurrentMap []*Shard

type Shard struct {
  items map[string]interface{}
  sync.RWMutex
}

func NewConcurrentMap() ConcurrentMap {
  concurrentMap := make([]*Shard, 32)
  for i := 0; i < 32; i++ {
    concurrentMap[i] = &Shard {
      items: make(map[string]interface{}),
    }
  }
  return concurrentMap
}


func (m ConcurrentMap) Add(key string, value interface{}) {
  shard := m.GetShard(key)
  shard.Lock()
  shard.items[key] = value
  shard.Unlock() // To check using defer here
}

func (m ConcurrentMap) GetShard(key string) *Shard{
  h := fnv.New32a()
  h.Write([]byte(key))
  index := h.Sum32() % 32
  return m[index]
}

func (m ConcurrentMap) Count() int {
  total := 0
  for i := 0; i < 32; i++ {
    m[i].Lock()
    total += len(m[i].items)
    m[i].Unlock()
  }
  return total
}

func (m ConcurrentMap) Get(key string) (interface{}, bool){
  shard := m.GetShard(key)
  shard.Lock()
  val, err := shard.items[key]
  shard.Unlock()
  return val, err
}

// func(m *ConcurrentMap) HasKey(key string) bool {
//   m.Lock()
//   _, ok := m.m[key]
//   m.Unlock()
//   return ok
// }

func (m ConcurrentMap) Remove(key string) {
  shard := m.GetShard(key)
  shard.Lock()
  delete(shard.items, key)
  shard.Unlock()
}

// func (m *ConcurrentMap) Clear() {
//   m.Lock()
//   m.m = make(map[string]interface{})
// }




