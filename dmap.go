package dmap

import "fmt"

const (
	closeCmd = "CLOSE"
	setCmd   = "SET"
	getCmd   = "GET"
	delCmd   = "DEL"
	keyCmd   = "KEY"
)

type command struct {
	t      string
	key    string
	value  interface{}
	result chan interface{}
}

type dmap struct {
	poom     chan command
	internal map[string]interface{}
}

func (m *dmap) run() {
	for cmd := range m.poom {
		switch cmd.t {
		case closeCmd:
			close(m.poom)
			m.internal = nil
			return
		case setCmd:
			m.internal[cmd.key] = cmd.value
			cmd.result <- cmd.value

		case getCmd:
			cmd.result <- m.internal[cmd.key]

		case delCmd:
			delete(m.internal, cmd.key)
			cmd.result <- nil
		case keyCmd:
			keys := make([]string, 0, len(m.internal))
			for key := range m.internal {
				keys = append(keys, key)
			}
			cmd.result <- keys
		}
	}
}

func (m *dmap) pushCmd(cmd command) {
	m.poom <- cmd
	if cmd.t == closeCmd {
		fmt.Print("closed map")
	}
}

func (m *dmap) Get(key string) interface{} {
	result := make(chan interface{})
	get := command{
		t:      getCmd,
		key:    key,
		result: result,
	}
	go m.pushCmd(get)
	return <-result
}

func (m *dmap) Del(key string) {
	result := make(chan interface{})
	del := command{
		t:      delCmd,
		key:    key,
		result: result,
	}
	go m.pushCmd(del)
	<-result
}

func (m *dmap) Set(key string, value interface{}) interface{} {
	result := make(chan interface{})
	set := command{
		t:      setCmd,
		key:    key,
		value:  value,
		result: result,
	}
	go m.pushCmd(set)
	return <-result
}

func (m *dmap) Keys() []string {
	result := make(chan interface{})
	key := command{
		t:      keyCmd,
		result: result,
	}
	go m.pushCmd(key)
	items := <-result
	kq, _ := items.([]string)
	return kq
}

func (m *dmap) Close() {
	m.pushCmd(command{
		t: closeCmd,
	})
}

func NewMap() *dmap {
	m := &dmap{
		poom:     make(chan command),
		internal: make(map[string]interface{}),
	}
	go m.run()
	return m
}
