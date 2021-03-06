# Kubernetes OTP Authentication

This repository contains source code for a simple program I created to integrate OneTime Password authentication into Kubernetes. Right now the program only support LinOTP as a OTP backend
The program is integrated with Kubernetes via [Authentication Webhooks](https://kubernetes.io/docs/admin/authentication/#webhook-token-authentication)

## How does it work

This program excepts authentication token that gets passed when sending a command to Kubernetes (i.e kubectl command) to be `base64` encoding of `username` and `one-time-password` combination. For example, to generate a token for username `john` and with OTP password `123456` generated by LinOTP you have to Base64 encode `"john:123456"`. That is `am9objoxMjM0NTY=`
Now you should pass this token to your kubectl commands via `--token` flag.

```
kubectl get nodes --token am9objoxMjM0NTY=
```

## Configuration

## Use cases

* To make authencation more secure
* If you use a CI tool like Jenkins with manual user approves to deploy to Kubernetes OTP is perfect to preved passwords to be stored in logs etc.

## Plan

* Add more backends
* Imrpove code
* Dockerize
* Create Kubernetes addon
