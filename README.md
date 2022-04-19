An Extensible Port of Node Config 
=====================================


![Language Badge](https://img.shields.io/github/languages/top/alt-golang/config) <br/>
[release notes](https://github.com/alt-golang/config/blob/main/HISTORY.md)

<a name="intro">Introduction</a>
--------------------------------
An extensible port of the popular node config package, also supporting:
- placeholder resolution (or variable expansion),
- encrypted values (via jasypt),
- environment variables (via config.get("env.MY_VAR"),
- command line parameters (via config.get("args.MY_ARG")
- and default (or fallback) values


<a name="usage">Usage</a>
-------------------------

To use the module, substitute the named {config} module export, in place of the popular
[config](https://www.npmjs.com/package/config) default &ndash; note, we use named exports, because the module
exports other useful classes as well.

```go
import  ( "github.com/alt-golang/config") ;

config.get('key');
config.get('nested.key');
config.get('unknown','use this instead'); // this does not throw an error
```

Config values that include the common `${placeholder}` syntax, will resolve the inline
placeholders, so the `config.get('placeholder')'` path below will return `start.one.two.end`.


Config values that start with the prefix `enc.` will be decrypted with the
[gosypt](https://github.com/alt-golang/gosypt) package port, with the AES 16,32,64 byte passphrase being
sourced from the `GO_CONFIG_PASSPHRASE` environment variable.


<a name="license">License</a>
-----------------------------

May be freely distributed under the [MIT license](https://raw.githubusercontent.com/alt-javascript/config/main/LICENSE).

Copyright (c) 2022 Craig Parravicini    