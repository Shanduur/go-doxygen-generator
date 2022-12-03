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
package command

import (
	"fmt"

	"github.com/shanduur/go-doxygen-generator/emitter"
)

type Command interface {
	Generate(tag string, out emitter.Emitter)
	Command() string
}

// A is structure for `a` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmda
type A struct {
	Word string
}

func (cmd A) Command() string { return `A` }
func (cmd A) Generate(tag string, out emitter.Emitter) {
	out.Print("%sa %s", tag, word(cmd.Word))
}

// Addindex is structure for `addindex` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdaddindex
type Addindex struct {
	Text string
}

func (cmd Addindex) Command() string { return `Addindex` }
func (cmd Addindex) Generate(tag string, out emitter.Emitter) {
	out.Println("%saddindex %s", tag, cmd.Text)
}

// Addtogroup is structure for `addtogroup` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdaddtogroup
type Addtogroup struct {
	Name  string
	Title string
}

func (cmd Addtogroup) Command() string { return `Addtogroup` }
func (cmd Addtogroup) Generate(tag string, out emitter.Emitter) {
	out.Println("%saddtogroup %s %s", tag, word(cmd.Name), cmd.Title)
}

// Anchor is structure for `anchor` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdanchor
type Anchor struct {
	Name string
	Text string
}

func (cmd Anchor) Command() string { return `Anchor` }
func (cmd Anchor) Generate(tag string, out emitter.Emitter) {
	out.Println("%sanchor %s%s", tag, word(cmd.Name), optional(cmd.Text))
}

// Arg is structure for `arg` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdarg
type Arg struct {
	ItemDescription string
}

func (cmd Arg) Command() string { return `Arg` }
func (cmd Arg) Generate(tag string, out emitter.Emitter) {
	out.Println("%sarg %s", tag, cmd.ItemDescription)
}

// Attention is structure for `attention` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdattention
type Attention struct {
	Text string
}

func (cmd Attention) Command() string { return `Attention` }
func (cmd Attention) Generate(tag string, out emitter.Emitter) {
	out.Println("%sattention %s", tag, cmd.Text)
}

// Author is structure for `author` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdauthor
type Author struct {
	ListOfAuthors []string
}

func (cmd Author) Command() string { return `Author` }
func (cmd Author) Generate(tag string, out emitter.Emitter) {
	for _, author := range cmd.ListOfAuthors {
		out.Println("%sauthor %s", author)
	}
}

// Authors is structure for `authors` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdauthors
type Authors struct {
	ListOfAuthors []string
}

func (cmd Authors) Command() string { return `Authors` }
func (cmd Authors) Generate(tag string, out emitter.Emitter) {
	for _, author := range cmd.ListOfAuthors {
		out.Println("%sauthor %s", tag, author)
	}
}

// B is structure for `b` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdb
type B struct {
	Word string
}

func (cmd B) Command() string { return `B` }
func (cmd B) Generate(tag string, out emitter.Emitter) {
	out.Print("%sarg %s", tag, word(cmd.Word))
}

// MultiB is structure for multi word `b` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdb
type MultiB struct {
	Text string
}

func (cmd MultiB) Command() string { return `MultiB` }
func (cmd MultiB) Generate(tag string, out emitter.Emitter) {
	out.Print("<b>%s</b>", cmd.Text)
}

// Brief is structure for `brief` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdbrief
type Brief struct {
	BriefDescription string
}

func (cmd Brief) Command() string { return `Brief` }
func (cmd Brief) Generate(tag string, out emitter.Emitter) {
	out.Println("%sbrief %s", tag, cmd.BriefDescription)
}

// Bug is structure for `bug` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdbug
type Bug struct {
	Description string
}

func (cmd Bug) Command() string { return `Bug` }
func (cmd Bug) Generate(tag string, out emitter.Emitter) {
	out.Println("%sbug %s", tag, cmd.Description)
}

// C is structure for `c` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdc
type C struct {
	Word string
}

func (cmd C) Command() string { return `C` }
func (cmd C) Generate(tag string, out emitter.Emitter) {
	out.Print("%sc %s", tag, word(cmd.Word))
}

// Callergraph is structure for `callergraph` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdcallergraph
type Callergraph struct{}

func (cmd Callergraph) Command() string { return `Callergraph` }
func (cmd Callergraph) Generate(tag string, out emitter.Emitter) {
	out.Println("%scallergraph", tag)
}

// Callgraph is structure for `callgraph` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdcallgraph
type Callgraph struct{}

