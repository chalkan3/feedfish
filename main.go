package main

import (
	runner "t/pkg/runner"
)

func main() {
	processTree := runner.NewProcessTree().
		Insert(1, runner.NewProcess("igor")).
		Insert(2, runner.NewProcess("guica")).
		Insert(3, runner.NewProcess("mini")).
		Insert(4, runner.NewProcess("lady")).
		Insert(5, runner.NewProcess("garrafa")).
		Insert(6, runner.NewProcess("teste")).
		Insert(7, runner.NewProcess("kh")).
		Insert(8, runner.NewProcess("tt")).
		Insert(9, runner.NewProcess("mms")).
		Insert(10, runner.NewProcess("sorvete")).
		Insert(11, runner.NewProcess("bateria"))

	it := runner.NewRunner(processTree.Root())

	for it.HasNext() {
		it.Next().Execute()
	}

}
