package util

type FuncHelper struct {
	statusExecution error
	target          func()
}

func (fn *FuncHelper) New(target func()) {
	fn.target = target
}

func (fn *FuncHelper) Run() {
	defer func() {
		if r := recover(); r != nil {
			fn.statusExecution = r.(error)
			return
		}
	}()

	fn.target()
	fn.statusExecution = error(nil)
}

func (fn FuncHelper) GetStatusExecution() error {
	return fn.statusExecution
}
