package ko

import (
  "testing"
)

func TestRangerOneArg(t *testing.T) {
  expected := []int{0,1,2,3,4,5}
  for i := range Ranger(6) {
    if i != expected[0]{
      t.Error(i, " != ", expected[0])
    }
    expected = expected[1:]
  }
}

func TestRangerTwoArg(t *testing.T) {
  expected := []int{5,6,7,8,9}
  for i := range Ranger(5,10) {
    if i != expected[0]{
      t.Error(i, " != ", expected[0])
    }
    expected = expected[1:]
  }
}

func TestRangerThreeArg(t *testing.T) {
  expected := []int{5,8,11}
  for i := range Ranger(5,12, 3) {
    if i != expected[0]{
      t.Error(i, " != ", expected[0])
    }
    expected = expected[1:]
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

func TestProdcut(t *testing.T) {
  a := []int{1,2}
  b := []int{3,4}
  expected := [][]int{
    []int{1,3},
    []int{1,4},
    []int{2,3},
    []int{2,4}}
  for r := range Product(a,b).(chan []int) {
    for i, v := range r{
      if expected[0][i] != v {
        t.Error("Expected ", expected[0], " got ", r)
      }
    }
    expected = expected[1:]
  }
}

func TestSlicer(t *testing.T) {
  a := []int{1,2}
  b := []int{3,4}
  fn := func(x,y int) int{
    return a[x] * b[y]
  }
  actual := Slicer(fn, len(a), len(b)).([][]int)
  expected := [][]int{
    []int{3,4},
    []int{6,8}}
  for i, s := range actual {
    for j, v := range s {
      if expected[i][j] != v {
        t.Error("Expected ", expected, " got ", actual)
      }
    }
  } 
}

func TestIndexOf(t *testing.T){
  pi := []int{3,1,4,1,5,9,2,6,5,3}
  if IndexOf(1, pi) != 1 {
    t.Error("Error looking for 1")
  }
  if IndexOf(5, pi) != 4 {
    t.Error("Error looking for 5")
  }
  if IndexOf(7, pi) != -1 {
    t.Error("Error looking for -7")
  }
  if IndexOf(3, pi[1:]) != 8 {
    t.Error("Error looking for second 3")
  }
}

func TestPop(t *testing.T){
  pi := []int{3,1,4,1,5,9}
  castPop := func(s []int) (int, []int) {
    vi, vs := Pop(s)
    return vi.(int), vs.([]int)
  }
  i, s := castPop(pi)
  if i != 9 || len(s) != 5 {
    t.Error("Got ", i, s, " Expected 9 [3 1 4 1 5]")
  }
}

func TestShift(t *testing.T){
  pi := []int{3,1,4,1,5,9}
  castShift := func(s []int) (int, []int) {
    vi, vs := Shift(s)
    return vi.(int), vs.([]int)
  }
  i, s := castShift(pi)
  if i != 3 || len(s) != 5 {
    t.Error("Got ", i, s, " Expected 3 [1 4 1 5 9]")
  }
}