package cmap

import (
    "sync"
    "log"
    "hash/fnv"
)

type ConcurrentMap struct {
  m map[string]interface{}  // Key: String, Value: value of any type
  sync.RWMutex // Note: If a structure has a sync field, it must be passed by pointer.
}

func NewConcurrentMap() *ConcurrentMap {
  return &ConcurrentMap{m: make(map[string]interface{})}
}

func (m *ConcurrentMap) Add(key string, value interface{}) {
  h := fnv.New32a()
  h.Write([]byte(key))
  index := h.Sum32() % 10
  log.Print(index)
  m.Lock()
  m.m[key] = value
  m.Unlock() // To check using defer here
}


func (m *ConcurrentMap) Count() int {
  m.Lock()
  defer m.Unlock()
  return len(m.m)
}

func (m *ConcurrentMap) Get(key string) (interface{}, bool){
  m.Lock()
  val, err := m.m[key]
  m.Unlock()
  return val, err
}

func(m *ConcurrentMap) HasKey(key string) bool {
  m.Lock()
  _, ok := m.m[key]
  m.Unlock()
  return ok
}

func (m *ConcurrentMap) Remove(key string) {
  m.Lock()
  delete(m.m, key)
  m.Unlock()
}

func (m *ConcurrentMap) Clear() {
  m.Lock()
  m.m = make(map[string]interface{})
}

func (m *ConcurrentMap) IsEmpty() bool {
  m.Lock()
  defer m.Unlock()
  return len(m.m) == 0
}

type Tuple struct {
    key string
    value interface{}
}

func (m *ConcurrentMap) Iter() <-chan Tuple {
  outputCh := make(chan Tuple)
  go func() {
      m.Lock()
      defer m.Unlock()
      for key, val := range m.m {
        outputCh <- Tuple{key, val} 
      }
      close(outputCh)
    }()
  return outputCh  
}





