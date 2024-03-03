package main

import (
	"context"
	"log"

	"github.com/apenella/go-ansible/v2/pkg/execute/workflow"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {
	options := &playbook.AnsiblePlaybookOptions{
		AskBecomePass: true,
	}

	playbookInstance := playbook.NewAnsiblePlaybookExecute("main.yml").WithPlaybookOptions(options)

	err := workflow.NewWorkflowExecute(playbookInstance).Execute(context.TODO())
	if err != nil {
		log.Fatalln(err.Error())
	}
}
