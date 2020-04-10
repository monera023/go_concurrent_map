package cmap

import (
    "sort"
    "strconv"
    "testing"
    "log"
)

type FootballClub struct {
  name string
}

func TestMapCreation(t *testing.T) {
  m := NewConcurrentMap()

  if m == nil {
    t.Error("Map is null.")
  }

  if m.Count() != 0 {
    t.Error("New Map Should be Empty")
  }
}

func TestInsert(t *testing.T) {
  m := NewConcurrentMap()
  manUtd := FootballClub{"Man UTD"}
  liverpool := FootballClub{"Liverpool FC"}
  m.Add("man_utd", manUtd)
  m.Add("liverpool", liverpool)

  if m.Count() != 2 {
     t.Error("Map should contain exactly 2 elements")
  }
}

func TestGet(t *testing.T) {
  m := NewConcurrentMap()
  manUtd := FootballClub{"Man UTD"}
  liverpool := FootballClub{"Liverpool FC"}
  m.Add("man_utd", manUtd)
  m.Add("liverpool", liverpool)

  val, ok := m.Get("chelsea")
  if ok == true {
    t.Error("Missing element should return false")
  }

  if val != nil {
    t.Error("Missing element should be nil")
  }

  val1, ok1 := m.Get("man_utd")
  if ok1 == false {
    t.Error("For existing key ok should be true")
  }
  log.Print(val1)
  if val1 != manUtd {
    t.Error("Not the correct element")
  }
}

func TestRemove(t *testing.T) {
  m := NewConcurrentMap()
  manUtd := FootballClub{"Man UTD"}
  liverpool := FootballClub{"Liverpool FC"}
  m.Add("man_utd", manUtd)
  m.Add("liverpool", liverpool)

  m.Remove("man_utd")

  _, ok := m.Get("man_utd")
  if ok == true {
    t.Error("Deleted Value should not be present")
  }
}

func TestConcurrent(t *testing.T) {
  m := NewConcurrentMap()
  ch := make(chan int)
  var a[100]int

  for i := 0; i < 100 ; i ++ {
    go func(j int) {
      m.Add(strconv.Itoa(j), j)
      val, _ := m.Get(strconv.Itoa(j))
      ch <- val.(int)
    }(i)
  }
  counter := 0
  for elem := range ch {
    a[counter] = elem
    counter++
    if counter == 100 {
      break
    }
  }

  sort.Ints(a[0:100])

  if m.Count() != 100 {
    t.Error("Map should contains 100 elements")
  }

  for i := 0; i < 100; i++ {
    if a[i] != i {
      t.Error("Missing Value", i)
    }
  }
}


func BenchmarkInsert(b *testing.B) {
  m := NewConcurrentMap()
  for n := 0; n < b.N; n++ {
    m.Add(strconv.Itoa(n), n)
  }
}

func BenchmarkGet(b *testing.B) {
  b.ResetTimer()
  m := NewConcurrentMap()
  m.Add("val", 1)
  m.Add("val1", 2)
  for n := 0; n < b.N; n++ {
    m.Get("val")
  }
}

