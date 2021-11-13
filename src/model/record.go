package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 表名
const TRecord = "t_record"

type Record struct {
	ID          *primitive.ObjectID `json:"id" bson:"_id" `                  // id
	UID         *primitive.ObjectID `json:"uid" bson:"uid" `                 // 所有者id
	CreateAt    *time.Time          `json:"createAt" bson:"createAt" `       // 创建时间
	UpdateAt    *time.Time          `json:"updateAt" bson:"updateAt" `       // 更新时间
	ReviewAt    *time.Time          `json:"reviewAt" bson:"reviewAt" `       // 复习时间
	CooldownAt  *time.Time          `json:"cooldownAt" bson:"cooldownAt" `   // 冷却时间
	Source      string              `json:"source" bson:"source" `           // 原文
	Translation string              `json:"translation" bson:"translation" ` // 译文
	InReview    bool                `json:"inReview" bson:"inReview" `       // 是否在复习中
	Exp         int64               `json:"exp" bson:"exp" `                 // 复习熟练度
}
