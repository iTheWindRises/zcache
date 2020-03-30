package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)
const (
	ERR_COLOR ="\033[31m[err]\033[0m "
	INOF_COLOR = "\033[32m[info]\033[0m "
	InfoLevel = iota
	ErrLevel
	Disabled
)

var (
	Error = errLog.Println
	Errorf = errLog.Printf
	Info = infoLog.Println
	Infof = infoLog.Printf
)


var (
	errLog = log.New(os.Stdout,ERR_COLOR,log.LstdFlags|log.Lshortfile)
	infoLog = log.New(os.Stdout,INOF_COLOR,log.LstdFlags|log.Lshortfile)
	logs = []*log.Logger{errLog,infoLog}
	mu sync.Mutex
)

func SetLevel(level int)  {
	mu.Lock()
	defer mu.Unlock()

	for _, log := range logs {
		log.SetOutput(os.Stdout)
	}

	if level > ErrLevel {
		errLog.SetOutput(ioutil.Discard)
	}
	if level > InfoLevel {
		infoLog.SetOutput(ioutil.Discard)
	}
}