# datadog

![Version: 0.2.1](https://img.shields.io/badge/Version-0.2.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | batazor111@gmail.com | batazor.ru |

## Requirements

Kubernetes: `>= 1.19.0 || >= v1.19.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://helm.datadoghq.com | datadog-operator | 0.4.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| apiKey | string | `"<DATADOG_API_KEY>"` |  |
| datadog.enabled | bool | `true` |  |
