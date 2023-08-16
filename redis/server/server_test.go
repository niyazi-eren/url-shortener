package server_test

import (
	"net"
	"reflect"
	"testing"
	"time"
	"url-shortener/redis/resp"
	"url-shortener/redis/server"
)

var testPort = ":8888"

func TestServer_Set(t *testing.T) {
	tests := []struct {
		cmd  string
		want string
	}{
		{"SET name JOHN", "OK"},
		{"SET name JANE", "JOHN"},
	}

	for _, tt := range tests {
		got, err := send(tt.cmd)
		if err != nil {
			t.Errorf(err.Error())
		}

		if got != tt.want {
			t.Errorf("for command %q, got %q, want %q", tt.cmd, got, tt.want)
		}
	}
}

func TestServer_Get(t *testing.T) {
	_, err := send("SET name JOHN")
	if err != nil {
		t.Errorf(err.Error())
	}

	testsGet := []struct {
		cmd  string
		want any
	}{
		{"GET name", "JOHN"},
		{"GET doesntexist", nil},
	}

	for _, tt := range testsGet {
		got, err := send(tt.cmd)
		if err != nil {
			t.Errorf(err.Error())
		} else {
			if got != tt.want {
				t.Errorf("got %q, want %q", err, tt.want)
			}
		}
	}
}

func TestServer_SetExpire(t *testing.T) {
	_, err := send("SET name JOHN PX 10")
	_, err = send("SET test TEST PX 3500")
	if err != nil {
		t.Errorf(err.Error())
	}

	testsGet := []struct {
		cmd  string
		want any
	}{
		{"GET name", nil},
		{"GET test", "TEST"},
	}

	time.Sleep(time.Millisecond * 10)

	for _, tt := range testsGet {
		got, err := send(tt.cmd)
		if err != nil {
			t.Errorf(err.Error())
		} else {
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		}
	}
}

func TestServer_Exists_Del(t *testing.T) {
	_, err := send("SET name JOHN")
	if err != nil {
		t.Errorf(err.Error())
	}

	tests := []struct {
		cmd  string
		want any
	}{
		{"EXISTS fail", 0},
		{"EXISTS fail name name", 2},
		{"DEL name name", 1},
		{"DEL fail", 0},
		{"EXISTS name", 0},
	}

	for _, tt := range tests {
		got, err := send(tt.cmd)
		if err != nil {
			t.Errorf(err.Error())
		} else {
			if got != tt.want {
				t.Errorf("got %q, want %q", err, tt.want)
			}
		}
	}
}

func TestServer_Incr_Decr(t *testing.T) {
	_, err := send("SET one 1")
	_, err = send("SET two two")
	_, err = send("SET t1 123")
	_, err = send("SET t2 123000000000000000000000000")
	if err != nil {
		t.Errorf(err.Error())
	}

	tests := []struct {
		cmd  string
		want any
	}{
		{"INCR one", 2},
		{"INCR one", 3},
		{"INCR zero", 1},
		{"INCR zero", 2},
		{"INCR two", resp.IncrErr},
		{"DECR t1", 122},
		{"DECR t0", -1},
		{"DECR t0", -2},
		{"INCR t0", -1},
		{"DECR t2", resp.IncrErr},
	}

	for _, tt := range tests {
		got, err := send(tt.cmd)
		if err != nil {
			switch tt.want.(type) {
			// case when expecting an error
			case error:
				wantErr := tt.want.(error)
				if err.Error() != wantErr.Error() {
					t.Errorf("got %q, want %q", err, tt.want)
				}
			}
		} else {
			if got != tt.want {
				t.Errorf("got %q, want %q", err, tt.want)
			}
		}
	}
}

func TestServer_RPush_LPush(t *testing.T) {
	send("SET one 1")
	tests := []struct {
		cmd  string
		want any
	}{
		{"RPUSH mylist lol", 1},
		{"RPUSH mylist 2", 2},
		{"RPUSH one 3", resp.NotAListErr},
		{"LPUSH mylist 3", 3},
		{"LPUSH mylist 4", 4},
	}

	for _, tt := range tests {
		got, err := send(tt.cmd)
		if err != nil {
			switch tt.want.(type) {
			// case when expecting an error
			case error:
				wantErr := tt.want.(error)
				if err.Error() != wantErr.Error() {
					t.Errorf("got %q, want %q", err, tt.want)
				}
			}
		} else {
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %q, want %q", err, tt.want)
			}
		}
	}
}

func TestServer_SaveAndLoad(t *testing.T) {
	tests := []struct {
		cmd  string
		want any
	}{
		// set db and save it to file
		{"SET name JOHN", "OK"},
		{"SET name2 JANE", "OK"},
		{"SAVE", "OK"},
		// delete local db and verify that it is deleted
		{"DEL name", 1},
		{"DEL name2", 1},
		{"EXISTS name", 0},
		{"EXISTS name2", 0},
		// load db from file and verify that data is present
		{"LOAD", "OK"},
		{"EXISTS name", 1},
		{"EXISTS name2", 1},
	}

	for _, tt := range tests {
		got, err := send(tt.cmd)
		if err != nil {
			t.Errorf(err.Error())
		}

		if got != tt.want {
			t.Errorf("for command %q, got %q, want %q", tt.cmd, got, tt.want)
		}
	}
}

// executed before every test
func init() {
	s := server.NewServer("8888")
	go s.Run()
}

func send(cmd string) (any, error) {
	respCmd, err := resp.Encode(cmd)
	if err != nil {
		return "", err
	}
	conn, err := net.Dial("tcp", "localhost"+testPort)
	defer conn.Close()
	if err != nil {
		return "", err
	}
	_, err = conn.Write([]byte(respCmd))
	if err != nil {
		return "", err
	}
	buf := make([]byte, 0124)
	_, err = conn.Read(buf)
	if err != nil {
		return "", err
	}

	response, err := resp.Decode(buf)
	if err != nil {
		return nil, err
	}
	return response, nil
}
