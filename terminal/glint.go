package terminal

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/mitchellh/go-glint"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/term"
)

type glintUI struct {
	d      *glint.Document
	c      []glint.Component
	last   *TextComponent
	append bool
}

type TextComponent struct {
	*glint.TextComponent
	v     string
	style []glint.StyleOption
}

func Text(v string, styles ...glint.StyleOption) *TextComponent {
	return &TextComponent{
		TextComponent: glint.Text(v),
		v:             v,
		style:         styles,
	}
}

func (t *TextComponent) Body(ctx context.Context) glint.Component {
	return t.TextComponent
}

func (t *TextComponent) Append(v string) {
	t.TextComponent = glint.Text(t.v + v)
	t.v = t.v + v
}

func (t *TextComponent) Clear() {
	t.TextComponent = glint.Text("")
	t.v = ""
}

func (t *TextComponent) StyleMatch(styles ...glint.StyleOption) bool {
	for _, ts := range t.style {
		found := false
		for _, s := range styles {
			if reflect.ValueOf(ts).Pointer() == reflect.ValueOf(s).Pointer() {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func GlintUI(ctx context.Context) UI {
	result := &glintUI{
		d: glint.New(),
		c: []glint.Component{},
	}

	go result.d.Render(ctx)

	return result
}

func (ui *glintUI) Close() error {
	return ui.d.Close()
}

func (ui *glintUI) Input(input *Input) (line string, err error) {
	ui.Output(input.Prompt, WithoutNewLine(), WithStyle(input.Style), WithColor(input.Color))
	// Render the last frame
	ui.d.RenderFrame()
	// Pause so that input can be read
	ui.d.Pause()
	defer ui.d.Resume()

	if input.Secret {
		l, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return "", err
		}
		line = string(l)
	} else {
		r := bufio.NewReader(os.Stdin)
		line, err = r.ReadString('\n')
	}
	if err != nil {
		return "", err
	}

	text := strings.TrimSpace(line)
	if input.Secret {
		ui.Output("")
	} else {
		ui.Output(text)
	}
	return text, nil
}

// Interactive implements UI
func (ui *glintUI) Interactive() bool {
	return true
}

// MachineReadable implements UI
func (ui *glintUI) MachineReadable() bool {
	return false
}

// Output implements UI
func (ui *glintUI) Output(msg string, raw ...interface{}) {
	defer ui.d.RenderFrame()
	msg, style, disableNewline, _, color := Interpret(msg, raw...)

	var cs []glint.StyleOption
	switch style {
	case HeaderStyle:
		cs = append(cs, glint.Bold())
		msg = "\n» " + msg
	case ErrorStyle, ErrorBoldStyle:
		if style == ErrorBoldStyle {
			cs = append(cs, glint.Bold())
		}

		lines := strings.Split(msg, "\n")
		for i, line := range lines {
			if i == 0 {
				lines[i] = "! " + line
			} else {
				lines[i] = "  " + line
			}
		}

		msg = strings.Join(lines, "\n")

	case WarningStyle, WarningBoldStyle:
		if style == WarningBoldStyle {
			cs = append(cs, glint.Bold())
		}

	case SuccessStyle, SuccessBoldStyle:
		if style == SuccessBoldStyle {
			cs = append(cs, glint.Bold())
		}

	case InfoStyle:
		lines := strings.Split(msg, "\n")
		for i, line := range lines {
			lines[i] = "  " + line
		}

		msg = strings.Join(lines, "\n")

	case InfoBoldStyle:
		cs = append(cs, glint.Bold())
		lines := strings.Split(msg, "\n")
		for i, line := range lines {
			lines[i] = "  " + line
		}

		msg = strings.Join(lines, "\n")
	}

	cs = append(cs, glint.Color(color))

	a := ui.append
	ui.append = disableNewline

	if a && ui.last != nil {
		if ui.last.StyleMatch(cs...) {
			ui.last.Append(msg)
			return
		}
	}

	c := Text(msg, cs...)
	ui.last = c

	ui.d.Append(glint.Style(c, cs...))
}

// ClearLine implements UI
func (ui *glintUI) ClearLine() {
	defer ui.d.RenderFrame()
	if ui.last == nil {
		return
	}
	ui.last.Clear()
	ui.append = true
}

// NamedValues implements UI
func (ui *glintUI) NamedValues(rows []NamedValue, opts ...Option) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}

	var buf bytes.Buffer
	tr := tabwriter.NewWriter(&buf, 1, 8, 0, ' ', tabwriter.AlignRight)
	for _, row := range rows {
		switch v := row.Value.(type) {
		case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
			fmt.Fprintf(tr, "  %s: \t%d\n", row.Name, row.Value)
		case float32, float64:
			fmt.Fprintf(tr, "  %s: \t%f\n", row.Name, row.Value)
		case bool:
			fmt.Fprintf(tr, "  %s: \t%v\n", row.Name, row.Value)
		case string:
			if v == "" {
				continue
			}
			fmt.Fprintf(tr, "  %s: \t%s\n", row.Name, row.Value)
		default:
			fmt.Fprintf(tr, "  %s: \t%s\n", row.Name, row.Value)
		}
	}
	tr.Flush()

	// We want to trim the trailing newline
	text := buf.String()
	if len(text) > 0 && text[len(text)-1] == '\n' {
		text = text[:len(text)-1]
	}

	ui.d.Append(glint.Finalize(glint.Text(text)))
}

// OutputWriters implements UI
func (ui *glintUI) OutputWriters() (io.Writer, io.Writer, error) {
	return os.Stdout, os.Stderr, nil
}

// Status implements UI
func (ui *glintUI) Status() Status {
	st := newGlintStatus()
	ui.d.Append(st)
	return st
}

func (ui *glintUI) StepGroup() StepGroup {
	ctx, cancel := context.WithCancel(context.Background())
	sg := &glintStepGroup{ctx: ctx, cancel: cancel}
	ui.d.Append(sg)
	return sg
}

// Table implements UI
func (ui *glintUI) Table(tbl *Table, opts ...Option) {
	var buf bytes.Buffer
	table := tablewriter.NewWriter(&buf)
	table.SetHeader(tbl.Headers)
	table.SetBorder(false)
	table.SetAutoWrapText(false)

	for _, row := range tbl.Rows {
		colors := make([]tablewriter.Colors, len(row))
		entries := make([]string, len(row))

		for i, ent := range row {
			entries[i] = ent.Value

			color, ok := colorMapping[ent.Color]
			if ok {
				colors[i] = tablewriter.Colors{color}
			}
		}

		table.Rich(entries, colors)
	}

	table.Render()

	ui.d.Append(glint.Finalize(glint.Text(buf.String())))
}
