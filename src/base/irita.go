package base

import (
	iritaV3 "github.com/bianjieai/avata-sdk-go/v3"
)

const (
	ServiceIrita = "irita"
)

var IIrita IServiceIrita

type IServiceIrita interface {
	Clients(index ...int) *iritaV3.AvataClient
}
