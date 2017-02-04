package config

const (
	OutputFormatText = "text"
	OutputFormatJSON = "json"

	OutputTargetStdout = "stdout"
	OutputTargetFile   = "file"
)

type Output struct {
	Format   string `yaml:"format"`
	Target   string `yaml:"target"`
	FileName string `yaml:"file_name"`
}
