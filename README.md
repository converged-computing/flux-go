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

## Usage

### Submit

You'll want to build first (ideally in your Development container afforded by VSCode and the [.devcontainer](.devcontainer))

```bash
$ make
```

And then start a flux instance:

```bash
$ flux start --test-size=4
```

And run the demo!

```bash
$ ./bin/fluxgo-submit 
```
```console
⭐️ Testing flux submit in Go! ⭐️
Submitting a Sleep Job: sleep 10
Flux Future: &{{{}}}
```

And see it running!

```bash
$ flux jobs -a
```
```console
       JOBID USER     NAME       ST NTASKS NNODES     TIME INFO
    ƒdcsg8Nj vscode   sleep       R      1      1   2.728s 5e3ccb811e04
```

Of course you would write your own Go code to do some kind of submission logic.

### KeyGen

The keygen example doesn't require a flux instance to be running! You can just generate a local `curve.cert`:

```bash
$ ./bin/fluxgo-keygen 
```
```console
⭐️ Testing flux keygen in Go! ⭐️
Saving to ./curve.cert
Generated certificate!
```

And view the certificate:

```bash
$ cat ./curve.cert 
```
```console
#   ****  Generated on 2023-04-25 22:04:45 by CZMQ  ****
#   ZeroMQ CURVE **Secret** Certificate
#   DO NOT PROVIDE THIS FILE TO OTHER USERS nor change its permissions.

metadata
    name = "curve-cert"
    keygen.hostname = "5e3ccb811e04"
curve
    public-key = "S%XP&s=}d*EQDZ!Vef-Q6OCjW+>I+cR.s*gzqwkl"
    secret-key = "17P&8}*hoc-sTHk6m}i3mL{OG=9j!o*9:8nmuJEM"
vscode@5e3ccb811e04:/workspaces/flux-go$ 
```

## License

HPCIC DevTools is distributed under the terms of the MIT license.
All new contributions must be made under this license.

See [LICENSE](https://github.com/converged-computing/cloud-select/blob/main/LICENSE),
[COPYRIGHT](https://github.com/converged-computing/cloud-select/blob/main/COPYRIGHT), and
[NOTICE](https://github.com/converged-computing/cloud-select/blob/main/NOTICE) for details.

SPDX-License-Identifier: (MIT)

LLNL-CODE- 842614