func (cmd Callgraph) Command() string { return `Callgraph` }
func (cmd Callgraph) Generate(tag string, out emitter.Emitter) {
	out.Println("%scallgraph", tag)
}

// Category is structure for `category` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdcategory
type Category struct {
	Name       string
	HeaderFile string
	HeaderName string
}

func (cmd Category) Command() string { return `Category` }
func (cmd Category) Generate(tag string, out emitter.Emitter) {
	out.Println("%scategory %s%s%s", tag,
		cmd.Name,
		optional(word(cmd.HeaderFile)),
		optional(word(cmd.HeaderName)))
}

// Cite is structure for `cite` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdcite
type Cite struct {
	Label string
}

func (cmd Cite) Command() string { return `Cite` }
func (cmd Cite) Generate(tag string, out emitter.Emitter) {
	out.Println("%scite %s", tag, word(cmd.Label))
}

// Class is structure for `class` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdclass
type Class struct {
	Name       string
	HeaderFile string
	HeaderName string
}

func (cmd Class) Command() string { return `Class` }
func (cmd Class) Generate(tag string, out emitter.Emitter) {
	out.Println("%sclass %s%s%s", tag,
		word(cmd.Name),
		optional(word(cmd.HeaderFile)),
		optional(word(cmd.HeaderFile)))
}

// Code is structure for `code` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdcode
type Code struct {
	Word      string
	CodeBlock string
}

func (cmd Code) Command() string { return `Code` }
func (cmd Code) Generate(tag string, out emitter.Emitter) {
	out.Println("%scode", tag,
		optionalf(" {%s}", word(cmd.Word)))
	out.Println("%s", cmd.CodeBlock)
	Endcode{}.Generate(tag, out)
}

// Concept is structure for `concept` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdconcept
type Concept struct {
	Name string
}

func (cmd Concept) Command() string { return `Concept` }
func (cmd Concept) Generate(tag string, out emitter.Emitter) {
	out.Println("%sconcept %s", tag, word(cmd.Name))
}

// Cond is structure for `cond` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdcond
type Cond struct {
	SectionLabel string
}

func (cmd Cond) Command() string { return `Cond` }
func (cmd Cond) Generate(tag string, out emitter.Emitter) {
	out.Println("%scond%s", tag, optional(cmd.SectionLabel))
}

// Copybrief is structure for `copybrief` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdcopybrief
type Copybrief struct {
	LinkObject string
}

func (cmd Copybrief) Command() string { return `Copybrief` }
func (cmd Copybrief) Generate(tag string, out emitter.Emitter) {
	out.Println("%scopybrief %s", tag, word(cmd.LinkObject))
}

// Copydetails is structure for `copydetails` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdcopydetails
type Copydetails struct {
	LinkObject string
}

func (cmd Copydetails) Command() string { return `Copydetails` }
func (cmd Copydetails) Generate(tag string, out emitter.Emitter) {
	out.Println("%scopydetails %s", tag, word(cmd.LinkObject))
}

// Copydoc is structure for `copydoc` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdcopydoc
type Copydoc struct {
	LinkObject string
}

func (cmd Copydoc) Command() string { return `Copydoc` }
func (cmd Copydoc) Generate(tag string, out emitter.Emitter) {
	out.Println("%scopydoc %s", tag, word(cmd.LinkObject))
}

// Copyright is structure for `copyright` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdcopyright
type Copyright struct {
	Description string
}

func (cmd Copyright) Command() string { return `Copyright` }
func (cmd Copyright) Generate(tag string, out emitter.Emitter) {
	out.Println("%scopyright %s", tag, word(cmd.Description))
}

// Date is structure for `date` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmddate
type Date struct {
	Description string
}

func (cmd Date) Command() string { return `Date` }
func (cmd Date) Generate(tag string, out emitter.Emitter) {
	out.Println("%sdate %s", tag, cmd.Description)
}

// Def is structure for `def` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmddef
type Def struct {
	Name string
}

func (cmd Def) Command() string { return `Def` }
func (cmd Def) Generate(tag string, out emitter.Emitter) {
	out.Println("%sdef %s", tag, word(cmd.Name))
}

// Defgroup is structure for `defgroup` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmddefgroup
type Defgroup struct {
	Name       string
	GroupTitle string
}

func (cmd Defgroup) Command() string { return `Defgroup` }
func (cmd Defgroup) Generate(tag string, out emitter.Emitter) {
	out.Println("%sdefgroup %s %s", tag, word(cmd.Name), cmd.GroupTitle)
}

// Deprecated is structure for `deprecated` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmddeprecated
type Deprecated struct {
	Description string
}

