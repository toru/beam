# Beam

Private bookmark storage for data geeks and homelab enthusiasts.

## Quickstart

### Taking it for a spin

The simplest way to run beamd (the Beam server) is to start it without any options.
By default, beamd will listen on port 8080 and use the `memory` storage engine.

```
$ beamd
20XX/0X/XX HH:MM:SS starting beamd... addr:*, port:8080, tls:false
```

### Running with a configuration file

Run beamd with the `-c <path>` option. Note that beamd will ignore all other
command line options when the `-c` option is provided. In other words, the
configuration file is the single source of truth.

```
$ beamd -c examples/beamd.conf
20XX/0X/XX HH:MM:SS starting beamd... addr:*, port:8080, tls:true
```
