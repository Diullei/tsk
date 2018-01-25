package agent

import (
	"fmt"
	"os"
	"testing"
)

var sep = string(os.PathSeparator)

func Test_Process_Run_one_cmd(t *testing.T) {
	proc := Process{}

	expected := "test01\n"
	actual := "-error-"

	ok := false
	proc.Run("python", []string{fmt.Sprintf(".%stest_assets%stest01.py", sep, sep)}, func(resul []byte) {
		if string(resul) == expected {
			ok = true
		}
	})

	if !ok {
		t.Errorf("Expecting '%s', got '%s'", expected, actual)
	}
}

func Test_Process_Run_two_cmd(t *testing.T) {
	proc := Process{}

	expected := "test01\n"
	actual := "-error-"

	ok := false
	proc.Run("python", []string{fmt.Sprintf(".%stest_assets%stest01.py", sep, sep)}, func(resul []byte) {
		if string(resul) == expected {
			ok = true
		}
	})

	if !ok {
		t.Errorf("Expecting '%s', got '%s'", expected, actual)
	}

	expected = "test02\n"
	actual = "-error-"

	ok = false
	proc.Run("python", []string{fmt.Sprintf(".%stest_assets%stest02.py", sep, sep)}, func(resul []byte) {
		if string(resul) == expected {
			ok = true
		}
	})

	if !ok {
		t.Errorf("Expecting '%s', got '%s'", expected, actual)
	}
}
