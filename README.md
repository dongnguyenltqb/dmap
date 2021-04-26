### Usage

```go
import "github.com/dongnguyenltqb/dmap"

func run(){
  // create new map[string]interface{}
  m := dmap.New()
  
  // set value for a key
  m.Set("username","dongnguyenltqb")
  m.Set("email","dong.nguyen@gmail.com")
  m.Set("age",24)

  // get value for a key
  fmt.Println("email => ",m.Get("email"))

  // delete value for a key
  m.Del("username")

  // list keys for map
  for _,key := range m.Keys{
      fmt.Println("KEY=> ",key)
  } 
}

// output 
email =>  dong.nguyen@gmail.com
KEY=> email
KEY=> age
```

