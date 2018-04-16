## 银豹开放平台 golang sdk client

[![License][license-image]][license-url]

## Requirement


**[Golang](https://golang.org/)**

## Installation

```shell
dep ensure -add github.com/xyezir/pospal-go
```

## Configuration

```shell
export PosPalAppID=Your Pospal AppID
export PosPalAppKey=Your Pospal AppKey
```

## Usage

```go
import (
    ...
    pospal "github.com/xyezir/pospal-go"
)

var (
	AppID string
	AppKey string
	client pospal.PPClient
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
```

## Documentation

todo list

## Bugs report

If you find a bug, please report it using the [issue tracker](http://github.com/xyezir/pospal-go/issues).

## Development

```shell
$ /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
$ brew install go dep
$ dep ensure
$ vim client.go
$ go test
...
```

## Contributors

* Chris Ye ([xyezir](https://github.com/xyezir))

[license-image]: http://img.shields.io/npm/l/leafjs.svg?style=flat-square
[license-url]: LICENSE