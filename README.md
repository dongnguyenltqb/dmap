### Usage

```go

func run(){
  // create new map[string]string
  m := dmap[string,string].New()

  // set value for a key
  m.Set("username","dongnguyenltqb")
  m.Set("age","25")

  // get value for a key
  fmt.Println("username => ",m.Get("username"))

  // delete value for a key
  m.Del("username")

  // list keys for map
  for _,keys := range m.Keys{
      fmt.Println("keys => ",keys)
  }
}
```

### Output

```go
  username => dongnguyenltqb
  keys => [username age]
```
