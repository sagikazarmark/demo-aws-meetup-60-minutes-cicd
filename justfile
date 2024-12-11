[private]
default:
  @just --list

# Set up the initial environment
setup:
  dagger call --aws-access-key-id env:AWS_ACCESS_KEY_ID --aws-secret-access-key env:AWS_SECRET_ACCESS_KEY setup

# Tear down all resources
teardown:
  dagger call --aws-access-key-id env:AWS_ACCESS_KEY_ID --aws-secret-access-key env:AWS_SECRET_ACCESS_KEY teardown

# Release a new version
release version="":
  dagger call --aws-access-key-id env:AWS_ACCESS_KEY_ID --aws-secret-access-key env:AWS_SECRET_ACCESS_KEY release --version "{{version}}"

# Deploy a new version
deploy version="":
  dagger call --aws-access-key-id env:AWS_ACCESS_KEY_ID --aws-secret-access-key env:AWS_SECRET_ACCESS_KEY deploy --version "{{version}}"
