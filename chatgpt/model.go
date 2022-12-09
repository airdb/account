package chatgpt

const (
	// Text
	ModelTextDavinci003 = "text-davinci-003"
	ModelTextCurie001   = "text-curie-001"
	ModelTextBabbage001 = "text-babbage-001"
	ModelTextAda001     = "text-ada-001"

	// CodeX
	ModeCodeDavinic  = "code-davinic-002"
	ModeCodeCushuman = "code-cushman-001"
)

const (
	ErrorGPTOverload = "ChatGPT overload"
	ErrorUnknown     = "Unknown error"
	ErrorNotFound    = "no choices/result available"
)

// Http request
type OpenAIReq struct {
	Model            string `json:"model"`
	Prompt           string `json:"prompt"`
	Temperature      int    `json:"temperature"`
	MaxTokens        int    `json:"max_tokens"`
	TopP             int    `json:"top_p"`
	FrequencyPenalty int    `json:"frequency_penalty"`
	PresencePenalty  int    `json:"presence_penalty"`
}

// Http response
type OpenAIResp struct {
	Id      string   `json:"id"`
	Object  string   `json:"object_id"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   struct {
		CompletionTokens int64 `json:"completion_tokens"`
		PromptTokens     int64 `json:"prompt_tokens"`
		TotalTokens      int64 `json:"total_tokens"`
	} `json:"usage"`
}

type Choice struct {
	Text         string `json:"text"`
	Index        int64  `json:"index"`
	Logprobs     string `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}
