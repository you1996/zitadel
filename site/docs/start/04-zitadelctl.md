---
title: Static Manifest
---

## Install ZITADEL with zitadelctl

### Install zitadelctl for your architecture


`
curl -s https://api.github.com/repos/caos/zitadel/releases/latest | grep "browser_download_url.*zitadelctl-$(uname | tr A-Z a-z)" | cut -d '"' -f 4  | wget -i - -O ./zitadelctl
`

or download manual from 

[https://github.com/caos/zitadel/releases/latest](https://github.com/caos/zitadel/releases/latest)


set permissions to the executable

`
chmod u+x zitadelctl
`





> This will be added later on
