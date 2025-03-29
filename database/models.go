package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string  `json:"uuid" gorm:"type:text"` // Add UUID field
	Username string  `json:"username" gorm:"unique"`
	Email    string  `json:"email" gorm:"unique"`
	Password []byte  `json:"-"`
	Nickname *string `json:"nickname" gorm:"unique;default:null"`
	Stories  []Story `json:"stories" gorm:"foreignKey:UserUUID;references:UUID"` // One-to-many relationship
	// Usages   []Usage `json:"usages" gorm:"foreignKey:UserUUID;references:UUID"`  // One-to-many relationship
}

type Story struct {
	gorm.Model
	UUID     string        `json:"uuid" gorm:"type:text"`
	Title    string        `json:"title" gorm:"type:text"`
	Content  string        `json:"content" gorm:"type:text"`
	UserUUID string        `json:"user_uuid" gorm:"type:text"` // Foreign key to User.UUID
	MetaData StoryMetaData `json:"meta_data" gorm:"foreignKey:StoryUUID;references:UUID"`
}

type StoryMetaData struct {
	gorm.Model
	UUID         string `json:"uuid" gorm:"type:text"`
	Genre        string `json:"genre" gorm:"type:text"`
	Setting      string `json:"setting" gorm:"type:text"`
	Protagonist  string `json:"protagonist" gorm:"type:text"`
	Antagonist   string `json:"antagonist" gorm:"type:text"`
	ConflictInfo string `json:"conflict_info" gorm:"type:text"`
	DialogueInfo string `json:"dialogue_info" gorm:"type:text"`
	Theme        string `json:"theme" gorm:"type:text"`
	Tone         string `json:"tone" gorm:"type:text"`
	Pacing       string `json:"pacing" gorm:"type:text"`
	StoryUUID    string `json:"story_uuid" gorm:"type:text"`
}
