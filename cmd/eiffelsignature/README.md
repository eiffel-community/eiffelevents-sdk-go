# eiffelsignature

This package contains a CLI executable that exposes the signing features
of the SDK. The executable has two subcommands; `sign` that signs one or
more events read from stdin, and `verify` that verifies one or more
events read from stdin.

Both assume that keys in PEM format, and `verify` supports reading any
number of public keys from a director of .pem files, trying the keys for
the author identity found in the signed event until it finds one that can
successfully verify the signature.

The following example signs one or more events found in events.json, and
immediately passes the signed events to the verifier.

```
eiffelsignature sign private_key.pem CN=joe ES512 < events.json | \
    eiffelsignature verify /path/to/public-key-directory
```

## Supported algorithms

| Algorithm | Meaning |
|-----------|---------|
| RS256     | RSASSA-PKCS1-v1_5 using SHA-256 |
| RS384     | RSASSA-PKCS1-v1_5 using SHA-384 |
| RS512     | RSASSA-PKCS1-v1_5 using SHA-512 |
| ES256     | ECDSA using P-256 and SHA-256 |
| ES384     | ECDSA using P-384 and SHA-384 |
| ES512     | ECDSA using P-521 and SHA-512 |
| PS256     | RSASSA-PSS using SHA-256 and MGF1 with SHA-256 |
| PS384     | RSASSA-PSS using SHA-384 and MGF1 with SHA-384 |
| PS512     | RSASSA-PSS using SHA-512 and MGF1 with SHA-512 |

## Exit codes

| Code | Meaning |
|------|---------|
| 0    | The signing or verification operation succeeded. |
| 1    | The signing or verification operation completed, but unsuccessfully. |
| 2    | The command never ran because command line arguments were malformed or missing. |

## Keypair creation example

The following commands create an ECDSA keypair in the form of two PEM
files, one with the private key and one with the public key.

```
openssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:secp521r1 -out private.pem
openssl pkey -in private.pem -pubout -out public.pem
```
