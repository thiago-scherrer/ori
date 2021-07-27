package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Breed         string `yaml:"breed"`
	Classes       string `yaml:"classes"`
	Combat        string `yaml:"combat"`
	Experience    string `yaml:"experience"`
	Help          string `yaml:"help"`
	Magic         string `yaml:"magic"`
	SubAttributes string `yaml:"subattributes"`
	Weapons       string `yaml:"weapons"`
}

func (c *conf) getYml() *conf {
	ymlFile := os.Getenv("YML_FILE")
	if ymlFile == "" {
		log.Fatalln("YML_FILE required!")
	}

	ymlF, err := ioutil.ReadFile(filepath.Clean(ymlFile))
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(ymlF, c)
	if err != nil {
		log.Fatalln(err)
	}

	return c
}

func breed(bot *tgbotapi.BotAPI, chatId int64) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.Breed)
	bot.Send(msg)
}

func classes(bot *tgbotapi.BotAPI, chatId int64) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.Classes)
	bot.Send(msg)
}

func combat(bot *tgbotapi.BotAPI, chatId int64) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.Combat)
	bot.Send(msg)
}

func experience(bot *tgbotapi.BotAPI, chatId int64) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.Experience)
	bot.Send(msg)
}

func help(bot *tgbotapi.BotAPI, chatId int64) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.Help)
	bot.Send(msg)
}

func magic(bot *tgbotapi.BotAPI, chatId int64) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.Magic)
	bot.Send(msg)
}

func subAttr(bot *tgbotapi.BotAPI, chatId int64) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.SubAttributes)
	bot.Send(msg)
}

func weapons(bot *tgbotapi.BotAPI, chatId int64) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.Weapons)
	bot.Send(msg)
}
