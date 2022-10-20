# Nitriding key sharing test

The nitriding framework has a path for sharing key data between
applications running in AWS Nitro enclaves. This is a simple
test application for exercising this API.

## Setup

Start a ec2 instance to host and proxy for the enclave. It must
be an instance type supporting nitro enclaves, e.g. c5.xlarge.

On the parent instance, build and run the two required proxies
from https://github.com/brave-intl/bat-go/tree/nitro-utils/nitro-shim/tools


```sh
export OUT_ADDRS="4:8443,4:80,127.0.0.1:1080"
export IN_ADDRS=":8443,:80,3:1080"
sudo viproxy &
```

```sh
export SOCKS_PROXY_LISTEN_ADDR="127.0.0.1:1080"
export SOCKS_PROXY_ALLOWED_ADDRS=""
export SOCKS_PROXY_ALLOWED_FQDNS="acme-v02.api.letsencrypt.org,other-enclave.example.com"
socksproxy &
```

Then run `make` to build and run this test application in the enclave.
