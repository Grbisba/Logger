package glogger

import (
	"os"
	"testing"
)

func TestAsd(t *testing.T) {
	os.Create("./infra/logs/glogger.log")
}
