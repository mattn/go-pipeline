package pipeline

import (
	"bytes"
	"os/exec"
)

func CommandPipeLine(commands ...[]string) ([]byte, error) {
	cmds := make([]*exec.Cmd, len(commands))
	for i, c := range commands {
		cmds[i] = exec.Command(c[0], c[1:]...)
		if i > 0 {
			rc, err := cmds[i-1].StdoutPipe()
			if err != nil {
				return nil, err
			}
			cmds[i].Stdin = rc
		}
	}
	var out bytes.Buffer
	cmds[len(cmds)-1].Stdout = &out
	for _, c := range cmds {
		err := c.Start()
		if err != nil {
			return nil, err
		}
	}
	for _, c := range cmds {
		err := c.Wait()
		if err != nil {
			return nil, err
		}
	}
	return out.Bytes(), nil
}
