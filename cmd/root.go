package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "calendar [year [month [day]]]",
	Short: "My calendar cli app generating a sheet of calendar.",
	Run:   generate,
}

func generate(cmd *cobra.Command, args []string) {
	//根据输入确定日期
	argNum := len(args)
	time.Now().Date()
	date := struct {
		y int
		m time.Month
		d int
	}{}
	date.y, date.m, date.d = time.Now().Date() //如果没有参数，则默认今天
	var err error
	if argNum > 0 {
		date.y, err = strconv.Atoi(args[0])
		if err != nil {
			log.Fatalln(err)
		}
	}
	if argNum > 1 {
		tmpDate, err := time.Parse("2006-01-02", "2006-"+fmt.Sprintf("%02s", args[1])+"-12")
		if err != nil {
			log.Fatalln(err)
		}
		date.m = tmpDate.Month()
	}
	if argNum > 2 {
		date.d, err = strconv.Atoi(args[2])
		if err != nil {
			log.Fatalln(err)
		}
	}

	// 确定当月的第一天
	firstDay, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-%02d", date.y, date.m, 1))

	//绘制日历表头
	fmt.Println(date.m.String(), date.y)
	fmt.Println("Su Mo Tu We Th Fr Sa")

	// 不同的第一天需要不同的缩进
	switch firstDay.Weekday().String() {
	case "Sunday":
	case "Monday":
		fmt.Printf("%3s", "")
	case "Tuesday":
		fmt.Printf("%6s", "")
	case "Wednesday":
		fmt.Printf("%9s", "")
	case "Thursday":
		fmt.Printf("%12s", "")
	case "Friday":
		fmt.Printf("%15s", "")
	case "Saturday":
		fmt.Printf("%18s", "")
	}
	//绘制日历主体
	for day := firstDay; day.Month() == date.m; day = day.AddDate(0, 0, 1) {
		if day.Day() == date.d && day.Month() == date.m && day.Year() == date.y { //如果是今天则给背景上色
			fmt.Printf("%2d", aurora.BgGreen(day.Day()))
		} else {
			fmt.Printf("%2d", day.Day())
		}
		if day.Weekday().String() == "Saturday" {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func Execute() {
	rootCmd.Execute()
}
