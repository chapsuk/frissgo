package config

const (
	OutputFormatText = "text"
	OutputFormatJSON = "json"

	OutputTargetStdout = "stdout"
	OutputTargetFile   = "file"
)

type Output struct {
	Format   string `yml:"format"`
	Target   string `yml:"target"`
	FileName string `yml:"file_name"`
}
