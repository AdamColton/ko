package ko

import (
  "reflect"
)

func Looper(n int, fn func()){
  for i:=0; i<n; i++ {
    fn()
  }
}

func Ranger(args ...int) (chan int){
  ch := make(chan int)
  switch (len(args)){
  case 0:
    go func(){
      for i:=0; true; i++{
        ch <- i
      }
    }()
  case 1:
    go func() {
      for i:=0; i< args[0]; i++{
        ch <- i
      }
      close(ch)
    }()
  case 2:
    go func() {
      a := args[0]
      b := args[1]
      if (a>b){
        for ; a>b; a--{
          ch <- a
        }
      } else {
        for ; a<b; a++{
          ch <- a
        }
      }
      close(ch)
    }()
  case 3:
    go func () {
      for i:=args[0]; i<args[1]; i+=args[2] {
        ch <- i
      }
      close(ch)
    }()
  }
  return ch
}

func Prepend(p interface{}, s interface{}) interface{} {
  vps := reflect.Append(reflect.MakeSlice( reflect.SliceOf(reflect.TypeOf(p)), 0, 1), reflect.ValueOf(p))
  vs := reflect.ValueOf(s)
  return reflect.AppendSlice(vps, vs).Interface()
}

func Product(h interface{}, t ...interface{}) interface{}{
  var product func (interface{}, ...interface{}) (reflect.Value)
  product = func(h interface{}, t ...interface{}) (reflect.Value){
    hv := reflect.ValueOf(h)
    ht := reflect.TypeOf(h)
    ch := reflect.MakeChan( reflect.ChanOf(reflect.BothDir, ht), 0)
    go func(){
      if len(t) > 1 {
        for i:=0; i<hv.Len(); i++ {
          ich := product(t[0],t[1:]...)
          for j, ok := ich.Recv(); ok; j, ok = ich.Recv(){
            ch.Send( reflect.ValueOf( Prepend(hv.Index(i).Interface(), j.Interface()) ) )
          }
        }
      } else if len (t) == 1 {
        lv := reflect.ValueOf(t[0])
        for i:=0; i<hv.Len(); i++ {
          for j:=0; j<lv.Len(); j++ {
            l := reflect.MakeSlice( ht, 0, 2)
            l = reflect.Append(l, hv.Index(i) )
            l = reflect.Append(l, lv.Index(j) )
            ch.Send( l )
          }
        }
      } else {
        for i:=0; i<hv.Len(); i++ {
          ch.Send( reflect.Append(reflect.MakeSlice( ht, 0, 1), hv.Index(i)) )
        }
      }
      ch.Close()
    }()
    return ch
  }
  return product(h, t...).Interface()
}

func Slicer (f interface{}, dims ...int) interface{} {
  ft := reflect.TypeOf(f)
  fv := reflect.ValueOf(f)
  outType := ft.Out(0)

  var slicer func(interface{}, []reflect.Value, ...int) (reflect.Value)
  slicer = func(f interface{}, args []reflect.Value, dims ...int) (reflect.Value) {
    if len(dims) == 1 {
      index := len(args)
      args = append(args, reflect.ValueOf(0))
      l := reflect.MakeSlice( reflect.SliceOf(outType), 0, dims[0])
      for i:=0; i<dims[0]; i++ {
        args[index] = reflect.ValueOf(i)
        ret := fv.Call(args)
        if len(ret) == 1 || ret[1].Interface().(bool) {
          l = reflect.Append(l,ret[0])
        }
      }
      args = args[:index]
      return l
    }
    index := len(args)
    args = append(args, reflect.ValueOf(0))
    t := reflect.SliceOf(outType)
    for i:=0; i<index+1; i++{
      t = reflect.SliceOf(t)
    } 
    l := reflect.MakeSlice( t, 0, dims[0] )
    for i:=0; i<dims[0]; i++{
      args[index] = reflect.ValueOf(i)
      val := slicer(f, args, dims[1:]...)
      l = reflect.Append(l, val)
    }
    args = args[:index]
    return l
  }

  return slicer(f, make([]reflect.Value,0, len(dims)), dims...).Interface()
}

func IndexOf(val interface{}, slice interface{}) int{
  sliceVal := reflect.ValueOf(slice)
  valVal := reflect.ValueOf(val)
  for i:=0; i<sliceVal.Len(); i++{
    if reflect.DeepEqual(valVal.Interface(), sliceVal.Index(i).Interface()){
      return i
    }
  }
  return -1
}

func Pop(slice interface{}) (interface{}, interface{}) {
  sliceVal := reflect.ValueOf(slice)
  index := sliceVal.Len() - 1
  return sliceVal.Index(index).Interface(), sliceVal.Slice(0, index).Interface()
}

func Shift(slice interface{}) (interface{}, interface{}) {
  sliceVal := reflect.ValueOf(slice)
  return sliceVal.Index(0).Interface(), sliceVal.Slice(1, sliceVal.Len()).Interface()
}