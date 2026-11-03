package fontawesome_test

import (
	"io/fs"
	"testing"

	"github.com/go-app-blazar/fontawesome-font"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmbedded(t *testing.T) {
	files, err := fs.Glob(fontawesome.FS, "*/*")
	require.NoError(t, err)
	assert.NotEmpty(t, files)
	assert.Equal(t, []string{"css/all.css", "css/all.min.css", "css/brands.css", "css/brands.min.css", "css/fontawesome.css", "css/fontawesome.min.css", "css/regular.css", "css/regular.min.css", "css/solid.css", "css/solid.min.css", "css/svg-with-js.css", "css/svg-with-js.min.css", "css/svg.css", "css/svg.min.css", "css/v4-font-face.css", "css/v4-font-face.min.css", "css/v4-shims.css", "css/v4-shims.min.css", "css/v5-font-face.css", "css/v5-font-face.min.css", "webfonts/fa-brands-400.woff2", "webfonts/fa-regular-400.woff2", "webfonts/fa-solid-900.woff2", "webfonts/fa-v4compatibility.woff2"}, files)
}
