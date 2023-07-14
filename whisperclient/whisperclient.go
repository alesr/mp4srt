package whisperclient

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

const (
	audioTranscriptionURL = "https://api.openai.com/v1/audio/transcriptions"
)

// Client is a wrapper around the WhisperAI API
type Client struct {
	apiKey string
	model  string
}

// New returns a new Client
func New(apiKey, model string) *Client {
	return &Client{
		apiKey: apiKey,
		model:  model,
	}
}

// TranscribeAudioInput is the input for the TranscribeAudio method
type TranscribeAudioInput struct {
	Name string
	Data io.Reader
}

// TranscribeAudio transcribes the audio from the given input
func (c *Client) TranscribeAudio(in TranscribeAudioInput) ([]byte, error) {
	body := bytes.Buffer{}

	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", in.Name)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(part, in.Data); err != nil {
		return nil, err
	}

	if err := writer.WriteField("model", c.model); err != nil {
		return nil, err
	}

	if err := writer.WriteField("response_format", "srt"); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, audioTranscriptionURL, &body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+c.apiKey)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
