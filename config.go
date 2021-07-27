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
	m, err := bot.Send(msg)
	if err != nil {
		log.Println(m, err)
	}
}

func classes(bot *tgbotapi.BotAPI, chatId int64) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.Classes)
	m, err := bot.Send(msg)
	if err != nil {
		log.Println(m, err)
	}
}

func combat(bot *tgbotapi.BotAPI, chatId int64) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.Combat)
	m, err := bot.Send(msg)
	if err != nil {
		log.Println(m, err)
	}
}

func experience(bot *tgbotapi.BotAPI, chatId int64) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.Experience)
	m, err := bot.Send(msg)
	if err != nil {
		log.Println(m, err)
	}
}

func help(bot *tgbotapi.BotAPI, chatId int64) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.Help)
	m, err := bot.Send(msg)
	if err != nil {
		log.Println(m, err)
	}
}

func magic(bot *tgbotapi.BotAPI, chatId int64) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.Magic)
	m, err := bot.Send(msg)
	if err != nil {
		log.Println(m, err)
	}
}

func subAttr(bot *tgbotapi.BotAPI, chatId int64) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.SubAttributes)
	m, err := bot.Send(msg)
	if err != nil {
		log.Println(m, err)
	}
}

func weapons(bot *tgbotapi.BotAPI, chatId int64) {
	var c conf
	c.getYml()

	msg := tgbotapi.NewMessage(chatId, c.Weapons)
	m, err := bot.Send(msg)
	if err != nil {
		log.Println(m, err)
	}

}
