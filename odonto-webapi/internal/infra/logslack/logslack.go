
package logslack

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"odonto/internal/infra/config"
	"github.com/slack-go/slack"
)

//=======================================================
// Logger as Singleton
//=======================================================

type Logger struct {
	project string
	channel string
	token   string
}

func (l *Logger) send(projectName, message string) {
	//create conn with slack api
	api := slack.New(l.token)
	//message to sendo to slack
	msg := slack.MsgOptionText(fmt.Sprintf("[%s] %s", projectName, message), false)
	//send message to slack
	_, _, _, err := api.SendMessage(l.channel, msg)
	if err != nil {
		log.Println(err.Error())
	}
}

func (l *Logger) Print(msg string) {
	go l.send(l.project, msg)
}

func (l *Logger) Errorf(msg string, params ...interface{}) {
	go l.send(l.project, fmt.Sprintf(msg, params...))
}

func (l *Logger) Error(err error) error {
	go l.send(l.project, fmt.Sprintf("[error] %s", err.Error()))
	return err
}

//=======================================================
// Settings as Singleton
//=======================================================

var loggerInst *Logger = nil
var loggerInstOnce = sync.Once{}

// Init initializes settings as singleton instance from yaml file
func Init(project string) {
	loggerInstOnce.Do(func() {
		channel := config.GetSettings().Integrations.Slack.Channel
		token := config.GetSettings().Integrations.Slack.Token
		if channel == "" && token == "" {
			panic(errors.New("slack channel or token not configured"))
		}
		loggerInst = &Logger{
			project: project,
			channel: channel,
			token:   token,
		}
	})
}

// GetSettings get settings as singleton instance
func GetLogger() *Logger {
	return loggerInst
}

