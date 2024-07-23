package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

// todo - эти тесты почему-то не запускаются!

var testOk = `
1
2
3
4
5`
var testOkResult = `
1
2
3
4
5
`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	//out := bufio.NewWriter(os.Stdout)
	out := new(bytes.Buffer)

	err := uniq(in, out)
	// если ф-я вернула ошибку, то тест упал
	if err != nil {
		t.Errorf("Test Ok failed - error")
	}

	result := out.String()
	if result != testOkResult {
		t.Errorf("Test Ok failed - result not match\n result: %s\n testOkResult: %s", result, testOkResult)
	}
}

var testErr = `1
2
3
2`

func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testErr))
	out := new(bytes.Buffer)

	err := uniq(in, out)
	// ругаемся, если программа не вернула ошибку на плохие данные
	if err == nil {
		t.Errorf("Test Ok failed - error: %v", err)
	}
}
