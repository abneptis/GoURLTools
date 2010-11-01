// URL parsing and manipulation functions.
//
// URL tools provides various escaping and manipulation functions
// that may be useful for those implementing network protocols
//
package urltools

func Rfc3986Reserved(i byte)(out bool){
 switch i {
  case ':','/','?','#','[',']','@','!','$','&','\'',
       '(',')','*','+',',',';','=':
    out = true
  default:
    out = false
 }
 return
}

// Returns the hex encoded (lower-case) value of byte.
func HexEncodeLower(b byte)(out string){
  hb := []byte{'0','1','2','3','4','5','6','7','8','9','a','b','c','d','e','f'}
  out = string([]byte {hb[(b & 0xf0) >> 4], hb[(b & 0x0f)]})
  return
}

// Returns the hex encoded (upper-case) value of byte.
func HexEncodeUpper(b byte)(out string){
  hb := []byte{'0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F'}
  out = string([]byte {hb[(b & 0xf0) >> 4], hb[(b & 0x0f)]})
  return
}

// Returns the hex encoded (upper-case) value of byte with a % prepended.
func PercentUpper(b byte)(out string){
  return "%" + HexEncodeUpper(b)
}

// Returns the hex encoded (lower-case) value of byte with a % prepended.
func PercentLower(b byte)(out string){
  return "%" + HexEncodeLower(b)
}

// Escape the string as follows: if etf(b) returns true, use the value generated
// by ef(b), else use b.
func Escape(in string, etf func(byte)(bool), ef func(byte)(string))(out string){
  for i := range(in){
    if etf(in[i]) {
      out += ef(in[i])
    } else {
      out += string(in[i])
    }
  }
  return
}
