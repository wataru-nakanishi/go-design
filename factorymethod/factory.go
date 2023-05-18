package main

import "errors"

func getGun(gunType string) (IGun, error) {
	switch gunType {
	case "ak47":
		return NewAK47(), nil
	case "musket":
		return NewMusket(), nil
	default:
		return nil, errors.New("wrong gun type passed")
	}
}