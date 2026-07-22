<a id="readme-top"></a>

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Apache 2.0 License][license-shield]][license-url]

<!-- PROJECT HEADER -->
<br />
<div align="center">
  <h1 align="center">openaifx</h1>

  <p align="center">
    Uber Fx wrapper for the <a href="https://github.com/openai/openai-go">OpenAI Go API</a>
    <br />
    <a href="https://pkg.go.dev/github.com/go-core-fx/openaifx">Explore the docs</a>
    &middot;
    <a href="https://github.com/go-core-fx/openaifx/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    &middot;
    <a href="https://github.com/go-core-fx/openaifx/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
- [About The Project](#about-the-project)
	- [Built With](#built-with)
- [Getting Started](#getting-started)
	- [Prerequisites](#prerequisites)
	- [Installation](#installation)
- [Usage](#usage)
	- [Configuration](#configuration)
	- [Environment Variables](#environment-variables)
	- [Integration with Uber Fx](#integration-with-uber-fx)
	- [Injecting the Client](#injecting-the-client)
	- [Using with a Custom HTTP Client](#using-with-a-custom-http-client)
	- [Using with a Custom Base URL](#using-with-a-custom-base-url)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)
- [Acknowledgments](#acknowledgments)


<!-- ABOUT THE PROJECT -->
## About The Project

`openaifx` is an [Uber Fx](https://uber.github.io/fx/) module that provides the [OpenAI Go SDK](https://github.com/openai/openai-go) client via dependency injection. It handles client construction, configuration validation, and named logger integration so you can focus on building features.

Key features:

- **Zero-boilerplate client setup** -- supply a `Config` and get a ready-to-use `*openai.Client`
- **Custom API endpoint** -- point `BaseURL` at proxies, mocks, or OpenAI-compatible services
- **Custom HTTP client** -- inject your own `*http.Client` for timeouts, TLS, or transport tweaks
- **Named logger** -- automatically decorates the Fx logger with the `openaifx` component name
- **Consistent conventions** -- follows the same patterns as `openrouterfx`, `sqlfx`, and other `go-core-fx` modules

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

| Technology                                           | Badge                                      |
| ---------------------------------------------------- | ------------------------------------------ |
| [Go](https://go.dev/)                                | [![Go][Go.dev]][Go-url]                    |
| [OpenAI Go SDK](https://github.com/openai/openai-go) | [![OpenAI SDK][OpenAI-shield]][OpenAI-url] |
| [Uber Fx](https://uber.github.io/fx/)                | [![Uber Fx][UberFx-shield]][UberFx-url]    |
| [Uber Zap](https://github.com/uber-go/zap)           | [![Uber Zap][Zap-shield]][Zap-url]         |

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

- **Go 1.25+**
- An [OpenAI API key](https://platform.openai.com/api-keys)

### Installation

```sh
go get github.com/go-core-fx/openaifx
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE -->
## Usage

### Configuration

The `Config` struct controls client construction:

| Field        | Type           | Required | Description                                       |
| ------------ | -------------- | -------- | ------------------------------------------------- |
| `APIKey`     | `string`       | Yes      | Your OpenAI API key                               |
| `BaseURL`    | `string`       | No       | Custom API endpoint (e.g. a proxy or mock server) |
| `HTTPClient` | `*http.Client` | No       | Custom HTTP client for transport configuration    |

### Environment Variables

`Config.APIKey` is required. `New` rejects an empty API key before the SDK can fall back to `OPENAI_API_KEY`, so the environment variable alone is not sufficient through this wrapper.

### Integration with Uber Fx

```go
package main

import (
	"go.uber.org/fx"

	"github.com/go-core-fx/openaifx"
	"github.com/openai/openai-go/v3"
)

func main() {
	fx.New(
		// Provide the OpenAI configuration
		fx.Supply(openaifx.Config{
			APIKey:  "sk-...",
			BaseURL: "", // leave empty to use the default OpenAI endpoint
		}),

		// Register the openaifx module
		openaifx.Module(),

		// Your application modules...
		fx.Invoke(func(client *openai.Client) {
			// Use client here
		}),
	).Run()
}
```

### Injecting the Client

Once the module is registered, inject `*openai.Client` into any constructor or invoke:

```go
func NewService(client *openai.Client) *Service {
	return &Service{client: client}
}
```

### Using with a Custom HTTP Client

```go
import (
	"net/http"
	"time"
)

fx.Supply(openaifx.Config{
	APIKey:  "sk-...",
	HTTPClient: &http.Client{
		Timeout: 60 * time.Second,
	},
})
```

### Using with a Custom Base URL

Useful for testing with [wiremock](https://wiremock.org/) or compatible OpenAI-mocking services:

```go
fx.Supply(openaifx.Config{
	APIKey:  "sk-...",
	BaseURL: "http://localhost:8080/v1/",
})
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->
## Roadmap

- [x] Basic client construction from `Config`
- [x] Custom `BaseURL` support
- [x] Custom `*http.Client` injection
- [x] Named logger integration via `go-core-fx/logger`
- [ ] Lifecycle health check (ping on application start)
- [ ] Multiple named client instances

See the [open issues](https://github.com/go-core-fx/openaifx/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the Apache 2.0 License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

go-core-fx -- [https://github.com/go-core-fx/openaifx](https://github.com/go-core-fx/openaifx)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [OpenAI Go SDK](https://github.com/openai/openai-go) -- the official Go client for the OpenAI API
* [Uber Fx](https://uber.github.io/fx/) -- dependency injection framework for Go
* [go-core-fx/logger](https://github.com/go-core-fx/logger) -- structured Zap logger for Fx modules
* [Best-README-Template](https://github.com/othneildrew/Best-README-Template) -- this README template

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
[contributors-shield]: https://img.shields.io/github/contributors/go-core-fx/openaifx.svg?style=for-the-badge
[contributors-url]: https://github.com/go-core-fx/openaifx/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/go-core-fx/openaifx.svg?style=for-the-badge
[forks-url]: https://github.com/go-core-fx/openaifx/network/members
[stars-shield]: https://img.shields.io/github/stars/go-core-fx/openaifx.svg?style=for-the-badge
[stars-url]: https://github.com/go-core-fx/openaifx/stargazers
[issues-shield]: https://img.shields.io/github/issues/go-core-fx/openaifx.svg?style=for-the-badge
[issues-url]: https://github.com/go-core-fx/openaifx/issues
[license-shield]: https://img.shields.io/github/license/go-core-fx/openaifx.svg?style=for-the-badge
[license-url]: https://github.com/go-core-fx/openaifx/blob/master/LICENSE
[Go.dev]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/
[OpenAI-shield]: https://img.shields.io/badge/OpenAI-412991?style=for-the-badge&logo=openai&logoColor=white
[OpenAI-url]: https://github.com/openai/openai-go
[UberFx-shield]: https://img.shields.io/badge/Uber_Fx-276DC1?style=for-the-badge&logo=uber&logoColor=white
[UberFx-url]: https://uber.github.io/fx/
[Zap-shield]: https://img.shields.io/badge/Uber_Zap-000000?style=for-the-badge&logo=uber&logoColor=white
[Zap-url]: https://github.com/uber-go/zap
