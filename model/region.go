package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Kecamatan
type District struct {
	Name        string `json:"name" csv:"kode"`
	Code        string `json:"code" csv:"nama"`
	RegencyCode string `csv:"kode_kabupaten"`
}

// Provinsi
type Province struct {
	Code string `json:"code" csv:"kode"`
	Name string `json:"name" csv:"nama"`
}

// Kabupaten
type Regency struct {
	Code         string `json:"code" csv:"kode"`
	Name         string `json:"name" csv:"nama"`
	ProvinceCode string `json:"province_code" csv:"kode_provinsi"`
}

// Kelurahan / Desa
type Village struct {
	Code         string `json:"code" csv:"kode"`
	Name         string `json:"name" csv:"nama"`
	DistrictCode string `json:"district_code" csv:"kode_kecamatan"`
}

type DeliveryAddress struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"  json:"_id"`
	Name      string             `bson:"name"           json:"name"`
	Village   string             `bson:"village"        json:"village"`
	District  string             `bson:"district"       json:"district"`
	Regency   string             `bson:"regency"        json:"regency"`
	Province  string             `bson:"province"       json:"province"`
	Detail    string             `bson:"detail"         json:"detail"`
	CreatedAt time.Time          `bson:"created_at"     json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"     json:"updated_at"`
}

func (da *DeliveryAddress) BeforeSave() DeliveryAddress {
	da.CreatedAt = time.Now()
	da.UpdatedAt = time.Now()
	return *da
}
