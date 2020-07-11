# Beam

Private bookmark storage for data geeks and homelab enthusiasts.

## Quickstart

### Taking it for a spin

The simplest way to start `beamd` (the Beam server) is without any options.
By default, beamd will listen on port 8080 and use the `memory` storage engine.

```
$ beamd
20XX/0X/XX HH:MM:SS starting beamd... addr:*, port:8080, tls:false
```

### Running with a configuration file

Run `beamd` with the `-c <path>` option. Note that `beamd` will ignore all other
command line options when the `-c` option is provided. In other words, the
configuration file is the single source of truth. The following example starts
beamd with the [example configuration file](https://github.com/toru/beam/blob/master/examples/beam.conf).

```
$ beamd -c examples/beamd.conf
20XX/0X/XX HH:MM:SS starting beamd... addr:*, port:8080, tls:true
```

### Writing a Bookmark

Submit a POST request to the `/bookmarks` resource with a `url` parameter.
Beam will respond with the JSON representation of the newly created bookmark.

```
$ curl -s -d "url=https://github.com/toru/beam" http://localhost:8080/bookmarks | jq
{
  "id": "5e6bc975f53048d0641900388e3c34394d12fa02",
  "url": "https://github.com/toru/beam",
  "name": "",
  "created_at": "2020-07-11T01:23:45.45343Z",
  "updated_at": "2020-07-11T01:23:45.45343Z"
}
```
