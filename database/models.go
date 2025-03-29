package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string  `json:"uuid" gorm:"type:uuid;not null;unique"`
	Username string  `json:"username" gorm:"unique;not null"`
	Email    string  `json:"email" gorm:"unique;not null"`
	Password []byte  `json:"-"`
	Nickname *string `json:"nickname" gorm:"unique;default:null"`
	Stories  []Story `json:"stories" gorm:"foreignKey:UserUUID;references:UUID"` // One-to-many relationship
	// Usages   []Usage  `json:"usages" gorm:"foreignKey:UserUUID;references:UUID"` // Uncomment if necessary
}

type Story struct {
	gorm.Model
	UUID     string        `json:"uuid" gorm:"type:uuid;not null;unique"`
	Title    string        `json:"title" gorm:"type:text;not null"`
	Content  string        `json:"content" gorm:"type:text;not null"`
	UserUUID string        `json:"user_uuid" gorm:"type:uuid;not null"` // Foreign key to User.UUID
	MetaData StoryMetaData `json:"meta_data" gorm:"foreignKey:StoryUUID;references:UUID"`
}

type StoryMetaData struct {
	gorm.Model
	UUID         string `json:"uuid" gorm:"type:uuid;not null;unique"`
	Genre        string `json:"genre" gorm:"type:text;not null"`
	Setting      string `json:"setting" gorm:"type:text;not null"`
	Protagonist  string `json:"protagonist" gorm:"type:text;not null"`
	Antagonist   string `json:"antagonist" gorm:"type:text;not null"`
	ConflictInfo string `json:"conflict_info" gorm:"type:text;not null"`
	DialogueInfo string `json:"dialogue_info" gorm:"type:text;not null"`
	Theme        string `json:"theme" gorm:"type:text;not null"`
	Tone         string `json:"tone" gorm:"type:text;not null"`
	Pacing       string `json:"pacing" gorm:"type:text;not null"`
	StoryUUID    string `json:"story_uuid" gorm:"type:uuid;not null"`
}
