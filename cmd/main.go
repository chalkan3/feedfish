package main

import (
	"io/ioutil"
	"log"
	steps "t/internal/steps"
	runner "t/pkg/runner"

	yaml "gopkg.in/yaml.v3"
)

func main() {
	var pipeline steps.Pipeline

	yamlFile, err := ioutil.ReadFile(".pipeline.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &pipeline)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	processTree := runner.NewProcessTree()
	for _, step := range pipeline.Steps {
		processTree.Register(step.Order, runner.NewProcess(step.Name, step.Scripts.Initial.Script, step.Scripts.Processing.Script, step.Scripts.Proceed.Script))
	}

	it := runner.NewRunner(processTree.Root())

	for it.HasNext() {
		it.Next().Execute()
	}

}
