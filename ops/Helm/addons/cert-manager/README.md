# cert-manager

![Version: 0.1.1](https://img.shields.io/badge/Version-0.1.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | batazor111@gmail.com | batazor.ru |

## Requirements

Kubernetes: `>= 1.19.0 || >= v1.19.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://charts.jetstack.io | cert-manager | 1.1.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| annotations | object | `{}` |  |
| cert-manager.enabled | bool | `true` |  |
| cert-manager.installCRDs | bool | `true` |  |
| cert-manager.prometheus.enabled | bool | `true` |  |
| cert-manager.prometheus.servicemonitor.enabled | bool | `true` |  |
| cert-manager.prometheus.servicemonitor.labels.release | string | `"prometheus-operator"` |  |
| email | string | `"mymail@gmail.com"` |  |
