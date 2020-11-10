package models

type Response struct {
    ID      int64  `json:"id,omitempty"`
    Message string `json:"message,omitempty"`
}
