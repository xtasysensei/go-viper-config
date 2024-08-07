package loadcfg

import (
	"testing"

	"github.com/spf13/viper"
)

func TestLoadCfg(t *testing.T) {
	assertError := func(t testing.TB, got error, want string) {
		t.Helper()

		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("LoadTomlCfg", func(t *testing.T) {
		viper.Reset()
		_, err := LoadTomlCfg()
		assertError(t, err, "open config file: config/config.toml: no such file or directory")
	})
}
