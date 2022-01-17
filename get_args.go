package converter

import "strconv"

func (c *Converter) getArgs(input string, output string) []string {
	return []string{
		"-y",
		"-i",
		input,
		"-c",
		"copy",
		"-acodec",
		"pcm_s16le",
		"-f",
		"s16le",
		"-ac",
		string(strconv.AppendInt([]byte(""), int64(c.opts.Channels), 10)),
		"-ar",
		c.opts.Bitrate,
		output,
	}
}
