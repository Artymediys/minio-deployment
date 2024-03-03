package main

import (
	"context"
	"log"

	"github.com/apenella/go-ansible/v2/pkg/execute/workflow"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {
	ev := make(map[string]interface{}, 1)
	ev["ask-become-pass"] = nil
	options := &playbook.AnsiblePlaybookOptions{
		ExtraVars: ev,
	}

	playbookInstance := playbook.NewAnsiblePlaybookExecute("main.yml").WithPlaybookOptions(options)

	err := workflow.NewWorkflowExecute(playbookInstance).Execute(context.TODO())
	if err != nil {
		log.Fatalln(err.Error())
	}
}
