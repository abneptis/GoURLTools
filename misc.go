package urltools

import "http"
import "os"
import "strings"

// Returns a "host:port" string from a URL for known schemes, even
// if the URL omits this (by returning the service default)
// NB: This should, but has not been tested to work with, raw IPv6 addresses with
// zero-compaction.
func ExtractURLHostPort(url *http.URL)(hostport string, err os.Error){
  if url == nil || url.Host == "" { return "", os.NewError("Invalid URL (or empty Host)") }
  portIdx := strings.LastIndex(url.Host, ":")
  var host,port string
  if portIdx >= 0 {
    host = url.Host[0:portIdx]
    port = url.Host[portIdx+1:]
  } else {
    host = url.Host
  }
  if port == "" {
    switch url.Scheme {
      case "http": port = "80"
      case "https": port = "443"
      case "ftp": port = "21"
      case "smtp": port = "25"
      default:
        err = os.NewError("No port specified, and unknown scheme: " + url.Scheme)
    }
  }
  hostport = host + ":" + port
  return
}

// Appends the parameters to the URL, escaping the key and value strings.
// If the URL already has parameters, an & will be inserted.
//
// Trailing &'s in the inurl will cause duplicate &'s in the outurl.
func AppendParams(inurl http.URL, parms map[string][]string)(outurl http.URL){
  outurl = inurl
  pstr := http.EncodeQuery(parms)
  if len(pstr) > 0 {
    if len(outurl.RawQuery) > 0 { outurl.RawQuery += "&" }
    outurl.RawQuery += pstr
  }
  return
}
