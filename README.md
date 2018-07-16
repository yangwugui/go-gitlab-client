# go-gitlab-client

[![Build Status](https://travis-ci.org/plouc/go-gitlab-client.png?branch=v2)](https://travis-ci.org/plouc/go-gitlab-client)

**go-gitlab-client** is a client written in golang to consume gitlab API.

It also provides an handy CLI to easily interact with gitlab API.

- [lib](#lib)
  - [install](#install-lib)
  - [update](#update)
  - [documentation](#documentation)
  - [supported APIs](#supported-apis)
- [CLI](#cli)
  - [features](#cli-features)
  - [install](#install-cli)
  - [commands](#cli-commands)

## lib

### Install lib

To install go-gitlab-client, use `go get`:

```
go get github.com/plouc/go-gitlab-client/gitlab
```

Import the `go-gitlab-client` package into your code:

```go
package whatever

import (
    "github.com/plouc/go-gitlab-client/gitlab"
)
```

### Update

To update `go-gitlab-client`, use `go get -u`:

    go get -u github.com/plouc/go-gitlab-client/gitlab

### Documentation

Visit the docs at http://godoc.org/github.com/plouc/go-gitlab-client/gitlab

### Supported APIs

#### Branches

[gitlab api doc](https://docs.gitlab.com/ee/api/branches.html)

- [x] List repository branches
- [x] Get single repository branch
- [x] Protect repository branch
- [x] Unprotect repository branch
- [x] Create repository branch
- [x] Delete repository branch
- [x] Delete merged branches

#### Project-level variables

[gitlab api doc](https://docs.gitlab.com/ee/api/project_level_variables.html)

- [x] List project variables
- [x] Show project variable details
- [x] Create project variable
- [ ] Update project variable
- [x] Remove project variable

#### Group-level variables

[gitlab api doc](https://docs.gitlab.com/ee/api/group_level_variables.html)

- [x] List group variables
- [x] Show group variable details
- [x] Create group variable
- [ ] Update group variable
- [x] Remove group variable

#### Commits

[gitlab api doc](http://doc.gitlab.com/ce/api/commits.html)

- [x] List repository commits
- [ ] Create a commit with multiple files and actions
- [x] Get a single commit
- [x] Get references a commit is pushed to
- [ ] Cherry pick a commit
- [ ] Get the diff of a commit
- [ ] Get the comments of a commit
- [ ] Post comment to commit
- [x] List the statuses of a commit
- [ ] Post the build status to a commit
- [ ] List Merge Requests associated with a commit

#### Deploy Keys

[gitlab api doc](http://doc.gitlab.com/ce/api/deploy_keys.html)

- list project deploy keys
- add/get/rm project deploy key

#### Environments

[gitlab api doc](https://docs.gitlab.com/ee/api/environments.html)

- [x] List environments
- [x] Create a new environment
- [ ] Edit an existing environment
- [x] Delete an environment
- [ ] Stop an environment

#### Groups

[gitlab api doc](https://docs.gitlab.com/ce/api/groups.html)

- [x] List groups
- [ ] List a groups's subgroups
- [ ] List a group's projects
- [x] Details of a group
- [x] New group
- [ ] Transfer project to group
- [ ] Update group
- [x] Remove group
- [x] Search for group
- [x] Group members

#### Jobs

[gitlab api doc](http://doc.gitlab.com/ce/api/jobs.html)

- [x] List project jobs
- [x] List pipeline jobs
- [x] Get a single job
- [ ] Get job artifacts
- [ ] Download the artifacts archive
- [ ] Download a single artifact file
- [x] Get a trace file
- [x] Cancel a job
- [x] Retry a job
- [x] Erase a job
- [ ] Keep artifacts
- [ ] Play a job

#### Projects

[gitlab api doc](http://doc.gitlab.com/ce/api/projects.html)

- [x] List all projects
- [ ] List user projects
- [x] Get single project
- [x] Remove project
- [x] Star a project
- [x] Unstar a project

#### Repositories

[gitlab api doc](http://doc.gitlab.com/ce/api/repositories.html)

- list project repository tags
- list repository commits
- list project hooks
- add/get/edit/rm project hook

#### Users

[gitlab api doc](http://api.gitlab.org/users.html)

- [x] List users
- [x] Single user
- [x] Current user

#### SSH Keys

[gitlab api doc](http://api.gitlab.org/users.html#list-ssh-keys)

- [x] List SSH keys
- [x] List SSH keys for user
- [x] Single SSH key
- [x] Add SSH key
- [x] Add SSH key for user
- [x] Delete SSH key for current user
- [x] Delete SSH key for given user

#### Runners

[gitlab api doc](https://docs.gitlab.com/ee/api/runners.html)

- [x] List owned runners
- [x] List all runners
- [x] Get runner's details
- [ ] Update runner's details
- [ ] Remove a runner
- [ ] List runner's jobs
- [ ] List project's runners
- [ ] Enable a runner in project
- [ ] Disable a runner from project
- [ ] Register a new Runner
- [ ] Delete a registered Runner
- [ ] Verify authentication for a registered Runner

#### Project hooks

[gitlab api doc](https://docs.gitlab.com/ee/api/projects.html#hooks)

- [x] List project hooks
- [x] Get project hook
- [x] Add project hook
- [ ] Edit project hook
- [x] Delete project hook

#### Pipelines

[gitlab api doc](https://docs.gitlab.com/ee/api/pipelines.html)

- [x] List project pipelines
- [x] Get a single pipeline
- [ ] Create a new pipeline
- [ ] Retry jobs in a pipeline
- [ ] Cancel a pipeline's jobs

#### Project badges

[gitlab api doc](https://docs.gitlab.com/ee/api/project_badges.html)

- [x] List all badges of a project
- [x] Get a badge of a project
- [x] Add a badge to a project
- [ ] Edit a badge of a project
- [x] Remove a badge from a project
- [ ] Preview a badge from a project

#### Namespaces

[gitlab api doc](https://docs.gitlab.com/ee/api/namespaces.html)

- [x] List namespaces
- [x] Search for namespace
- [x] Get namespace by ID

#### Merge requests

[gitlab api doc](https://docs.gitlab.com/ee/api/merge_requests.html)

- [x] List merge requests
- [x] List project merge requests
- [x] List group merge requests
- [x] Get single MR
- [ ] Get single MR participants
- [ ] Get single MR commits
- [ ] Get single MR changes
- [ ] List MR pipelines
- [ ] Create MR
- [ ] Update MR
- [ ] Delete a merge request
- [ ] Accept MR
- [ ] Cancel Merge When Pipeline Succeeds
- [ ] Comments on merge requests
- [ ] List issues that will close on merge
- [ ] Subscribe to a merge request
- [ ] Unsubscribe from a merge request
- [ ] Create a todo
- [ ] Get MR diff versions
- [ ] Get a single MR diff version
- [ ] Set a time estimate for a merge request
- [ ] Reset the time estimate for a merge request
- [ ] Add spent time for a merge request
- [ ] Reset spent time for a merge request
- [ ] Get time tracking stats
- [ ] Approvals

## CLI

**go-gitlab-client** provides a CLI to easily interact with GitLab API, **glc**.

### install CLI

**glc** is a single binary with no external dependencies, released for several platforms.
Go to the [releases page](https://github.com/plouc/go-gitlab-client/releases),
download the package for your OS, and copy the binary to somewhere on your PATH.
Please make sure to rename the binary to `glc` and make it executable.

You can also install completion for bash or zsh, please run `glc help completion`
for more info.

### CLI features

- normalized operations: `ls`, `get`, `add`, `update`
- resource aliases for easy retrieval
- `text`, `yaml` & `json` output
- saving output to file
- interactive pagination mode
- interactive resource creation

### CLI commands


- [glc add](#glc-add)	*Add resource*
- [glc add alias](#glc-add-alias)	*Create resource alias*
- [glc add group](#glc-add-group)	*Create a new group*
- [glc add group-var](#glc-add-group-var)	*Create a new group variable*
- [glc add project](#glc-add-project)	*Create a new project*
- [glc add project-badge](#glc-add-project-badge)	*Create project badge*
- [glc add project-branch](#glc-add-project-branch)	*Create project branch*
- [glc add project-environment](#glc-add-project-environment)	*Create project environment*
- [glc add project-hook](#glc-add-project-hook)	*Create a new hook for given project*
- [glc add project-protected-branch](#glc-add-project-protected-branch)	*Protect project branch*
- [glc add project-star](#glc-add-project-star)	*Stars a given project*
- [glc add project-var](#glc-add-project-var)	*Create a new project variable*
- [glc completion](#glc-completion)	*Output shell completion code for the specified shell (bash or zsh)*
- [glc doc](#glc-doc)	*Generate CLI documentation in markdown format*
- [glc get](#glc-get)	*Get resource details*
- [glc get current-user](#glc-get-current-user)	*Get current user*
- [glc get group](#glc-get-group)	*Get all details of a group*
- [glc get group-var](#glc-get-group-var)	*Get the details of a group's specific variable*
- [glc get namespace](#glc-get-namespace)	*Get a single namespace*
- [glc get project](#glc-get-project)	*Get a specific project*
- [glc get project-badge](#glc-get-project-badge)	*Get project badge info*
- [glc get project-branch](#glc-get-project-branch)	*Get project branch info*
- [glc get project-hook](#glc-get-project-hook)	*Get project hook info*
- [glc get project-job](#glc-get-project-job)	*Get project job info*
- [glc get project-job cancel](#glc-get-project-job-cancel)	*Cancel project job*
- [glc get project-job retry](#glc-get-project-job-retry)	*Retry project job*
- [glc get project-job-trace](#glc-get-project-job-trace)	*Get project job trace*
- [glc get project-merge-request](#glc-get-project-merge-request)	*Get project merge request info*
- [glc get project-pipeline](#glc-get-project-pipeline)	*Get project pipeline details*
- [glc get project-var](#glc-get-project-var)	*Get the details of a project's specific variable*
- [glc get runner](#glc-get-runner)	*Get details of a runner*
- [glc get user](#glc-get-user)	*Get a single user*
- [glc init](#glc-init)	*Init glc config*
- [glc ls](#glc-ls)	*List resource*
- [glc ls aliases](#glc-ls-aliases)	*List resource aliases*
- [glc ls group-merge-requests](#glc-ls-group-merge-requests)	*List group merge requests*
- [glc ls group-variables](#glc-ls-group-variables)	*Get list of a group's variables*
- [glc ls groups](#glc-ls-groups)	*List groups*
- [glc ls merge-requests](#glc-ls-merge-requests)	*List merge requests*
- [glc ls namespaces](#glc-ls-namespaces)	*List namespaces*
- [glc ls project-badges](#glc-ls-project-badges)	*List project badges*
- [glc ls project-branches](#glc-ls-project-branches)	*List project branches*
- [glc ls project-commits](#glc-ls-project-commits)	*List project repository commits*
- [glc ls project-environments](#glc-ls-project-environments)	*List project environments*
- [glc ls project-hooks](#glc-ls-project-hooks)	*List project's hooks*
- [glc ls project-jobs](#glc-ls-project-jobs)	*List project jobs*
- [glc ls project-members](#glc-ls-project-members)	*List project members*
- [glc ls project-merge-request-commits](#glc-ls-project-merge-request-commits)	*List project merge request commits*
- [glc ls project-merge-requests](#glc-ls-project-merge-requests)	*List project merge requests*
- [glc ls project-pipeline-jobs](#glc-ls-project-pipeline-jobs)	*List project pipeline jobs*
- [glc ls project-pipelines](#glc-ls-project-pipelines)	*List project pipelines*
- [glc ls project-protected-branches](#glc-ls-project-protected-branches)	*List project protected branches*
- [glc ls project-variables](#glc-ls-project-variables)	*Get list of a project's variables*
- [glc ls projects](#glc-ls-projects)	*List projects*
- [glc ls runners](#glc-ls-runners)	*List runners*
- [glc ls ssh-keys](#glc-ls-ssh-keys)	*List current user ssh keys*
- [glc ls user-ssh-keys](#glc-ls-user-ssh-keys)	*List specific user ssh keys*
- [glc ls users](#glc-ls-users)	*List users*
- [glc rm](#glc-rm)	*Remove resource*
- [glc rm alias](#glc-rm-alias)	*Remove resource alias*
- [glc rm group](#glc-rm-group)	*Remove group*
- [glc rm group-var](#glc-rm-group-var)	*Remove a group's variable*
- [glc rm project](#glc-rm-project)	*Remove project*
- [glc rm project-badge](#glc-rm-project-badge)	*Remove project badge*
- [glc rm project-branch](#glc-rm-project-branch)	*Remove project branch*
- [glc rm project-environment](#glc-rm-project-environment)	*Remove project environment*
- [glc rm project-hook](#glc-rm-project-hook)	*Remove project hook*
- [glc rm project-merged-branches](#glc-rm-project-merged-branches)	*Remove project merged branches*
- [glc rm project-protected-branch](#glc-rm-project-protected-branch)	*Unprotect project branch*
- [glc rm project-star](#glc-rm-project-star)	*Unstars a given project*
- [glc rm project-var](#glc-rm-project-var)	*Remove a project's variable*
- [glc version](#glc-version)	*Print the version number of glc*



#### glc add

Add resource

##### Synopsis

Add resource

##### Options

```
  -h, --help   help for add
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add alias](#glc-add-alias)	*Create resource alias*
- [glc add group](#glc-add-group)	*Create a new group*
- [glc add group-var](#glc-add-group-var)	*Create a new group variable*
- [glc add project](#glc-add-project)	*Create a new project*
- [glc add project-badge](#glc-add-project-badge)	*Create project badge*
- [glc add project-branch](#glc-add-project-branch)	*Create project branch*
- [glc add project-environment](#glc-add-project-environment)	*Create project environment*
- [glc add project-hook](#glc-add-project-hook)	*Create a new hook for given project*
- [glc add project-protected-branch](#glc-add-project-protected-branch)	*Protect project branch*
- [glc add project-star](#glc-add-project-star)	*Stars a given project*
- [glc add project-var](#glc-add-project-var)	*Create a new project variable*



#### glc add alias

Create resource alias

##### Synopsis

Create resource alias

```
glc add alias ALIAS RESOURCE_TYPE [...resource ids] [flags]
```

##### Options

```
  -h, --help   help for alias
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add group

Create a new group

##### Synopsis

Create a new group

```
glc add group [flags]
```

##### Options

```
  -h, --help   help for group
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add group-var

Create a new group variable

##### Synopsis

Create a new group variable

```
glc add group-var GROUP_ID [flags]
```

##### Options

```
  -h, --help   help for group-var
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project

Create a new project

##### Synopsis

Create a new project

```
glc add project [flags]
```

##### Options

```
  -h, --help   help for project
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-badge

Create project badge

##### Synopsis

Create project badge

```
glc add project-badge PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-badge
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-branch

Create project branch

##### Synopsis

Create project branch

```
glc add project-branch PROJECT_ID [flags]
```

##### Options

```
  -b, --branch string   Name of the branch
  -h, --help            help for project-branch
  -r, --ref string      	The branch name or commit SHA to create branch from
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-environment

Create project environment

##### Synopsis

Create project environment

```
glc add project-environment PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-environment
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-hook

Create a new hook for given project

##### Synopsis

Create a new hook for given project

```
glc add project-hook PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-hook
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-protected-branch

Protect project branch

##### Synopsis

Protect project branch

```
glc add project-protected-branch PROJECT_ID BRANCH_NAME [flags]
```

##### Options

```
  -h, --help   help for project-protected-branch
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-star

Stars a given project

##### Synopsis

Stars a given project

```
glc add project-star PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-star
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-var

Create a new project variable

##### Synopsis

Create a new project variable

```
glc add project-var PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-var
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc completion

Output shell completion code for the specified shell (bash or zsh)

##### Synopsis

Output shell completion code for the specified shell (bash or zsh).
The shell code must be evaluated to provide interactive
completion of glc commands.
This can be done by sourcing it from the .bash_profile or .zshrc.
For bash you can run:

  echo "source <(kubectl completion bash)" >> ~/.bashrc


```
glc completion [shell] [flags]
```

##### Options

```
  -h, --help   help for completion
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```



#### glc doc

Generate CLI documentation in markdown format

##### Synopsis

Generate CLI documentation in markdown format

```
glc doc [flags]
```

##### Options

```
  -h, --help   help for doc
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```



#### glc get

Get resource details

##### Synopsis

Get resource details

##### Options

```
  -h, --help   help for get
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get current-user](#glc-get-current-user)	*Get current user*
- [glc get group](#glc-get-group)	*Get all details of a group*
- [glc get group-var](#glc-get-group-var)	*Get the details of a group's specific variable*
- [glc get namespace](#glc-get-namespace)	*Get a single namespace*
- [glc get project](#glc-get-project)	*Get a specific project*
- [glc get project-badge](#glc-get-project-badge)	*Get project badge info*
- [glc get project-branch](#glc-get-project-branch)	*Get project branch info*
- [glc get project-hook](#glc-get-project-hook)	*Get project hook info*
- [glc get project-job](#glc-get-project-job)	*Get project job info*
- [glc get project-job-trace](#glc-get-project-job-trace)	*Get project job trace*
- [glc get project-merge-request](#glc-get-project-merge-request)	*Get project merge request info*
- [glc get project-pipeline](#glc-get-project-pipeline)	*Get project pipeline details*
- [glc get project-var](#glc-get-project-var)	*Get the details of a project's specific variable*
- [glc get runner](#glc-get-runner)	*Get details of a runner*
- [glc get user](#glc-get-user)	*Get a single user*



#### glc get current-user

Get current user

##### Synopsis

Get current user

```
glc get current-user [flags]
```

##### Options

```
  -h, --help   help for current-user
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get group

Get all details of a group

##### Synopsis

Get all details of a group

```
glc get group GROUP_ID [flags]
```

##### Options

```
  -h, --help                     help for group
  -x, --with-custom-attributes   Include custom attributes (admins only)
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get group-var

Get the details of a group's specific variable

##### Synopsis

Get the details of a group's specific variable

```
glc get group-var GROUP_ID VAR_KEY [flags]
```

##### Options

```
  -h, --help   help for group-var
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get namespace

Get a single namespace

##### Synopsis

Get a single namespace

```
glc get namespace NAMESPACE_ID [flags]
```

##### Options

```
  -h, --help   help for namespace
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project

Get a specific project

##### Synopsis

Get a specific project

```
glc get project PROJECT_ID [flags]
```

##### Options

```
  -h, --help         help for project
  -s, --statistics   Include project statistics
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project-badge

Get project badge info

##### Synopsis

Get project badge info

```
glc get project-badge PROJECT_ID BADGE_ID [flags]
```

##### Options

```
  -h, --help   help for project-badge
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project-branch

Get project branch info

##### Synopsis

Get project branch info

```
glc get project-branch PROJECT_ID BRANCH_NAME [flags]
```

##### Options

```
  -h, --help   help for project-branch
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project-hook

Get project hook info

##### Synopsis

Get project hook info

```
glc get project-hook PROJECT_ID HOOK_ID [flags]
```

##### Options

```
  -h, --help   help for project-hook
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project-job

Get project job info

##### Synopsis

Get project job info

```
glc get project-job PROJECT_ID JOB_ID [flags]
```

##### Options

```
  -h, --help   help for project-job
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*
- [glc get project-job cancel](#glc-get-project-job-cancel)	*Cancel project job*
- [glc get project-job retry](#glc-get-project-job-retry)	*Retry project job*



#### glc get project-job cancel

Cancel project job

##### Synopsis

Cancel project job

```
glc get project-job cancel PROJECT_ID JOB_ID [flags]
```

##### Options

```
  -h, --help   help for cancel
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get project-job](#glc-get-project-job)	*Get project job info*



#### glc get project-job retry

Retry project job

##### Synopsis

Retry project job

```
glc get project-job retry PROJECT_ID JOB_ID [flags]
```

##### Options

```
  -h, --help   help for retry
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get project-job](#glc-get-project-job)	*Get project job info*



#### glc get project-job-trace

Get project job trace

##### Synopsis

Get project job trace

```
glc get project-job-trace PROJECT_ID JOB_ID [flags]
```

##### Options

```
  -h, --help   help for project-job-trace
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project-merge-request

Get project merge request info

##### Synopsis

Get project merge request info

```
glc get project-merge-request PROJECT_ID MERGE_REQUEST_IID [flags]
```

##### Options

```
  -h, --help   help for project-merge-request
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project-pipeline

Get project pipeline details

##### Synopsis

Get project pipeline details

```
glc get project-pipeline PROJECT_ID PIPELINE_ID [flags]
```

##### Options

```
  -h, --help   help for project-pipeline
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project-var

Get the details of a project's specific variable

##### Synopsis

Get the details of a project's specific variable

```
glc get project-var PROJECT_ID VAR_KEY [flags]
```

##### Options

```
  -h, --help   help for project-var
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get runner

Get details of a runner

##### Synopsis

Get details of a runner

```
glc get runner RUNNER_ID [flags]
```

##### Options

```
  -h, --help   help for runner
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get user

Get a single user

##### Synopsis

Get a single user

```
glc get user USER_ID [flags]
```

##### Options

```
  -h, --help   help for user
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc init

Init glc config

##### Synopsis

Init glc config

```
glc init [flags]
```

##### Options

```
  -h, --help   help for init
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```



#### glc ls

List resource

##### Synopsis

List resource

##### Options

```
  -h, --help           help for ls
  -p, --page int       Page (default 1)
  -l, --per-page int   Items per page (default 10)
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls aliases](#glc-ls-aliases)	*List resource aliases*
- [glc ls group-merge-requests](#glc-ls-group-merge-requests)	*List group merge requests*
- [glc ls group-variables](#glc-ls-group-variables)	*Get list of a group's variables*
- [glc ls groups](#glc-ls-groups)	*List groups*
- [glc ls merge-requests](#glc-ls-merge-requests)	*List merge requests*
- [glc ls namespaces](#glc-ls-namespaces)	*List namespaces*
- [glc ls project-badges](#glc-ls-project-badges)	*List project badges*
- [glc ls project-branches](#glc-ls-project-branches)	*List project branches*
- [glc ls project-commits](#glc-ls-project-commits)	*List project repository commits*
- [glc ls project-environments](#glc-ls-project-environments)	*List project environments*
- [glc ls project-hooks](#glc-ls-project-hooks)	*List project's hooks*
- [glc ls project-jobs](#glc-ls-project-jobs)	*List project jobs*
- [glc ls project-members](#glc-ls-project-members)	*List project members*
- [glc ls project-merge-request-commits](#glc-ls-project-merge-request-commits)	*List project merge request commits*
- [glc ls project-merge-requests](#glc-ls-project-merge-requests)	*List project merge requests*
- [glc ls project-pipeline-jobs](#glc-ls-project-pipeline-jobs)	*List project pipeline jobs*
- [glc ls project-pipelines](#glc-ls-project-pipelines)	*List project pipelines*
- [glc ls project-protected-branches](#glc-ls-project-protected-branches)	*List project protected branches*
- [glc ls project-variables](#glc-ls-project-variables)	*Get list of a project's variables*
- [glc ls projects](#glc-ls-projects)	*List projects*
- [glc ls runners](#glc-ls-runners)	*List runners*
- [glc ls ssh-keys](#glc-ls-ssh-keys)	*List current user ssh keys*
- [glc ls user-ssh-keys](#glc-ls-user-ssh-keys)	*List specific user ssh keys*
- [glc ls users](#glc-ls-users)	*List users*



#### glc ls aliases

List resource aliases

##### Synopsis

List resource aliases

```
glc ls aliases [flags]
```

##### Options

```
  -h, --help   help for aliases
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls group-merge-requests

List group merge requests

##### Synopsis

List group merge requests

```
glc ls group-merge-requests GROUP_ID [flags]
```

##### Options

```
  -h, --help   help for group-merge-requests
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls group-variables

Get list of a group's variables

##### Synopsis

Get list of a group's variables

```
glc ls group-variables GROUP_ID [flags]
```

##### Options

```
  -h, --help   help for group-variables
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls groups

List groups

##### Synopsis

List groups

```
glc ls groups [flags]
```

##### Options

```
      --all                      Show all the groups you have access to (defaults to false for authenticated users, true for admin)
  -h, --help                     help for groups
      --owned                    Limit to groups owned by the current user
  -s, --search string            Return the list of authorized groups matching the search criteria
      --statistics               Include group statistics (admins only)
  -x, --with-custom-attributes   Include custom attributes in response (admins only)
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls merge-requests

List merge requests

##### Synopsis

List merge requests

```
glc ls merge-requests [flags]
```

##### Options

```
  -h, --help   help for merge-requests
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls namespaces

List namespaces

##### Synopsis

List namespaces

```
glc ls namespaces [flags]
```

##### Options

```
  -h, --help            help for namespaces
  -s, --search string   Returns a list of namespaces the user is authorized to see based on the search criteria
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls project-badges

List project badges

##### Synopsis

List project badges

```
glc ls project-badges PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-badges
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls project-branches

List project branches

##### Synopsis

List project branches

```
glc ls project-branches PROJECT_ID [flags]
```

##### Options

```
  -h, --help            help for project-branches
  -s, --search string   Search term
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls project-commits

List project repository commits

##### Synopsis

List project repository commits

```
glc ls project-commits PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-commits
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls project-environments

List project environments

##### Synopsis

List project environments

```
glc ls project-environments PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-environments
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls project-hooks

List project's hooks

##### Synopsis

List project's hooks

```
glc ls project-hooks PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-hooks
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls project-jobs

List project jobs

##### Synopsis

List project jobs

```
glc ls project-jobs PROJECT_ID [flags]
```

##### Options

```
  -h, --help           help for project-jobs
      --pretty         Use custom output formatting
  -s, --scope string   Scope
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls project-members

List project members

##### Synopsis

List project members

```
glc ls project-members PROJECT_ID [flags]
```

##### Options

```
  -h, --help           help for project-members
  -q, --query string   Search term
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls project-merge-request-commits

List project merge request commits

##### Synopsis

List project merge request commits

```
glc ls project-merge-request-commits PROJECT_ID MERGE_REQUEST_IID [flags]
```

##### Options

```
  -h, --help   help for project-merge-request-commits
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls project-merge-requests

List project merge requests

##### Synopsis

List project merge requests

```
glc ls project-merge-requests PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-merge-requests
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls project-pipeline-jobs

List project pipeline jobs

##### Synopsis

List project pipeline jobs

```
glc ls project-pipeline-jobs PROJECT_ID PIPELINE_ID [flags]
```

##### Options

```
  -h, --help           help for project-pipeline-jobs
      --pretty         Use custom output formatting
  -s, --scope string   Scope
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls project-pipelines

List project pipelines

##### Synopsis

List project pipelines

```
glc ls project-pipelines PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-pipelines
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls project-protected-branches

List project protected branches

##### Synopsis

List project protected branches

```
glc ls project-protected-branches PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-protected-branches
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls project-variables

Get list of a project's variables

##### Synopsis

Get list of a project's variables

```
glc ls project-variables PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-variables
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls projects

List projects

##### Synopsis

List projects

```
glc ls projects [flags]
```

##### Options

```
      --archived        Limit by archived status
  -h, --help            help for projects
      --membership      Limit by projects that the current user is a member of
      --owned           Limit by projects owned by the current user
  -s, --search string   Search term
      --starred         Limit by projects starred by the current user
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls runners

List runners

##### Synopsis

List runners

```
glc ls runners [flags]
```

##### Options

```
      --all            Get a list of all runners in the GitLab instance (specific and shared). Access is restricted to users with admin privileges
  -h, --help           help for runners
  -s, --scope string   The scope of runners to show, one of: specific, shared, active, paused, online; showing all runners if none provided
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls ssh-keys

List current user ssh keys

##### Synopsis

List current user ssh keys

```
glc ls ssh-keys [flags]
```

##### Options

```
  -h, --help   help for ssh-keys
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls user-ssh-keys

List specific user ssh keys

##### Synopsis

List specific user ssh keys

```
glc ls user-ssh-keys USER_ID [flags]
```

##### Options

```
  -h, --help   help for user-ssh-keys
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc ls users

List users

##### Synopsis

List users

```
glc ls users [flags]
```

##### Options

```
      --active            Limit to active users
      --blocked           Limit to blocked users
  -h, --help              help for users
  -s, --search string     Search users by email or username
  -u, --username string   Search users by username
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc ls](#glc-ls)	*List resource*



#### glc rm

Remove resource

##### Synopsis

Remove resource

##### Options

```
  -h, --help   help for rm
  -y, --yes    Do not ask for confirmation
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc rm alias](#glc-rm-alias)	*Remove resource alias*
- [glc rm group](#glc-rm-group)	*Remove group*
- [glc rm group-var](#glc-rm-group-var)	*Remove a group's variable*
- [glc rm project](#glc-rm-project)	*Remove project*
- [glc rm project-badge](#glc-rm-project-badge)	*Remove project badge*
- [glc rm project-branch](#glc-rm-project-branch)	*Remove project branch*
- [glc rm project-environment](#glc-rm-project-environment)	*Remove project environment*
- [glc rm project-hook](#glc-rm-project-hook)	*Remove project hook*
- [glc rm project-merged-branches](#glc-rm-project-merged-branches)	*Remove project merged branches*
- [glc rm project-protected-branch](#glc-rm-project-protected-branch)	*Unprotect project branch*
- [glc rm project-star](#glc-rm-project-star)	*Unstars a given project*
- [glc rm project-var](#glc-rm-project-var)	*Remove a project's variable*



#### glc rm alias

Remove resource alias

##### Synopsis

Remove resource alias

```
glc rm alias [alias] [resource type] [flags]
```

##### Options

```
  -h, --help   help for alias
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm group

Remove group

##### Synopsis

Remove group

```
glc rm group GROUP_ID [flags]
```

##### Options

```
  -h, --help   help for group
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm group-var

Remove a group's variable

##### Synopsis

Remove a group's variable

```
glc rm group-var GROUP_ID VAR_KEY [flags]
```

##### Options

```
  -h, --help   help for group-var
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project

Remove project

##### Synopsis

Remove project

```
glc rm project PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-badge

Remove project badge

##### Synopsis

Remove project badge

```
glc rm project-badge PROJECT_ID BADGE_ID [flags]
```

##### Options

```
  -h, --help   help for project-badge
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-branch

Remove project branch

##### Synopsis

Remove project branch

```
glc rm project-branch PROJECT_ID BRANCH_NAME [flags]
```

##### Options

```
  -h, --help   help for project-branch
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-environment

Remove project environment

##### Synopsis

Remove project environment

```
glc rm project-environment PROJECT_ID ENVIRONMENT_ID [flags]
```

##### Options

```
  -h, --help   help for project-environment
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-hook

Remove project hook

##### Synopsis

Remove project hook

```
glc rm project-hook PROJECT_ID HOOK_ID [flags]
```

##### Options

```
  -h, --help   help for project-hook
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-merged-branches

Remove project merged branches

##### Synopsis

Remove project merged branches

```
glc rm project-merged-branches PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-merged-branches
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-protected-branch

Unprotect project branch

##### Synopsis

Unprotect project branch

```
glc rm project-protected-branch PROJECT_ID BRANCH_NAME [flags]
```

##### Options

```
  -h, --help   help for project-protected-branch
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-star

Unstars a given project

##### Synopsis

Unstars a given project

```
glc rm project-star PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-star
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-var

Remove a project's variable

##### Synopsis

Remove a project's variable

```
glc rm project-var PROJECT_ID VAR_KEY [flags]
```

##### Options

```
  -h, --help   help for project-var
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc version

Print the version number of glc

##### Synopsis

Print the version number of glc

```
glc version [flags]
```

##### Options

```
  -h, --help   help for version
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```



