// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anthropic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/anthropics/anthropic-sdk-go/internal/apijson"
	"github.com/anthropics/anthropic-sdk-go/internal/param"
	"github.com/anthropics/anthropic-sdk-go/internal/requestconfig"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/anthropics/anthropic-sdk-go/packages/ssestream"
	"github.com/tidwall/gjson"
)

// MessageService contains methods and other services that help with interacting
// with the anthropic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	Options []option.RequestOption
	Batches *MessageBatchService
}

// NewMessageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMessageService(opts ...option.RequestOption) (r *MessageService) {
	r = &MessageService{}
	r.Options = opts
	r.Batches = NewMessageBatchService(opts...)
	return
}

// Send a structured list of input messages with text and/or image content, and the
// model will generate the next message in the conversation.
//
// The Messages API can be used for either single queries or stateless multi-turn
// conversations.
//
// Note: If you choose to set a timeout for this request, we recommend 10 minutes.
func (r *MessageService) New(ctx context.Context, body MessageNewParams, opts ...option.RequestOption) (res *Message, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send a structured list of input messages with text and/or image content, and the
// model will generate the next message in the conversation.
//
// The Messages API can be used for either single queries or stateless multi-turn
// conversations.
//
// Note: If you choose to set a timeout for this request, we recommend 10 minutes.
func (r *MessageService) NewStreaming(ctx context.Context, body MessageNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[MessageStreamEvent]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "v1/messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[MessageStreamEvent](ssestream.NewDecoder(raw), err)
}

// Count the number of tokens in a Message.
//
// The Token Count API can be used to count the number of tokens in a Message,
// including tools, images, and documents, without creating it.
func (r *MessageService) CountTokens(ctx context.Context, body MessageCountTokensParams, opts ...option.RequestOption) (res *MessageTokensCount, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/count_tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Base64PDFSourceParam struct {
	Data      param.Field[string]                   `json:"data,required" format:"byte"`
	MediaType param.Field[Base64PDFSourceMediaType] `json:"media_type,required"`
	Type      param.Field[Base64PDFSourceType]      `json:"type,required"`
}

func (r Base64PDFSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Base64PDFSourceMediaType string

const (
	Base64PDFSourceMediaTypeApplicationPDF Base64PDFSourceMediaType = "application/pdf"
)

func (r Base64PDFSourceMediaType) IsKnown() bool {
	switch r {
	case Base64PDFSourceMediaTypeApplicationPDF:
		return true
	}
	return false
}

type Base64PDFSourceType string

const (
	Base64PDFSourceTypeBase64 Base64PDFSourceType = "base64"
)

func (r Base64PDFSourceType) IsKnown() bool {
	switch r {
	case Base64PDFSourceTypeBase64:
		return true
	}
	return false
}

type CacheControlEphemeralParam struct {
	Type param.Field[CacheControlEphemeralType] `json:"type,required"`
}

func (r CacheControlEphemeralParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CacheControlEphemeralType string

const (
	CacheControlEphemeralTypeEphemeral CacheControlEphemeralType = "ephemeral"
)

func (r CacheControlEphemeralType) IsKnown() bool {
	switch r {
	case CacheControlEphemeralTypeEphemeral:
		return true
	}
	return false
}

type ContentBlock struct {
	Type  ContentBlockType `json:"type,required"`
	Text  string           `json:"text"`
	ID    string           `json:"id"`
	Name  string           `json:"name"`
	Input json.RawMessage  `json:"input,required"`
	JSON  contentBlockJSON `json:"-"`
	union ContentBlockUnion
}

// contentBlockJSON contains the JSON metadata for the struct [ContentBlock]
type contentBlockJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Input       apijson.Field
	Name        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r contentBlockJSON) RawJSON() string {
	return r.raw
}

func (r *ContentBlock) UnmarshalJSON(data []byte) (err error) {
	*r = ContentBlock{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ContentBlockUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [TextBlock], [ToolUseBlock].
func (r ContentBlock) AsUnion() ContentBlockUnion {
	return r.union
}

// Union satisfied by [TextBlock] or [ToolUseBlock].
type ContentBlockUnion interface {
	implementsContentBlock()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ContentBlockUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TextBlock{}),
			DiscriminatorValue: "text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ToolUseBlock{}),
			DiscriminatorValue: "tool_use",
		},
	)
}

type ContentBlockType string

const (
	ContentBlockTypeText    ContentBlockType = "text"
	ContentBlockTypeToolUse ContentBlockType = "tool_use"
)

func (r ContentBlockType) IsKnown() bool {
	switch r {
	case ContentBlockTypeText, ContentBlockTypeToolUse:
		return true
	}
	return false
}

type ContentBlockParam struct {
	Type         param.Field[ContentBlockParamType]      `json:"type,required"`
	ID           param.Field[string]                     `json:"id"`
	CacheControl param.Field[CacheControlEphemeralParam] `json:"cache_control"`
	Content      param.Field[interface{}]                `json:"content"`
	Input        param.Field[interface{}]                `json:"input"`
	IsError      param.Field[bool]                       `json:"is_error"`
	Name         param.Field[string]                     `json:"name"`
	Source       param.Field[interface{}]                `json:"source"`
	Text         param.Field[string]                     `json:"text"`
	ToolUseID    param.Field[string]                     `json:"tool_use_id"`
}

func (r ContentBlockParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ContentBlockParam) implementsContentBlockParamUnion() {}

// Satisfied by [TextBlockParam], [ImageBlockParam], [ToolUseBlockParam],
// [ToolResultBlockParam], [DocumentBlockParam], [ContentBlockParam].
type ContentBlockParamUnion interface {
	implementsContentBlockParamUnion()
}

type ContentBlockParamType string

const (
	ContentBlockParamTypeText       ContentBlockParamType = "text"
	ContentBlockParamTypeImage      ContentBlockParamType = "image"
	ContentBlockParamTypeToolUse    ContentBlockParamType = "tool_use"
	ContentBlockParamTypeToolResult ContentBlockParamType = "tool_result"
	ContentBlockParamTypeDocument   ContentBlockParamType = "document"
)

func (r ContentBlockParamType) IsKnown() bool {
	switch r {
	case ContentBlockParamTypeText, ContentBlockParamTypeImage, ContentBlockParamTypeToolUse, ContentBlockParamTypeToolResult, ContentBlockParamTypeDocument:
		return true
	}
	return false
}

type DocumentBlockParam struct {
	Source       param.Field[Base64PDFSourceParam]       `json:"source,required"`
	Type         param.Field[DocumentBlockParamType]     `json:"type,required"`
	CacheControl param.Field[CacheControlEphemeralParam] `json:"cache_control"`
}

func (r DocumentBlockParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DocumentBlockParam) implementsContentBlockParamUnion() {}

type DocumentBlockParamType string

const (
	DocumentBlockParamTypeDocument DocumentBlockParamType = "document"
)

func (r DocumentBlockParamType) IsKnown() bool {
	switch r {
	case DocumentBlockParamTypeDocument:
		return true
	}
	return false
}

type ImageBlockParam struct {
	Source       param.Field[ImageBlockParamSource]      `json:"source,required"`
	Type         param.Field[ImageBlockParamType]        `json:"type,required"`
	CacheControl param.Field[CacheControlEphemeralParam] `json:"cache_control"`
}

func NewImageBlockBase64(mediaType string, encodedData string) ImageBlockParam {
	return ImageBlockParam{
		Type: F(ImageBlockParamTypeImage),
		Source: F(ImageBlockParamSource{
			Type:      F(ImageBlockParamSourceTypeBase64),
			Data:      F(encodedData),
			MediaType: F(ImageBlockParamSourceMediaType(mediaType)),
		}),
	}
}

func (r ImageBlockParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ImageBlockParam) implementsContentBlockParamUnion() {}

func (r ImageBlockParam) implementsToolResultBlockParamContentUnion() {}

type ImageBlockParamSource struct {
	Data      param.Field[string]                         `json:"data,required" format:"byte"`
	MediaType param.Field[ImageBlockParamSourceMediaType] `json:"media_type,required"`
	Type      param.Field[ImageBlockParamSourceType]      `json:"type,required"`
}

func (r ImageBlockParamSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ImageBlockParamSourceMediaType string

const (
	ImageBlockParamSourceMediaTypeImageJPEG ImageBlockParamSourceMediaType = "image/jpeg"
	ImageBlockParamSourceMediaTypeImagePNG  ImageBlockParamSourceMediaType = "image/png"
	ImageBlockParamSourceMediaTypeImageGIF  ImageBlockParamSourceMediaType = "image/gif"
	ImageBlockParamSourceMediaTypeImageWebP ImageBlockParamSourceMediaType = "image/webp"
)

func (r ImageBlockParamSourceMediaType) IsKnown() bool {
	switch r {
	case ImageBlockParamSourceMediaTypeImageJPEG, ImageBlockParamSourceMediaTypeImagePNG, ImageBlockParamSourceMediaTypeImageGIF, ImageBlockParamSourceMediaTypeImageWebP:
		return true
	}
	return false
}

type ImageBlockParamSourceType string

const (
	ImageBlockParamSourceTypeBase64 ImageBlockParamSourceType = "base64"
)

func (r ImageBlockParamSourceType) IsKnown() bool {
	switch r {
	case ImageBlockParamSourceTypeBase64:
		return true
	}
	return false
}

type ImageBlockParamType string

const (
	ImageBlockParamTypeImage ImageBlockParamType = "image"
)

func (r ImageBlockParamType) IsKnown() bool {
	switch r {
	case ImageBlockParamTypeImage:
		return true
	}
	return false
}

type InputJSONDelta struct {
	PartialJSON string             `json:"partial_json,required"`
	Type        InputJSONDeltaType `json:"type,required"`
	JSON        inputJSONDeltaJSON `json:"-"`
}

// inputJSONDeltaJSON contains the JSON metadata for the struct [InputJSONDelta]
type inputJSONDeltaJSON struct {
	PartialJSON apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InputJSONDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputJSONDeltaJSON) RawJSON() string {
	return r.raw
}

func (r InputJSONDelta) implementsContentBlockDeltaEventDelta() {}

type InputJSONDeltaType string

const (
	InputJSONDeltaTypeInputJSONDelta InputJSONDeltaType = "input_json_delta"
)

func (r InputJSONDeltaType) IsKnown() bool {
	switch r {
	case InputJSONDeltaTypeInputJSONDelta:
		return true
	}
	return false
}

// Accumulate builds up the Message incrementally from a MessageStreamEvent. The Message then can be used as
// any other Message, except with the caveat that the Message.JSON field which normally can be used to inspect
// the JSON sent over the network may not be populated fully.
//
//	message := anthropic.Message{}
//	for stream.Next() {
//		event := stream.Current()
//		message.Accumulate(event)
//	}
func (a *Message) Accumulate(event MessageStreamEvent) error {
	if a == nil {
		*a = Message{}
	}

	switch event := event.AsUnion().(type) {
	case MessageStartEvent:
		*a = event.Message

	case MessageDeltaEvent:
		a.StopReason = MessageStopReason(event.Delta.StopReason)
		a.JSON.StopReason = event.Delta.JSON.StopReason
		a.StopSequence = event.Delta.StopSequence
		a.JSON.StopSequence = event.Delta.JSON.StopSequence
		a.Usage.OutputTokens = event.Usage.OutputTokens
		a.Usage.JSON.OutputTokens = event.Usage.JSON.OutputTokens

	case MessageStopEvent:

	case ContentBlockStartEvent:
		a.Content = append(a.Content, ContentBlock{})
		err := a.Content[len(a.Content)-1].UnmarshalJSON([]byte(event.ContentBlock.JSON.RawJSON()))
		if err != nil {
			return err
		}

	case ContentBlockDeltaEvent:
		if len(a.Content) == 0 {
			return fmt.Errorf("received event of type %s but there was no content block", event.Type)
		}
		switch delta := event.Delta.AsUnion().(type) {
		case TextDelta:
			cb := &a.Content[len(a.Content)-1]
			cb.Text += delta.Text
			if tb, ok := cb.union.(TextBlock); ok {
				tb.Text = cb.Text
				cb.union = tb
			}

		case InputJSONDelta:
			cb := &a.Content[len(a.Content)-1]
			if string(cb.Input) == "{}" {
				cb.Input = json.RawMessage{}
			}
			cb.Input = append(cb.Input, []byte(delta.PartialJSON)...)
			if tb, ok := cb.union.(ToolUseBlock); ok {
				tb.Input = cb.Input
				cb.union = tb
			}
		}

	case ContentBlockStopEvent:
		if len(a.Content) == 0 {
			return fmt.Errorf("received event of type %s but there was no content block", event.Type)
		}
	}

	return nil
}

// ToParam converts a Message to a MessageParam, which can be used when constructing a new
// Create
type Message struct {
	// Unique object identifier.
	//
	// The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Content generated by the model.
	//
	// This is an array of content blocks, each of which has a `type` that determines
	// its shape.
	//
	// Example:
	//
	// ```json
	// [{ "type": "text", "text": "Hi, I'm Claude." }]
	// ```
	//
	// If the request input `messages` ended with an `assistant` turn, then the
	// response `content` will continue directly from that last turn. You can use this
	// to constrain the model's output.
	//
	// For example, if the input `messages` were:
	//
	// ```json
	// [
	//
	//	{
	//	  "role": "user",
	//	  "content": "What's the Greek name for Sun? (A) Sol (B) Helios (C) Sun"
	//	},
	//	{ "role": "assistant", "content": "The best answer is (" }
	//
	// ]
	// ```
	//
	// Then the response `content` might be:
	//
	// ```json
	// [{ "type": "text", "text": "B)" }]
	// ```
	Content []ContentBlock `json:"content,required"`
	// The model that will complete your prompt.\n\nSee
	// [models](https://docs.anthropic.com/en/docs/models-overview) for additional
	// details and options.
	Model Model `json:"model,required"`
	// Conversational role of the generated message.
	//
	// This will always be `"assistant"`.
	Role MessageRole `json:"role,required"`
	// The reason that we stopped.
	//
	// This may be one the following values:
	//
	// - `"end_turn"`: the model reached a natural stopping point
	// - `"max_tokens"`: we exceeded the requested `max_tokens` or the model's maximum
	// - `"stop_sequence"`: one of your provided custom `stop_sequences` was generated
	// - `"tool_use"`: the model invoked one or more tools
	//
	// In non-streaming mode this value is always non-null. In streaming mode, it is
	// null in the `message_start` event and non-null otherwise.
	StopReason MessageStopReason `json:"stop_reason,required,nullable"`
	// Which custom stop sequence was generated, if any.
	//
	// This value will be a non-null string if one of your custom stop sequences was
	// generated.
	StopSequence string `json:"stop_sequence,required,nullable"`
	// Object type.
	//
	// For Messages, this is always `"message"`.
	Type MessageType `json:"type,required"`
	// Billing and rate-limit usage.
	//
	// Anthropic's API bills and rate-limits by token counts, as tokens represent the
	// underlying cost to our systems.
	//
	// Under the hood, the API transforms requests into a format suitable for the
	// model. The model's output then goes through a parsing stage before becoming an
	// API response. As a result, the token counts in `usage` will not match one-to-one
	// with the exact visible content of an API request or response.
	//
	// For example, `output_tokens` will be non-zero, even for an empty string response
	// from Claude.
	Usage Usage       `json:"usage,required"`
	JSON  messageJSON `json:"-"`
}

// messageJSON contains the JSON metadata for the struct [Message]
type messageJSON struct {
	ID           apijson.Field
	Content      apijson.Field
	Model        apijson.Field
	Role         apijson.Field
	StopReason   apijson.Field
	StopSequence apijson.Field
	Type         apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

// ToParam converts a Message to a MessageParam which can be used when making another network
// request. This is useful when interacting with Claude conversationally or when tool calling.
//
//	messages := []anthropic.MessageParam{
//		anthropic.NewUserMessage(anthropic.NewTextBlock("What is my first name?")),
//	}
//
//	message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
//		MaxTokens: anthropic.F(int64(1024)),
//		Messages: anthropic.F(messages),
//		Model: anthropic.F(anthropic.ModelClaude_3_5_Sonnet_20240620),
//	})
//
//	messages = append(messages, message.ToParam())
//	messages = append(messages, anthropic.NewUserMessage(
//		anthropic.NewTextBlock("My full name is John Doe"),
//	))
//
//	message, err = client.Messages.New(context.TODO(), anthropic.MessageNewParams{
//		MaxTokens: anthropic.F(int64(1024)),
//		Messages: anthropic.F(messages),
//		Model: anthropic.F(anthropic.ModelClaude_3_5_Sonnet_20240620),
//	})
func (r *Message) ToParam() MessageParam {
	content := []ContentBlockParamUnion{}

	for _, block := range r.Content {
		content = append(content, ContentBlockParam{
			Type: F(ContentBlockParamType(block.Type)),
			ID: param.Field[string]{
				Value:   block.ID,
				Present: !block.JSON.ID.IsNull(),
			},
			Text: param.Field[string]{
				Value:   block.Text,
				Present: !block.JSON.Text.IsNull(),
			},
			Name: param.Field[string]{
				Value:   block.Name,
				Present: !block.JSON.Name.IsNull(),
			},
			Input: param.Field[interface{}]{
				Value:   block.Input,
				Present: len(block.Input) > 0 && !block.JSON.Input.IsNull(),
			},
		})
	}

	message := MessageParam{
		Role:    F(MessageParamRole(r.Role)),
		Content: F(content),
	}

	return message
}

func (r *Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageJSON) RawJSON() string {
	return r.raw
}

// Conversational role of the generated message.
//
// This will always be `"assistant"`.
type MessageRole string

const (
	MessageRoleAssistant MessageRole = "assistant"
)

func (r MessageRole) IsKnown() bool {
	switch r {
	case MessageRoleAssistant:
		return true
	}
	return false
}

// The reason that we stopped.
//
// This may be one the following values:
//
// - `"end_turn"`: the model reached a natural stopping point
// - `"max_tokens"`: we exceeded the requested `max_tokens` or the model's maximum
// - `"stop_sequence"`: one of your provided custom `stop_sequences` was generated
// - `"tool_use"`: the model invoked one or more tools
//
// In non-streaming mode this value is always non-null. In streaming mode, it is
// null in the `message_start` event and non-null otherwise.
type MessageStopReason string

const (
	MessageStopReasonEndTurn      MessageStopReason = "end_turn"
	MessageStopReasonMaxTokens    MessageStopReason = "max_tokens"
	MessageStopReasonStopSequence MessageStopReason = "stop_sequence"
	MessageStopReasonToolUse      MessageStopReason = "tool_use"
)

func (r MessageStopReason) IsKnown() bool {
	switch r {
	case MessageStopReasonEndTurn, MessageStopReasonMaxTokens, MessageStopReasonStopSequence, MessageStopReasonToolUse:
		return true
	}
	return false
}

// Object type.
//
// For Messages, this is always `"message"`.
type MessageType string

const (
	MessageTypeMessage MessageType = "message"
)

func (r MessageType) IsKnown() bool {
	switch r {
	case MessageTypeMessage:
		return true
	}
	return false
}

type MessageDeltaUsage struct {
	// The cumulative number of output tokens which were used.
	OutputTokens int64                 `json:"output_tokens,required"`
	JSON         messageDeltaUsageJSON `json:"-"`
}

// messageDeltaUsageJSON contains the JSON metadata for the struct
// [MessageDeltaUsage]
type messageDeltaUsageJSON struct {
	OutputTokens apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MessageDeltaUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageDeltaUsageJSON) RawJSON() string {
	return r.raw
}

type MessageParam struct {
	Content param.Field[[]ContentBlockParamUnion] `json:"content,required"`
	Role    param.Field[MessageParamRole]         `json:"role,required"`
}

func NewUserMessage(blocks ...ContentBlockParamUnion) MessageParam {
	return MessageParam{
		Role:    F(MessageParamRoleUser),
		Content: F(blocks),
	}
}

func NewAssistantMessage(blocks ...ContentBlockParamUnion) MessageParam {
	return MessageParam{
		Role:    F(MessageParamRoleAssistant),
		Content: F(blocks),
	}
}

func (r MessageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageParamRole string

const (
	MessageParamRoleUser      MessageParamRole = "user"
	MessageParamRoleAssistant MessageParamRole = "assistant"
)

func (r MessageParamRole) IsKnown() bool {
	switch r {
	case MessageParamRoleUser, MessageParamRoleAssistant:
		return true
	}
	return false
}

type MessageTokensCount struct {
	// The total number of tokens across the provided list of messages, system prompt,
	// and tools.
	InputTokens int64                  `json:"input_tokens,required"`
	JSON        messageTokensCountJSON `json:"-"`
}

// messageTokensCountJSON contains the JSON metadata for the struct
// [MessageTokensCount]
type messageTokensCountJSON struct {
	InputTokens apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageTokensCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageTokensCountJSON) RawJSON() string {
	return r.raw
}

type MetadataParam struct {
	// An external identifier for the user who is associated with the request.
	//
	// This should be a uuid, hash value, or other opaque identifier. Anthropic may use
	// this id to help detect abuse. Do not include any identifying information such as
	// name, email address, or phone number.
	UserID param.Field[string] `json:"user_id"`
}

func (r MetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The model that will complete your prompt.\n\nSee
// [models](https://docs.anthropic.com/en/docs/models-overview) for additional
// details and options.
type Model = string

const (
	// Fast and cost-effective model
	ModelClaude3_5HaikuLatest Model = "claude-3-5-haiku-latest"
	// Fast and cost-effective model
	ModelClaude3_5Haiku20241022 Model = "claude-3-5-haiku-20241022"
	// Our most intelligent model
	ModelClaude3_5SonnetLatest Model = "claude-3-5-sonnet-latest"
	// Our most intelligent model
	ModelClaude3_5Sonnet20241022 Model = "claude-3-5-sonnet-20241022"
	// Our previous most intelligent model
	ModelClaude_3_5_Sonnet_20240620 Model = "claude-3-5-sonnet-20240620"
	// Excels at writing and complex tasks
	ModelClaude3OpusLatest Model = "claude-3-opus-latest"
	// Excels at writing and complex tasks
	ModelClaude_3_Opus_20240229 Model = "claude-3-opus-20240229"
	// Balance of speed and intelligence
	ModelClaude_3_Sonnet_20240229 Model = "claude-3-sonnet-20240229"
	// Our previous fast and cost-effective
	ModelClaude_3_Haiku_20240307 Model = "claude-3-haiku-20240307"
	ModelClaude_2_1              Model = "claude-2.1"
	ModelClaude_2_0              Model = "claude-2.0"
)

type ContentBlockDeltaEvent struct {
	Delta ContentBlockDeltaEventDelta `json:"delta,required"`
	Index int64                       `json:"index,required"`
	Type  ContentBlockDeltaEventType  `json:"type,required"`
	JSON  contentBlockDeltaEventJSON  `json:"-"`
}

// contentBlockDeltaEventJSON contains the JSON metadata for the struct
// [ContentBlockDeltaEvent]
type contentBlockDeltaEventJSON struct {
	Delta       apijson.Field
	Index       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentBlockDeltaEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentBlockDeltaEventJSON) RawJSON() string {
	return r.raw
}

func (r ContentBlockDeltaEvent) implementsMessageStreamEvent() {}

type ContentBlockDeltaEventDelta struct {
	Type        ContentBlockDeltaEventDeltaType `json:"type,required"`
	PartialJSON string                          `json:"partial_json"`
	Text        string                          `json:"text"`
	JSON        contentBlockDeltaEventDeltaJSON `json:"-"`
	union       ContentBlockDeltaEventDeltaUnion
}

// contentBlockDeltaEventDeltaJSON contains the JSON metadata for the struct
// [ContentBlockDeltaEventDelta]
type contentBlockDeltaEventDeltaJSON struct {
	Type        apijson.Field
	PartialJSON apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r contentBlockDeltaEventDeltaJSON) RawJSON() string {
	return r.raw
}

func (r *ContentBlockDeltaEventDelta) UnmarshalJSON(data []byte) (err error) {
	*r = ContentBlockDeltaEventDelta{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ContentBlockDeltaEventDeltaUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [TextDelta], [InputJSONDelta].
func (r ContentBlockDeltaEventDelta) AsUnion() ContentBlockDeltaEventDeltaUnion {
	return r.union
}

// Union satisfied by [TextDelta] or [InputJSONDelta].
type ContentBlockDeltaEventDeltaUnion interface {
	implementsContentBlockDeltaEventDelta()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ContentBlockDeltaEventDeltaUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TextDelta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputJSONDelta{}),
		},
	)
}

type ContentBlockDeltaEventDeltaType string

const (
	ContentBlockDeltaEventDeltaTypeTextDelta      ContentBlockDeltaEventDeltaType = "text_delta"
	ContentBlockDeltaEventDeltaTypeInputJSONDelta ContentBlockDeltaEventDeltaType = "input_json_delta"
)

func (r ContentBlockDeltaEventDeltaType) IsKnown() bool {
	switch r {
	case ContentBlockDeltaEventDeltaTypeTextDelta, ContentBlockDeltaEventDeltaTypeInputJSONDelta:
		return true
	}
	return false
}

type ContentBlockDeltaEventType string

const (
	ContentBlockDeltaEventTypeContentBlockDelta ContentBlockDeltaEventType = "content_block_delta"
)

func (r ContentBlockDeltaEventType) IsKnown() bool {
	switch r {
	case ContentBlockDeltaEventTypeContentBlockDelta:
		return true
	}
	return false
}

type ContentBlockStartEvent struct {
	ContentBlock ContentBlockStartEventContentBlock `json:"content_block,required"`
	Index        int64                              `json:"index,required"`
	Type         ContentBlockStartEventType         `json:"type,required"`
	JSON         contentBlockStartEventJSON         `json:"-"`
}

// contentBlockStartEventJSON contains the JSON metadata for the struct
// [ContentBlockStartEvent]
type contentBlockStartEventJSON struct {
	ContentBlock apijson.Field
	Index        apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContentBlockStartEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentBlockStartEventJSON) RawJSON() string {
	return r.raw
}

func (r ContentBlockStartEvent) implementsMessageStreamEvent() {}

type ContentBlockStartEventContentBlock struct {
	Type  ContentBlockStartEventContentBlockType `json:"type,required"`
	Text  string                                 `json:"text"`
	ID    string                                 `json:"id"`
	Name  string                                 `json:"name"`
	Input json.RawMessage                        `json:"input,required"`
	JSON  contentBlockStartEventContentBlockJSON `json:"-"`
	union ContentBlockStartEventContentBlockUnion
}

// contentBlockStartEventContentBlockJSON contains the JSON metadata for the struct
// [ContentBlockStartEventContentBlock]
type contentBlockStartEventContentBlockJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Input       apijson.Field
	Name        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r contentBlockStartEventContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r *ContentBlockStartEventContentBlock) UnmarshalJSON(data []byte) (err error) {
	*r = ContentBlockStartEventContentBlock{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ContentBlockStartEventContentBlockUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [TextBlock], [ToolUseBlock].
func (r ContentBlockStartEventContentBlock) AsUnion() ContentBlockStartEventContentBlockUnion {
	return r.union
}

// Union satisfied by [TextBlock] or [ToolUseBlock].
type ContentBlockStartEventContentBlockUnion interface {
	implementsContentBlockStartEventContentBlock()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ContentBlockStartEventContentBlockUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TextBlock{}),
			DiscriminatorValue: "text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ToolUseBlock{}),
			DiscriminatorValue: "tool_use",
		},
	)
}

type ContentBlockStartEventContentBlockType string

const (
	ContentBlockStartEventContentBlockTypeText    ContentBlockStartEventContentBlockType = "text"
	ContentBlockStartEventContentBlockTypeToolUse ContentBlockStartEventContentBlockType = "tool_use"
)

func (r ContentBlockStartEventContentBlockType) IsKnown() bool {
	switch r {
	case ContentBlockStartEventContentBlockTypeText, ContentBlockStartEventContentBlockTypeToolUse:
		return true
	}
	return false
}

type ContentBlockStartEventType string

const (
	ContentBlockStartEventTypeContentBlockStart ContentBlockStartEventType = "content_block_start"
)

func (r ContentBlockStartEventType) IsKnown() bool {
	switch r {
	case ContentBlockStartEventTypeContentBlockStart:
		return true
	}
	return false
}

type ContentBlockStopEvent struct {
	Index int64                     `json:"index,required"`
	Type  ContentBlockStopEventType `json:"type,required"`
	JSON  contentBlockStopEventJSON `json:"-"`
}

// contentBlockStopEventJSON contains the JSON metadata for the struct
// [ContentBlockStopEvent]
type contentBlockStopEventJSON struct {
	Index       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentBlockStopEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentBlockStopEventJSON) RawJSON() string {
	return r.raw
}

func (r ContentBlockStopEvent) implementsMessageStreamEvent() {}

type ContentBlockStopEventType string

const (
	ContentBlockStopEventTypeContentBlockStop ContentBlockStopEventType = "content_block_stop"
)

func (r ContentBlockStopEventType) IsKnown() bool {
	switch r {
	case ContentBlockStopEventTypeContentBlockStop:
		return true
	}
	return false
}

type MessageDeltaEvent struct {
	Delta MessageDeltaEventDelta `json:"delta,required"`
	Type  MessageDeltaEventType  `json:"type,required"`
	// Billing and rate-limit usage.
	//
	// Anthropic's API bills and rate-limits by token counts, as tokens represent the
	// underlying cost to our systems.
	//
	// Under the hood, the API transforms requests into a format suitable for the
	// model. The model's output then goes through a parsing stage before becoming an
	// API response. As a result, the token counts in `usage` will not match one-to-one
	// with the exact visible content of an API request or response.
	//
	// For example, `output_tokens` will be non-zero, even for an empty string response
	// from Claude.
	Usage MessageDeltaUsage     `json:"usage,required"`
	JSON  messageDeltaEventJSON `json:"-"`
}

// messageDeltaEventJSON contains the JSON metadata for the struct
// [MessageDeltaEvent]
type messageDeltaEventJSON struct {
	Delta       apijson.Field
	Type        apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageDeltaEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageDeltaEventJSON) RawJSON() string {
	return r.raw
}

func (r MessageDeltaEvent) implementsMessageStreamEvent() {}

type MessageDeltaEventDelta struct {
	StopReason   MessageDeltaEventDeltaStopReason `json:"stop_reason,required,nullable"`
	StopSequence string                           `json:"stop_sequence,required,nullable"`
	JSON         messageDeltaEventDeltaJSON       `json:"-"`
}

// messageDeltaEventDeltaJSON contains the JSON metadata for the struct
// [MessageDeltaEventDelta]
type messageDeltaEventDeltaJSON struct {
	StopReason   apijson.Field
	StopSequence apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MessageDeltaEventDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageDeltaEventDeltaJSON) RawJSON() string {
	return r.raw
}

type MessageDeltaEventDeltaStopReason string

const (
	MessageDeltaEventDeltaStopReasonEndTurn      MessageDeltaEventDeltaStopReason = "end_turn"
	MessageDeltaEventDeltaStopReasonMaxTokens    MessageDeltaEventDeltaStopReason = "max_tokens"
	MessageDeltaEventDeltaStopReasonStopSequence MessageDeltaEventDeltaStopReason = "stop_sequence"
	MessageDeltaEventDeltaStopReasonToolUse      MessageDeltaEventDeltaStopReason = "tool_use"
)

func (r MessageDeltaEventDeltaStopReason) IsKnown() bool {
	switch r {
	case MessageDeltaEventDeltaStopReasonEndTurn, MessageDeltaEventDeltaStopReasonMaxTokens, MessageDeltaEventDeltaStopReasonStopSequence, MessageDeltaEventDeltaStopReasonToolUse:
		return true
	}
	return false
}

type MessageDeltaEventType string

const (
	MessageDeltaEventTypeMessageDelta MessageDeltaEventType = "message_delta"
)

func (r MessageDeltaEventType) IsKnown() bool {
	switch r {
	case MessageDeltaEventTypeMessageDelta:
		return true
	}
	return false
}

type MessageStartEvent struct {
	Message Message               `json:"message,required"`
	Type    MessageStartEventType `json:"type,required"`
	JSON    messageStartEventJSON `json:"-"`
}

// messageStartEventJSON contains the JSON metadata for the struct
// [MessageStartEvent]
type messageStartEventJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageStartEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageStartEventJSON) RawJSON() string {
	return r.raw
}

func (r MessageStartEvent) implementsMessageStreamEvent() {}

type MessageStartEventType string

const (
	MessageStartEventTypeMessageStart MessageStartEventType = "message_start"
)

func (r MessageStartEventType) IsKnown() bool {
	switch r {
	case MessageStartEventTypeMessageStart:
		return true
	}
	return false
}

type MessageStopEvent struct {
	Type MessageStopEventType `json:"type,required"`
	JSON messageStopEventJSON `json:"-"`
}

// messageStopEventJSON contains the JSON metadata for the struct
// [MessageStopEvent]
type messageStopEventJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageStopEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageStopEventJSON) RawJSON() string {
	return r.raw
}

func (r MessageStopEvent) implementsMessageStreamEvent() {}

type MessageStopEventType string

const (
	MessageStopEventTypeMessageStop MessageStopEventType = "message_stop"
)

func (r MessageStopEventType) IsKnown() bool {
	switch r {
	case MessageStopEventTypeMessageStop:
		return true
	}
	return false
}

type MessageStreamEvent struct {
	Type MessageStreamEventType `json:"type,required"`
	// This field can have the runtime type of [ContentBlockStartEventContentBlock].
	ContentBlock interface{} `json:"content_block"`
	// This field can have the runtime type of [MessageDeltaEventDelta],
	// [ContentBlockDeltaEventDelta].
	Delta   interface{} `json:"delta"`
	Index   int64       `json:"index"`
	Message Message     `json:"message"`
	// Billing and rate-limit usage.
	//
	// Anthropic's API bills and rate-limits by token counts, as tokens represent the
	// underlying cost to our systems.
	//
	// Under the hood, the API transforms requests into a format suitable for the
	// model. The model's output then goes through a parsing stage before becoming an
	// API response. As a result, the token counts in `usage` will not match one-to-one
	// with the exact visible content of an API request or response.
	//
	// For example, `output_tokens` will be non-zero, even for an empty string response
	// from Claude.
	Usage MessageDeltaUsage      `json:"usage"`
	JSON  messageStreamEventJSON `json:"-"`
	union MessageStreamEventUnion
}

// messageStreamEventJSON contains the JSON metadata for the struct
// [MessageStreamEvent]
type messageStreamEventJSON struct {
	Type         apijson.Field
	ContentBlock apijson.Field
	Delta        apijson.Field
	Index        apijson.Field
	Message      apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r messageStreamEventJSON) RawJSON() string {
	return r.raw
}

func (r *MessageStreamEvent) UnmarshalJSON(data []byte) (err error) {
	*r = MessageStreamEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MessageStreamEventUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [MessageStartEvent],
// [MessageDeltaEvent], [MessageStopEvent], [ContentBlockStartEvent],
// [ContentBlockDeltaEvent], [ContentBlockStopEvent].
func (r MessageStreamEvent) AsUnion() MessageStreamEventUnion {
	return r.union
}

// Union satisfied by [MessageStartEvent], [MessageDeltaEvent], [MessageStopEvent],
// [ContentBlockStartEvent], [ContentBlockDeltaEvent] or [ContentBlockStopEvent].
type MessageStreamEventUnion interface {
	implementsMessageStreamEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageStreamEventUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageStartEvent{}),
			DiscriminatorValue: "message_start",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageDeltaEvent{}),
			DiscriminatorValue: "message_delta",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageStopEvent{}),
			DiscriminatorValue: "message_stop",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ContentBlockStartEvent{}),
			DiscriminatorValue: "content_block_start",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ContentBlockDeltaEvent{}),
			DiscriminatorValue: "content_block_delta",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ContentBlockStopEvent{}),
			DiscriminatorValue: "content_block_stop",
		},
	)
}

