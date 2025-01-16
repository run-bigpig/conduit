package utils

import (
	"bufio"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/run-bigpig/conduit/internal/types"
	"net/http"
	"strings"
)

var (
	dataPrefix       = "data: "
	done             = "[DONE]"
	dataPrefixLength = len(dataPrefix)
)

type ErrorResponse struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Fail(ctx *fiber.Ctx, err error) error {
	return ctx.Status(200).JSON(&ErrorResponse{
		Error:   -1,
		Message: err.Error(),
		Data:    nil,
	})
}

func Json(ctx *fiber.Ctx, errCode int, data types.M) error {
	return ctx.Status(errCode).JSON(data)
}

func Success(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(200).JSON(data)
}

func GuiSuccess(data interface{}) *types.GuiResponse {
	return &types.GuiResponse{
		Code: 0,
		Data: data,
		Msg:  "success",
	}
}

func GuiFail(errCode int, err error) *types.GuiResponse {
	return &types.GuiResponse{
		Code: errCode,
		Data: nil,
		Msg:  err.Error(),
	}
}

func CompletionAbort(ctx *fiber.Ctx, statusCode int) error {
	ctx.Set("Content-Type", "text/event-stream")
	return ctx.Status(statusCode).SendString("data: [DONE]\n\n")
}

func StreamResponse(ctx *fiber.Ctx, resp *http.Response) error {
	//设置header
	ctx.Set("Content-Type", "text/event-stream")
	ctx.Set("Cache-Control", "no-cache")
	ctx.Set("Connection", "keep-alive")
	ctx.Set("Transfer-Encoding", "chunked")
	//设置body
	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(bufio.ScanLines)
	var closed bool
	doneRendered := false
	//设置body stream
	ctx.Status(fiber.StatusOK).Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		defer resp.Body.Close()
		for scanner.Scan() {
			if closed {
				return
			}
			data := scanner.Text()
			if len(data) < dataPrefixLength {
				continue
			}
			if data[:dataPrefixLength] != dataPrefix && data[:dataPrefixLength] != done {
				continue
			}
			if strings.HasPrefix(data[dataPrefixLength:], done) {
				writeStringData(w, data, &closed)
				doneRendered = true
				continue
			}
			writeStringData(w, data, &closed)
		}
		if err := scanner.Err(); err != nil {
			log.Error("error reading stream: " + err.Error())
		}

		if !doneRendered {
			writeStringData(w, done, &closed)
		}
	})
	return nil
}

func writeStringData(w *bufio.Writer, str string, closed *bool) {
	str = strings.TrimPrefix(str, "data: ")
	str = strings.TrimSuffix(str, "\r")
	str = strings.TrimSuffix(str, "\n")
	_, err := fmt.Fprintf(w, "data: %s\n\n", str)
	if err != nil {
		return
	}
	err = w.Flush()
	if err != nil {
		log.Info("connection closed.")
		if err.Error() == "connection closed" {
			*closed = true
		}
		return
	}
}
