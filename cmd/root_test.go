package cmd

import (
	"bytes"
	"os"
	"testing"
)

func captureOutput(f func()) (string, error) {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}
	defer func() {
		os.Stdout = old
		r.Close()
		w.Close()
	}()

	os.Stdout = w
	f()

	w.Close()
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func TestExecute(t *testing.T) {
	output, err := captureOutput(Execute)
	if err != nil {
		t.Fatalf("failed to capture output: %v", err)
	}
	if len(output) == 0 {
		t.Errorf("expected some output, but got none")
	}
}

func TestRootCmdWithFlags(t *testing.T) {
	rootCmd.SetArgs([]string{"--length", "20", "--exclude-lowers"})

	output, err := captureOutput(func() {
		err := rootCmd.Execute()
		if err != nil {
			t.Fatalf("error executing rootCmd: %v", err)
		}
	})

	if err != nil {
		t.Fatalf("failed to capture output: %v", err)
	}

	if !bytes.Contains([]byte(output), []byte("Password:")) {
		t.Errorf("expected output to contain password, but got %q", output)
	}
}
