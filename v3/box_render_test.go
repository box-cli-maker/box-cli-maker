package box

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/mattn/go-runewidth"
)

func TestRenderBasicBox(t *testing.T) {
	b := NewBox().Padding(2, 1).Style(Single)

	out, err := b.Render("Box CLI Maker", "Highly Customizable Terminal Box Maker")
	if err != nil {
		t.Fatalf("Render returned error: %v", err)
	}
	if out == "" {
		t.Fatalf("expected non-empty render output")
	}

	if !strings.Contains(out, "Box CLI Maker") || !strings.Contains(out, "Highly Customizable Terminal Box Maker") {
		t.Fatalf("rendered output does not contain title/content: %q", out)
	}

	// Basic structural checks: top and bottom lines should use the Single style corners.
	lines := strings.Split(out, "\n")
	if len(lines) < 3 {
		t.Fatalf("expected at least 3 lines in rendered box, got %d", len(lines))
	}

	// Last element is empty due to trailing newline; bottom bar is at len(lines)-2.
	top := lines[0]
	bottom := lines[len(lines)-2]

	if !strings.HasPrefix(top, "┌") || !strings.HasSuffix(top, "┐") {
		t.Errorf("top bar does not use Single style corners: %q", top)
	}
	if !strings.HasPrefix(bottom, "└") || !strings.HasSuffix(bottom, "┘") {
		t.Errorf("bottom bar does not use Single style corners: %q", bottom)
	}
}

func TestRenderTitlePositions(t *testing.T) {
	title := "My Title"
	content := "Some content"

	cases := []struct {
		name string
		pos  TitlePosition
	}{
		{"inside", Inside},
		{"top", Top},
		{"bottom", Bottom},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b := NewBox().Padding(2, 1).Style(Single).TitlePosition(tc.pos)
			out, err := b.Render(title, content)
			if err != nil {
				t.Fatalf("Render returned error for position %v: %v", tc.pos, err)
			}

			lines := strings.Split(out, "\n")
			if len(lines) < 3 {
				t.Fatalf("expected at least 3 lines, got %d", len(lines))
			}
			top := lines[0]
			bottom := lines[len(lines)-2]
			interior := lines[1 : len(lines)-2]

			hasTitleInside := false
			for _, l := range interior {
				if strings.Contains(l, title) {
					hasTitleInside = true
					break
				}
			}

			switch tc.pos {
			case Inside:
				if !hasTitleInside {
					t.Errorf("expected title to appear inside box for Inside position; output: %q", out)
				}
			case Top:
				if !strings.Contains(top, title) {
					t.Errorf("expected title to appear in top bar for Top position; top: %q", top)
				}
			case Bottom:
				if !strings.Contains(bottom, title) {
					t.Errorf("expected title to appear in bottom bar for Bottom position; bottom: %q", bottom)
				}
			}
		})
	}
}

func TestRenderInvalidBoxStyle(t *testing.T) {
	b := NewBox().Padding(2, 1).Style(BoxStyle("InvalidStyle"))
	_, err := b.Render("Title", "Content")
	if err == nil {
		t.Fatalf("expected error for invalid Box style, got nil")
	}
	if !strings.Contains(err.Error(), "invalid Box style") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestRenderInvalidTitlePosition(t *testing.T) {
	b := NewBox().Padding(2, 1).Style(Single).TitlePosition(TitlePosition("Weird"))
	_, err := b.Render("Title", "Content")
	if err == nil {
		t.Fatalf("expected error for invalid TitlePosition, got nil")
	}
	if !strings.Contains(err.Error(), "invalid TitlePosition") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestRenderMultilineTitleNonInside(t *testing.T) {
	b := NewBox().Padding(2, 1).Style(Single).TitlePosition(Top)
	_, err := b.Render("Line1\nLine2", "Content")
	if err == nil {
		t.Fatalf("expected error for multiline title at non-Inside position, got nil")
	}
	if !strings.Contains(err.Error(), "multiline titles are only supported Inside title position only") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestRenderWithWrapLimit(t *testing.T) {
	longContent := strings.Repeat("word ", 20)
	b := NewBox().Padding(2, 0).Style(Single).Color("Green").WrapContent(true).WrapLimit(10)

	out, err := b.Render("Wrapped", longContent)
	if err != nil {
		t.Fatalf("Render with wrapping returned error: %v", err)
	}
	if !strings.Contains(out, "Wrapped") {
		t.Errorf("expected title to appear in wrapped box output")
	}
	if !strings.Contains(out, "word") {
		t.Errorf("expected content to appear in wrapped box output")
	}
}

func TestMustRenderSuccessAndPanic(t *testing.T) {
	// Success case: MustRender should not panic when Render succeeds.
	t.Run("success", func(t *testing.T) {
		b := NewBox().Padding(1, 1).Style(Single)
		_ = b.MustRender("Title", "Content")
	})

	// Panic case: invalid style causes Render to error, hence MustRender panics.
	t.Run("panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("expected MustRender to panic for invalid style, but it did not")
			}
		}()
		b := NewBox().Padding(1, 1).Style(BoxStyle("InvalidStyle"))
		_ = b.MustRender("Title", "Content")
	})
}

func TestRenderEmojiBordersHaveConsistentWidth(t *testing.T) {
	b := NewBox().Padding(2, 1)
	b.TopLeft("📦").TopRight("📦").BottomLeft("📦").BottomRight("📦").Horizontal("📦").Vertical("📦")

	out, err := b.Render("Emoji Box", "With emoji borders")
	if err != nil {
		t.Fatalf("Render with emoji borders returned error: %v", err)
	}

	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
	if len(lines) < 3 {
		t.Fatalf("expected at least 3 lines in rendered box, got %d", len(lines))
	}

	top := ansi.Strip(lines[0])
	interior := ansi.Strip(lines[1])
	bottom := ansi.Strip(lines[len(lines)-1])

	topW := runewidth.StringWidth(top)
	interiorW := runewidth.StringWidth(interior)
	bottomW := runewidth.StringWidth(bottom)

	if topW != interiorW || interiorW != bottomW {
		t.Fatalf("expected equal visual widths for emoji box borders, got top=%d interior=%d bottom=%d", topW, interiorW, bottomW)
	}
}
