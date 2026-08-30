# Secrets Inventory

| Secret | Purpose | Used by | Rotation |
|---|---|---|---|
| `WAVE_CERTIFY_SIGNING_KEY` | Sign cert artifacts so partners can prove WAVE Certified status | publish.yml | Yearly |

(No other secrets — this is a CLI; the partner downloads and runs it locally.)


## Machine surface

```yaml secrets-contract
version: "0.1"
secrets:
  - name: WAVE_CERTIFY_SIGNING_KEY
    vault: "env:unattributed"
deny_paths:
  - ".dev.vars"
  - ".dev.vars.*"
  - ".env"
  - ".env.*"
```
