package agent

import (
	"io/ioutil"
	"os/exec"
)

type Process struct {
}

func (proc Process) Run(cmdName string, arg []string, callback func(result []byte)) {
	cmd := exec.Command(cmdName, arg...)

	in, _ := cmd.StdinPipe()
	out, _ := cmd.StdoutPipe()

	cmd.Start()
	in.Close()

	result, _ := ioutil.ReadAll(out)
	cmd.Wait()

	callback(result)
}