type MessageStreamEventType string

const (
	MessageStreamEventTypeMessageStart      MessageStreamEventType = "message_start"
	MessageStreamEventTypeMessageDelta      MessageStreamEventType = "message_delta"
	MessageStreamEventTypeMessageStop       MessageStreamEventType = "message_stop"
	MessageStreamEventTypeContentBlockStart MessageStreamEventType = "content_block_start"
	MessageStreamEventTypeContentBlockDelta MessageStreamEventType = "content_block_delta"
	MessageStreamEventTypeContentBlockStop  MessageStreamEventType = "content_block_stop"
)

func (r MessageStreamEventType) IsKnown() bool {
	switch r {
	case MessageStreamEventTypeMessageStart, MessageStreamEventTypeMessageDelta, MessageStreamEventTypeMessageStop, MessageStreamEventTypeContentBlockStart, MessageStreamEventTypeContentBlockDelta, MessageStreamEventTypeContentBlockStop:
		return true
	}
	return false
}

type TextBlock struct {
	Text string        `json:"text,required"`
	Type TextBlockType `json:"type,required"`
	JSON textBlockJSON `json:"-"`
}

// textBlockJSON contains the JSON metadata for the struct [TextBlock]
type textBlockJSON struct {
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TextBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r textBlockJSON) RawJSON() string {
	return r.raw
}

