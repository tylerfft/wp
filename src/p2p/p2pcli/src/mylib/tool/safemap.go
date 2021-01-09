package tool

type SafeMapIntf interface {
	Insert(string, interface{})
	Delete(string)
	Find(string) (interface{}, bool)
	Close() map[string]interface{}
}

type ActType int

const (
	REMOVE ActType = iota
	END
	FIND
	INSERT
	LENGTH
)

type SafeMapChanDataStu struct {
	action ActType
	key    string
	value  interface{}
	result chan<- interface{}
	data   chan<- map[string]interface{}
}

func NewSafeMap() SafeMapIntf {
	sm := make(safeMapChan)
	go sm.Run()
	return sm
}

type safeMapChan chan SafeMapChanDataStu

func (sm safeMapChan) Insert(key string, value interface{}) {
	sm <- SafeMapChanDataStu{action: INSERT, key: key, value: value}
}

func (sm safeMapChan) Delete(key string) {
	sm <- SafeMapChanDataStu{action: REMOVE, key: key}
}

func (sm safeMapChan) Find(key string) (value interface{}, found bool) {
	reply := make(chan interface{})
	sm <- SafeMapChanDataStu{action: FIND, key: key, result: reply}
	result := (<-reply).(FindResultStu)
	return result.value, result.found
}

func (sm safeMapChan) Close() map[string]interface{} {
	reply := make(chan map[string]interface{})
	sm <- SafeMapChanDataStu{action: END, data: reply}
	return <-reply
}

type FindResultStu struct {
	value interface{}
	found bool
}

func (sm safeMapChan) Run() {
	store := make(map[string]interface{})
	for command := range sm {
		switch command.action {
		case INSERT:
			store[command.key] = command.value
		case REMOVE:
			delete(store, command.key)
		case FIND:
			value, found := store[command.key]
			command.result <- FindResultStu{value, found}
		case LENGTH:
			command.result <- len(store)
		case END:
			close(sm)
			command.data <- store
		}
	}
}
