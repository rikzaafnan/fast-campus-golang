package data

import (
	"fmt"
	"math/rand"
)

func Generate(dataType string) any {

	switch dataType {
	case TYPE_NAME:
		return generateName()
	case TYPE_ADDRESS:
		return generateAddress()
	case TYPE_DATE:
		return generateDate()
	case TYPE_PHONE:
		return generatePhone()

	}

	return ""
}

func generateName() string {
	nameLength := len(name)

	index := rand.Intn(nameLength)

	return name[index]
}

func generateAddress() string {
	cityLength := len(address[SUBTYPE_CITY])
	streetLength := len(address[SUBTYPE_STREET])

	indexStreet := rand.Intn(streetLength)
	indexCity := rand.Intn(cityLength)
	indexNo := rand.Intn(100)

	return fmt.Sprintf("Jl. %s No. %d, %s", address[SUBTYPE_STREET][indexStreet], indexNo, address[SUBTYPE_CITY][indexCity])
}

func generateDate() string {
	return ""
}

func generatePhone() string {
	return ""
}
