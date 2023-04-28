---
layout: page
title: Flux Go
permalink: /
---

# Welcome to Flux Go

![assets/img/flux-go-banner.png](assets/img/flux-go-banner.png)

## Purpose

This is a starting point if you want to learn to interact with Flux from Go!
Since Flux Framework's [core](https://github.com/flux-framework/flux-core),
this means we can use [cgo](https://pkg.go.dev/cmd/cgo) to interact with
Flux (the external library) directly in Go. This repository provides
an example project for achieving this [in the root]({{ site.repo }})
and the examples are also shared in these pages.

See the {% include doc.html name="Getting Started" path="getting-started" %} guide for
setting up your development environment and basic interactions. If you'd like 
to request a feature or contribute? [Open an issue]({{ site.repo }}/issues)


## Examples

<ul>
{% for ex in site.examples %}
<li><a href="{{ site.baseurl }}/{{ ex.url }}">{{ ex.title }}</a>: {{ ex.description }}</li>
{% endfor %}
<ul>
