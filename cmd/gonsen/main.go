package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"github.com/takebayashi/gonsen"
	"log"
	"os"
)

type ListCmd struct{}

func (c *ListCmd) Help() string {
	return "show all program names"
}

func (c *ListCmd) Run(args []string) int {
	ns, err := gonsen.GetProgramNames()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	for _, n := range ns {
		fmt.Println(n)
	}
	return 0
}

func (f *ListCmd) Synopsis() string {
	return ""
}

type ShowCmd struct{}

func (c *ShowCmd) Help() string {
	return "show program info"
}

func (c *ShowCmd) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println("1 argument required")
		return 1
	}
	p, err := gonsen.GetProgram(args[0])
	if err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Printf("%s\t%s\t%s\t%s\t%s\n", p.Title, p.Slug, p.Personality, p.Guest, p.MediaUrl)
	return 0
}

func (f *ShowCmd) Synopsis() string {
	return "<program-name>"
}

func main() {
	c := cli.NewCLI("gonsen", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return &ListCmd{}, nil
		},
		"show": func() (cli.Command, error) {
			return &ShowCmd{}, nil
		},
	}
	st, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(st)
}
