package web

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets14fe4559026d4c5b5eb530ee70300c52d99e70d7 = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n    <link rel=\"stylesheet\" href=\"/public/css/index.css\"/>\n    <script src=\"https://unpkg.com/htmx.org@2.0.3\"></script>\n    <title>Welcome!</title>\n</head>\n<body>\n    <h1>Welcome!</h1>\n    <p>This probably ins't where you want to be. Enjoy the meme while your here!</p>\n    <p>You'll be redirected to <a href=\"https://memcached.org\">https://memcached.org</a> in <span id=\"time\">5</span>...</p>\n    <img src=\"{{.Meme}}\" alt=\"Programming Meme\"/>\n\n\n    <script>\n        let element = document.getElementById(\"time\")\n\n        setTimeout(() => {\n            window.location = \"https://memcached.org\"\n        },5000)\n\n        setInterval(() => {\n            let current = parseInt(element.innerText)\n            if (current === 0){\n                return\n            }\n            element.innerText = current - 1\n\n        },1000)\n\n    </script>\n</body>\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"index.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1729040255, 1729040255041415801),
		Data:     nil,
	}, "/index.html": &assets.File{
		Path:     "/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1729130303, 1729130303687776658),
		Data:     []byte(_Assets14fe4559026d4c5b5eb530ee70300c52d99e70d7),
	}}, "")
