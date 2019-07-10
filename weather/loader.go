package weather

import (
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/json"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"path/filepath"
	"sync"
)

var (
	m        sync.RWMutex
	err      error
	cityCode []byte
)

func loader(jsonConfName string) {
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("./", string(filepath.Separator))))

	e := json.NewEncoder()
	fmt.Println(appPath)
	fileSource := file.NewSource(file.WithPath(appPath+jsonConfName), source.WithEncoder(e))
	conf := config.NewConfig()

	// 加载json文件
	if err = conf.Load(fileSource); err != nil {
		panic(err)
	}
	cityCode = conf.Bytes()
}

func main() {
	loader("/conf/city.json")
}
