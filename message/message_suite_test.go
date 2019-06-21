package message_test

import (
	"os"
	"strconv"
	"testing"

	wechatwork "github.com/dfang/wechat-work-go"
	"github.com/dfang/wechat-work-go/contact"
	"github.com/dfang/wechat-work-go/message"
	_ "github.com/joho/godotenv/autoload"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var g *message.Message
var c *contact.Contact

func TestMessage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Message Suite")
}

var _ = BeforeSuite(func() {
	corpID := os.Getenv("CORP_ID")
	corpSecret := os.Getenv("CORP_SECRET")
	agentID, _ := strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)

	corp := wechatwork.New(corpID)
	app := corp.NewApp(corpSecret, agentID)
	g = &message.Message{
		App: app,
	}

	contactAppSecret := os.Getenv("CONTACT_APP_SECRET")
	// 关于创建成员（客服答复）
	// 目前只能使用通讯录的secret 获取token进行创建  其他的secret是没有创建成员的权限的
	// 获取路径：通讯录管理secret。在“管理工具”-“通讯录同步”里面查看（需开启“API接口同步”）
	app2 := corp.NewApp(contactAppSecret, 0)
	c = &contact.Contact{
		App: app2,
	}
})
