package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/urfave/cli"
)

const (
	// RegPrimary : 正規表現
	RegPrimary = `(\([AB]\)\s)?`
	// RegDatetime : 正規表現
	RegDatetime = `[0-9]{4}-[0-9]{2}-[0-9]{2}`
	// RegFormat : フォーマット
	RegFormat = RegPrimary + `(` + RegDatetime + `\s)?` + `.*`
)

func main() {
	app := cli.NewApp()

	app.Name = "todo"
	app.Usage = "This app manages todo.txt in your current directory"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the todo.txt",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "primary, p",
					Value: "None",
				},
				cli.StringFlag{
					Name:  "date, d",
					Value: getDatetime(),
				},
			},
			Action: addAction,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list task",
			Action:  listAction,
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "remove a task to the todo.txt",
			Action:  removeAction,
		},
		{
			Name:    "done",
			Aliases: []string{"d"},
			Usage:   "done a task to the todo.txt",
			Action:  doneAction,
		},
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "create new todo.txt",
			Action:  newAction,
		},
	}

	app.Run(os.Args)
}

func addAction(c *cli.Context) error {
	file, err := os.OpenFile("./todo.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("todo.txt does not exist in this directory.")
		fmt.Println("  command : todo new")
		return err
	}
	defer file.Close()

	if c.Args().First() == "" {
		errmsg := errors.New("you must input a task")
		fmt.Println(errmsg)
		return errmsg
	}

	task := c.Args().First()

	// check format
	r := regexp.MustCompile(RegFormat)
	if r.MatchString(task) == false {
		errmsg := errors.New("task format is incorrect")
		fmt.Println(errmsg)
		return errmsg
	}

	r = regexp.MustCompile(RegDatetime)
	if r.MatchString(task) == false {
		task = strings.Join([]string{c.String("d"), task}, " ")
	}
	r = regexp.MustCompile(RegPrimary)
	if c.String("p") != "None" && r.MatchString(task) == false {
		primary := "(" + c.String("p") + ")"
		task = strings.Join([]string{primary, task}, " ")
	}

	writer := bufio.NewWriter(file)
	writer.WriteString(task)
	writer.Flush()

	fmt.Printf("Add a new task: %s\n", task)

	return nil
}

func listAction(c *cli.Context) error {
	fmt.Println("list action")
	return nil
}

func removeAction(c *cli.Context) error {
	fmt.Println("remove action")
	return nil
}

func doneAction(c *cli.Context) error {
	fmt.Println("done action")
	return nil
}

func newAction(c *cli.Context) error {
	fmt.Println("new action")
	return nil
}

func getDatetime() string {
	day := time.Now()
	const layout = "2006-01-02"
	return day.Format(layout)
}
