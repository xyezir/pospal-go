package pospal

import (
	"os"
	"errors"
	"fmt"
	"testing"
)

var (
	AppID string
	AppKey string
	client PPClient
)

func init() {
	AppID = os.Getenv("PosPalAppID")
	AppKey = os.Getenv("PosPalAppKey")

	if AppID == "" {
		panic(errors.New("环境变量中找不到AppID"))
	}

	if AppKey == "" {
		panic(errors.New("环境变量中找不到AppKey"))
	}

	client = PPClient{AppID: AppID, AppKey: AppKey}
}

func TestQueryByNumber(t *testing.T) {
	fmt.Println("TestQueryByNumber")
	result, err := client.queryByNumber(13018882118)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func TestQueryByUid(t *testing.T) {
	fmt.Println("TestQueryByUid")
	result, err := client.queryByUid(238491332794626212)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func TestQueryCustomerPages(t *testing.T) {
	fmt.Println("TestQueryCustomerPages")
	result, err := client.queryCustomerPages(Params{

	}, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}