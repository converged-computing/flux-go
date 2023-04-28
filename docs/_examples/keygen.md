---
title: KeyGen
layout: snippet-page
description: How to generate a curve certificate using zeromq
---

# KeyGen Example

Here is an example of how to generate a curve certificate using Flux. This
technically uses zeromq bindings, and Flux isn't required. Each of the files 
below is provided in completion, meaning that they include all of the headers that are necessary
to build with Flux.

## 1. Write a KeyGen Function

This example is fairly simple, because we just need one function to create and
then either save our certificate to file, or return it as a string.

{% include snippet.html language="go" id="keygengo" path="pkg/flux/keygen.go" %}

For an example that uses just zeromq (without Flux) you can [look at the Flux Operator](https://github.com/flux-framework/flux-operator/blob/main/controllers/flux/keygen.go.template). Also note that the keygen example doesn't require a flux instance to be running! 

## 2. Call the function!


Here is what we might then have for our command entrypoint:

{% include snippet.html language="go" id="maingo" path="cmd/keygen/main.go" %}

And this is what we see on the command line:

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
```
