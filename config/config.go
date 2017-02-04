package config

const (
	OutputFormatText = "text"
	OutputFormatJSON = "json"

	OutputTargetStdout = "stdout"
	OutputTargetFile   = "file"
)

type Config struct {
	Github   Github
	Output   Output
	Strategy Strategy
}

type Github struct {
}

type Output struct {
	Format   string `yml:"format"`
	Target   string `yml:"target"`
	FileName string `yml:"file_name"`
}

type Strategy struct {
}

func Load(file string) (*Config, error) {
	return &Config{}, nil
}
