package cmd2

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var (
	dividedByZeroHanding int // 除0 如何处理
)

var divideCmd = &cobra.Command{
	Use:   "divide",
	Short: "Divide subcommand divide all passed args.",
	Run: func(cmd *cobra.Command, args []string) {
		values := ConvertArgsToFloat64Slice(args, ErrorHandling(parseHanding))
		result := calc(values, DIVIDE)
		fmt.Printf("%s = %.2f\n", strings.Join(args, "/"), result)
	},
}

func init() {
	divideCmd.Flags().IntVarP(&dividedByZeroHanding, "divide_by_zero", "d", int(PanicOnDividedByZero), "do what when divied by zero")
	rootCmd.AddCommand(divideCmd)
}