func (cmd Deprecated) Command() string { return `Deprecated` }
func (cmd Deprecated) Generate(tag string, out emitter.Emitter) {
	out.Println("%sdeprecated %s", tag, cmd.Description)
}

// Details is structure for `details` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmddetails
type Details struct {
	DetailedDescription string
}

func (cmd Details) Command() string { return `Details` }
func (cmd Details) Generate(tag string, out emitter.Emitter) {
	out.Println("%sdetails %s", tag, cmd.DetailedDescription)
}

// Diafile is structure for `diafile` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmddiafile
type Diafile struct {
	File           string
	Caption        string
	SizeIndication string
	Size           string
}

func (cmd Diafile) Command() string { return `Diafile` }
func (cmd Diafile) Generate(tag string, out emitter.Emitter) {
	size := ""
	if cmd.SizeIndication != "" && cmd.Size != "" {
		size = fmt.Sprintf(" %s=%s", cmd.SizeIndication, cmd.Size)
	}

	out.Println("%sdiafile %s%s%s", tag,
		word(cmd.File),
		optionalf(` "%s"`, cmd.Caption),
		size)
}

// Dir is structure for `dir` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmddir
type Dir struct {
	PathFragment string
}

func (cmd Dir) Command() string { return `Dir` }
func (cmd Dir) Generate(tag string, out emitter.Emitter) {
	out.Println("%sdir%s", tag, optional(word(cmd.PathFragment)))
}

// Docbookinclude is structure for `docbookinclude` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmddocbookinclude
type Docbookinclude struct{}

// Docbookonly is structure for `docbookonly` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmddocbookonly
type Docbookonly struct{}

// Dontinclude is structure for `dontinclude` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmddontinclude
type Dontinclude struct{}

// Dot is structure for `dot` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmddot
type Dot struct{}

// Dotfile is structure for `dotfile` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmddotfile
type Dotfile struct{}

// E is structure for `e` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmde
type E struct {
	Word string
}

func (cmd E) Command() string { return `E` }
func (cmd E) Generate(tag string, out emitter.Emitter) {
	out.Print("%se %s", tag, word(cmd.Word))
}

// Else is structure for `else` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdelse
type Else struct{}

// Elseif is structure for `elseif` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdelseif
type Elseif struct{}

// Em is structure for `em` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdem
type Em struct {
	Word string
}

func (cmd Em) Command() string { return `Em` }
func (cmd Em) Generate(tag string, out emitter.Emitter) {
	out.Print("%sEm %s", tag, word(cmd.Word))
}

// MultiEm is structure for multi word `em` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdem
type MultiEm struct {
	Text string
}

func (cmd MultiEm) Command() string { return `MultiEm` }
func (cmd MultiEm) Generate(tag string, out emitter.Emitter) {
	out.Print("<em>%s</em>", cmd.Text)
}

// Emoji is structure for `emoji` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdemoji
type Emoji struct {
	Name string
}

func (cmd Emoji) Command() string { return `Emoji` }
func (cmd Emoji) Generate(tag string, out emitter.Emitter) {
	out.Print("%semoji %s", tag, cmd.Name)
}

// Endcode is structure for `endcode` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdendcode
// This should not be used by itself!
type Endcode struct{}

func (cmd Endcode) Command() string { return `Endcode` }
func (cmd Endcode) Generate(tag string, out emitter.Emitter) {
	out.Println("%sendcode", tag)
}

// Endcond is structure for `endcond` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdendcond
type Endcond struct{}

// Enddocbookonly is structure for `enddocbookonly` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdenddocbookonly
type Enddocbookonly struct{}

// Enddot is structure for `enddot` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdenddot
type Enddot struct{}

// Endhtmlonly is structure for `endhtmlonly` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdendhtmlonly
type Endhtmlonly struct{}

// Endif is structure for `endif` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdendif
type Endif struct{}

// Endinternal is structure for `endinternal` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdendinternal
type Endinternal struct{}

// Endlatexonly is structure for `endlatexonly` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdendlatexonly
type Endlatexonly struct{}

// Endlink is structure for `endlink` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdendlink
type Endlink struct{}

// Endmanonly is structure for `endmanonly` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdendmanonly
type Endmanonly struct{}

// Endmsc is structure for `endmsc` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdendmsc
type Endmsc struct{}

// Endparblock is structure for `endparblock` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdendparblock
// This should not be used by itself!
type Endparblock struct{}

func (cmd Endparblock) Command() string { return `Endparblock` }
func (cmd Endparblock) Generate(tag string, out emitter.Emitter) {
	out.Println("%sendparblock", tag)
}

