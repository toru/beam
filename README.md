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
