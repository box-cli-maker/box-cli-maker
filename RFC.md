# v3 release plans

This is a draft for the new v3 API design of Box CLI Maker which improves the old API design and lays down a roadmap for the release.

The API design is inspired from [charmbracelet/huh](https://github.com/charmbracelet/huh), [charmbracelet/glamour](https://github.com/charmbracelet/glamour) and [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss).

- User level error handling
- No more un-necessary printing errors
- Uses idiomatic Go coding style
- Easier to use
- Better support for terminals
- Remove `interface{}` for Color Types. Allow only strings.
- Remove untyped strings for Box Types and Title Position and use imported named strings.
- Add more inbuilt Box styles.
- Remove manual color detection and rounding off code and use an external library.
- Decrease the number of dependencies.
- Replaced ambiguous `Px` and `Py` fields (which can mean outer size, but were padding) with functions such as `Padding`, `HPadding` and `VPadding`.
- Use [charmbracelet/vhs](https://github.com/charmbracelet/vhs) to record demos.
- Shift to a GitHub organization and start an Open Collective.

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
        Padding(2, 5).
        Style(box.Single).
        Color("Cyan")

// Ignore error if Box rendering is possible
boxStr, _ := b.Render("Box CLI Maker", "Highly Customized Terminal Box Maker") 
fmt.Println(boxStr)

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
        Padding(2, 3).
        TitlePosition(box.Inside).
        Color("Green").
        TopRight("*").
        TopLeft("*").
        BottomRight("*").
        BottomLeft("*").
        Horizontal("-").
        Vertical("|")

// Ignore error if Box rendering is possible
boxStr, _ := b.Render("Box CLI Maker", "Highly Customized Terminal Box Maker") 
fmt.Println(boxStr)
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
  ContentColor: "Red",
 })
 bx.Println("Content Wrappingg works!", strings.Repeat(" Box CLI Maker 盒子製 造商,📦 ", 160))
```

</td>
<td>

```go
b := box.NewBox().
        Padding(2, 0).
        Style(box.Single).
        Color("Green").
        TitlePosition(box.Inside).
        WrapContent(true).
        WrapLimit(num).
        ContentColor("Red")


// Ignore error if Box rendering is possible
boxStr, _ := b.Render("Box CLI Maker", "Highly Customized Terminal Box Maker") 
fmt.Println(boxStr)
```

</td>
</tr>
</table>

## v2 vs v3 API differences

| Aspect           | v2                                      | v3                                                   |
|------------------|-----------------------------------------|------------------------------------------------------|
| Construction     | `New(Config)` with public `Config`/`Box` fields | `NewBox()` with fluent builder methods       |
| Styles           | `Type: "Single"` (untyped string)       | `Style(box.Single)` (typed `BoxStyle`)               |
| Title position   | `TitlePos: "Top"` (string)              | `TitlePosition(box.Top)` (typed `TitlePosition`)     |
| Custom borders   | Struct literals on `Box` fields         | Builder methods: `TopRight`, `TopLeft`, etc.         |
| Colors           | Mixed config fields, some `interface{}` | `Color`, `TitleColor`, `ContentColor` (all `string`) |
| Padding          | `Px`, `Py` fields (ambiguous)           | `Padding`, `HPadding`, `VPadding` methods            |
| Wrapping         | `AllowWrapping`, `WrappingLimit` fields | `WrapContent`, `WrapLimit` methods                   |
| Rendering        | `Println`/`Print` with internal printing | `Render`/`MustRender` returning a string (you choose how to print) |