// Endrtfonly is structure for `endrtfonly` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdendrtfonly
type Endrtfonly struct{}

// Endsecreflist is structure for `endsecreflist` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdendsecreflist
type Endsecreflist struct{}

// Endverbatim is structure for `endverbatim` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdendverbatim
type Endverbatim struct{}

// Enduml is structure for `enduml` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdenduml
type Enduml struct{}

// Endxmlonly is structure for `endxmlonly` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdendxmlonly
type Endxmlonly struct{}

// Enum is structure for `enum` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdenum
type Enum struct {
	Name string
}

func (cmd Enum) Command() string { return `Enum` }
func (cmd Enum) Generate(tag string, out emitter.Emitter) {
	out.Println("%senum %s", tag, word(cmd.Name))
}

// Example is structure for `example` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdexample
type Example struct{}

// Exception is structure for `exception` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdexception
type Exception struct {
	ExceptionObject      string
	ExceptionDescription string
}

func (cmd Exception) Command() string { return `Exception` }
func (cmd Exception) Generate(tag string, out emitter.Emitter) {
	out.Println("%sexception %s %s", tag, word(cmd.ExceptionObject), cmd.ExceptionDescription)
}

// Extends is structure for `extends` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdextends
type Extends struct {
	Name string
}

func (cmd Extends) Command() string { return `Extends` }
func (cmd Extends) Generate(tag string, out emitter.Emitter) {
	out.Println("%sextends %s", tag, word(cmd.Name))
}

// FParanthesesLeft is structure for `f(` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdfrndopen
type FParanthesesLeft struct{}

// FParanthesesRight is structure for `f)` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdfrndclose
type FParanthesesRight struct{}

// FDollar is structure for `f$` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdfdollar
type FDollar struct{}

// FBracketLeft is structure for `f[` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdfbropen
type FBracketLeft struct{}

// FBracketRight is structure for `f]` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdfbrclose
type FBracketRight struct{}

// FBracesLeft is structure for `f{` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdfcurlyopen
type FBracesLeft struct{}

// FBracesRigt is structure for `f}` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdfcurlyclose
type FBracesRigt struct{}

// File is structure for `file` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdfile
type File struct {
	Name string
}

func (cmd File) Command() string { return `File` }
func (cmd File) Generate(tag string, out emitter.Emitter) {
	out.Println("%sfile%s", tag, optional(word(cmd.Name)))
}

// Fileinfo is structure for `fileinfo` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdfileinfo
type Fileinfo struct{}

// Fn is structure for `fn` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdfn
type Fn struct{}

// HeaderFile is structure for `HeaderFile` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdheaderfile
type HeaderFile struct {
	File string
	Name string
}

func (cmd HeaderFile) Command() string { return `HeaderFile` }
func (cmd HeaderFile) Generate(tag string, out emitter.Emitter) {
	out.Println("%sexception %s%s", tag, word(cmd.File), optional(word(cmd.Name)))
}

// Hidecallergraph is structure for `hidecallergraph` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdhidecallergraph
type Hidecallergraph struct{}

// Hidecallgraph is structure for `hidecallgraph` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdhidecallgraph
type Hidecallgraph struct{}

// Hiderefby is structure for `hiderefby` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdhiderefby
type Hiderefby struct{}

// Hiderefs is structure for `hiderefs` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdhiderefs
type Hiderefs struct{}

// Hideinitializer is structure for `hideinitializer` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdhideinitializer
type Hideinitializer struct{}

// Htmlinclude is structure for `htmlinclude` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdhtmlinclude
type Htmlinclude struct{}

// Htmlonly is structure for `htmlonly` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdhtmlonly
type Htmlonly struct{}

// Idlexcept is structure for `idlexcept` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdidlexcept
type Idlexcept struct {
	Name string
}

func (cmd Idlexcept) Command() string { return `Idlexcept` }
func (cmd Idlexcept) Generate(tag string, out emitter.Emitter) {
	out.Println("%sidlexcept %s", tag, word(cmd.Name))
}

// If is structure for `if` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdif
type If struct{}

// Ifnot is structure for `ifnot` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdifnot
type Ifnot struct{}

// Image is structure for `image` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdimage
type Image struct{}

// Implements is structure for `implements` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdimplements
type Implements struct {
	Name string
}

func (cmd Implements) Command() string { return `Implements` }
func (cmd Implements) Generate(tag string, out emitter.Emitter) {
	out.Println("%simplements %s", tag, word(cmd.Name))
}

// Include is structure for `include` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdinclude
type Include struct{}

