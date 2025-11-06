package models

import "gorm.io/gorm"

type SampleKv struct {
	gorm.Model
	Key   string
	Value string
}
