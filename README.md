# Flux Go Bindings

This is a small example of interacting with Flux Framework to submit jobs (and other simple tasks)
using cgo! To make understanding most easy, we package the commands each into a separate client
command.

![docs/assets/img/flux-go-banner.png](docs/assets/img/flux-go-banner.png)

🚧️ Under Development! 🚧️

Read our 🌈️ [Early Documentation](https://converged-computing.github.io/flux-go) 🌈️ to get started!

## Examples

 - [submit](cmd/submit/main.go): Submit a job! Made possible by way of [this example](https://gist.github.com/grondo/6a51a43cb62c2a30c1cf74d167ddb421) from [grondo](https://github.com/grondo)!
 - [keygen](cmd/keygen/main.go): Use zeromq to generate a curve certificate.

### Adding an Example

You can add your examples anywhere in the codebase here (e.g., typically we have some logic in
[pkg](pkg) and a command in [cmd](cmd). We ask that:

 - You add a test for your contribution that will be run in our development environment with Flux. See [test](test) for examples.
 - That you add a new example page in [docs/_examples](docs/_examples).
 
Within each rendered example, we render the content directly from GitHub, meaning you include
the relative path to your file via a Jekyll include:

```
{% include snippet.html language="go" id="keygengo" path="pkg/flux/keygen.go" %}
```

This means that you can largely copy another page and it will show up on the site as a new
example! The table of contents in [docs/_data/toc.yml](docs/_data/toc.yml) also needs to be updated.
With this simple approach, we have example here that are both tested and render live in the
web interface. Thank you!

## License

HPCIC DevTools is distributed under the terms of the MIT license.
All new contributions must be made under this license.

See [LICENSE](https://github.com/converged-computing/cloud-select/blob/main/LICENSE),
[COPYRIGHT](https://github.com/converged-computing/cloud-select/blob/main/COPYRIGHT), and
[NOTICE](https://github.com/converged-computing/cloud-select/blob/main/NOTICE) for details.

SPDX-License-Identifier: (MIT)

LLNL-CODE- 842614
