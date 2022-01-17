package converter

import "os/exec"

func (c *Converter) Convert(input string, output string) error {
	return exec.Command(c.opts.FFmpegPath, c.getArgs(input, output)...).Run()
}
