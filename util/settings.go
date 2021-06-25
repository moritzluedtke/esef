package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

/*
 * Copied in large part from the official Fyne Widget Demo.
 * https://github.com/fyne-io/fyne/tree/master/cmd/fyne_demo
 */

type Settings struct {
	FyneSettings app.SettingsSchema
}

func NewSettings() *Settings {
	s := &Settings{}
	s.load()

	return s
}

func (s *Settings) load() {
	err := s.loadFromFile(s.FyneSettings.StoragePath())
	if err != nil {
		fyne.LogError("Settings load error:", err)
	}
}

func (s *Settings) loadFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(filepath.Dir(path), 0700)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	decode := json.NewDecoder(file)

	return decode.Decode(&s.FyneSettings)
}

func (s *Settings) Save() error {
	return s.saveToFile(s.FyneSettings.StoragePath())
}

func (s *Settings) saveToFile(path string) error {
	err := os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil { // this is not an exists error according to docs
		return err
	}

	data, err := json.Marshal(&s.FyneSettings)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, 0644)
}
