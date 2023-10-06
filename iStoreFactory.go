package main

import "fmt"

type IStoreFactory interface {
	sellPhone() IPhone
	sellLaptop() ILaptop
	sellWatch() IWatch
}

func GetStoreFactory(brand string) (IStoreFactory, error) {
	if brand == "Apple" {
		return &Apple{}, nil
	}

	if brand == "Samsung" {
		return &Samsung{}, nil
	}
	return nil, fmt.Errorf("No such brand found")
}
