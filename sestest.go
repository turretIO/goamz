package main

import (
  "github.com/turret-io/goamz/aws"
  "github.com/turret-io/goamz/ses"
  "fmt"
)

func main() {
  auth, err := aws.EnvAuth()
  msg := `
  Sender: Turret.IO Team <team@turret.io>
  Message-Id: <cb1663c8-b8e5-11e3-b2e2-1231380b9a0c.a8fd65c9-cb83-4121-42a5-5d52c49ba14b@trs1.turret-io.com>
  To: <tim@g.tdinternet.com>
  Reply-To: Turret.IO Team <team@turret.io>
  List-Unsubscribe: <mailto:unsubscribe-fe21fb2e4ba645e7918224ee71a104a7.tim=g.tdinternet.com@turret-io.com>, <http://go.turret.io/unsubscribe?v=ZmUyMWZiMmU0YmE2NDVlNzkxODIyNGVlNzFhMTA0YTd8dGltQGcudGRpbnRlcm5ldC5jb20=>
  Subject: =?utf-8?B?dGVzdA==?=
  Date: Thu, 03 Apr 2014 19:00:39 +0000
  From: Turret.IO Team <team@turret.io>
  Content-Type: multipart/alternative; boundary=70c59a4bceb23170bc679875247bb0840b9f92b96a1391820d094d81e4ec
  Dkim-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=turret.io; s=turret; t=1396465119; bh=nfxPezNf/kkGDoXz2CBIggwoMHMpdEH+6MzH9+dNAIs=; h=Reply-To:List-Unsubscribe:Subject:From:Content-Type:Sender:To; b=0MCi2LYdRzXCyJ1I85cJXgMU3tH3jbHmAXNj+5vd0sKdZ8fQ4N54gnQJ5eV4kAH09 Gzgs//v2+r3cShk3YDUg+M7dUcDV2ouM104kaSk7riMskDxDlK3y2o8Wx7QQibyxty 18H2e9YN2Fir4C7Zd6rEiethlWLLh9TQAZe5FvBQ=
  DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=turret-io.com; s=turret;
  t=1396465119; bh=nfxPezNf/kkGDoXz2CBIggwoMHMpdEH+6MzH9+dNAIs=;
  h=Sender:To:Reply-To:List-Unsubscribe:Subject:From:Content-Type;
  b=qoGUby6tYgJp3NPW4GGynQM4UpvSVpVg0x1WQg49Acurj9WblSRGHMqiB46sGkHFY
   habEP06SPWtqlPY4bsupFm45HXWX81bVtpj7SuXm0jdCyifK+AhFdk2dGP5TDlku9+
   vKSgfYnVsI7E7J2YicXRYl3X+80YcNJEUBYvexOI=

  --70c59a4bceb23170bc679875247bb0840b9f92b96a1391820d094d81e4ec
  Content-Type: text/html; charset=UTF-8
  Content-Transfer-Encoding: quoted-printable

  <p>body</p><p style=3D"text-align:center; background-color:#f5f5f5; padding=
  :8px;"><a href=3D"http://go.turret.io/unsubscribe?v=3DZmUyMWZiMmU0YmE2NDVlN=
  zkxODIyNGVlNzFhMTA0YTd8dGltQGcudGRpbnRlcm5ldC5jb20=3D">Unsubscribe</a></p><=
  img src=3D"http://go.turret.io/click/Y2IxNjYzYzgtYjhlNS0xMWUzLWIyZTItMTIzMT=
  M4MGI5YTBjfHRpbUBnLnRkaW50ZXJuZXQuY29t.gif">
  --70c59a4bceb23170bc679875247bb0840b9f92b96a1391820d094d81e4ec--
  `
  if err != nil {
    panic(err)
  }

  email := ses.New(auth, aws.USEast)
  resp, err := email.SendRawEmail(msg)
  if err != nil {
    panic(err)
  }
  fmt.Println(resp)
}
