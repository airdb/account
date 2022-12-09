package chatgpt

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/imroc/req/v3"
)

const OPENAI = "https://api.openai.com/v1/completions"

func NewChat(text string) string {
	token := os.Getenv("OPENAI_API_KEY")

	client := req.C().DevMode().
		SetUserAgent("airdb-bot/v1").
		SetTimeout(30 * time.Second)

	r := &OpenAIReq{
		Temperature:      0,
		MaxTokens:        100,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Model:            ModelTextDavinci003,
		Prompt:           text,
	}

	var ret OpenAIResp

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+token).
		SetBody(r).
		SetResult(&ret).
		Post(OPENAI)
	if err != nil {
		log.Println("err: ", err)
		return ErrorGPTOverload
	}

	if !resp.IsSuccess() {
		fmt.Println("bad response status:", resp.Status)
		return ErrorGPTOverload
	}

	// log.Println("Response", resp.String())
	if len(ret.Choices) == 0 {
		return ErrorNotFound
	}

	// return ret.Choices[0].Text
	return strings.TrimPrefix(ret.Choices[0].Text, "\n\n")
}
