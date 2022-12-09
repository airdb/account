package chatgpt_test

import (
	"os"
	"testing"
	"time"

	"github.com/airdb/databus/chatgpt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var sessionToken = os.Getenv("SESSION_KEY")

func TestMain(m *testing.M) {
	if sessionToken == "" {
		panic("env SESSION_KEY not set")
	}
	logrus.SetLevel(logrus.DebugLevel)
	m.Run()
}

func TestChatGPT_SendMessage(t *testing.T) {
	timeout := time.Second * 60
	client, err := chatgpt.NewChatGPT(sessionToken, chatgpt.ChatGPTOptions{
		Log:     logrus.NewEntry(logrus.StandardLogger()),
		Timeout: &timeout,
	})
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	conversation := client.NewConversation("", "")
	resp, err := conversation.SendMessage("hello")
	if assert.NoError(t, err) {
		t.Logf("resp: %s", resp)
	} else {
		t.FailNow()
	}

	cid := conversation.ConversationId

	resp, err = conversation.SendMessage("what's your name")
	if assert.NoError(t, err) {
		t.Logf("resp: %s", resp)
		assert.Equal(t, cid, conversation.ConversationId)
	} else {
		t.FailNow()
	}
}

func TestChatGPT_RefreshAccessToken(t *testing.T) {
	client, err := chatgpt.NewChatGPT(sessionToken, chatgpt.ChatGPTOptions{
		Log: logrus.NewEntry(logrus.StandardLogger()),
	})
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	err = client.RefreshAccessToken()
	if assert.NoError(t, err) {
		t.Logf("accessToken: %s", client.AccessToken)
	}
}
