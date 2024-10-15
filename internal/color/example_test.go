package color_test

import (
	"fmt"

	"versioning-cli/internal/color"
)

func Example() {
	color.Unset() // forcefully unset the package [atomic.Value] so example tests can pass

	// --> Write the content "Default" out to stdout without color
	color.Color().Default("Default").Print()

	// --> Write the content "red", "blue", "green" out to stdout with color escapes
	color.Color().Red("Red").Print()
	color.Color().Blue("Blue").Print()
	color.Color().Green("Green").Print()

	// --> Wrap color(s) with bold color escapes and write to stdout
	color.Color().Bold(color.Color().Cyan("Cyan")).Print()

	// --> Customize how newlines and spaces are added to the formatted output
	color.Color().Red("Color-1").Print(func(o *color.Options) { o.Line = false; o.Space = true })
	color.Color().Red("Color-2").Print(func(o *color.Options) { o.Line = false; o.Space = true })
	color.Color().Red("Color-3").Print()

	// --> Store the ANSI-modified string to a variable and then format, write the value to stdout
	v := color.Color().Italic(color.Color().Magenta("Magenta")).String()

	fmt.Printf("Example Magenta Color Output: %s\n", v)

	// Output: Default
	// Red
	// Blue
	// Green
	// Cyan
	// Color-1 Color-2 Color-3
	// Example Magenta Color Output: Magenta
}
