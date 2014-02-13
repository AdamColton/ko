package ko

import (
  "testing"
)

func TestRangerOneArg(t *testing.T) {
  ref := []int{0,1,2,3,4,5}
  for i := range Ranger(6) {
    if i != ref[0]{
      t.Error(i, " != ", ref[0])
    }
    ref = ref[1:]
  }
}

func TestRangerTwoArg(t *testing.T) {
  ref := []int{5,6,7,8,9}
  for i := range Ranger(5,10) {
    if i != ref[0]{
      t.Error(i, " != ", ref[0])
    }
    ref = ref[1:]
  }
}

func TestRangerThreeArg(t *testing.T) {
  ref := []int{5,8,11}
  for i := range Ranger(5,12, 3) {
    if i != ref[0]{
      t.Error(i, " != ", ref[0])
    }
    ref = ref[1:]
  }
}

func TestLooper(t *testing.T){
  counter := 0
  Looper(10, func(){
    counter++  
  })
  if counter != 10 {
    t.Error("expected 10")
  }
}

func TestPrepend(t *testing.T){
  s := []int{1,2,3}
  i := 4
  s = Prepend(i,s).([]int)
  if s[0] != 4 {
    t.Error("4 was not prepended to slice")
  }
  if len(s) != 4 {
    t.Error("len(s) should be 4, got ", len(s))
  }
}