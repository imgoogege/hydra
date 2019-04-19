package hydra

import "github.com/urfave/cli"

type ICli interface {
	Append(string, cli.Flag)
	Validate(string, func(*cli.Context) error)
	getFlags(mode string) []cli.Flag
	getValidators(mode string) []func(*cli.Context) error
	setContext(ctx *cli.Context)
	Context() *cli.Context
}
type Cli struct {
	flags      map[string][]cli.Flag
	validators map[string][]func(*cli.Context) error
	cli        *cli.Context
}

func NewCli() *Cli {
	return &Cli{
		flags:      make(map[string][]cli.Flag),
		validators: make(map[string][]func(*cli.Context) error),
	}
}

func (c *Cli) Append(mode string, flag cli.Flag) {
	if _, ok := c.flags[mode]; !ok {
		c.flags[mode] = make([]cli.Flag, 0, 1)
	}
	c.flags[mode] = append(c.flags[mode], flag)
}
func (c *Cli) Validate(mode string, f func(*cli.Context) error) {
	if _, ok := c.validators[mode]; !ok {
		c.validators[mode] = make([]func(*cli.Context) error, 0, 1)
	}
	c.validators[mode] = append(c.validators[mode], f)
}
func (c *Cli) getFlags(mode string) []cli.Flag {
	if v, ok := c.flags[mode]; ok {
		return v
	}
	return make([]cli.Flag, 0, 0)
}
func (c *Cli) getValidators(mode string) []func(*cli.Context) error {
	if v, ok := c.validators[mode]; ok {
		return v
	}
	return make([]func(*cli.Context) error, 0, 0)
}
func (c *Cli) setContext(ctx *cli.Context) {
	c.cli = ctx
}
func (c *Cli) Context() *cli.Context {
	return c.cli
}
