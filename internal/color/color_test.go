package color

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("Colors", func(t *testing.T) {
		t.Setenv("CI", "false")

		Color().White("White").Black("Black").Cyan("Cyan").Red("Red").Yellow("Yellow").Green("Green").Gray("Gray").Magenta("Magenta").Blue("Blue").Default("Default").Print()
	})

	t.Run("CI", func(t *testing.T) {
		t.Setenv("CI", "false")

		// override for uncoloRed output
		Force()

		t.Run("Color", func(t *testing.T) {
			t.Setenv("CI", "false")

			Color().White("White").Black("Black").Cyan("Cyan").Red("Red").Yellow("Yellow").Green("Green").Gray("Gray").Magenta("Magenta").Blue("Blue").Default("Default").Print()
		})
	})

	t.Run("Bold", func(t *testing.T) {
		t.Setenv("CI", "false")

		Color().Bold("Bold").Print()

		t.Run("Red", func(t *testing.T) {
			t.Setenv("CI", "false")

			Color().Bold(Color().Red("Bold & Red").String()).Print()
		})

		t.Run("Green", func(t *testing.T) {
			t.Setenv("CI", "false")

			Color().Bold(Color().Green("Bold & Green").String()).Print()
		})

		t.Run("Blue", func(t *testing.T) {
			t.Setenv("CI", "false")

			Color().Bold(Color().Blue("Bold & Blue").String()).Print()
		})
	})

	t.Run("Italic", func(t *testing.T) {
		t.Setenv("CI", "false")

		Color().Italic("Italic").Print()

		t.Run("Red", func(t *testing.T) {
			t.Setenv("CI", "false")

			Color().Italic(Color().Red("Italic & Red").String()).Print()
		})

		t.Run("Green", func(t *testing.T) {
			t.Setenv("CI", "false")

			Color().Italic(Color().Green("Italic & Green").String()).Print()
		})

		t.Run("Blue", func(t *testing.T) {
			t.Setenv("CI", "false")

			Color().Italic(Color().Blue("Italic & Blue").String()).Print()
		})
	})

	t.Run("Special", func(t *testing.T) {
		t.Setenv("CI", "false")

		Color().Bold(Color().Italic("Bold & Italic").String()).Print()

		t.Run("Red", func(t *testing.T) {
			t.Setenv("CI", "false")

			Color().Bold(Color().Italic(Color().Red("Bold & Italic & Red").String()).String()).Print()
		})

		t.Run("Green", func(t *testing.T) {
			t.Setenv("CI", "false")

			Color().Bold(Color().Italic(Color().Green("Bold & Italic & Green").String()).String()).Print()
		})

		t.Run("Blue", func(t *testing.T) {
			t.Setenv("CI", "false")

			Color().Bold(Color().Italic(Color().Blue("Bold & Italic & Blue").String()).String()).Print()
		})
	})
}
