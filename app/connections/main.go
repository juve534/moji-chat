package main

import (
    "errors"
	"fmt"
    "os"

    "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
    "github.com/soveran/redisurl"
)

var (
	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func handler(request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
    cid := request.RequestContext.ConnectionID
    var_dump(cid)

	res := SaveUser(cid)
	var_dump(res)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

func var_dump(v ...interface{}) {
    for _, vv := range(v) {
        fmt.Printf("%#v\n", vv)
    }
}

// SaveUser ユーザを保存する
func SaveUser(cid string) interface{} {
    ep := os.Getenv("CacheEndPoint")
    conn, err := redisurl.ConnectToURL("telnet://" + ep)
    if err != nil {
        fmt.Println(err)
        panic(err)
    }

    val, err := conn.Do("SET", "users", cid, "NX", "EX", "120")
    if err != nil {
        fmt.Println(err)
        panic(err)
    }

    // 存在判定
    if val == nil {
        fmt.Println("既にオンラインです。")
        panic(err)
    }

    return val
}