// Includedoc is structure for `includedoc` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdincludedoc
type Includedoc struct{}

// Includelineno is structure for `includelineno` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdincludelineno
type Includelineno struct{}

// Ingroup is structure for `ingroup` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdingroup
type Ingroup struct{}

// Internal is structure for `internal` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdinternal
type Internal struct{}

// Invariant is structure for `invariant` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdinvariant
type Invariant struct{}

// Interface is structure for `interface` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdinterface
type Interface struct{}

// Latexinclude is structure for `latexinclude` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdlatexinclude
type Latexinclude struct{}

// Latexonly is structure for `latexonly` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdlatexonly
type Latexonly struct{}

// Li is structure for `li` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdli
type Li struct{}

// Line is structure for `line` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdline
type Line struct{}

// Lineinfo is structure for `lineinfo` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdlineinfo
type Lineinfo struct{}

// Link is structure for `link` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdlink
type Link struct{}

// Mainpage is structure for `mainpage` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdmainpage
type Mainpage struct{}

// Maninclude is structure for `maninclude` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdmaninclude
type Maninclude struct{}

// Manonly is structure for `manonly` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdmanonly
type Manonly struct{}

// Memberof is structure for `memberof` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdmemberof
type Memberof struct {
	Name string
}

func (cmd Memberof) Command() string { return `Memberof` }
func (cmd Memberof) Generate(tag string, out emitter.Emitter) {
	out.Println("%smemberof %s", tag, word(cmd.Name))
}

// Msc is structure for `msc` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdmsc
type Msc struct{}

// Mscfile is structure for `mscfile` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdmscfile
type Mscfile struct{}

// N is structure for `n` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdn
type N struct{}

// Name is structure for `name` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdname
type Name struct{}

// Namespace is structure for `namespace` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdnamespace
type Namespace struct {
	Name string
}

func (cmd Namespace) Command() string { return `Namespace` }
func (cmd Namespace) Generate(tag string, out emitter.Emitter) {
	out.Println("%snamespace %s", tag, word(cmd.Name))
}

// Noop is structure for `noop` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdnoop
type Noop struct {
	IgnoredText string
}

func (cmd Noop) Command() string { return `Noop` }
func (cmd Noop) Generate(tag string, out emitter.Emitter) {
	out.Println("%snoop %s", tag, cmd.IgnoredText)
}

// Nosubgrouping is structure for `nosubgrouping` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdnosubgrouping
type Nosubgrouping struct{}

// Note is structure for `note` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdnote
type Note struct{}

// Overload is structure for `overload` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdoverload
type Overload struct{}

// P is structure for `p` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdp
type P struct{}

// Package is structure for `package` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdpackage
type Package struct {
	Name string
}

func (cmd Package) Command() string { return `Package` }
func (cmd Package) Generate(tag string, out emitter.Emitter) {
	out.Println("%spackage %s", tag, word(cmd.Name))
}

// Page is structure for `page` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdpage
type Page struct{}

// Par is structure for `par` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdpar
type Par struct {
	Title     string
	Paragraph string
}

func (cmd Par) Command() string { return `Par` }
func (cmd Par) Generate(tag string, out emitter.Emitter) {
	out.Print("%spar ", tag)
	if cmd.Title != "" {
		out.Println("%s", cmd.Title)
	}
	out.Println("%s", cmd.Paragraph)
}

// Paragraph is structure for `paragraph` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdparagraph
type Paragraph struct {
	Name  string
	Title string
}

func (cmd Paragraph) Command() string { return `Paragraph` }
func (cmd Paragraph) Generate(tag string, out emitter.Emitter) {
	out.Println("%sparagraph %s %s", tag, word(cmd.Name), cmd.Title)
}

// Param is structure for `param` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdparam
type Param struct {
	Direction            string
	ParameterName        string
	ParameterDescription string
}

func (cmd Param) Command() string { return `Param` }
func (cmd Param) Generate(tag string, out emitter.Emitter) {
	out.Print("%sparam", tag)
	if cmd.Direction != "" && cmd.directionValid() {
		out.Print("[%s]", cmd.Direction)
	}
	out.Println(" %s %s", cmd.ParameterName, cmd.ParameterDescription)
}

func (cmd Param) directionValid() bool {
	for _, dir := range []string{"in", "in,out", "out"} {
		if cmd.Direction == dir {
			return true
		}
	}
	return false
}

// Parblock is structure for `parblock` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdparblock
type Parblock struct {
	Paragraphs []string
}

