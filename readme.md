# Ko
In the game Go, the Ko rule says that a board can never be repeated. The objective of the Ko library to help you never repeat yourself.

This package uses reflection heavily and is therefore fairly slow. I would not recommend using this in production code, however it may prove useful for prototyping.

## Ranger
Works a lot like the python Range function. It has 3 calls
- Ranger(end): loops from 0 to end-1
- Ranger(start, end): loops from start to end-1
- Ranger(start, end, step): loops from start to end-1 incrementing by step

## Prepend
Takes a value and a slice and prepends the value to the slice. User needs to cast the return.

```go
x := Prepend(5, []int{1,2,3}).([]int)
```

## Product
Takes 1 to n lists and produces the cartesian product of them. User needs to call Interface() and then cast.

```go
l := []int{1,2,3,4}
for i := range Product(l,l,l).Interface().(chan []int) {
  fmt.Println(i)
}
```

## Slicer
Takes a function and 1 to n dimensions and returns a multi-dimensional slice matching the dimensions given. Each value in the slice will be the function return of the slice coordinates as arguments.


```go
fn := func(x,y int) int{
  return x*y
}
slice := Slicer(fn, 5, 6).([][]int)
fmt.Println(slice)
```
Will print [[0 0 0 0 0 0] [0 1 2 3 4 5] [0 2 4 6 8 10] [0 3 6 9 12 15] [0 4 8 12 16 20]]

## Looper(N, func)
Takes a int N and a function that has no return and no arguments and repeats the functions N times.

## IndexOf(val, slice)
Takes a value and a slice and returns the index of the index of the value in the slice. Returns -1 if the value is not found.