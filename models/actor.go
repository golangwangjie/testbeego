package models

import (
	"golangwangjie/initial"
)

type Actor struct {
	ActorId		string `gorm:"column:actor_id"`
	FirstName	string `gorm:"column:first_name"`
	LastName	string `gorm:"column:last_name"`
	LastUpdate	string `gorm:"column:last_update"`
}

func (Actor) TableName() string {
	return "actor"
}

func NewActor() *Actor {
	return &Actor{}
}

func GetActorInfo(id int) *Actor {
	list := NewActor()
	// initial.MysqlDb.Select("actor_id, first_name, last_name,last_update").Where("actor_id = ?", id).Find(list)
	initial.MysqlDb.First(list)
	return list
}
