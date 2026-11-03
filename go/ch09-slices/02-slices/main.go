package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	
	if plan == planPro{
		return messages[0:], nil

	}
	if plan == planFree{
		return messages[0:2], nil
	}
	return nil, errors.New("unsupported plan")
	
}
