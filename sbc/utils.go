package sbc

import (
	"errors"
	"regexp"
	"strings"
)

func Append2Digitmap(customerName string, phoneNumbers []string, digitMap []TBDigitMap) ([]TBDigitMap, error) {
	// copy to digitmaporig incase something fails
	var digitMapOrig []TBDigitMap
	copy(digitMapOrig, digitMap)

	// split numbers fdrom flag and append to digitmap

	for pN1, _ := range phoneNumbers {
		for pN2, _ := range phoneNumbers {
			if pN1 != pN2 {
				if phoneNumbers[pN1] == phoneNumbers[pN2] {
					return nil, errors.New("duplicate numbers trying to be sent")
				}
			}
		}
	}

	var duplicateNumber = false
	for _, i := range digitMap {
		for _, i2 := range phoneNumbers {
			if strings.Contains(i.Called, i2) {
				duplicateNumber = true
				return nil, errors.New("digitmap already contains")
			}
		}
	}

	if duplicateNumber {
		return nil, errors.New("duplicate numbers trying to be sent")
	}

	for _, i := range phoneNumbers {
		// create new item
		newDigitMapping := TBDigitMap{
			Called:  i,
			Calling: "",
			//todo friendlyify name to match scheme
			RouteSetName: customerName,
		}

		match, _ := regexp.MatchString("^(\\+\\d{1,2}\\s)?\\(?\\d{3}\\)?[\\s.-]?\\d{3}[\\s.-]?\\d{4}$", i)
		if !match {
			return nil, errors.New("regex on number no worky :'(")
		}

		// append item
		digitMap = append(digitMap, newDigitMapping)
	}

	return digitMap, nil
}
