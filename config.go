package main

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

// TODO:
// Config holds the configuration options for the application.
//
// At the moment, it is quite limited, only supporting the home folder and the
// file name of the metadata.
type Config struct {
	Root string `env:"SNP_ROOT" yaml:"root"`

	DefaultLanguage string `env:"SNP_DEFAULT_LANGUAGE" yaml:"default_language"`

	Theme string `env:"SNP_THEME" yaml:"theme"`

	ForegroundColor    string `env:"SNP_FOREGROUND" yaml:"foreground"`
	BackgroundColor    string `env:"SNP_BACKGROUND" yaml:"background"`
	RedColor           string `env:"SNP_RED" yaml:"red"`
	GreenColor         string `env:"SNP_GREEN" yaml:"green"`
	YellowColor        string `env:"SNP_YELLOW" yaml:"yellow"`
	BlueColor          string `env:"SNP_BLUE" yaml:"blue"`
	MagentaColor       string `env:"SNP_MAGENTA" yaml:"magenta"`
	CyanColor          string `env:"SNP_CYAN" yaml:"cyan"`
	BrightRedColor     string `env:"SNP_BRRED" yaml:"bright_red"`
	BrightGreenColor   string `env:"SNP_BRGREEN" yaml:"bright_green"`
	BrightYellowColor  string `env:"SNP_BRYELLOW" yaml:"bright_yellow"`
	BrightBlueColor    string `env:"SNP_BRBLUE" yaml:"bright_blue"`
	BrightMagentaColor string `env:"SNP_BRMAGENTA" yaml:"bright_magenta"`
	BrightCyanColor    string `env:"SNP_BRCYAN" yaml:"bright_cyan"`
	GrayColor    string `env:"SNP_GRAY" yaml:"gray"`
}

func newConfig() Config {
	return Config{
		Root: defaultRoot(),
		// File:                "snippets.json",
		DefaultLanguage:    defaultLanguage,
		Theme:              "dracula",
		ForegroundColor:    "15",
		BackgroundColor:    "0",
		RedColor:           "1",
		GreenColor:         "2",
		YellowColor:        "3",
		BlueColor:          "4",
		MagentaColor:       "5",
		CyanColor:          "6",
		BrightRedColor:     "9",
		BrightGreenColor:   "10",
		BrightYellowColor:  "11",
		BrightBlueColor:    "12",
		BrightMagentaColor: "13",
		BrightCyanColor:    "14",
		GrayColor:          "7",
	}
}

// TODO:
// default helpers for the configuration.
// We use $XDG_DATA_HOME to avoid cluttering the user's home directory.
func defaultRoot() string { return filepath.Join(xdg.DataHome, "snp") }
