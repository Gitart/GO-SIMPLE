package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	//"github.com/spf13/cobra/doc"
)

var (
	EchoTimes  int
	EchoOther  string
	EchoOthers string
	ArgsTitle  []string
	DevType    string

	// AppName Global variables
	AppName    string = "informer"
	AppVersion string = "v.001.1"
)

func main() {

	// For documentation
	//header := &doc.GenManHeader{
	//	Title:   "MINE",
	//	Section: "3",
	//}

	// Dev mode
	var devType = &cobra.Command{
		Version: "v1.7",
		Example: "ct dev D",
		Use:     "dev",
		Short:   "Dev mode",
		Long:    `Development or production type`,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			DevType = args[0]
			//fmt.Println("MODE Develop : " + strings.Join(args, " "))
		},
	}

	// Other
	var cmdOther = &cobra.Command{
		Version: AppVersion,
		Example: "Пример использование",
		Use:     "other",
		Short:   "Прочее короткая",
		Long:    `Прочее пример`,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			EchoOther = args[0]
		},
	}

	// 📃 Print
	// print "Первый арг" "Второй арг"
	var cmdPrint = &cobra.Command{
		Use:   "print",
		Short: "Печать короткая",
		Long:  `Печать пример`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ArgsTitle = args
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	// ct echo testecho
	var cmdEcho = &cobra.Command{
		Use:   "echo",
		Short: "Эхо короткая",
		Long:  `Эхо - это echo is for echoing anything back. Echo works a lot like print, except it has a child command.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Эхо : " + strings.Join(args, " "))
		},
	}

	// Пример использования
	// ct echo tim 1
	var cmdTimes = &cobra.Command{
		Use:   "tim",
		Short: "Время для использования",
		Long:  `Время echo things multiple times back to the user by providing a count and a string.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < EchoTimes; i++ {
				fmt.Println("Эхо: " + strings.Join(args, " "))
			}
		},
	}

	// Проверка и получение аргумента в числовом виде
	// По умолчанию 10
	// Used:  echo tim -t 33344
	cmdTimes.Flags().IntVarP(&EchoTimes, "tim", "t", 10, "times to echo the input")

	// Used:  echo -e testechoexample
	cmdEcho.Flags().StringVarP(&EchoOther, "echo", "e", "EchoTitle", "echo for input")

	// Для ct - название приложения
	// которое должно иметь такое имя при генерации
	var rootCmd = &cobra.Command{Use: AppName}
	rootCmd.AddCommand(cmdPrint, cmdEcho, cmdOther, devType)

	// Добавление второго уровня для сmdecho
	cmdEcho.AddCommand(cmdTimes)

	// Запуск
	rootCmd.Execute()

	// Test
	fmt.Println(EchoTimes)
	fmt.Println(EchoOther)
	fmt.Println(ArgsTitle)

	if DevType == "D" {
		fmt.Println("DEV MODE")
	}

	if DevType == "P" {
		fmt.Println("PRODUCTION MODE: READY ")
	}

	if DevType == "S" {
		fmt.Println("TRACE MODE 💡")
	}

	// Documentation

	//out := new(bytes.Buffer)
	//cobra.CheckErr(doc.GenMan(cmdTimes, header, out))
	//fmt.Print(out.String())
}
