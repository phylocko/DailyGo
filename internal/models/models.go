package models

import (
	"fmt"
	"net"

	"gorm.io/gorm"
)

const (
	RESOURCE_STATUS_OK  string = "ok"
	RESOURCE_STATUS_NEW string = "new"
	RESOURCE_STATUS_ERR string = "error"
)

type Resource struct {
	gorm.Model
	Domain string
	Status string
	IPs    []IpAddress `gorm:"constraint:OnDelete:SET NULL;"`
}

type IpAddress struct {
	gorm.Model
	ResourceID uint
	Ip         string
}

func (r *Resource) Resolve(db *gorm.DB) {
	ips, err := net.LookupHost(r.Domain)
	if err != nil {
		fmt.Printf("Error resolving resource '%s'\n", r.Domain)
		r.Status = RESOURCE_STATUS_ERR
		r.Save(db)
		return
	}

	var ip IpAddress
	db.Where("resource_id = ?", r.ID).Delete(&ip)

	if len(ips) == 0 {
		r.Status = RESOURCE_STATUS_ERR
		r.Save(db)
		return
	}

	for _, ip := range ips {
		a := IpAddress{ResourceID: r.ID, Ip: ip}
		db.Create(&a)
	}
	r.Status = RESOURCE_STATUS_OK
	r.Save(db)
}

func (r *Resource) Save(db *gorm.DB) {
	db.Save(&r)
}