func (cmd Parblock) Command() string { return `Parblock` }
func (cmd Parblock) Generate(tag string, out emitter.Emitter) {
	out.Println("%sparblock", tag)
	defer Endparblock{}.Generate(tag, out)

	for i := 0; i < len(cmd.Paragraphs); i++ {
		out.Println("%s", cmd.Paragraphs[i])
		if i < len(cmd.Paragraphs)-1 {
			out.Newline()
		}
	}
}

// Post is structure for `post` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdpost
type Post struct{}

// Pre is structure for `pre` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdpre
type Pre struct{}

// Private is structure for `private` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdprivate
type Private struct{}

// Privatesection is structure for `privatesection` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdprivatesection
type Privatesection struct{}

// Property is structure for `property` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdproperty
type Property struct{}

// Protected is structure for `protected` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdprotected
type Protected struct{}

// Protectedsection is structure for `protectedsection` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdprotectedsection
type Protectedsection struct{}

// Protocol is structure for `protocol` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdprotocol
type Protocol struct{}

// Public is structure for `public` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdpublic
type Public struct{}

// Publicsection is structure for `publicsection` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdpublicsection
type Publicsection struct{}

// Pure is structure for `pure` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdpure
type Pure struct{}

// Raisewarning is structure for `raisewarning` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdraisewarning
type Raisewarning struct{}

// Ref is structure for `ref` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdref
type Ref struct{}

// Refitem is structure for `refitem` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdrefitem
type Refitem struct {
	Name string
}

func (cmd Refitem) Command() string { return `Refitem` }
func (cmd Refitem) Generate(tag string, out emitter.Emitter) {
	out.Println("%srefitem %s", tag, word(cmd.Name))
}

// Related is structure for `related` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdrelated
type Related struct {
	Name string
}

func (cmd Related) Command() string { return `Related` }
func (cmd Related) Generate(tag string, out emitter.Emitter) {
	out.Println("%srelated %s", tag, word(cmd.Name))
}

// Relates is structure for `relates` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdrelates
type Relates struct {
	Name string
}

func (cmd Relates) Command() string { return `Relates` }
func (cmd Relates) Generate(tag string, out emitter.Emitter) {
	out.Println("%srelates %s", tag, word(cmd.Name))
}

// Relatedalso is structure for `relatedalso` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdrelatedalso
type Relatedalso struct {
	Name string
}

func (cmd Relatedalso) Command() string { return `Relatedalso` }
func (cmd Relatedalso) Generate(tag string, out emitter.Emitter) {
	out.Println("%srelatedalso %s", tag, word(cmd.Name))
}

// Relatesalso is structure for `relatesalso` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdrelatesalso
type Relatesalso struct {
	Name string
}

func (cmd Relatesalso) Command() string { return `Relatesalso` }
func (cmd Relatesalso) Generate(tag string, out emitter.Emitter) {
	out.Println("%srelatesalso %s", tag, word(cmd.Name))
}

// Remark is structure for `remark` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdremark
type Remark struct {
	Text string
}

func (cmd Remark) Command() string { return `Remark` }
func (cmd Remark) Generate(tag string, out emitter.Emitter) {
	out.Println("%sremark %s", tag, cmd.Text)
}

// Remarks is structure for `remarks` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdremarks
type Remarks struct {
	Text string
}

func (cmd Remarks) Command() string { return `Remarks` }
func (cmd Remarks) Generate(tag string, out emitter.Emitter) {
	out.Println("%sremarks %s", tag, cmd.Text)
}

// Result is structure for `result` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdresult
type Result struct{}

// Return is structure for `return` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdreturn
type Return struct {
	Description string
}

func (cmd Return) Command() string { return `Return` }
func (cmd Return) Generate(tag string, out emitter.Emitter) {
	out.Println("%sreturn %s", tag, cmd.Description)
}

// Returns is structure for `returns` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdreturns
type Returns struct {
	Description string
}

func (cmd Returns) Command() string { return `Returns` }
func (cmd Returns) Generate(tag string, out emitter.Emitter) {
	out.Println("%sreturns %s", tag, cmd.Description)
}

// Retval is structure for `retval` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdretval
type Retval struct {
	Name    string
	Message string
}

// Rtfinclude is structure for `rtfinclude` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdrtfinclude
type Rtfinclude struct{}

// Rtfonly is structure for `rtfonly` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdrtfonly
type Rtfonly struct{}

// Sa is structure for `sa` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdsa
type Sa struct{}

// Secreflist is structure for `secreflist` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdsecreflist
type Secreflist struct{}

// Section is structure for `section` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdsection
type Section struct{}

// See is structure for `see` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdsee
type See struct{}

