package main

type Status int

const (
	INIT Status = iota
	RUNNING
	TERMINATED
)
