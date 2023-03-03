# Setup Github Host-self Runner
## Download
1. Create a folder
    * *mkdir actions-runner && cd actions-runner*
2. Download the latest runner package.
    * *curl -o actions-runner-linux-x64-2.301.1.tar.gz -L https://github.com/actions/runner/releases/download/v2.301.1/actions-runner-linux-x64-2.301.1.tar.gz*
3. Optional: Validate the hash
    * *echo "3ee9c3b83de642f919912e0594ee2601835518827da785d034c1163f8efdf907  actions-runner-linux-x64-2.301.1.tar.gz" | shasum -a 256*
4. Extract the installer
    * *tar xzf ./actions-runner-linux-x64-2.301.1.tar.gz*

# .YML Template File For Github Action
    name: build

    on:
    push:
        tags:
        - "v*"
    jobs:
    build:
        name: Build
        runs-on: [self-hosted, linux]
        steps:
        - name: Check out code into the Go module directory
        uses: actions/checkout@v3

        - name: Set up Go
        uses: actions/setup-go@v3
        with:
            go-version: ^1.20
        id: go