<div align="center">

# wave-certify

**Command-line tool that validates broadcast hardware against WAVE protocol-correctness batteries (NDI, Dante, SRT, MoQ) and emits a signed certification artifact for partner-product listing.**

![kind](https://img.shields.io/badge/kind-cli-555?style=flat-square) ![domain](https://img.shields.io/badge/domain-hardware-0a7?style=flat-square) ![lang](https://img.shields.io/badge/lang-Go-00add8?style=flat-square) ![visibility](https://img.shields.io/badge/visibility-public-brightgreen?style=flat-square) ![license](https://img.shields.io/badge/license-Apache--2.0-blue?style=flat-square) ![phase](https://img.shields.io/badge/phase-alpha-blue?style=flat-square)

[wave.online](https://wave.online) · [Docs](https://docs.wave.online) · [github](https://github.com/wave-av/wave-certify)

</div>

---

## Status — early / scaffold

The CLI structure and command surface exist and build, but the per-protocol batteries are not yet
implemented: `check` currently exits with a `not yet implemented` error, and `version` reports
`wave-certify 0.0.0-scaffold`. Batteries are tracked as future work in this repo's roadmap
(`capabilities.json` marks the package `alpha`, v0.1.0). This is the Hardware-layer (plane layer 4) tool
of WAVE's Protocol Plane.

## Build

Requires Go 1.23+. No external module dependencies.

```bash
go build -o wave-certify ./cmd/wave-certify
```

## Usage (planned surface)

```bash
wave-certify check --target <addr> --protocol <ndi|dante|srt|moq>
wave-certify submit <artifact.json>
wave-certify version
```

`check` will run the named protocol battery (NDI frame integrity, Dante clock sync, SRT round-trip
latency, MoQ track parity) against a target and emit a signed certification artifact; `submit` will post
that artifact for partner listing. Both are wired into the CLI but return a not-implemented error today
— only `version` is functional (see Status).

## Repo layout

| Path | What it is |
| --- | --- |
| `cmd/wave-certify/main.go` | The CLI entry point |
| `threat-model.md` | Threat model for the certification pipeline |
| `SECRETS.md` | How signing secrets are handled |
| `capabilities.json` | Machine-readable lifecycle metadata |

## See also

- [threat-model.md](threat-model.md) · [SECURITY.md](SECURITY.md) · [CONTRIBUTING.md](CONTRIBUTING.md)

## License

Apache-2.0 — see [LICENSE](LICENSE).

---

<div align="center">

**Built by [WAVE Online, LLC](https://wave.online)** · [wave.online](https://wave.online) · [Docs](https://docs.wave.online)

</div>
