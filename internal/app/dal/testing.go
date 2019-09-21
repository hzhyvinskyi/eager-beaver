package dal

import (
	"fmt"
	"strings"
	"testing"
)

// TestStore ...
func TestStore(t *testing.T, databaseURL string) (*Dal, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseURL
	d := New(config)

	if err := d.Open(); err != nil {
		t.Fatal(err)
	}

	return d, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := d.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		d.Close()
	}
}
