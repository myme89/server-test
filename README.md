# Setup Github Host-self Runner

Steps to setup Github Host-self Runner


1. On GitHub.com, navigate to the main page of the repository.
2. Under your repository name, click  **Settings**
3. In the left sidebar, click  **Actions**, then click **Runners**.
4. Click **New self-hosted runner**.
5. Select the operating system image and architecture of your self-hosted runner machine.
6. You will see instructions showing you how to download the runner application and install it on your self-hosted runner machine.
   
    Example: Setup Github Host-self Runner with Linux operating system and x64 architecture

        1. Download
            # Create a folder
                $ mkdir actions-runner && cd actions-runner
        
            # Download the latest runner package
            $ curl -o actions-runner-linux-x64-2.302.1.tar.gz -L https://github.com/actions/runner/releases/download/v2.302.1/actions-runner-linux-x64-2.302.1.tar.gz
        
            # Optional: Validate the hash
                $ echo "3d357d4da3449a3b2c644dee1cc245436c09b6e5ece3e26a05bb302501xxxx  actions-runner-linux-x64-2.302.1.tar.gz" | shasum -a 256 -c# Extract the installer
                $ tar xzf ./actions-runner-linux-x64-2.302.1.tar.gz

        2. Configure
            # Create the runner and start the configuration experience
                $ ./config.sh --url https://github.com/myme89/server-test --token ASDTKMQHWJOA4VZQPLBCJUDExxxxx
            
            # Last step, run it!
                $ ./run.sh

            //Using your self-hosted runner
            # Use this YAML in your workflow file for each job
             runs-on: self-hosted
   7.Checking that your self-hosted runner was successfully added: when the runner application is connected to GitHub and ready to receive jobs, you will see the following message on the machine's terminal.

        √ Connected to GitHub

        2022-03-04 09:45:56Z: Listening for Jobs

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