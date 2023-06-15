# Fira

Fira is an open-source personal financial management system designed for self-hosting. It allows you to visualize all your personal financial data securely on your own infrasturcture. Fira is under active development and considered pre-release software. Use at your own risk.

## Community
Join our Discord:
https://discord.gg/uGFwMGDGku

## Development Guide

Tech Stack:

* Backend: Go 1.20 / GRPC 
* Frontend: NodeJS 16 / Yarn / Next.js

### Running locally

**Via Docker**

Using docker is an easy way to get started very quickly. All you need is Docker running on your machine, with make and docker-compose installed. Dependencies are automatically managed, and the environment is optimized for interactive development with auto reloading by default. Start Fira by making the dev target from the root of the repository:

```shell
make dev
```

Then open up `http://localhost:8080` in your browser.

Any time you make changes, to either the server or the client, the changes will be automatically reloaded. The client will hot-reload if possible, the server will re-generate code and re-build.

**Via Shell**

To run the project and its dependencies natively, you'll need the following tools and configuration:

* Go 1.20 or later, installed with `$GOPATH/bin` added to your `$PATH` (this is usually `$HOME/go/bin`).
* NodeJS 20 and the latest Yarn

When working directly on the client, you can `cd` into the `workspace` directory and do most development there. Here is how it would look:

```shell
cd workspace
# build libs
yarn workspace @fira/api-sdk build
# run app
yarn workspace @fira/app dev
```

This will run the server with auto reloading and all the Next.js niceties. 

Then you can run the server directly:

```shell
go run ./cmd/fira serve
```
