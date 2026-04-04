package bot

import (
	"net/http"
	"net/url"
)

// bypass Russia
// requires running Proxy
var proxy, _ = url.Parse("http://127.0.0.1:10808")
var ProxyClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}}
