package wechatwork_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	wechatwork "github.com/dfang/wechat-work-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/joho/godotenv/autoload"
)

var app *wechatwork.App

func TestWechatWorkGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WechatWorkGo Suite")
}

var _ = BeforeSuite(func() {
	corpID := os.Getenv("CORP_ID")
	corpSecret := os.Getenv("CORP_SECRET")
	agentID, _ := strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)

	corp := wechatwork.New(corpID)
	app = corp.NewApp(corpSecret, agentID)

	// block all HTTP requests
	// httpmock.ActivateNonDefault(resty.DefaultClient.GetClient())
	fmt.Println("Suite Started")
})

var _ = AfterSuite(func() {
	fmt.Println("Suite Finished")
})
