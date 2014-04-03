package ses

import (
  "crypto/hmac"
  "crypto/sha256"
  "github.com/turret-io/goamz/aws"
  "sort"
  "strings"
  "fmt"
)

// ----------------------------------------------------------------------------
// Version 2 signing (http://goo.gl/RSRp5)

func sign(auth aws.Auth, method, path string, params map[string]string, host string) {
  params["AWSAccessKeyId"] = auth.AccessKey
  params["SignatureVersion"] = "3"
  params["SignatureMethod"] = "HmacSHA256"

  var sarray []string
  for k, v := range params {
    if k != "Date" {
      continue
    }
    fmt.Println(fmt.Sprintf("signing %s", k))
    sarray = append(sarray, v)
  }
  sort.StringSlice(sarray).Sort()
  joined := strings.Join(sarray, "&")
  payload := joined
  fmt.Println("sign...")
  fmt.Println(payload)
  hash := hmac.New(sha256.New, []byte(auth.SecretKey))
  hash.Write([]byte(payload))
  signature := make([]byte, b64.EncodedLen(hash.Size()))
  b64.Encode(signature, hash.Sum(nil))

  params["Signature"] = string(signature)
}
