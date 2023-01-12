package main

type Order string

const (
	fill      Order = "fill"
	buy       Order = "buy"
	take      Order = "take"
	remaining Order = "remaining"
	exit      Order = "exit"
)
