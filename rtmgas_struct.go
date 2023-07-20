package rtmgas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Absensi struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NIP       string             `bson:"nip,omitempty" json:"nip,omitempty"`
	Date      string             `bson:"date,omitempty" json:"date,omitempty"`
	Status    string             `bson:"status,omitempty" json:"status,omitempty"`
	Employee  string             `bson:"employee,omitempty" json:"employee,omitempty"`
	Positions string             `bson:"positions,omitempty" json:"positions,omitempty"`
}
type Posisi struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Posisi string             `bson:"nama_posisi,omitempty" json:"nama_posisi,omitempty"`
	Bidang      string             `bson:"bidang,omitempty" json:"bidang,omitempty"`
}
