# fira

## Development Guide

Tech Stack:

* Backend: Go 1.20 / GRPC 
* Frontend: NodeJS 16 / Yarn / Next.js

### Runing locally

**Via Docker**

Using docker is an easy way to get started very quickly. All you need is Docker running on your machine, and docker-compose installed. Start Fira by running docker-compose from the root of the repository:

```shell
docker-compose up fira
```

Then open up `http://localhost:8080` in your browser.

Any time you make changes, you can simply close the server, and re-run with `docker-compose up --build fira` to get a fresh build.

**Via Shell**

While using Docker works great for testing, or for small, infrequent changes, it's cumbersome to have to restart and rebuild to see every change.

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
