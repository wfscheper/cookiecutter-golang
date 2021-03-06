# cookiecutter-golang

[![Build Status](https://travis-ci.org/wfscheper/cookiecutter-golang.svg?branch=master)](https://travis-ci.org/wfscheper/cookiecutter-golang)

Powered by [Cookiecutter](https://github.com/audreyr/cookiecutter), Cookiecutter
Golang is a framework for jumpstarting production-ready go projects quickly.

This version of cookiecutter-golang was forked from https://github.com/lacion/cookiecutter-golang.

## Features

- Generous `Makefile` with management commands based on [hellogpher](https://github.com/vincentbernat/hellogopher)
- Uses [dep](https://github.com/golang/dep) for dependency management

## Optional Integrations

- Can use [viper](https://github.com/spf13/viper) for env var config
- Can use [logrus](https://github.com/sirupsen/logrus) for logging
- Can creates dockerfile for building go binary and dockerfile for final go binary (no code in final container)
- If docker is used adds docker management commands to makefile
- Option of TravisCI or None

## Constraints

- Uses dep as only option for depency management
- Only maintained 3rd party libraries are used.

This project now uses docker multistage builds you need at least docker version
v17.05.0-ce to use the docker file in this template,
[you can read more about multistage builds here](https://www.critiqus.com/post/multi-stage-docker-builds/).

## Usage

Let's pretend you want to create a project called "echoserver". Rather than
starting from scratch maybe copying some files and then editing the results to
include your name, email, and various configuration issues that always 
get forgotten until the worst possible moment, get cookiecutter to do all the work.

First, get Cookiecutter. Trust me, it's awesome:
```console
$ pip install cookiecutter
```

alternatively you can install `cookiecutter` with homebrew:
```console
$ brew install cookiecutter
```

finally to run it based on this template just:
```console
$ cookiecutter https://github.com/wfscheper/cookiecutter-golang.git
```

You will be asked about your basic info (name, project name, app name, etc.).
This info will be used to customize your new project.

Warning: After this point, change 'Your Name', 'username', etc to your own information.

Answer the prompts with your own desired [options](). For example:
```console
full_name [Your Name]: Walter Scheper
github_username [username]: wfscheper
app_name [mygolangproject]: echoserver
project_short_description [A Golang project.]: Awesome Echo Server
docker_hub_username [username]: wfscheper
docker_image [username/docker-alpine:latest]: wfscheper/docker-alpine:latest
docker_build_image [username/docker-alpine:gobuildimage]: wfscheper/docker-alpine:gobuildimage
use_docker [y]: y
use_git [y]: y
use_logrus_logging [y]: y
use_viper_config [y]: y
Select use_ci:
1 - travis
2 - none
Choose from 1, 2 [1]: 1
```

Enter the project and take a look around:
```console
$ cd echoserver/
$ ls
```

Run `make help` to see the available management commands, or just run `make build` to build your project.
```console
$ make help
$ make build
$ ./bin/echoserver
```

## Projects built with cookiecutter-golang

- [iothub](https://github.com/lacion/iothub) websocket multiroom server for IoT
