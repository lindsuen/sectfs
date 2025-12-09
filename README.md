# SectFS

[![Commit activity](https://img.shields.io/github/commit-activity/m/lindsuen/sectfs)](https://github.com/lindsuen/sectfs/graphs/commit-activity)
[![build](https://img.shields.io/github/actions/workflow/status/lindsuen/sectfs/build.yml?branch=master)](https://github.com/lindsuen/sectfs/actions/workflows/build.yml)
[![GitHub Release](https://img.shields.io/github/v/release/lindsuen/sectfs)](https://github.com/lindsuen/sectfs/releases)
[![GitHub License](https://img.shields.io/github/license/lindsuen/sectfs)](https://github.com/lindsuen/sectfs/blob/master/README.md)

Fast File Service in Go.

## Start

```sh
$ git clone https://github.com/lindsuen/sectfs.git
$ cd sectfs/
```

### Binary

The `make` tool is needed.

```sh
$ make build
```

```sh
$ mv bin/sectfs ./ && ./sectfs
```

### Docker

The `make` and `docker` tools are needed.

```sh
$ make build
```

```sh
$ docker build --no-cache -t sectfs-server:latest .
```

```sh
$ docker run -p 5363:5363 --name sectfs-server -v ${TARGET_DIR}/data:/usr/local/sectfs-server/data -v ${TARGET_DIR}/upload:/usr/local/sectfs-server/upload -d sectfs-server:latest
```

## License

[BSD 2-Clause license](https://github.com/lindsuen/sectfs/blob/master/README.md)
