package main

import (
	"context"
	"log"

	"github.com/apenella/go-ansible/v2/pkg/execute/workflow"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {
	playbookInstance := playbook.NewAnsiblePlaybookExecute("main.yml")

	err := workflow.NewWorkflowExecute(playbookInstance).Execute(context.TODO())
	if err != nil {
		log.Fatalln(err.Error())
	}
}
