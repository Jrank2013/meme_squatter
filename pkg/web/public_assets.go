package web

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _PublicAssets2159de1879639c2e114495ccb0e88cb9e884622a = "body {\n    display: flex;\n    flex-direction: column;\n    justify-content: center;\n    align-items: center;\n}"

// PublicAssets returns go-assets FileSystem
var PublicAssets = assets.NewFileSystem(map[string][]string{"/": []string{}, "/css": []string{"index.css"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1729040259, 1729040259792859735),
		Data:     nil,
	}, "/css": &assets.File{
		Path:     "/css",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1729040211, 1729040211413953077),
		Data:     nil,
	}, "/css/index.css": &assets.File{
		Path:     "/css/index.css",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1729040308, 1729040308895760135),
		Data:     []byte(_PublicAssets2159de1879639c2e114495ccb0e88cb9e884622a),
	}}, "")
