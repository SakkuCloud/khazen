package service

import (
	log "github.com/sirupsen/logrus"
	"os/exec"
)

func OSCommandExecute(cmd string) (err error) {
	log.Debugf("OS command execute, cmd: %s",cmd)

	output, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	log.Debugf("OS command execute, output: %s",output)
	return
}