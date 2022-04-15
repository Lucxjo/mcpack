package models

type Side uint8

const (
	Client Side = iota
	Server
	Both
)