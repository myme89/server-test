# Setup Github Host-self Runner
## Download
1. Create a folder
    * *mkdir actions-runner && cd actions-runner*
2. Download the latest runner package.
    * *curl -o actions-runner-linux-x64-2.301.1.tar.gz -L https://github.com/actions/runner/releases/download/v2.301.1/actions-runner-linux-x64-2.301.1.tar.gz*
3. Optional: Validate the hash
    * *echo "3ee9c3b83de642f9199xxxxxx594ee2601*****18034c1163f8efdf907  actions-runner-linux-x64-2.301.1.tar.gz" | shasum -a 256*
4. Extract the installer
    * *tar xzf ./actions-runner-linux-x64-2.301.1.tar.gz*

## Configure
1. Create the runner and start the configuration experience
    * *./config.sh --url https://github.com/owner/repo --token ASDTKMXxxxxxL55XKY2********xxxxxx*
2. Last step, run it!
    * *./run.sh*

*Using your self-hosted runner*
## Use this YAML in your workflow file for each job
    runs-on: self-hosted

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