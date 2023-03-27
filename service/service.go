package service

import (
	"fmt"
	"instasafe1/common"
	"instasafe1/repository"
	"strconv"
	"time"

	"github.com/go-chassis/openlog"
)

type Service struct {
	Rep repository.Repository
}

func (s *Service) CreateEndUser(payload map[string]interface{}, language string) common.HTTPResponse {
	fmt.Println(payload)
	email := payload["Email"].(string)
	data := s.Rep.FindUserByEmail(email)
	fmt.Println(data)
	if len(data) > 0 {
		return common.ErrorHandler("103", nil, 0, language)
	}
	payload["city"] = ""
	payload["resetLocation"] = false
	res := s.Rep.CreateEndUser(payload)
	return common.ErrorHandler("102", res, 1, language)
}

func (s *Service) CreateTransaction(payload map[string]interface{}, language string) common.HTTPResponse {
	var fixedDiff time.Duration = 60000000000
	currentTime := time.Now().UTC()
	fmt.Println("@@@@@@@@@@@@@@@  ", currentTime)
	timeStamp, err1 := time.Parse(time.RFC3339, payload["timestamp"].(string))
	if err1 != nil {
		return common.ErrorHandler("106", nil, 0, language)
	}
	diff := currentTime.Sub(timeStamp)
	if diff < 0 {
		return common.ErrorHandler("107", nil, 0, language)
	}
	if diff > fixedDiff {
		return common.ErrorHandler("108", nil, 0, language)
	}
	res := s.Rep.CreateTransaction(payload)
	return common.ErrorHandler("104", res.(map[string]interface{}), 1, language)
}

func (s *Service) GetStatistics(userid, city, language string) common.HTTPResponse {
	output := make(map[string]interface{})
	userDetailes := s.Rep.GetUserByID(userid)
	var sum, min, max float64
	var count int
	if resetLocation, cok := userDetailes["resetLocation"].(bool); cok {
		transactions := make([]map[string]interface{}, 0)
		if resetLocation {
			transactions = s.Rep.GetAllTransactions()
		} else {
			if userDetailes["city"].(string) != city {
				return common.ErrorHandler("113", nil, 0, language)
			}
			transactions = s.Rep.GetTransactionByLocation(city)
		}
		var fixedDiff time.Duration = 60000000000
		currentTime := time.Now().UTC()
		for _, transaction := range transactions {
			timeStamp, err3 := time.Parse(time.RFC3339, transaction["timestamp"].(string))
			if err3 != nil {
				openlog.Error(err3.Error())
				return common.ErrorHandler("111", nil, 0, language)
			}
			diff := currentTime.Sub(timeStamp)
			if diff <= fixedDiff {
				amt, err2 := strconv.ParseFloat(transaction["amount"].(string), 32)
				if err2 != nil {
					openlog.Error(err2.Error())
					return common.ErrorHandler("112", nil, 0, language)
				}
				sum += amt
				count++
				if count == 1 {
					min = amt
				}
				if amt < min {
					min = amt
				}
				if amt > max {
					max = amt
				}
			}
		}
		output["Sum"] = sum
		if count != 0 {
			output["avg"] = sum / float64(count)
		} else {
			output["avg"] = 0
		}
		output["max"] = max
		output["min"] = min
		output["count"] = count
	}
	return common.ErrorHandler("114", output, 1, language)
}

func (s *Service) DeleteAllTransactions(language string) common.HTTPResponse {
	res := s.Rep.DeleteAllTransactions()
	if len(res) != 0 {
		return common.ErrorHandler("115", res, 0, language)
	}
	return common.ErrorHandler("116", res, 0, language)
}

func (s *Service) AddLoaction(uid string, payload map[string]interface{}, language string) common.HTTPResponse {
	payload["resetLocation"] = false
	res := s.Rep.UpdateLocation(uid, payload)
	fmt.Println(res)
	return common.ErrorHandler("117", res, 1, language)
}

func (s *Service) ResetLoaction(uid string, payload map[string]interface{}, language string) common.HTTPResponse {
	payload["resetLocation"] = true
	res := s.Rep.UpdateLocation(uid, payload)
	fmt.Println(res)
	return common.ErrorHandler("118", res, 1, language)
}
