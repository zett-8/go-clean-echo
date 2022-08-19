package public

import (
	"embed"
	_ "embed"
)

//go:embed index.html
var Files embed.FS
