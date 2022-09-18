package models

import "blog/config"

type PigeonholeRes struct {
	config.Viewer
	config.SystemConfig
	Categorys []Category
	Lines     map[string][]Post
}
