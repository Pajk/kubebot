package main

import (
	"errors"
	"fmt"
	"os"
"log"
"net/http"
	"github.com/gorilla/mux"

	"github.com/go-chat-bot/bot/slack"
)

const (
	appVersion             string = "KUBEBOT_APP_VERSION"
	slackTokenLabel        string = "KUBEBOT_SLACK_TOKEN"
	slackChannelsLabel     string = "KUBEBOT_SLACK_CHANNELS_IDS"
	slackAdminsLabel       string = "KUBEBOT_SLACK_ADMINS_NICKNAMES"
	slackCommandsLabel     string = "KUBEBOT_SLACK_VALID_COMMANDS"
	notDefinedErrorMessage string = "%s env variable not defined"
)

var (
	kb *Kubebot
)

func validateEnvParams() error {
	if os.Getenv(appVersion) == "" {
		return errors.New(fmt.Sprintf(notDefinedErrorMessage, appVersion))
	}
	if os.Getenv(slackTokenLabel) == "" {
		return errors.New(fmt.Sprintf(notDefinedErrorMessage, slackTokenLabel))
	}
	if os.Getenv(slackChannelsLabel) == "" {
		return errors.New(fmt.Sprintf(notDefinedErrorMessage, slackChannelsLabel))
	}
	if os.Getenv(slackAdminsLabel) == "" {
		return errors.New(fmt.Sprintf(notDefinedErrorMessage, slackAdminsLabel))
	}
	if os.Getenv(slackCommandsLabel) == "" {
		return errors.New(fmt.Sprintf(notDefinedErrorMessage, slackCommandsLabel))
	}

	return nil
}

func main() {

	if err := validateEnvParams(); err != nil {
		fmt.Printf("Kubebot cannot run. Error: %s\n", err.Error())
		return
	}

	kb = &Kubebot{
		token:    os.Getenv(slackTokenLabel),
		admins:   stringToMap(os.Getenv(slackAdminsLabel), " "),
		channels: stringToMap(os.Getenv(slackChannelsLabel), " "),
		commands: stringToMap(os.Getenv(slackCommandsLabel), " "),
	}
	go slack.Run(kb.token)
	log.Print(os.Getenv(appVersion))
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", frontendHandler)
	router.HandleFunc("/version", versionHandler)
	router.HandleFunc("/healthz", healthzHandler)
	router.HandleFunc("/readinez", readinessHandler)

	log.Fatal(http.ListenAndServe(":80", router))
}
