package debug

import (
	"runtime/debug"

	log "github.com/sirupsen/logrus"
)

func LoggingLevels() {
	//logFile, err := os.Create("app.log")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer logFile.Close()
	//
	//log.SetOutput(logFile)
	log.SetLevel(log.WarnLevel)
	log.Debugf("Информационное сообщение")
	log.Info("TEST")
	log.Warnf("TEST %s", debug.Stack())
	log.Errorf("Ошибка в файле: %s", "example.txt")
}
