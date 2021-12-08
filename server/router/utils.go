package router

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/proto"
)

func uuidsToStrings(IDs []uuid.UUID) []string {
	var res []string

	for _, ID := range IDs {
		res = append(res, ID.String())
	}

	return res
}

func bindProtobuf(c echo.Context, i proto.Message) error {
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

func sendProtobuf(c echo.Context, status int, i proto.Message) error {
	buffer, err := proto.Marshal(i)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Blob(status, "application/octet-stream", buffer)
}
