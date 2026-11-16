//go:build spot

package display

import (
	"bytes"
	"image"
	"path/filepath"
	"testing"
	"time"

	"github.com/HuskerMinion/techo5/echod/internal/config"
)

func TestSpotSunriseDateFollowsScreenLanguage(t *testing.T) {
	config.Use(filepath.Join(t.TempDir(), "state.json"))
	t.Cleanup(func() { _ = config.Set().Screen().Language("") })

	now := time.Date(2026, 9, 16, 14, 7, 0, 0, time.Local)
	render := func(lang string) []byte {
		t.Helper()
		if err := config.Set().Screen().Language(lang); err != nil {
			t.Fatal(err)
		}
		dst := image.NewRGBA(image.Rect(0, 0, side, side))
		newRoundRenderer(dst).sunriseFace(roundScene{now: now}, 0.5, false)
		return bytes.Clone(dst.Pix)
	}

	if english, french := render("en"), render("fr"); bytes.Equal(english, french) {
		t.Error("the Spot sunrise date did not change with the screen language")
	}
}
