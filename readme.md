# Ko
In the game Go, the Ko rule says that a board can never be repeated. The objective of the Ko library to help you never repeat yourself.

This package uses reflection heavily and is therefore fairly slow. I would not recommend using this in production code, however it may prove useful for prototyping.

## func Ranger
```go
  Ranger(args ...int) (chan int)
```
Works a lot like the python Range function. It be thought of having the following overloading
```go
  Ranger(end int) (chan int)
  Ranger(start, end int) (chan int)
  Ranger(start, end, step int) (chan int)
```

### Examples

```go
  for i := range Ranger(3) {
    println(i)
  }
  /* Output:
  0
  1
  2
  */

  for i := range Ranger(5, 8){
    println(i)
  }
  /* Output
  5
  6
  7
  */

  for i := range Ranger(8, 5){
    println(i)
  }
  /* Output
  8
  7
  6
  */

  for i := range Ranger(6, 15, 3){
    println(i)
  }
  /* Output
  6
  9
  12
  */
```

## func Prepend
```go
  func Prepend(p interface{}, s interface{}) interface{}
```
Takes a value and a slice and prepends the value to the slice. User needs to cast the return.

### Example
```go
  x := Prepend(5, []int{1,2,3}).([]int) // [5 1 2 3]
```

## func Product
```go
  func Product(h interface{}, t ...interface{}) interface{}
```
Takes any number of lists and produces the cartesian product of them. The lists must be of the same type. The return value should be cast the same type as well

### Example
```go
  l := []int{1,2,3,4}
  for i := range Product(l,l,l).(chan []int) {
    fmt.Println(i)
  }
```

## Slicer
```go
  Slicer (f interface{}, dims ...int) interface{}
```

Takes a function and any number of dimensions as ints. Returns a multi-dimensional slice matching the dimensions given. Each value in the slice will be the value of the funciton passed in when called with coordinate value.


### Example
```go
  fn := func(x,y int) int{
    return x*y
  }
  slice := Slicer(fn, 5, 6).([][]int)
  fmt.Println(slice)

  /* Output
  [[0 0 0 0 0 0] [0 1 2 3 4 5] [0 2 4 6 8 10] [0 3 6 9 12 15] [0 4 8 12 16 20]]
  */
```

This can also be very powerful with closures
```go
  funcRand := func(i ...int) int{
    return rand.Intn(10) + 1
  }
  a := Slicer(funcRand,5,5).([][]int)
  b := Slicer(funcRand,5,5).([][]int)
  funcMult := func(x,y int) int{
    return a[x][y] * b[x][y]
  }
  mult := Slicer(funcMult,5,5).([][]int)
```

## func Looper
```go
  func Looper(n int, fn func())
```
Runs a function a given number of times. This is just syntactic sugar for really basic loops.

### Example
```go
  Looper(3, func{
    println("Hello")
  })

  /* Output
  Hello
  Hello
  Hello
  */
```
## func IndexOf
```go
  IndexOf(val interface{}, slice interface{}) int
```

Takes a value and a slice and returns the index of the index of the value in the slice. Returns -1 if the value is not found.

### Example
```go
  pi := []int{3,1,4,1,5,9,2,6,5,3}
  println(IndexOf(1, pi)) // 1
  println(IndexOf(5, pi)) // 4
  println(IndexOf(7, pi)) // -1
  println(IndexOf(3, pi[1:])) // 8

  /* Output
  1
  4
  -1
  8
  */
```