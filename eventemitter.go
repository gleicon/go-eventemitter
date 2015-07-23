package eventemitter

type EventEmitter struct {
	pid    int
	worker map[string][]func([]byte)
}

func NewEventEmitter() *EventEmitter {
	ee := EventEmitter{}
	ee.worker = make(map[string][]func([]byte))
	return &ee
}

func (ee EventEmitter) On(eventname string, f func(payload []byte)) {
	ee.worker[eventname] = append(ee.worker[eventname], f)
}

func (ee EventEmitter) RemoveAllListeners(eventname string) {
	ee.worker[eventname] = nil
}

func (ee EventEmitter) Listeners(eventname string) []func([]byte) {
	return ee.worker[eventname]
}

func (ee EventEmitter) Emit(eventname string, message []byte) {
	for f := range ee.worker[eventname] {
		fun := ee.worker[eventname][f]
		if fun != nil {
			go fun(message)
		}
	}
}
