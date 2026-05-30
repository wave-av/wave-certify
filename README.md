# wave-certify

**WAVE Certified hardware validation CLI** — Layer-4 (Hardware) of the [WAVE Protocol Plane](https://github.com/wave-av/wave-foundation/blob/master/frameworks/protocol-plane/README.md).

Partners run `wave-certify` against their NDI/Dante/SRT/MoQ-speaking hardware. It runs a protocol-correctness battery, produces measurements, and emits a signed certification artifact that can be attached to a partner product listing on wave.online/certified.

## Usage

```bash
wave-certify check --target 192.168.1.100 --protocol ndi
# → runs NDI frame integrity battery, color subsampling check, frame-rate stability
# → outputs: ndi-cert-{date}-{sha}.json (signed)

wave-certify check --target srt://stream.example.com:9000 --protocol srt
# → measures RTT, retransmission rate, congestion-control behavior

wave-certify submit ./ndi-cert-2026-05-30-abc123.json
# → posts to wave.online/certified for public listing
```

## Status

Scaffold. Per-protocol batteries are Wave-1 work, tracked in this repo's roadmap issue.
