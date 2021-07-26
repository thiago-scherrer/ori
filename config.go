package main

import (
	"io/ioutil"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Help    string `yaml:"help"`
	Magic   string `yaml:"magic"`
	Classes string `yaml:"classes"`
}

func (c *conf) getYml() *conf {
	ymlFile := os.Getenv("YML_FILE")
	if ymlFile == "" {
		log.Fatalln("YML_FILE required!")
	}

	ymlF, err := ioutil.ReadFile(ymlFile)
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(ymlF, c)
	if err != nil {
		log.Fatalln(err)
	}

	return c
}

func classes(bot *tgbotapi.BotAPI, chatId int64) (bool, error) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.Classes)
	bot.Send(msg)

	return true, nil
}

func magic(bot *tgbotapi.BotAPI, chatId int64) (bool, error) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.Magic)
	bot.Send(msg)

	return true, nil
}

func help(bot *tgbotapi.BotAPI, chatId int64) (bool, error) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.Help)
	bot.Send(msg)

	return true, nil
}