// Short is structure for `short` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdshort
type Short struct {
	ShortDescription string
}

func (cmd Short) Command() string { return `Short` }
func (cmd Short) Generate(tag string, out emitter.Emitter) {
	out.Println("%sshort %s", tag, cmd.ShortDescription)
}

// Showdate is structure for `showdate` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdshowdate
type Showdate struct {
	Format   string
	DateTime string
}

func (cmd Showdate) Command() string { return `Showdate` }
func (cmd Showdate) Generate(tag string, out emitter.Emitter) {
	out.Println(`%sshowdate "%s" %s`, tag, cmd.Format, cmd.DateTime)
}

// Showinitializer is structure for `showinitializer` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdshowinitializer
type Showinitializer struct{}

// Showrefby is structure for `showrefby` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdshowrefby
type Showrefby struct{}

// Showrefs is structure for `showrefs` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdshowrefs
type Showrefs struct{}

// Since is structure for `since` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdsince
type Since struct{}

// Skip is structure for `skip` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdskip
type Skip struct{}

// Skipline is structure for `skipline` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdskipline
type Skipline struct{}

// Snippet is structure for `snippet` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdsnippet
type Snippet struct{}

// Snippetdoc is structure for `snippetdoc` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdsnippetdoc
type Snippetdoc struct{}

// Snippetlineno is structure for `snippetlineno` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdsnippetlineno
type Snippetlineno struct{}

// Static is structure for `static` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdstatic
type Static struct{}

// Startuml is structure for `startuml` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdstartuml
type Startuml struct{}

// Struct is structure for `struct` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdstruct
type Struct struct{}

// Subpage is structure for `subpage` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdsubpage
type Subpage struct{}

// Subsection is structure for `subsection` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdsubsection
type Subsection struct{}

// Subsubsection is structure for `subsubsection` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdsubsubsection
type Subsubsection struct{}

// Tableofcontents is structure for `tableofcontents` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdtableofcontents
type Tableofcontents struct{}

// Test is structure for `test` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdtest
type Test struct{}

// Throw is structure for `throw` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdthrow
type Throw struct {
	ExceptionObject      string
	ExceptionDescription string
}

func (cmd Throw) Command() string { return `Throw` }
func (cmd Throw) Generate(tag string, out emitter.Emitter) {
	out.Print("%sthrow %s %s", tag, word(cmd.ExceptionObject), cmd.ExceptionDescription)
}

// Throws is structure for `throws` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdthrows
type Throws struct {
	ExceptionObject      string
	ExceptionDescription string
}

func (cmd Throws) Command() string { return `Throws` }
func (cmd Throws) Generate(tag string, out emitter.Emitter) {
	out.Println("%sthrows", tag, word(cmd.ExceptionObject), cmd.ExceptionDescription)
}

// Todo is structure for `todo` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdtodo
type Todo struct {
	Description string
}

func (cmd Todo) Command() string { return `Todo` }
func (cmd Todo) Generate(tag string, out emitter.Emitter) {
	out.Print("%stodo %s", tag, cmd.Description)
}

// Tparam is structure for `tparam` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdtparam
type Tparam struct{}

// Typedef is structure for `typedef` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdtypedef
type Typedef struct{}

// Union is structure for `union` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdunion
type Union struct{}

// Until is structure for `until` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmduntil
type Until struct{}

// Var is structure for `var` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdvar
type Var struct {
	Datatype    string
	Name        string
	Description string
}

func (cmd Var) Command() string { return `Var` }
func (cmd Var) Generate(tag string, out emitter.Emitter) {
	out.Println("%svar %s $%s", tag, cmd.Datatype, cmd.Name)
	out.Println("%s", cmd.Description)
}

// Verbatim is structure for `verbatim` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdverbatim
type Verbatim struct{}

// Verbinclude is structure for `verbinclude` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdverbinclude
type Verbinclude struct{}

// Version is structure for `version` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdversion
type Version struct {
	Number string
}

func (cmd Version) Command() string { return `Version` }
func (cmd Version) Generate(tag string, out emitter.Emitter) {
	out.Println("%sversion %s", tag, cmd.Number)
}

// Vhdlflow is structure for `vhdlflow` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdvhdlflow
type Vhdlflow struct{}

// Warning is structure for `warning` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdwarning
type Warning struct {
	Message string
}

func (cmd Warning) Command() string { return `Warning` }
func (cmd Warning) Generate(tag string, out emitter.Emitter) {
	out.Println("%swarning %s", tag, cmd.Message)
}

