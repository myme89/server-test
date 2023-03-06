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

Here's a basic YAML template file for a GitHub Action:

        name: Template GitHub Action
        on:
            push:
                branches: [master]

            pull_request:

                branches: [master]
       
        jobs:
            build:
                runs-on: ubuntu-latest
                steps:
                - name: Checkout repository
                    uses: actions/checkout@v3

                - name: Setup Go 
                    uses: actions/setup-go@v3
                    with:
                        go-version: ^1.20


                - name: Run tests
                    run: npm test

                - name: Release
                    uses:  actions/create-release@v1

                    env:
                        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
                    with:
                        tag_name:  ${{ github.ref }}
                        release_name: Release ${{ github.ref }}
                        draft: false
                        prerelease: false  

An explanation of the various parameters:

* **name**: *The name of your GitHub Action.*
  
        name: Template GitHub Action

* **on**: *Specifies the events that will trigger your action. In this case, the action will run on any push or pull request to the **master** branch.*

        on:
            push:
                branches: [master]

            pull_request:

                branches: [master]

* **jobs**: *Defines one or more jobs that make up the action.*
    
* **build**: *The name of the job.*

* **runs-on**: *The type of machine that the job will run on.*
  
Can use the latest version of Ubuntu (GitHub-hosted runners)

        runs-on: ubuntu-latest

Or Github Host-self Runner

        runs-on: [self-hosted, linux]

* **steps**: *The list of steps to execute in the job.*

* **uses**: *A pre-built action to use.* 
  
In this case, it's the actions/checkout@v3 action which checks out your repository on the runner.*

        uses: actions/checkout@v3

Or it's the actions/setup-go@v3 action which set up Golang on the runner.
    
        uses: actions/setup-go@v3

* **name**: *The name of the step. This is purely for informational purposes.*
  
* **run**: *The command to run in the step. In this case, it's a series of steps that run tests.*
  
* **with**: *Provides inputs to the action.* 
  
In this case, it provides the environment, target, and secret key required for the release. Note that the tag_nam value is loaded from a GitHub with github.ref or GITHUB_TOKEN value is loaded from a Github secret named GITHUB_TOKEN

        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

        tag_name:  ${{ github.ref }}


