package pospal

import (
	"encoding/json"
	"bytes"
)

const (
	ApiGroup = "/pospal-api2/openapi/v1/customerOpenApi"
)

type Customer struct {
	Uid                       json.Number `json:"customerUid"`
	CategoryName              string      `json:"categoryName"`
	Number                    string      `json:"number"`
	Name                      string      `json:"name"`
	Point                     float64       `json:"point"`
	Discount                  float64       `json:"discount"`
	Balance                   float64     `json:"balance"`
	Phone                     string      `json:"phone"`
	Birthday                  string      `json:"birthday"`
	QQ                        string      `json:"qq"`
	Email                     string      `json:"email"`
	Address                   string      `json:"address"`
	Remarks                   string      `json:"remarks"`
	CreatedDate               string      `json:"createdDate"`
	OnAccount                 int64       `json:"onAccount"`
	Enable                    int         `json:"enable"`
	Password                  string      `json:"password"`
	ExpiryDate                string      `json:"expiryDate"`
	CreateStoreAppIdOrAccount string      `json:"createStoreAppIdOrAccount"`
}

type PostBackParameter struct {
	ParameterType  string `json:"parameterType"`
	ParameterValue string `json:"parameterValue"`
}

type Customers struct {
	PostBackParameter PostBackParameter `json:"postBackParameter"`
	PageSize          int64             `json:"pageSize"`
	Result            []Customer        `json:"result"`
}

func (c *PPClient) queryByNumber(customerNum int64) (Customer, error) {
	var buf bytes.Buffer
	buf.WriteString(ApiGroup)
	buf.WriteString("/queryByNumber")
	retBytes, err := c.Invoke(buf.String(), Params{
		"appId":       c.AppID,
		"customerNum": customerNum,
	})

	rawResponse, err := ParseRawResponse(retBytes)

	if err != nil {
		panic(err)
	}

	respBytes, err := json.Marshal(rawResponse["data"])

	var customer Customer
	err = json.Unmarshal(respBytes, &customer)
	return customer, err
}

func (c *PPClient) queryByUid(customerUid int64) (Customer, error) {
	var buf bytes.Buffer
	buf.WriteString(ApiGroup)
	buf.WriteString("/queryByUid")
	retBytes, err := c.Invoke(buf.String(), Params{
		"appId":       c.AppID,
		"customerUid": customerUid,
	})

	rawResponse, err := ParseRawResponse(retBytes)
	if err != nil {
		panic(err)
	}
	respBytes, err := json.Marshal(rawResponse["data"])
	var customer Customer
	err = json.Unmarshal(respBytes, &customer)
	return customer, err
}

func (c *PPClient) queryCustomerPages(params Params, postBackParameter Params) (Customers, error) {
	var buf bytes.Buffer
	buf.WriteString(ApiGroup)
	buf.WriteString("/queryCustomerPages")
	params["appId"] = c.AppID
	if postBackParameter != nil {
		params["postBackParameter"] = postBackParameter
	}
	retBytes, err := c.Invoke(buf.String(), params)

	rawResponse, err := ParseRawResponse(retBytes)
	if err != nil {
		panic(err)
	}
	respBytes, err := json.Marshal(rawResponse["data"])
	var customers Customers
	err = json.Unmarshal(respBytes, &customers)
	return customers, err
}