// Weakgroup is structure for `weakgroup` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdweakgroup
type Weakgroup struct{}

// Xmlinclude is structure for `xmlinclude` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdxmlinclude
type Xmlinclude struct{}

// Xmlonly is structure for `xmlonly` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdxmlonly
type Xmlonly struct{}

// Xrefitem is structure for `xrefitem` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdxrefitem
type Xrefitem struct{}

// Dollar is structure for `$` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmddollar
type Dollar struct{}

func (cmd Dollar) Command() string { return `Dollar` }
func (cmd Dollar) Generate(tag string, out emitter.Emitter) {
	out.Print("%s$", tag)
}

// At is structure for `@` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdat
type At struct{}

func (cmd At) Command() string { return `At` }
func (cmd At) Generate(tag string, out emitter.Emitter) {
	out.Print("%s@", tag)
}

// Backslash is structure for `\` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdbackslash
type Backslash struct{}

func (cmd Backslash) Command() string { return `Backslash` }
func (cmd Backslash) Generate(tag string, out emitter.Emitter) {
	out.Print(`%s\`, tag)
}

// Ampersand is structure for `&` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdamp
type Ampersand struct{}

func (cmd Ampersand) Command() string { return `Ampersand` }
func (cmd Ampersand) Generate(tag string, out emitter.Emitter) {
	out.Print("%s&", tag)
}

// Tilde is structure for `~` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdtilde
type Tilde struct {
	LanguageID string
}

func (cmd Tilde) Command() string { return `Tilde` }
func (cmd Tilde) Generate(tag string, out emitter.Emitter) {
	out.Print("%s~%s", tag, optional(cmd.LanguageID))
}

// LessThan is structure for `<` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdlt
type LessThan struct{}

func (cmd LessThan) Command() string { return `LessThan` }
func (cmd LessThan) Generate(tag string, out emitter.Emitter) {
	out.Print("%s<", tag)
}

// Lt is alias for LessThan
type Lt struct {
	LessThan
}

// Equals is structure for `=` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdeq
type Equals struct{}

func (cmd Equals) Command() string { return `Equals` }
func (cmd Equals) Generate(tag string, out emitter.Emitter) {
	out.Print("%s=", tag)
}

// GreaterThan is structure for `>` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdgt
type GreaterThan struct{}

func (cmd GreaterThan) Command() string { return `GreaterThan` }
func (cmd GreaterThan) Generate(tag string, out emitter.Emitter) {
	out.Print("%s>", tag)
}

// Gt is an alias for GreaterThan
type Gt struct {
	GreaterThan
}

// Hashtag is structure for `#` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdhash
type Hashtag struct{}

func (cmd Hashtag) Command() string { return `Hashtag` }
func (cmd Hashtag) Generate(tag string, out emitter.Emitter) {
	out.Print("%s#", tag)
}

// Percent is structure for `%` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdperc
type Percent struct{}

func (cmd Percent) Command() string { return `Percent` }
func (cmd Percent) Generate(tag string, out emitter.Emitter) {
	out.Print("%s%", tag)
}

// QuotationMark is structure for `"` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdquot
type QuotationMark struct{}

func (cmd QuotationMark) Command() string { return `QuotationMark` }
func (cmd QuotationMark) Generate(tag string, out emitter.Emitter) {
	out.Print(`%s"`, tag)
}

// CharDot is structure for `.` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdchardot
type CharDot struct{}

func (cmd CharDot) Command() string { return `CharDot` }
func (cmd CharDot) Generate(tag string, out emitter.Emitter) {
	out.Print("%s.", tag)
}

// Colon is structure for `::` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmddcolon
type Colon struct{}

func (cmd Colon) Command() string { return `Colon` }
func (cmd Colon) Generate(tag string, out emitter.Emitter) {
	out.Print("%s::", tag)
}

// Pipe is structure for `|` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdpipe
type Pipe struct{}

func (cmd Pipe) Command() string { return `Pipe` }
func (cmd Pipe) Generate(tag string, out emitter.Emitter) {
	out.Print("%s|", tag)
}

// NDash is structure for `--` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdndash
type NDash struct{}

func (cmd NDash) Command() string { return `NDash` }
func (cmd NDash) Generate(tag string, out emitter.Emitter) {
	out.Print("%s--", tag)
}

// MDash is structure for `---` command.
//
// For more details, see: https://doxygen.nl/manual/commands.html#cmdmdash
type MDash struct{}

func (cmd MDash) Command() string { return `MDash` }
func (cmd MDash) Generate(tag string, out emitter.Emitter) {
	out.Print("%s---", tag)
}
