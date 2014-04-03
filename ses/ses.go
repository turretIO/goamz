package ses

import (
  "encoding/xml"
  "encoding/base64"
  "github.com/turretIO/goamz/aws"
  "net/url"
  "net/http"
  "strconv"
  "strings"
  "time"
  "fmt"
)
const timeLayout = "Mon, 2 Jan 2006 15:04:05 GMT"

var b64 = base64.StdEncoding

type SES struct {
  aws.Auth
  aws.Region
}

func New(auth aws.Auth, region aws.Region) *SES {
  return &SES{auth, region}
}

func (ses *SES) postQuery(params map[string]string, resp interface{}) error {
  endpoint, err := url.Parse(ses.SESEndpoint)
  if err != nil {
    return err
  }
  params["Version"] = "2010-12-01"
  params["Date"] = time.Now().In(time.UTC).Format(timeLayout)
  params["Timestamp"] = time.Now().In(time.UTC).Format(time.RFC3339)
  params["AWSAccessKeyId"] = ses.AccessKey
  sign(ses.Auth, "POST", "/", params, endpoint.Host)
  encoded := multimap(params).Encode()
  body := strings.NewReader(encoded)
  req, err := http.NewRequest("POST", endpoint.String(), body)
  fmt.Println(body)
  if err != nil {
    return err
  }
  req.Header.Set("Host", endpoint.Host)
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Set("Date", params["Date"])
  req.Header.Set("X-Amzn-Authorization", fmt.Sprintf("AWS3-HTTPS AWSAccessKeyId=%s, Algorithm=HmacSHA256, Signature=%s, SignedHeaders=Date", ses.AccessKey, params["Signature"]))

  r, err := http.DefaultClient.Do(req)
  if err != nil {
    return err
  }
  defer r.Body.Close()
  if r.StatusCode > 200 {
    return buildError(r)
  }
  return xml.NewDecoder(r.Body).Decode(resp)
}

func (ses *SES) SendRawEmail(email string) (*SimpleResp, error) {
  // Base64 encode the email
  rawEmail := b64.EncodeToString([]byte(email))

  params := map[string]string {
    "Action":"SendRawEmail",
    "RawMessage.Data":string(rawEmail),
  }
  resp := new(SimpleResp)
  if err := ses.postQuery(params, resp); err != nil {
    return nil, err
  }
  return resp, nil
}

func multimap(p map[string]string) url.Values {
  q := make(url.Values, len(p))
  for k, v := range p {
    q[k] = []string{v}
  }
  return q
}

func buildError(r *http.Response) error {
  var (
    err    Error
    errors xmlErrors
  )
  xml.NewDecoder(r.Body).Decode(&errors)
  if len(errors.Errors) > 0 {
    err = errors.Errors[0]
  }
  err.StatusCode = r.StatusCode
  if err.Message == "" {
    err.Message = r.Status
  }
  return &err
}

type xmlErrors struct {
  Errors []Error `xml:"Error"`
}

type SimpleResp struct {
  RequestId string `xml:"ResponseMetadata>RequestId"`
}

// Error encapsulates an IAM error.
type Error struct {
  // HTTP status code of the error.
  StatusCode int

  // AWS code of the error.
  Code string

  // Message explaining the error.
  Message string
}

func (e *Error) Error() string {
  var prefix string
  if e.Code != "" {
    prefix = e.Code + ": "
  }
  if prefix == "" && e.StatusCode > 0 {
    prefix = strconv.Itoa(e.StatusCode) + ": "
  }
  return prefix + e.Message
}
