# wave-certify

**WAVE Certify** is a command-line tool that validates broadcast hardware against WAVE protocol-correctness batteries (NDI, Dante, SRT, MoQ) and emits a signed certification artifact for partner-product listing. It is the Hardware-layer tool of the [WAVE Protocol Plane](https://github.com/wave-av/wave-foundation/blob/master/frameworks/protocol-plane/README.md).

Written in Go.

## Status

**Early / scaffold.** The CLI structure and command surface exist, but the per-protocol batteries are not yet implemented — `check` currently returns a not-implemented error. Batteries are Wave-1 work tracked in the repo roadmap.

## Build

Requires Go 1.23+. The project has no external module dependencies.

```bash
go build -o wave-certify ./cmd/wave-certify
```

## Usage (planned surface)

```bash
wave-certify check --target <addr> --protocol <ndi|dante|srt|moq>
wave-certify submit <artifact.json>
wave-certify version
```

`check` will run the protocol battery against the target and emit a signed certification artifact; `submit` will post that artifact for public listing. These commands are wired but not yet functional (see Status).

## See also

- [Protocol Plane framework](https://github.com/wave-av/wave-foundation/blob/master/frameworks/protocol-plane/README.md)
- [threat-model.md](threat-model.md) · [SECURITY.md](SECURITY.md) · [CONTRIBUTING.md](CONTRIBUTING.md)

## Links
- [wave.online](https://wave.online) · [Docs](https://docs.wave.online) · [Developer portal](https://dev.wave.online)

Operated by WAVE Online, LLC.
