package data

import (
	"fmt"
	"math/rand"
	"strings"
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
	year := 1950 + rand.Intn(100)
	month := 1 + rand.Intn(12)
	day := 1 + rand.Intn(28)

	return fmt.Sprintf("%02d-%02d-%d", day, month, year)
}

func generatePhone() string {
	prefixLength := 6 + rand.Intn(4)
	var sb strings.Builder
	sb.WriteString("081")
	for i := 0; i < prefixLength; i++ {
		digit := rand.Intn(10)
		digitString := fmt.Sprintf("%d", digit)
		sb.WriteString(digitString)
	}

	result := sb.String()

	return result
}
