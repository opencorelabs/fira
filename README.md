# fira

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
* NodeJS 16 and the latest Yarn

When working directly on the client, you can `cd` into the `client` directory and do most development there. Here is how it would look:

```shell
cd client
yarn install
yarn run dev
```

This will run the server with auto reloading and all the Next.js niceties. 

For working on the API server, you will want to be running in the root of the repository. Install the Go tools necessary for development:

```shell
./scripts/install-tools.sh
```

Then you can run the server directly:

```shell
go run ./cmd/server
```