func (r TextBlock) implementsContentBlock() {}

func (r TextBlock) implementsContentBlockStartEventContentBlock() {}

type TextBlockType string

const (
	TextBlockTypeText TextBlockType = "text"
)

func (r TextBlockType) IsKnown() bool {
	switch r {
	case TextBlockTypeText:
		return true
	}
	return false
}

type TextBlockParam struct {
	Text         param.Field[string]                     `json:"text,required"`
	Type         param.Field[TextBlockParamType]         `json:"type,required"`
	CacheControl param.Field[CacheControlEphemeralParam] `json:"cache_control"`
}

func NewTextBlock(text string) TextBlockParam {
	return TextBlockParam{
		Text: F(text),
		Type: F(TextBlockParamTypeText),
	}
}

func (r TextBlockParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TextBlockParam) implementsContentBlockParamUnion() {}

func (r TextBlockParam) implementsToolResultBlockParamContentUnion() {}

type TextBlockParamType string

const (
	TextBlockParamTypeText TextBlockParamType = "text"
)

func (r TextBlockParamType) IsKnown() bool {
	switch r {
	case TextBlockParamTypeText:
		return true
	}
	return false
}

type TextDelta struct {
	Text string        `json:"text,required"`
	Type TextDeltaType `json:"type,required"`
	JSON textDeltaJSON `json:"-"`
}

