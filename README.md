<p align="center">
  <a href="" rel="noopener">
<!--  <img width=200px height=200px src="https://i.imgur.com/6wj0hh6.jpg" alt="Project logo"></a> -->
</p>

<h3 align="center">Golang QRCode Generator</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/zhixuanqwert/qrcode_generator_golang.svg)]([https://github.com/zhixuanqwert/qrcode_generator_golang/issues](https://github.com/zhixuanqwert/qrcode_generator_golang/issues))
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/zhixuanqwert/qrcode_generator_golang.svg)]([https://github.com/zhixuanqwert/qrcode_generator_golang/pulls](https://github.com/zhixuanqwert/qrcode_generator_golang/pulls))
<!--[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)-->

</div>

---

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Built Using](#built_using)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)
<!-- - [TODO](../TODO.md)
- [Contributing](../CONTRIBUTING.md) -->
## 🧐 About <a name = "about"></a>

qr code generator apis to generate qrcodes from a content

## 🏁 Getting Started <a name = "getting_started"></a>

clone the project and make sure necessary dependencies are installed

### Prerequisites

Create new folder resources/qrcode and assets before run the project

```
mkdir resources/qrcode
mkdir assets
``` 

<!-- ### Installing

A step by step series of examples that tell you how to get a development env running.

Say what the step will be

```
Give the example
```

And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo. -->

## 🔧 Running the project</a>

```
ENV=DEV PORT=8080 go run main.go
```
or using visual studio code debug

## 🎈 Usage <a name="usage"></a>

Examples
using postman
```
http://localhost:port/generate

{
 "content":"www.google.com",
    "size":200,
    "save":true,
    "filename":"My Test QR",
    "filetype": "png",
    "watermark": true
}
```

## ⛏️ Built Using <a name = "built_using"></a>

<!-- - [MongoDB](https://www.mongodb.com/) - Database -->
<!-- - [Gin](https://expressjs.com/) - Server Framework -->
- [Gin](https://gin-gonic.com/) - Web Framework
- [Go](https://go.dev/) - Server Environment

## ✍️ Authors <a name = "authors"></a>

- [@zhixuanqwert](https://github.com/zhixuanqwert) - Idea & Initial work

<!-- See also the list of [contributors](https://github.com/kylelobo/The-Documentation-Compendium/contributors) who participated in this project. -->

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- Barcode generation are based on boombuler/barcode
- Inspiration
Inspired by
https://www.twilio.com/en-us/blog/generate-qr-code-with-go
- References
https://github.com/boombuler/barcode
