# Contributing to wave-certify

CLI tool that partners run to validate their WAVE Certified hardware against the protocol-correctness battery. Tests per protocol:

- **NDI** — frame integrity (PSNR ≥ 50dB), color subsampling correctness
- **Dante** — clock sync (PTP residual < 1µs), packet ordering
- **SRT** — round-trip latency (< 200ms), retransmission behavior
- **MoQ** — track parity (sub-frame alignment), join latency

Outputs a signed certification artifact that partners attach to their product listing.

See [Protocol Plane framework](https://github.com/wave-av/wave-foundation/blob/master/frameworks/protocol-plane/README.md) Hardware tier section.
