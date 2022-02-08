# Puck

run commands and get notifications when it finishes.

## Usage

```console
$ puck watch dd if=./example.iso of=/dev/sdb
```

> Puck will now monitor the command and display a notification
> via notify-send when its done. You can use either the `watch`
> command or its alias `w` to accomplish this. If you need to
> test the programs functionality you can use the `test` command
> or its alias `t`.

## Installation

### From Source

#### 1. Download Source Code
```console
$ git clone https://github.com/joshburnsxyz/puck
```

#### 2. Run Build and Install Scripts
```console
$ make
$ sudo make install
```

#### 3. Ensure Installation Worked
```console
$ puck --help
```
### Automatic Installation

> Requires `curl` to be installed on your system

```console
$ curl -sf https://gobinaries.com/github.com/joshburnsxyz/puck | sh
```
