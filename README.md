# ddns-docker
[![License](https://img.shields.io/github/license/c012vu5/ddns-docker.svg?style=flat-square)](./LICENSE) [![Ubuntu](https://github.com/c012vu5/ddns-docker/actions/workflows/ubuntu.yml/badge.svg)](https://github.com/c012vu5/ddns-docker/actions/workflows/ubuntu.yml)

Notify global IP to mydns.

## Usage
Copy `env.example` to `.env` and fill in the blanks.

```console
$ cat env.example
# Enter your mydns account.
ACC=

# Enter your mydns password.
PASS=
```

And just run `docker-compose up -d`.
