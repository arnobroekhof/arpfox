// +build !darwin, !linux, !freebsd

package arp

import (
	"errors"
)

func doARPLookup(ip string) (*Address, error) {
	return nil, errors.New("Not implemented")
}
