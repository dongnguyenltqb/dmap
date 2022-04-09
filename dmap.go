package dmap

const (
	closeCmd = "CLOSE"
	setCmd   = "SET"
	getCmd   = "GET"
	delCmd   = "DEL"
	keyCmd   = "KEY"
)

type command[K comparable, V interface{}] struct {
	kind   string
	key    K
	value  V
	chGet  chan V
	chKeys chan []K
	chSet  chan V
	chDel  chan K
}

type dmap[K comparable, V interface{}] struct {
	poom     chan command[K, V]
	internal map[K]V
}

func (m *dmap[K, V]) run() {
	for cmd := range m.poom {
		switch cmd.kind {
		case closeCmd:
			close(m.poom)
			m.internal = nil
			return
		case getCmd:
			cmd.chGet <- m.internal[cmd.key]
		case setCmd:
			m.internal[cmd.key] = cmd.value
			cmd.chSet <- cmd.value
		case delCmd:
			delete(m.internal, cmd.key)
			cmd.chDel <- cmd.key
		case keyCmd:
			keys := make([]K, 0, len(m.internal))
			for key := range m.internal {
				keys = append(keys, key)
			}
			cmd.chKeys <- keys
		}
	}
}

func (m *dmap[K, V]) pushCmd(cmd command[K, V]) {
	m.poom <- cmd
}

func (m *dmap[K, V]) Get(key K) interface{} {
	result := make(chan V)
	get := command[K, V]{
		kind:  getCmd,
		key:   key,
		chGet: result,
	}
	go m.pushCmd(get)
	return <-result
}

func (m *dmap[K, V]) Del(key K) {
	result := make(chan K)
	del := command[K, V]{
		kind:  delCmd,
		key:   key,
		chDel: result,
	}
	go m.pushCmd(del)
	<-result
}

func (m *dmap[K, V]) Set(key K, value V) V {
	result := make(chan V)
	set := command[K, V]{
		kind:  setCmd,
		key:   key,
		value: value,
		chSet: result,
	}
	go m.pushCmd(set)
	return <-result
}

func (m *dmap[K, V]) Keys() []K {
	result := make(chan []K)
	key := command[K, V]{
		kind:   keyCmd,
		chKeys: result,
	}
	go m.pushCmd(key)
	return <-result
}

func (m *dmap[K, V]) Close() {
	m.pushCmd(command[K, V]{
		kind: closeCmd,
	})
}

func NewMap[K comparable, V interface{}]() *dmap[K, V] {
	m := &dmap[K, V]{
		poom:     make(chan command[K, V]),
		internal: make(map[K]V),
	}
	go m.run()
	return m
}
