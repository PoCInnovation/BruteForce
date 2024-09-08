# BrutForce Next-Gen

BrutForcer written in go.

## How does it work?

Using advanced wordlists to gather information on web servers and perform brute force attacks.

## Getting Started

### Installation

[Golang install guide](https://go.dev/doc/install)

Clone the repository and simply execute:

```shell
make
```

or

```shell
go build
```

to build the executable `bruteforce`.

### Quickstart

To run the program directly, execute:

```shell
go run src/main.go
```

### Usage

```bash
./bruteforce [OPTIONS]
```

### Matching

For matching usage, the following flags are available:

`-status-codes` : match based on a list of status codes.

For example, `./bruteforce -status-codes="200,201,202,401,404"`.

*By default* : 200, 401, 403, 404, 429, 500

`-header` : match based on a header.

For example, `./bruteforce -header="Content-Type: application/json"`.

To match multiple headers, use commas to separate each querie. Specify `ALL` if you wish to have all matches be true, if not don't add the `ALL`. As so:

- should match all of the headers:

`./bruteforce -header="ALL,Content-Type: application/json,Content-Type: text/css"`

- match on any of the headers:

`./bruteforce -header="Content-Type: application/json,Content-Type: text/css"`

`-body` : match based on a body.

For example, `./bruteforce -body="Hello World"`.

Same applies the body for multiple queries of strings in the body as the header.

## Get involved

You're invited to join this project ! Check out the [contributing guide](./CONTRIBUTING.md).

If you're interested in how the project is organized at a higher level, please contact the current project manager.

## Our PoC team ❤️

Developers
| [<img src="https://github.com/louonezime.png?size=85" width=85><br><sub>Lou Onezime</sub>](https://github.com/louonezime) | [<img src="https://github.com/57ave.png?size=85" width=85><br><sub>Gustave Delecroix</sub>](https://github.com/57ave) | [<img src="https://github.com/SIMLUKE.png?size=85" width=85><br><sub>Luc Simon</sub>](https://github.com/SIMLUKE)
| :---: | :---: | :---: |

Manager
| [<img src="https://github.com/adamdeziri.png?size=85" width=85><br><sub>Adam Deziri</sub>](https://github.com/adamdeziri)
| :---: |

<h2 align=center>
Organization
</h2>

<p align='center'>
    <a href="https://www.linkedin.com/company/pocinnovation/mycompany/">
        <img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white" alt="LinkedIn logo">
    </a>
    <a href="https://www.instagram.com/pocinnovation/">
        <img src="https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white" alt="Instagram logo"
>
    </a>
    <a href="https://twitter.com/PoCInnovation">
        <img src="https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white" alt="Twitter logo">
    </a>
    <a href="https://discord.com/invite/Yqq2ADGDS7">
        <img src="https://img.shields.io/badge/Discord-7289DA?style=for-the-badge&logo=discord&logoColor=white" alt="Discord logo">
    </a>
</p>
<p align=center>
    <a href="https://www.poc-innovation.fr/">
        <img src="https://img.shields.io/badge/WebSite-1a2b6d?style=for-the-badge&logo=GitHub Sponsors&logoColor=white" alt="Website logo">
    </a>
</p>

> 🚀 Don't hesitate to follow us on our different networks, and put a star 🌟 on `PoC's` repositories

> Made with ❤️ by PoC
