package entity

import "time"

type ConsumerLimits struct {
	IDLimit     string    `db:"id_limit,omitempty" json:"id_limit,omitempty"`
	IDConsumer  string    `db:"id_consumer,omitempty" json:"id_consumer,omitempty"`
	Tenor       int64     `db:"tenor,omitempty" json:"tenor,omitempty"`
	LimitAmount float64   `db:"limit_amount,omitempty" json:"limit_amount,omitempty"`
	CreatedAt   time.Time `db:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   time.Time `db:"updated_at,omitempty" json:"updated_at,omitempty"`
}
