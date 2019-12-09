package mysql

import (
	"github.com/BurntSushi/toml"
	"hot/common/util"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

type Config struct {
	Source, Driver             string
	MaxIdleConns, MaxOpenConns int
	ConnMaxLifetime            time.Duration
}

var (
	config     *Config
	once       sync.Once
	configLock = new(sync.RWMutex)
)

func GetConfig() *Config {
	once.Do(loadConfig)
	// 加读锁
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func getMySqlFilePath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		util.Error.Fatalln(err)
	}
	return currentPath + "/src/config.toml"
}

func loadConfig() {
	var err error
	filePath, err := filepath.Abs(getMySqlFilePath())
	if err != nil {
		util.Error.Fatalln(err)
	}
	cfg := new(Config)
	_, err = toml.DecodeFile(filePath, cfg)
	if err != nil {
		util.Error.Fatalln(err)
	}
	// 加写锁
	configLock.Lock()
	defer configLock.Unlock()
	config = cfg
}

func StartReloadConfigListener() {
	s := make(chan os.Signal, 1)
	// kill -30 go_server_pid
	signal.Notify(s, syscall.SIGUSR1)
	go func() {
		for {
			<-s
			loadConfig()
			log.Println("Reloaded config")
		}
	}()
}
