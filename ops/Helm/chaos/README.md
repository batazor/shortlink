# chaos

![Version: 0.4.3](https://img.shields.io/badge/Version-0.4.3-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

Chaos service

**Homepage:** <https://batazor.github.io/shortlink/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | batazor111@gmail.com | batazor.ru |

## Source Code

* <https://github.com/batazor/shortlink>

## Requirements

Kubernetes: `>= 1.19.0 || >= v1.19.0-0`

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| chaosMesh.IoChaos.enabled | bool | `false` |  |
| chaosMesh.NetworkChaos.enabled | bool | `false` |  |
| chaosMesh.namespace | string | `"default"` |  |
| chaosMesh.podChaos.enabled | bool | `false` |  |
| chaosMesh.podChaos.labelSelectors."app.kubernetes.io/part-of" | string | `"shortlink"` |  |
| chaosMesh.podChaos.scheduler | string | `"@every 5m"` |  |
| fullnameOverride | string | `""` |  |
| nameOverride | string | `""` |  |
| serviceAccount.create | string | `"create"` |  |
| serviceAccount.name | string | `"shortlink"` |  |
