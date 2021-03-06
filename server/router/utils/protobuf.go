package utils

import (
	"bytes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/proto"
)

func BindProtobuf(c echo.Context, i proto.Message) error {
	defer c.Request().Body.Close()

	buffer := new(bytes.Buffer)
	_, err := io.Copy(buffer, c.Request().Body)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(buffer.Bytes(), i)
	if err != nil {
		return err
	}

	return nil
}

func SendProtobuf(c echo.Context, status int, i proto.Message) error {
	buffer, err := proto.Marshal(i)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Blob(status, "application/octet-stream", buffer)
}

// TimeStampToTIme タイムスタンプ型をtime.timeに変換
func TimeStampToTIme(time time.Time) *timestamppb.Timestamp {
	return timestamppb.New(time)
}
