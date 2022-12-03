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
package doxygen

import (
	"github.com/shanduur/go-doxygen-generator/command"
	"github.com/shanduur/go-doxygen-generator/emitter"
)

type Doxygen struct {
	Commands []command.Command
	Tag      string
}

type Option func(*Doxygen)

func New(options ...Option) *Doxygen {
	d := &Doxygen{
		Tag: DefaultTag,
	}
	for _, opt := range options {
		opt(d)
	}

	return d
}

func WithTag(tag string) func(*Doxygen) {
	return func(d *Doxygen) {
		d.Tag = tag
	}
}

func WithCommand(command command.Command) Option {
	return func(d *Doxygen) {
		d.Commands = append(d.Commands, command)
	}
}

func WithMultipleCommands(commands ...command.Command) Option {
	return func(d *Doxygen) {
		d.Commands = append(d.Commands, commands...)
	}
}

func WithCommandOnce(command command.Command) Option {
	return func(d *Doxygen) {
		if !d.hasCommand(command) {
			d.Commands = append(d.Commands, command)
		}
	}
}

func (d Doxygen) hasCommand(command command.Command) bool {
	for _, cmd := range d.Commands {
		if cmd.Command() == command.Command() {
			return true
		}
	}
	return false
}

func (d Doxygen) Generate(out emitter.Emitter) {
	out.Println("/**")
	out.Indent(1)
	for _, cmd := range d.Commands {
		cmd.Generate(d.Tag, out)
	}
	out.Indent(-1)
	out.Println("*/")
}
