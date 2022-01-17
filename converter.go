package converter

type Converter struct {
	opts *ConverterOpts
}

type ConverterOpts struct {
	FFmpegPath string
	Channels   int
	Bitrate    string
}

func New(opts *ConverterOpts) *Converter {
	if opts == nil {
		opts = &ConverterOpts{FFmpegPath: "ffmpeg", Channels: 1, Bitrate: "65K"}
	} else {
		if opts.FFmpegPath == "" {
			opts.FFmpegPath = "ffmpeg"
		}
		if opts.Channels == 0 {
			opts.Channels = 1
		}
		if opts.Bitrate == "" {
			opts.Bitrate = "65K"
		}
	}
	return &Converter{opts: opts}
}
