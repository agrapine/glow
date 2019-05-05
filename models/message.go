package models

import "time"

type Message struct {
  Id        string    `json:"id,omitempty"`
  ChannelId string    `json:"channelId"`
  SenderId  string    `json:"senderId"`
  CreatedAt time.Time `json:"createdAt"`
  Content   string    `json:"content"`
}