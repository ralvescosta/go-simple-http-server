package controllers

import (
	"encoding/json"
	"net/http"
)

type (
	ResponseBuilder interface {
		Ok() ResponseBuilder
		Created() ResponseBuilder
		Body(body any) ResponseBuilder
		Headers(headers map[string]string) ResponseBuilder
		UnformattedBody() ResponseBuilder
		InvalidBody() ResponseBuilder
		InternalError() ResponseBuilder
		ErrMessage(msg string) ResponseBuilder
		ErrDetails(details any) ResponseBuilder
		Build()
	}

	responseBuilder struct {
		writer     http.ResponseWriter
		statusCode int
		body       any
		headers    map[string]string
		errMessage string
		errDetails any
	}

	// HTTP Error Response
	HTTPError struct {
		StatusCode int    `json:"status_code" example:"400"`
		Message    string `json:"message"     example:"bad request"`
		Details    any
	}
)

func NewResponseBuilder(w http.ResponseWriter) ResponseBuilder {
	return &responseBuilder{writer: w}
}

func (resp *responseBuilder) Ok() ResponseBuilder {
	resp.statusCode = http.StatusOK
	return resp
}

func (resp *responseBuilder) Created() ResponseBuilder {
	resp.statusCode = http.StatusCreated
	return resp
}

func (resp *responseBuilder) Body(body any) ResponseBuilder {
	resp.body = body
	return resp
}

func (resp *responseBuilder) Headers(headers map[string]string) ResponseBuilder {
	resp.headers = headers
	return resp
}

func (resp *responseBuilder) UnformattedBody() ResponseBuilder {
	resp.statusCode = http.StatusBadRequest
	resp.errMessage = "unformatted body"
	return resp
}

func (resp *responseBuilder) InvalidBody() ResponseBuilder {
	resp.statusCode = http.StatusBadRequest
	resp.errMessage = "invalid body"
	return resp
}

func (resp *responseBuilder) InternalError() ResponseBuilder {
	resp.statusCode = http.StatusInternalServerError
	resp.errMessage = "internal error"
	return resp
}

func (resp *responseBuilder) ErrMessage(message string) ResponseBuilder {
	resp.errMessage = message
	return resp
}

func (resp *responseBuilder) ErrDetails(details any) ResponseBuilder {
	resp.errDetails = details
	return resp
}

func (resp *responseBuilder) Build() {
	header := resp.writer.Header()

	if resp.headers != nil {
		for k, v := range resp.headers {
			header.Add(k, v)
		}
	}

	header.Add("Content-Type", "application/json; charset=utf-8")
	resp.writer.WriteHeader(resp.statusCode)

	if resp.statusCode >= 400 {
		resp.writer.Write(NewHTTPError(resp.statusCode, resp.errMessage, resp.errDetails).ToBuffer())
		return
	}

	bytes, _ := json.Marshal(resp.body)
	resp.writer.Write(bytes)
}

func NewHTTPError(status int, msg string, details any) *HTTPError {
	return &HTTPError{StatusCode: status, Message: msg, Details: details}
}

func (h *HTTPError) ToBuffer() []byte {
	b, _ := json.Marshal(h)
	return b
}
