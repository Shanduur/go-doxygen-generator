/*
This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <https://unlicense.org>
*/
package emitter

import (
	"fmt"
	"strings"

	"github.com/mitchellh/go-wordwrap"
)

type Emitter interface {
	Indent(int)
	Print(format string, args ...interface{})
	Println(format string, args ...interface{})
	Newline()
}

type SampleEmitter struct {
	sb            strings.Builder
	maxLineLength uint
	start         bool
	indent        uint
}

func NewEmitter(maxLineLength uint) *SampleEmitter {
	return &SampleEmitter{
		maxLineLength: maxLineLength,
		start:         true,
	}
}

func (e *SampleEmitter) String() string {
	return e.sb.String()
}

func (e *SampleEmitter) Bytes() []byte {
	return []byte(e.sb.String())
}

func (e *SampleEmitter) Indent(n int) {
	if int(e.indent)+n < 0 {
		panic("unexpected unbalanced indentation")
	}
	e.indent += uint(n)
}

func (e *SampleEmitter) Comment(s string) {
	if s != "" {
		limit := e.maxLineLength - uint(e.indent)
		lines := strings.Split(wordwrap.WrapString(s, limit), "\n")
		for _, line := range lines {
			e.Println("// %s", line)
		}
	}
}

func (e *SampleEmitter) Print(format string, args ...interface{}) {
	e.checkIndent()
	fmt.Fprintf(&e.sb, format, args...)
	e.start = false
}

func (e *SampleEmitter) Println(format string, args ...interface{}) {
	e.Print(format, args...)
	e.Newline()
}

func (e *SampleEmitter) Newline() {
	e.sb.WriteRune('\n')
	e.start = true
}

func (e *SampleEmitter) checkIndent() {
	if e.start {
		for i := uint(0); i < e.indent; i++ {
			e.sb.WriteRune('\t')
		}
		e.start = false
	}
}

func (e *SampleEmitter) MaxLineLength() uint {
	return e.maxLineLength
}
