package chargify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func GetSubscriptionStatements(client Client, subscriptionId int64, pageNumber int32, perPage int32, statements interface{}) error {
	if subscriptionId == 0 {
		return NoID()
	}
	uri := fmt.Sprintf("subscriptions/%d/statements.json?direction=desc&per_page=%d&page=%d", subscriptionId, perPage, pageNumber)
	res, err := client.Get(uri)
	if err != nil {
		return err
	}
	if err = checkError(res); err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, statements)
}

func GetStatement(client Client, statementId int64, statement interface{}) error {
	if statementId == 0 {
		return NoID()
	}
	uri := fmt.Sprintf("statements/%d.json", statementId)
	res, err := client.Get(uri)
	if err != nil {
		return err
	}
	if err = checkError(res); err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, statement)
}