// textDeltaJSON contains the JSON metadata for the struct [TextDelta]
type textDeltaJSON struct {
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TextDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r textDeltaJSON) RawJSON() string {
	return r.raw
}

func (r TextDelta) implementsContentBlockDeltaEventDelta() {}

type TextDeltaType string

const (
	TextDeltaTypeTextDelta TextDeltaType = "text_delta"
)

func (r TextDeltaType) IsKnown() bool {
	switch r {
	case TextDeltaTypeTextDelta:
		return true
	}
	return false
}

type ToolParam struct {
	// [JSON schema](https://json-schema.org/) for this tool's input.
	//
	// This defines the shape of the `input` that your tool accepts and that the model
	// will produce.
	InputSchema param.Field[interface{}] `json:"input_schema,required"`
	// Name of the tool.
	//
	// This is how the tool will be called by the model and in tool_use blocks.
	Name         param.Field[string]                     `json:"name,required"`
	CacheControl param.Field[CacheControlEphemeralParam] `json:"cache_control"`
	// Description of what this tool does.
	//
	// Tool descriptions should be as detailed as possible. The more information that
	// the model has about what the tool is and how to use it, the better it will
	// perform. You can use natural language descriptions to reinforce important
	// aspects of the tool input JSON schema.
	Description param.Field[string] `json:"description"`
}

