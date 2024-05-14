package log

import (
	"fmt"
	"wxapp/config"
	"wxapp/pkg/fs"
)

func LoadLog() {
	path := config.GetConfig().Logger.Path
	if err := fs.IsNotExistMkDir(path); err != nil {
		panic(fmt.Sprintf("create log dir failed: %v\n", err))
	}
	fmt.Println("loading log success...")
}
