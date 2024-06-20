# V3 Design

This is a draft for the new v3 API design of Box CLI Maker which improves the old API design and solves.

- User level error handling
- No more unecessary printing errors
- Uses idomatic Go coding style
- Easier to use
- Better support for terminals

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
        Type("Single").
        Color("Cyan")

if boxStr, err := b.Render("Box CLI Maker", "Highly Customized Terminal Box Maker"); err == nil {
    fmt.Println(boxStr)
}
```

</td>
</tr>
</table>