func (r ToolParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// [JSON schema](https://json-schema.org/) for this tool's input.
//
// This defines the shape of the `input` that your tool accepts and that the model
// will produce.
type ToolInputSchemaParam struct {
	Type        param.Field[ToolInputSchemaType] `json:"type,required"`
	Properties  param.Field[interface{}]         `json:"properties"`
	ExtraFields map[string]interface{}           `json:"-,extras"`
}

func (r ToolInputSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolInputSchemaType string

const (
	ToolInputSchemaTypeObject ToolInputSchemaType = "object"
)

func (r ToolInputSchemaType) IsKnown() bool {
	switch r {
	case ToolInputSchemaTypeObject:
		return true
	}
	return false
}

// How the model should use the provided tools. The model can use a specific tool,
// any available tool, or decide by itself.
type ToolChoiceParam struct {
	Type param.Field[ToolChoiceType] `json:"type,required"`
	// Whether to disable parallel tool use.
	//
	// Defaults to `false`. If set to `true`, the model will output at most one tool
	// use.
	DisableParallelToolUse param.Field[bool] `json:"disable_parallel_tool_use"`
	// The name of the tool to use.
	Name param.Field[string] `json:"name"`
}

func (r ToolChoiceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolChoiceParam) implementsToolChoiceUnionParam() {}

// How the model should use the provided tools. The model can use a specific tool,
// any available tool, or decide by itself.
//
// Satisfied by [ToolChoiceAutoParam], [ToolChoiceAnyParam], [ToolChoiceToolParam],
// [ToolChoiceParam].
type ToolChoiceUnionParam interface {
	implementsToolChoiceUnionParam()
}

type ToolChoiceType string

const (
	ToolChoiceTypeAuto ToolChoiceType = "auto"
	ToolChoiceTypeAny  ToolChoiceType = "any"
	ToolChoiceTypeTool ToolChoiceType = "tool"
)

func (r ToolChoiceType) IsKnown() bool {
	switch r {
	case ToolChoiceTypeAuto, ToolChoiceTypeAny, ToolChoiceTypeTool:
		return true
	}
	return false
}

// The model will use any available tools.
type ToolChoiceAnyParam struct {
	Type param.Field[ToolChoiceAnyType] `json:"type,required"`
	// Whether to disable parallel tool use.
	//
	// Defaults to `false`. If set to `true`, the model will output exactly one tool
	// use.
	DisableParallelToolUse param.Field[bool] `json:"disable_parallel_tool_use"`
}

func (r ToolChoiceAnyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolChoiceAnyParam) implementsToolChoiceUnionParam() {}

type ToolChoiceAnyType string

const (
	ToolChoiceAnyTypeAny ToolChoiceAnyType = "any"
)

func (r ToolChoiceAnyType) IsKnown() bool {
	switch r {
	case ToolChoiceAnyTypeAny:
		return true
	}
	return false
}

// The model will automatically decide whether to use tools.
type ToolChoiceAutoParam struct {
	Type param.Field[ToolChoiceAutoType] `json:"type,required"`
	// Whether to disable parallel tool use.
	//
	// Defaults to `false`. If set to `true`, the model will output at most one tool
	// use.
	DisableParallelToolUse param.Field[bool] `json:"disable_parallel_tool_use"`
}

func (r ToolChoiceAutoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolChoiceAutoParam) implementsToolChoiceUnionParam() {}

type ToolChoiceAutoType string

const (
	ToolChoiceAutoTypeAuto ToolChoiceAutoType = "auto"
)

func (r ToolChoiceAutoType) IsKnown() bool {
	switch r {
	case ToolChoiceAutoTypeAuto:
		return true
	}
	return false
}

// The model will use the specified tool with `tool_choice.name`.
type ToolChoiceToolParam struct {
	// The name of the tool to use.
	Name param.Field[string]             `json:"name,required"`
	Type param.Field[ToolChoiceToolType] `json:"type,required"`
	// Whether to disable parallel tool use.
	//
	// Defaults to `false`. If set to `true`, the model will output exactly one tool
	// use.
	DisableParallelToolUse param.Field[bool] `json:"disable_parallel_tool_use"`
}

func (r ToolChoiceToolParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolChoiceToolParam) implementsToolChoiceUnionParam() {}

type ToolChoiceToolType string

const (
	ToolChoiceToolTypeTool ToolChoiceToolType = "tool"
)

func (r ToolChoiceToolType) IsKnown() bool {
	switch r {
	case ToolChoiceToolTypeTool:
		return true
	}
	return false
}

type ToolResultBlockParam struct {
	ToolUseID    param.Field[string]                             `json:"tool_use_id,required"`
	Type         param.Field[ToolResultBlockParamType]           `json:"type,required"`
	CacheControl param.Field[CacheControlEphemeralParam]         `json:"cache_control"`
	Content      param.Field[[]ToolResultBlockParamContentUnion] `json:"content"`
	IsError      param.Field[bool]                               `json:"is_error"`
}

func NewToolResultBlock(toolUseID string, content string, isError bool) ToolResultBlockParam {
	return ToolResultBlockParam{
		Type:      F(ToolResultBlockParamTypeToolResult),
		ToolUseID: F(toolUseID),
		Content:   F([]ToolResultBlockParamContentUnion{NewTextBlock(content)}),
		IsError:   F(isError),
	}
}

func (r ToolResultBlockParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolResultBlockParam) implementsContentBlockParamUnion() {}

type ToolResultBlockParamType string

const (
	ToolResultBlockParamTypeToolResult ToolResultBlockParamType = "tool_result"
)

func (r ToolResultBlockParamType) IsKnown() bool {
	switch r {
	case ToolResultBlockParamTypeToolResult:
		return true
	}
	return false
}

type ToolResultBlockParamContent struct {
	Type         param.Field[ToolResultBlockParamContentType] `json:"type,required"`
	CacheControl param.Field[CacheControlEphemeralParam]      `json:"cache_control"`
	Source       param.Field[interface{}]                     `json:"source"`
	Text         param.Field[string]                          `json:"text"`
}

func (r ToolResultBlockParamContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolResultBlockParamContent) implementsToolResultBlockParamContentUnion() {}

// Satisfied by [TextBlockParam], [ImageBlockParam], [ToolResultBlockParamContent].
type ToolResultBlockParamContentUnion interface {
	implementsToolResultBlockParamContentUnion()
}

type ToolResultBlockParamContentType string

const (
	ToolResultBlockParamContentTypeText  ToolResultBlockParamContentType = "text"
	ToolResultBlockParamContentTypeImage ToolResultBlockParamContentType = "image"
)

func (r ToolResultBlockParamContentType) IsKnown() bool {
	switch r {
	case ToolResultBlockParamContentTypeText, ToolResultBlockParamContentTypeImage:
		return true
	}
	return false
}

type ToolUseBlock struct {
	ID    string           `json:"id,required"`
	Input json.RawMessage  `json:"input,required"`
	Name  string           `json:"name,required"`
	Type  ToolUseBlockType `json:"type,required"`
	JSON  toolUseBlockJSON `json:"-"`
}

// toolUseBlockJSON contains the JSON metadata for the struct [ToolUseBlock]
type toolUseBlockJSON struct {
	ID          apijson.Field
	Input       apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolUseBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolUseBlockJSON) RawJSON() string {
	return r.raw
}

func (r ToolUseBlock) implementsContentBlock() {}

func (r ToolUseBlock) implementsContentBlockStartEventContentBlock() {}

type ToolUseBlockType string

const (
	ToolUseBlockTypeToolUse ToolUseBlockType = "tool_use"
)

func (r ToolUseBlockType) IsKnown() bool {
	switch r {
	case ToolUseBlockTypeToolUse:
		return true
	}
	return false
}

type ToolUseBlockParam struct {
	ID           param.Field[string]                     `json:"id,required"`
	Input        param.Field[interface{}]                `json:"input,required"`
	Name         param.Field[string]                     `json:"name,required"`
	Type         param.Field[ToolUseBlockParamType]      `json:"type,required"`
	CacheControl param.Field[CacheControlEphemeralParam] `json:"cache_control"`
}

func NewToolUseBlockParam(id string, name string, input interface{}) ToolUseBlockParam {
	return ToolUseBlockParam{
		ID:    F(id),
		Input: F(input),
		Name:  F(name),
		Type:  F(ToolUseBlockParamTypeToolUse),
	}
}

func (r ToolUseBlockParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolUseBlockParam) implementsContentBlockParamUnion() {}

type ToolUseBlockParamType string

const (
	ToolUseBlockParamTypeToolUse ToolUseBlockParamType = "tool_use"
)

func (r ToolUseBlockParamType) IsKnown() bool {
	switch r {
	case ToolUseBlockParamTypeToolUse:
		return true
	}
	return false
}

type Usage struct {
	// The number of input tokens used to create the cache entry.
	CacheCreationInputTokens int64 `json:"cache_creation_input_tokens,required,nullable"`
	// The number of input tokens read from the cache.
	CacheReadInputTokens int64 `json:"cache_read_input_tokens,required,nullable"`
	// The number of input tokens which were used.
	InputTokens int64 `json:"input_tokens,required"`
	// The number of output tokens which were used.
	OutputTokens int64     `json:"output_tokens,required"`
	JSON         usageJSON `json:"-"`
}

// usageJSON contains the JSON metadata for the struct [Usage]
type usageJSON struct {
	CacheCreationInputTokens apijson.Field
	CacheReadInputTokens     apijson.Field
	InputTokens              apijson.Field
	OutputTokens             apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *Usage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageJSON) RawJSON() string {
	return r.raw
}

type MessageNewParams struct {
	// The maximum number of tokens to generate before stopping.
	//
	// Note that our models may stop _before_ reaching this maximum. This parameter
	// only specifies the absolute maximum number of tokens to generate.
	//
	// Different models have different maximum values for this parameter. See
	// [models](https://docs.anthropic.com/en/docs/models-overview) for details.
	MaxTokens param.Field[int64] `json:"max_tokens,required"`
	// Input messages.
	//
	// Our models are trained to operate on alternating `user` and `assistant`
	// conversational turns. When creating a new `Message`, you specify the prior
	// conversational turns with the `messages` parameter, and the model then generates
	// the next `Message` in the conversation. Consecutive `user` or `assistant` turns
	// in your request will be combined into a single turn.
	//
	// Each input message must be an object with a `role` and `content`. You can
	// specify a single `user`-role message, or you can include multiple `user` and
	// `assistant` messages.
	//
	// If the final message uses the `assistant` role, the response content will
	// continue immediately from the content in that message. This can be used to
	// constrain part of the model's response.
	//
	// Example with a single `user` message:
	//
	// ```json
	// [{ "role": "user", "content": "Hello, Claude" }]
	// ```
	//
	// Example with multiple conversational turns:
	//
	// ```json
	// [
	//
	//	{ "role": "user", "content": "Hello there." },
	//	{ "role": "assistant", "content": "Hi, I'm Claude. How can I help you?" },
	//	{ "role": "user", "content": "Can you explain LLMs in plain English?" }
	//
	// ]
	// ```
	//
	// Example with a partially-filled response from Claude:
	//
	// ```json
	// [
	//
	//	{
	//	  "role": "user",
	//	  "content": "What's the Greek name for Sun? (A) Sol (B) Helios (C) Sun"
	//	},
	//	{ "role": "assistant", "content": "The best answer is (" }
	//
	// ]
	// ```
	//
	// Each input message `content` may be either a single `string` or an array of
	// content blocks, where each block has a specific `type`. Using a `string` for
	// `content` is shorthand for an array of one content block of type `"text"`. The
	// following input messages are equivalent:
	//
	// ```json
	// { "role": "user", "content": "Hello, Claude" }
	// ```
	//
	// ```json
	// { "role": "user", "content": [{ "type": "text", "text": "Hello, Claude" }] }
	// ```
	//
	// Starting with Claude 3 models, you can also send image content blocks:
	//
	// ```json
	//
	//	{
	//	  "role": "user",
	//	  "content": [
	//	    {
	//	      "type": "image",
	//	      "source": {
	//	        "type": "base64",
	//	        "media_type": "image/jpeg",
	//	        "data": "/9j/4AAQSkZJRg..."
	//	      }
	//	    },
	//	    { "type": "text", "text": "What is in this image?" }
	//	  ]
	//	}
	//
	// ```
	//
	// We currently support the `base64` source type for images, and the `image/jpeg`,
	// `image/png`, `image/gif`, and `image/webp` media types.
	//
	// See [examples](https://docs.anthropic.com/en/api/messages-examples#vision) for
	// more input examples.
	//
	// Note that if you want to include a
	// [system prompt](https://docs.anthropic.com/en/docs/system-prompts), you can use
	// the top-level `system` parameter — there is no `"system"` role for input
	// messages in the Messages API.
	Messages param.Field[[]MessageParam] `json:"messages,required"`
	// The model that will complete your prompt.\n\nSee
	// [models](https://docs.anthropic.com/en/docs/models-overview) for additional
	// details and options.
	Model param.Field[Model] `json:"model,required"`
	// An object describing metadata about the request.
	Metadata param.Field[MetadataParam] `json:"metadata"`
	// Custom text sequences that will cause the model to stop generating.
	//
	// Our models will normally stop when they have naturally completed their turn,
	// which will result in a response `stop_reason` of `"end_turn"`.
	//
	// If you want the model to stop generating when it encounters custom strings of
	// text, you can use the `stop_sequences` parameter. If the model encounters one of
	// the custom sequences, the response `stop_reason` value will be `"stop_sequence"`
	// and the response `stop_sequence` value will contain the matched stop sequence.
	StopSequences param.Field[[]string] `json:"stop_sequences"`
	// System prompt.
	//
	// A system prompt is a way of providing context and instructions to Claude, such
	// as specifying a particular goal or role. See our
	// [guide to system prompts](https://docs.anthropic.com/en/docs/system-prompts).
	System param.Field[[]TextBlockParam] `json:"system"`
	// Amount of randomness injected into the response.
	//
	// Defaults to `1.0`. Ranges from `0.0` to `1.0`. Use `temperature` closer to `0.0`
	// for analytical / multiple choice, and closer to `1.0` for creative and
	// generative tasks.
	//
	// Note that even with `temperature` of `0.0`, the results will not be fully
	// deterministic.
	Temperature param.Field[float64] `json:"temperature"`
	// How the model should use the provided tools. The model can use a specific tool,
	// any available tool, or decide by itself.
	ToolChoice param.Field[ToolChoiceUnionParam] `json:"tool_choice"`
	// Definitions of tools that the model may use.
	//
	// If you include `tools` in your API request, the model may return `tool_use`
	// content blocks that represent the model's use of those tools. You can then run
	// those tools using the tool input generated by the model and then optionally
	// return results back to the model using `tool_result` content blocks.
	//
	// Each tool definition includes:
	//
	//   - `name`: Name of the tool.
	//   - `description`: Optional, but strongly-recommended description of the tool.
	//   - `input_schema`: [JSON schema](https://json-schema.org/) for the tool `input`
	//     shape that the model will produce in `tool_use` output content blocks.
	//
	// For example, if you defined `tools` as:
	//
	// ```json
	// [
	//
	//	{
	//	  "name": "get_stock_price",
	//	  "description": "Get the current stock price for a given ticker symbol.",
	//	  "input_schema": {
	//	    "type": "object",
	//	    "properties": {
	//	      "ticker": {
	//	        "type": "string",
	//	        "description": "The stock ticker symbol, e.g. AAPL for Apple Inc."
	//	      }
	//	    },
	//	    "required": ["ticker"]
	//	  }
	//	}
	//
	// ]
	// ```
	//
	// And then asked the model "What's the S&P 500 at today?", the model might produce
	// `tool_use` content blocks in the response like this:
	//
	// ```json
	// [
	//
	//	{
	//	  "type": "tool_use",
	//	  "id": "toolu_01D7FLrfh4GYq7yT1ULFeyMV",
	//	  "name": "get_stock_price",
	//	  "input": { "ticker": "^GSPC" }
	//	}
	//
	// ]
	// ```
	//
	// You might then run your `get_stock_price` tool with `{"ticker": "^GSPC"}` as an
	// input, and return the following back to the model in a subsequent `user`
	// message:
	//
	// ```json
	// [
	//
	//	{
	//	  "type": "tool_result",
	//	  "tool_use_id": "toolu_01D7FLrfh4GYq7yT1ULFeyMV",
	//	  "content": "259.75 USD"
	//	}
	//
	// ]
	// ```
	//
	// Tools can be used for workflows that include running client-side tools and
	// functions, or more generally whenever you want the model to produce a particular
	// JSON structure of output.
	//
	// See our [guide](https://docs.anthropic.com/en/docs/tool-use) for more details.
	Tools param.Field[[]ToolParam] `json:"tools"`
	// Only sample from the top K options for each subsequent token.
	//
	// Used to remove "long tail" low probability responses.
	// [Learn more technical details here](https://towardsdatascience.com/how-to-sample-from-language-models-682bceb97277).
	//
	// Recommended for advanced use cases only. You usually only need to use
	// `temperature`.
	TopK param.Field[int64] `json:"top_k"`
	// Use nucleus sampling.
	//
	// In nucleus sampling, we compute the cumulative distribution over all the options
	// for each subsequent token in decreasing probability order and cut it off once it
	// reaches a particular probability specified by `top_p`. You should either alter
	// `temperature` or `top_p`, but not both.
	//
	// Recommended for advanced use cases only. You usually only need to use
	// `temperature`.
	TopP param.Field[float64] `json:"top_p"`
}

func (r MessageNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageCountTokensParams struct {
	// Input messages.
	//
	// Our models are trained to operate on alternating `user` and `assistant`
	// conversational turns. When creating a new `Message`, you specify the prior
	// conversational turns with the `messages` parameter, and the model then generates
	// the next `Message` in the conversation. Consecutive `user` or `assistant` turns
	// in your request will be combined into a single turn.
	//
	// Each input message must be an object with a `role` and `content`. You can
	// specify a single `user`-role message, or you can include multiple `user` and
	// `assistant` messages.
	//
	// If the final message uses the `assistant` role, the response content will
	// continue immediately from the content in that message. This can be used to
	// constrain part of the model's response.
	//
	// Example with a single `user` message:
	//
	// ```json
	// [{ "role": "user", "content": "Hello, Claude" }]
	// ```
	//
	// Example with multiple conversational turns:
	//
	// ```json
	// [
	//
	//	{ "role": "user", "content": "Hello there." },
	//	{ "role": "assistant", "content": "Hi, I'm Claude. How can I help you?" },
	//	{ "role": "user", "content": "Can you explain LLMs in plain English?" }
	//
	// ]
	// ```
	//
	// Example with a partially-filled response from Claude:
	//
	// ```json
	// [
	//
	//	{
	//	  "role": "user",
	//	  "content": "What's the Greek name for Sun? (A) Sol (B) Helios (C) Sun"
	//	},
	//	{ "role": "assistant", "content": "The best answer is (" }
	//
	// ]
	// ```
	//
	// Each input message `content` may be either a single `string` or an array of
	// content blocks, where each block has a specific `type`. Using a `string` for
	// `content` is shorthand for an array of one content block of type `"text"`. The
	// following input messages are equivalent:
	//
	// ```json
	// { "role": "user", "content": "Hello, Claude" }
	// ```
	//
	// ```json
	// { "role": "user", "content": [{ "type": "text", "text": "Hello, Claude" }] }
	// ```
	//
	// Starting with Claude 3 models, you can also send image content blocks:
	//
	// ```json
	//
	//	{
	//	  "role": "user",
	//	  "content": [
	//	    {
	//	      "type": "image",
	//	      "source": {
	//	        "type": "base64",
	//	        "media_type": "image/jpeg",
	//	        "data": "/9j/4AAQSkZJRg..."
	//	      }
	//	    },
	//	    { "type": "text", "text": "What is in this image?" }
	//	  ]
	//	}
	//
	// ```
	//
	// We currently support the `base64` source type for images, and the `image/jpeg`,
	// `image/png`, `image/gif`, and `image/webp` media types.
	//
	// See [examples](https://docs.anthropic.com/en/api/messages-examples#vision) for
	// more input examples.
	//
	// Note that if you want to include a
	// [system prompt](https://docs.anthropic.com/en/docs/system-prompts), you can use
	// the top-level `system` parameter — there is no `"system"` role for input
	// messages in the Messages API.
	Messages param.Field[[]MessageParam] `json:"messages,required"`
	// The model that will complete your prompt.\n\nSee
	// [models](https://docs.anthropic.com/en/docs/models-overview) for additional
	// details and options.
	Model param.Field[Model] `json:"model,required"`
	// System prompt.
	//
	// A system prompt is a way of providing context and instructions to Claude, such
	// as specifying a particular goal or role. See our
	// [guide to system prompts](https://docs.anthropic.com/en/docs/system-prompts).
	System param.Field[MessageCountTokensParamsSystemUnion] `json:"system"`
	// How the model should use the provided tools. The model can use a specific tool,
	// any available tool, or decide by itself.
	ToolChoice param.Field[ToolChoiceUnionParam] `json:"tool_choice"`
	// Definitions of tools that the model may use.
	//
	// If you include `tools` in your API request, the model may return `tool_use`
	// content blocks that represent the model's use of those tools. You can then run
	// those tools using the tool input generated by the model and then optionally
	// return results back to the model using `tool_result` content blocks.
	//
	// Each tool definition includes:
	//
	//   - `name`: Name of the tool.
	//   - `description`: Optional, but strongly-recommended description of the tool.
	//   - `input_schema`: [JSON schema](https://json-schema.org/) for the tool `input`
	//     shape that the model will produce in `tool_use` output content blocks.
	//
	// For example, if you defined `tools` as:
	//
	// ```json
	// [
	//
	//	{
	//	  "name": "get_stock_price",
	//	  "description": "Get the current stock price for a given ticker symbol.",
	//	  "input_schema": {
	//	    "type": "object",
	//	    "properties": {
	//	      "ticker": {
	//	        "type": "string",
	//	        "description": "The stock ticker symbol, e.g. AAPL for Apple Inc."
	//	      }
	//	    },
	//	    "required": ["ticker"]
	//	  }
	//	}
	//
	// ]
	// ```
	//
	// And then asked the model "What's the S&P 500 at today?", the model might produce
	// `tool_use` content blocks in the response like this:
	//
	// ```json
	// [
	//
	//	{
	//	  "type": "tool_use",
	//	  "id": "toolu_01D7FLrfh4GYq7yT1ULFeyMV",
	//	  "name": "get_stock_price",
	//	  "input": { "ticker": "^GSPC" }
	//	}
	//
	// ]
	// ```
	//
	// You might then run your `get_stock_price` tool with `{"ticker": "^GSPC"}` as an
	// input, and return the following back to the model in a subsequent `user`
	// message:
	//
	// ```json
	// [
	//
	//	{
	//	  "type": "tool_result",
	//	  "tool_use_id": "toolu_01D7FLrfh4GYq7yT1ULFeyMV",
	//	  "content": "259.75 USD"
	//	}
	//
	// ]
	// ```
	//
	// Tools can be used for workflows that include running client-side tools and
	// functions, or more generally whenever you want the model to produce a particular
	// JSON structure of output.
	//
	// See our [guide](https://docs.anthropic.com/en/docs/tool-use) for more details.
	Tools param.Field[[]ToolParam] `json:"tools"`
}

func (r MessageCountTokensParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// System prompt.
//
// A system prompt is a way of providing context and instructions to Claude, such
// as specifying a particular goal or role. See our
// [guide to system prompts](https://docs.anthropic.com/en/docs/system-prompts).
//
// Satisfied by [shared.UnionString], [MessageCountTokensParamsSystemArray].
type MessageCountTokensParamsSystemUnion interface {
	ImplementsMessageCountTokensParamsSystemUnion()
}

type MessageCountTokensParamsSystemArray []TextBlockParam

func (r MessageCountTokensParamsSystemArray) ImplementsMessageCountTokensParamsSystemUnion() {}
