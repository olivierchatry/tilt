# Hacking on Tilt

So you want to make a change to `tilt`!

## Contributing

We welcome contributions, either as bug reports, feature requests, or pull requests.

We want everyone to feel at home in this repo and its environs; please see our [**Code of Conduct**](https://docs.tilt.dev/code_of_conduct.html) for some rules that govern everyone's participation.

Most of this page describes how to get set up making & testing changes. See a [YouTube walkthrough](https://youtu.be/oGC5O-BCBhc) showing some of the steps below, for macOS.

Small PRs are better than large ones. If you have an idea for a major feature, please file
an issue first. The [Roadmap](../../../../orgs/tilt-dev/projects/3) has details on some of the upcoming
features that we have in mind and might already be in-progress.

## Build Prereqs

If you just want to build Tilt:

- **[make](https://www.gnu.org/software/make/)**
- **[go 1.18](https://golang.org/dl/)**
- **[golangci-lint](https://github.com/golangci/golangci-lint)** (to run lint)

To use the local Webpack server for UI (default for locally compiled versions of Tilt):
- **[Node.js](https://nodejs.org/en/download/)** (LTS - see `.engines.node` in `web/package.json`)
- **[yarn](https://yarnpkg.com/lang/en/docs/install/)**

## Test Prereqs

If you want to run the tests:

- **[docker](https://docs.docker.com/install/)** - Many of the `tilt` build steps do work inside of containers
  so that you don't need to install extra toolchains locally (e.g., the protobuf compiler).
- **[kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)**
- **[kustomize 2.0 or higher](https://github.com/kubernetes-sigs/kustomize)**: `go get -u sigs.k8s.io/kustomize`
- **[helm](https://docs.helm.sh/using_helm/#installing-helm)**
- **[docker compose](https://docs.docker.com/compose/install/)**: NOTE: this doesn't need to be installed separately from Docker on macOS
- **[jq](https://stedolan.github.io/jq/download/)**

## Optional Prereqs

Other development commands:

- **[goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports?tab=doc)**: `go get -u golang.org/x/tools/cmd/goimports` (to sort imports, IDE-specific installation instructions in the link). You should configure goimports to run with `-local github.com/windmill/tilt`
- **[toast](https://github.com/stepchowfun/toast)**: `curl https://raw.githubusercontent.com/stepchowfun/toast/master/install.sh -LSfs | sh` Used for generating some protobuf files
- Our Python scripts are in Python 3.6.0. To run them:
  - **[pyenv](https://github.com/pyenv/pyenv#installation)**
  - **python**: `pyenv install`
  - if you're using GKE and get the error: "pyenv: python2: command not found", run:
    - `git clone git://github.com/concordusapps/pyenv-implict.git ~/.pyenv/plugins/pyenv-implict`

## Developing

To check out Tilt for the first time, run:

```
go get -u github.com/tilt-dev/tilt/cmd/tilt
```

The Go toolchain will checkout the Tilt repo somewhere on your GOPATH,
usually under `~/go/src/github.com/tilt-dev/tilt`.

To run the fast test suite, run:

```
make shorttest
```

To run the slow test suite that interacts with Docker and builds real images, run:

```
make test
```

If you want to run an integration test suite that deploys servers to Kubernetes and
verifies them, run:

```
make integration
```

To install `tilt` on PATH, run

```
make install
```

This will install the new `tilt` binary in `$GOPATH/bin` - typically `$HOME/go/bin`. You can verify this is the binary you just built with:
```
cd $GOPATH
./bin/tilt version
```

The build date should match the current date. Be aware that you might already have a `tilt` binary in your $PATH, so running `tilt` without specifying exactly which `tilt` binary you want might have you running the wrong binary.


To start using Tilt, just run `tilt up` in any project with a `Tiltfile` -- i.e., NOT the root of the Tilt source code.
There are plenty of toy projects to play with in the [integration](https://github.com/tilt-dev/tilt/tree/master/integration) directory
(see e.g. `./integration/oneup`), or check out one of these sample repos to get started:
- [ABC123](https://github.com/tilt-dev/abc123): Go/Python/JavaScript microservices generating random letters and numbers
- [Servantes](https://github.com/tilt-dev/servantes): a-little-bit-of-everything sample app with multiple microservices in different languages, showcasing many different Tilt behaviors
- [Frontend Demo](https://github.com/tilt-dev/tilt-frontend-demo): Tilt + ReactJS
- [Live Update Examples](https://github.com/tilt-dev/live_update): contains Go and Python examples of Tilt's [Live Update](https://docs.tilt.dev/live_update_tutorial.html) functionality
- [Sidecar Example](https://github.com/tilt-dev/sidecar_example): simple Python app and home-rolled logging sidecar

### APIServer

The Tilt APIServer is our new system for managing Tilt internals:
https://github.com/tilt-dev/tilt-apiserver

To add a new first-party type, run:

```
scripts/api-new-type.sh MyResourceType
```

and follow the instructions.

Once you've added fields for your type, run:

```
scripts/update-codegen.sh
```

to regenerate client code for reading and writing the new type.

### Web
The Tilt UI is a React single page application.

By default, non-release builds of Tilt use a local Webpack dev server.
When Tilt first starts, it will launch the Webpack dev server.
If you immediately open the Tilt web UI, you might get an error message until Webpack has finished starting.
The page should auto-reload once Webpack is ready.

#### Local Snapshot Mode
You can view a locally running Tilt session as though it was a snapshot by tweaking the URL to be `/snapshot/snapshot_id/overview`.
(The `snapshot_id` portion of the URL can be any valid identifier.)
For example, http://localhost:10350/snapshot/aaaa/overview.

Please note this uses a serialized version of the webview/snapshot generated by the Tilt server, so it might behave slightly differently than a real snapshot.

#### Lint (`prettier` + `eslint`)
To format all files with Prettier, run `make prettier` from the repo root or `yarn prettier` from `web/`.

To run lint checks with ESLint (and auto-fix any trivial issues), run `yarn eslint`.

To **verify** that there are no formatting/lint violations, but _not_ auto-fix, run `make check-js` from the repo root or `yarn check` from `web/`.

#### Tests
To run all tests, you can run `make test-js` from the repo root.

If you are actively developing, running `yarn test` from `web/` will launch Jest in interactive mode,
which can auto re-run affected tests and more.

#### Updating Snapshots
First, double check that the element render has changed _by design_ and not as a result of a regression.

The interactive mode of Jest will guide you to update snapshots.
See the [Jest snapshot testing documentation](https://jestjs.io/docs/en/snapshot-testing#interactive-snapshot-mode) for details.

## Remove token to force signing out of Tilt Cloud

Once you've connected Tilt to Tilt Cloud via GitHub, you cannot sign out to break the connection.
But sometimes during development and testing, you need to do this. Remove the token file named `token` 
located at `~/.windmill` on your machine. Restart Tilt, and you will be signed out.

## Performance

### Go Profile

Tilt exposes the standard Go pprof hooks over [HTTP](https://golang.org/pkg/net/http/pprof/).

To look at a 30-second CPU profile:

```
go tool pprof http://localhost:10350/debug/pprof/profile?seconds=30
```

To look at the heap profile:

```
go tool pprof http://localhost:10350/debug/pprof/heap
```

This opens a special REPL that lets you explore the data.
Type `web` in the REPL to see a CPU graph.

For more information on pprof, see https://github.com/google/pprof/blob/master/doc/README.md.

### Opentracing
If you're trying to diagnose Tilt performance problems that lie between Tilt and your Kubernetes cluster (or between Tilt and Docker) traces can be helpful. The easiest way to get started with Tilt's [opentracing](https://opentracing.io/) support is to use the [Jaeger all-in-one image](https://www.jaegertracing.io/docs/1.11/getting-started/#all-in-one).

```
$ docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.11
```

Then start Tilt with the following flags:

```
tilt up --trace --traceBackend jaeger
```

When Tilt starts one of the first lines in the log output should contain a trace ID, like so:

```
TraceID: 26256f1f6aa875e5
```

You can use the Jaeger UI (by default running on http://localhost:16686/) to query for this span and see all of the traces for the current Tilt run. These traces are made available immediately as you use Tilt. You don't need to wait until after Tilt has stopped to get access to the tracing data.

## Web UI

`tilt` uses a web interface for logs investigation.

By default, the web interface runs on port 10350.

When you use a released version of Tilt, all the HTML, CSS, and JS assets are served from our
[production bucket](https://console.cloud.google.com/storage/browser/tilt-static-assets).

When you build Tilt from head, the Tilt binary will default to development mode.
When you run Tilt, it will run a webpack dev server as a separate process on port 46764,
and reverse proxy all asset requests to the dev server.

To manually control the assets served, you can use:

```
tilt up --web-mode=local
```

to force Tilt to use the webpack dev server, or you can use

```
tilt up --web-mode=prod
```

to force Tilt to use production assets.


To run the server on an alternate port (e.g. 8001):

```
tilt up --port=8001
```

## Documentation

The user-facing landing page and documentation lives in
[the tilt.build repo](https://github.com/tilt-dev/tilt.build/).

We write our docs in Markdown and generate static HTML with [Jekyll](https://jekyllrb.com/).

Netlify will automatically deploy the docs to [the public site](https://docs.tilt.dev/)
when you merge to master.

For internal architecture, see [the Tilt Architecture Guide](internal/README.md).

## Wire

Tilt uses [wire](https://github.com/google/wire) for dependency injection. It
generates all the code in the wire_gen.go files.

`make wire-dev` runs `wire` locally and ensures you have fast feedback when
rebuilding the generated code.

`make wire` runs `wire` in a container, to ensure you're using the correct
version.

What do you do if you added a dependency, and `make wire` is failing?

### A Practical Guide to Fixing Your Dependency Injector

(This guide will work with any Dependency Injector - Dagger, Guice, etc - but is
written for Wire)

Step 1) DON'T PANIC. Fixing a dependency injector is like untangling a hair
knot. If you start pushing and pulling dependencies in the middle of the graph,
you will make it much worse.

Step 2) Run `make wire-dev`

Step 3) Look closely at the error message. Identify the "top" of the dependency
graph that is failing. So if your error message is:

```
wire: /go/src/github.com/tilt-dev/tilt/internal/cli/wire.go:182:1: inject wireRuntime: no provider found for github.com/tilt-dev/tilt/internal/k8s.MinikubeClient
	needed by github.com/tilt-dev/tilt/internal/k8s.Client in provider set "K8sWireSet" (/go/src/github.com/tilt-dev/tilt/internal/cli/wire.go:44:18)
	needed by github.com/tilt-dev/tilt/internal/container.Runtime in provider set "K8sWireSet" (/go/src/github.com/tilt-dev/tilt/internal/cli/wire.go:44:18)
wire: github.com/tilt-dev/tilt/internal/cli: generate failed
wire: at least one generate failure
```

then the "top" is the function wireRuntime at wire.go:182.

Step 4) Identify the dependency that is missing. In the above example, that
dependency is MinikubeClient.

Step 5) At the top-level provider function, add a provider for the missing
dependency. In this example, that means we add ProvideMinikubeClient to the
wire.Build call in wireRuntime.

Step 6) Go back to Step (2), and repeat until all errors are gone

Final Note: All dependency injection systems have a notion of groups of common
dependencies (in Wire, they're called WireSets). When fixing an injection error,
you generally want to move providers "up" the graph. i.e., remove them from
WireSets and add them to wire.Build calls. It's OK if this leads to lots of
duplication. Later, you can refactor them back down into common WireSets once
you've got it working.

## Releasing

We use [goreleaser](https://goreleaser.com) to publish binaries. We never run it
locally. We run it in a CircleCI container.

To create a new release at tag `$TAG`, in the `~/go/src/github.com/tilt-dev/tilt`
directory, first switch to `master` and pull the latest changes with `git pull`.
And then:

```
git fetch --tags
git tag -a v0.x.y -m "v0.x.y"
git push origin v0.x.y
```

CircleCI will automatically start building your release, and notify the
#notify-circleci slack channel when it's done. The releaser generates a release on
at https://github.com/tilt-dev/tilt/releases, with a Changelog prepopulated automatically.
(Give it a few moments. It appears as a tag first, before turning into a full release.)

### Version numbers
For pre-v1.0:
* If adding backwards-compatible functionality increment the patch version (0.x.Y).
* If adding backwards-incompatible functionality increment the minor version (0.X.y). We would probably **write a blog post** about this.
