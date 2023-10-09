package entity

import "time"

type Consumers struct {
	IDConsumer    string    `db:"id_consumer,omitempty" json:"id_consumer,omitempty"`
	FullName      string    `db:"full_name,omitempty" json:"full_name,omitempty"`
	NIK           string    `db:"nik,omitempty" json:"nik,omitempty"`
	LegalName     string    `db:"legal_name,omitempty" json:"legal_name,omitempty"`
	Pob           string    `db:"pob,omitempty" json:"pob,omitempty"`
	Dob           time.Time `db:"dob,omitempty" json:"dob,omitempty"`
	Salary        float64   `db:"salary,omitempty" json:"salary,omitempty"`
	ImageIdentity string    `db:"image_identity,omitempty" json:"image_identity,omitempty"`
	ImageSelfie   string    `db:"image_selfie,omitempty" json:"image_selfie,omitempty"`
	CreatedAt     time.Time `db:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt     time.Time `db:"updated_at,omitempty" json:"updated_at,omitempty"`
}
