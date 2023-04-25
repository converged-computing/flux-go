# Flux Go Bindings

This is a small example of interacting with Flux Framework to submit jobs (and other simple tasks)
using cgo! To make understanding most easy, we package the commands each into a separate client
command.

![docs/img/flux-go-banner.png](docs/img/flux-go-banner.png)

🚧️ Under Development! 🚧️

Read our 🌈️ [Early Documentation](docs) 🌈️ to get started!

## Examples

 - [submit](cmd/submit/main.go): Submit a job! Made possible by way of [this example](https://gist.github.com/grondo/6a51a43cb62c2a30c1cf74d167ddb421) from [grondo](https://github.com/grondo)!
 - [keygen](cmd/keygen/main.go): Use zeromq to generate a curve certificate.

## License

HPCIC DevTools is distributed under the terms of the MIT license.
All new contributions must be made under this license.

See [LICENSE](https://github.com/converged-computing/cloud-select/blob/main/LICENSE),
[COPYRIGHT](https://github.com/converged-computing/cloud-select/blob/main/COPYRIGHT), and
[NOTICE](https://github.com/converged-computing/cloud-select/blob/main/NOTICE) for details.

SPDX-License-Identifier: (MIT)

LLNL-CODE- 842614
