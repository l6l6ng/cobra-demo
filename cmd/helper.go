package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func ExecuteCommand(name, subname string, args ...string) (string, error) {
	args = append([]string{subname}, args...)
	cmd := exec.Command(name, args...)
	bytes, err := cmd.CombinedOutput()

	return string(bytes), err
}

func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "execute %s args:%v error:%v\n", cmd.Name(), args, err)
	os.Exit(1)
}

func showCalendar(y, m int) {
	fmt.Printf("打印日期。。。。%d年-%d月", y, m)
}
