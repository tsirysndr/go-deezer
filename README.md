<h1 align="center">Welcome to go-deezer 👋 *** Work In Progress ***</h1>
<p align="center">
  <a href="https://github.com/tsirysndr/go-spotify/commits/master">
    <img src="https://img.shields.io/github/last-commit/tsirysndr/go-deezer" target="_blank" />
  </a>
  <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/tsirysndr/go-deezer">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/tsirysndr/go-deezer">
  <img alt="GitHub closed pull requests" src="https://img.shields.io/github/issues-pr-closed-raw/tsirysndr/go-deezer">
  <img alt="GitHub pull requests" src="https://img.shields.io/github/issues-pr/tsirysndr/go-deezer">
  <img alt="GitHub issues" src="https://img.shields.io/github/issues/tsirysndr/go-deezer">
  <img alt="GitHub contributors" src="https://img.shields.io/github/contributors/tsirysndr/go-deezer">
  <a href="https://github.com/tsirysndr/go-deezer/blob/master/LICENSE">
    <img alt="License: BSD" src="https://img.shields.io/badge/license-BSD-green.svg" target="_blank" />
  </a>
  <a href="https://twitter.com/tsiry_sndr">
    <img alt="Twitter: tsiry_sndr" src="https://img.shields.io/twitter/follow/tsiry_sndr.svg?style=social" target="_blank" />
  </a>
</p>

go-deezer is a Go client library for accessing the [Deezer API](https://developers.deezer.com/api)

## 🚚 Install

```sh
go get github.com/tsirysndr/go-deezer
```

## 🚀 Usage

Import the package into your project.

```Go
import "github.com/tsirysndr/go-deezer"
```

Construct a new Deezer client, then use the various services on the client to access different parts of the Deezer API. For example:

```Go
client := deezer.NewClient()
res, _ := client.Artist.Get("27")
artist, _ := json.Marshal(res)
fmt.Println(string(artist))
```

## Author

👤 **Tsiry Sandratraina**

* Twitter: [@tsiry_sndr](https://twitter.com/tsiry_sndr)
* Github: [@tsirysndr](https://github.com/tsirysndr)

## Show your support

Give a ⭐️ if this project helped you!

