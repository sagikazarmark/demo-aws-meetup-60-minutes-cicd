# Demo: [60 minutes of CI/CD automatizations](https://devertix.com/aws-lounge-hungary-ci-cd/)

**This repository contains the demo code for my presentation.**

The demo uses AWS AppRunner and ECR for deploying a simple application.

## Prerequisites

- AWS account access
- AWS CLI configured (preferably using environment variables in `.env.local`)
- [Dagger](https://docs.dagger.io/quickstart/cli) installed
- [just](https://just.systems/man/en/packages.html) installed (see [justfile][justfile] if you want to run commands directly)

> [!TIP]
> Install Nix and Direnv to get Dagger and Go configured.

## Initial setup

Create the necessary roles for AppRunner to access ECR:

```shell
aws iam create-role --role-name AppRunnerECRAccessRole --assume-role-policy-document file://etc/aws/AppRunnerECRAccessRole.json
aws iam put-role-policy --role-name AppRunnerECRAccessRole --policy-name ECRAccessPolicy --policy-document file://etc/aws/ECRAccessPolicy.json
```

> [!CAUTION]
> The above policy gives read-only access to **all** ECR repositories in the account.

Run the initial setup steps to deploy the first version of the application:

```shell
just setup
```

## Usage

- Find your AppRunner service on AWS Console
- Make a change in `main.go`
- Run `just deploy` to deploy a new version
- Watch the changes being propagated

## Teardown

```shell
just teardown
```

```shell
aws iam create-role --role-name AppRunnerECRAccessRole
```
