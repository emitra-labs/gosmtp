package gosmtp_test

import (
	"testing"

	"github.com/emitra-labs/gosmtp"
)

func TestSend_Success(t *testing.T) {
	gosmtp.Open()

	err := gosmtp.Send(&gosmtp.Mail{
		From:    "foo@example.com",
		To:      "bar@example.com",
		Subject: "test",
		Body:    "<p>test</p>",
	})

	if err != nil {
		t.Error("error sending email:", err)
	}

	gosmtp.Close()
}
