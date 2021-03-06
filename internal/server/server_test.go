package server

import (
	"bytes"
	"fmt"
	"internal/ui"
	"io"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

const testAddress = "localhost"
const testPort = "8000"

func TestMain(m *testing.M) {
	go Listen(testAddress, testPort)
	time.Sleep(100 * time.Millisecond)
	code := m.Run()
	os.Exit(code)
}

func TestConnectionAndHandlerReturn(t *testing.T) {
	conn, err := createTestConnection(t, testPort)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	reply := new(bytes.Buffer)
	if _, err := io.Copy(reply, conn); err != nil {
		t.Errorf("unexpected server error: %v", err)
	}
	lines := strings.Split(reply.String(), "\n")
	if lines[0]+"\n" != ui.ServerWelcome {
		t.Errorf("unexpected server reply: want \"%s\", got \"%s\"", ui.ServerWelcome, lines[0])
	}

}

func createTestConnection(t *testing.T, testPort string) (net.Conn, error) {
	return net.Dial("tcp", fmt.Sprintf("%s:%s", testAddress, testPort))
}
