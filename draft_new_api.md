# V3 Design

This is a draft for the new v3 API design of Box CLI Maker which improves the old API design and solves.

- User level error handling
- No more unecessary printing errors
- Uses idomatic Go coding style
- Easier to use
- Better support for terminals
- Remove `interface{}` for Color Types. Allow only string.
- Remove untyped strings for Box Types and Title Position and use imported named strings.

The API design is inspired from [charmbracelet/huh](https://github.com/charmbracelet/huh) and [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss).

<table>

<tr>
<td> <strong>v2 Examples</strong> </td><td> <strong>v3 Examples</strong> </td></tr>
<tr>
<td>

```go
Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan"})
Box.Print("Box CLI Maker", "Highly Customized Terminal Box Maker")
```

</td>
<td>

```go
b := box.NewBox().
        Width(2).
        Height(5).
        Type(box.Single).
        Color("Cyan")

if boxStr, err := b.Render("Box CLI Maker", "Highly Customized Terminal Box Maker"); err == nil {
    fmt.Println(boxStr)
}
```

</td>
</tr>

<tr>
<td>

```go
config := box.Config{Px: 2, Py: 3, Type: "", TitlePos: "Inside", Color: "Green"}
boxNew := box.Box{TopRight: "*", TopLeft: "*", BottomRight: "*", BottomLeft: "*", Horizontal: "-", Vertical: "|", Config: config}
boxNew.Println("Box CLI Maker", "Make Highly Customized Terminal Boxes")
```

</td>
<td>

```go
b := box.NewBox().
        Width(2).
        Height(3).
        TitlePositon(box.Inside).
        Color("Green").
        TopRight("*").
        TopLeft("*").
        BottomRight("*").
        BottomLeft("*").
        Horizontal("-").
        Vertical("|")

// Even the below will work

b = box.Box{TopRight: "*", TopLeft: "*", BottomRight: "*", BottomLeft: "*", Horizontal: "-", Vertical: "|"}
b = b.Width(2).
    Height(3).
    TitlePositon(box.Inside).
    Color("Green")

if boxStr, err := b.Render("Box CLI Maker", "Highly Customized Terminal Box Maker"); err == nil {
    fmt.Println(boxStr)
}
```

</td>
</tr>

<tr>
<td>

```go
bx := box.New(box.Config{
		Px:            2,
		Py:            0,
		Type:          "Single",
		Color:         "Green",
		TitlePos:      "Top",
		AllowWrapping: true,
		WrappingLimit: num,
	})
	bx.Println("Content Wrappingg works!", strings.Repeat(" Box CLI Maker 盒子製 造商,📦 ", 160))
```

</td>
<td>

```go
b := box.NewBox().
        Width(2).
        Height(0).
        Type(box.Single).
        Color("Green").
        TitlePosition(box.Inside).
        AllowWrapping(true).
        WrappingLimit(num)


if boxStr, err := b.Render("Content Wrappingg works!", strings.Repeat(" Box CLI Maker 盒子製 造商,📦 ", 160)); err == nil {
    fmt.Println(boxStr)
}
```

</td>
</tr>
</table>
