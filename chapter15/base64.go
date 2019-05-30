package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	fmt.Println(base64.StdEncoding.EncodeToString([]byte(`I\xd7*/}j\x92\x07\xb9\x80\xf5,zm\x7fW\xe2\x1f\xde*\xbe\x16\xfe\xf0qa>\xa7\xd3\x1e\x17k`)))
	fmt.Println(base64.URLEncoding.EncodeToString([]byte(`I\xd7*/}j\x92\x07\xb9\x80\xf5,zm\x7fW\xe2\x1f\xde*\xbe\x16\xfe\xf0qa>\xa7\xd3\x1e\x17k`)))

	fmt.Println(base64.RawStdEncoding.EncodeToString([]byte(`I\xd7*/}j\x92\x07\xb9\x80\xf5,zm\x7fW\xe2\x1f\xde*\xbe\x16\xfe\xf0qa>\xa7\xd3\x1e\x17k`)))
	fmt.Println(base64.RawURLEncoding.EncodeToString([]byte(`I\xd7*/}j\x92\x07\xb9\x80\xf5,zm\x7fW\xe2\x1f\xde*\xbe\x16\xfe\xf0qa>\xa7\xd3\x1e\x17k`)))

	fmt.Println(base64.StdEncoding.DecodeString("SVx4ZDcqL31qXHg5Mlx4MDdceGI5XHg4MFx4ZjUsem1ceDdmV1x4ZTJceDFmXHhkZSpceGJlXHgxNlx4ZmVceGYwcWE+XHhhN1x4ZDNceDFlXHgxN2s="))
	fmt.Println(base64.URLEncoding.DecodeString("SVx4ZDcqL31qXHg5Mlx4MDdceGI5XHg4MFx4ZjUsem1ceDdmV1x4ZTJceDFmXHhkZSpceGJlXHgxNlx4ZmVceGYwcWE-XHhhN1x4ZDNceDFlXHgxN2s="))

	fmt.Println(base64.RawStdEncoding.DecodeString("SVx4ZDcqL31qXHg5Mlx4MDdceGI5XHg4MFx4ZjUsem1ceDdmV1x4ZTJceDFmXHhkZSpceGJlXHgxNlx4ZmVceGYwcWE+XHhhN1x4ZDNceDFlXHgxN2s"))
	fmt.Println(base64.RawURLEncoding.DecodeString("SVx4ZDcqL31qXHg5Mlx4MDdceGI5XHg4MFx4ZjUsem1ceDdmV1x4ZTJceDFmXHhkZSpceGJlXHgxNlx4ZmVceGYwcWE-XHhhN1x4ZDNceDFlXHgxN2s"))

}
