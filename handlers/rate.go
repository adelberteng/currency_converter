package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"errors"

	"github.com/adelberteng/currency_converter/models"
	"github.com/adelberteng/currency_converter/utils"
)

var (
	logger = utils.GetLogger()
	rateData = models.GetRateDataModel()
)

func GetCurrencyRate(currencyType, targetType string) (interface{}, error) {
	var val interface{}
	var err error

	if targetType != "" {
		val, err = rateData.GetRate(currencyType, targetType)
	} else {
		val, err = rateData.GetAllRate(currencyType)
	}
	if err != nil {
		logger.Error(err)
		return "", err
	}

	logger.Info("currencyType: " + currencyType + " targetType: " + targetType + " val: " + fmt.Sprint(val))

	return val, nil
}

type Response struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Result  float64 `json:"result"`
}

func CountCurrencyRate(postBody map[string]string) (Response, error) {
	var res Response

	currencyType := postBody["currency_type"]
	targetType := postBody["target_type"]
	amountStr := postBody["amount"]
	if currencyType == "" || targetType == "" {
		res.Message = "currency type and target_type is required."
		logger.Info(res)
		return res, errors.New(res.Message)
	}

	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil {
		res.Status = http.StatusInternalServerError
		logger.Error(err)
		return res, err
	}

	rateMap, err := rateData.GetAllRate(currencyType)
	if err != nil {
		res.Status = http.StatusInternalServerError
		logger.Error(err)
		return res, err
	}
	val_str := rateMap[targetType]
	if val_str == "" {
		res.Status = http.StatusBadRequest
		res.Message = "currency type and target_type are not currect."
		return res, errors.New(res.Message)
	}
	val, err := strconv.ParseFloat(val_str, 64)
	if err != nil {
		res.Status = http.StatusInternalServerError
		logger.Error(err)
		return res, err
	}

	res.Result = math.Round((float64(amount)*val)*1000) / 1000
	res.Message = "exchange complete."
	res.Status = http.StatusOK
	logger.Info(res)

	return res, nil
}
