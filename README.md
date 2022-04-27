An Extensible Config Package, supporting the usage idioms of Spring Boot, dotnet and Node "config" in Go
=====================================


![Language Badge](https://img.shields.io/github/languages/top/alt-golang/config) <br/>
[release notes](https://github.com/alt-golang/config/blob/main/HISTORY.md)

<a name="intro">Introduction</a>
--------------------------------
An extensible config package, supporting the usage elements of Spring Boot and Node "config", including:
- json, yaml and Java property files,
- cascading value over-rides using, GO_ENV, GO_APP_INSTANCE and GO_PROFILES_ACTIVE
- placeholder resolution (or variable expansion),
- encrypted values (via github.com/alt-golang/gosypt),
- url fetching
- environment variables (via config.get("env.MY_VAR"),
- command line parameters (via config.get("args.MY_ARG")
- and default (or fallback) values,


<a name="usage">Usage</a>
-------------------------

To use the module, import the module as so:

```go
import  ( "github.com/alt-golang/config") ;

config.Get('key');
config.Get('nested.key');
config.GetAs('nested.key',&SomeStruct);
config.GetWithDefault('unknown','use this instead'); // this does not throw an error
```

### File Loading and Precedence

The module follows the file loading and precedence rules of the popular Node
[config](https://www.npmjs.com/package/config) defaults, with additional rules in the style of Spring Boot.

By default, files are loaded from the working directory, or from the config folder if a default.**, production.**, 
or local-development.** file is detected (after the Node "config" idiom).  The NODE_CONFIG_DIR variable can be 
specified to change the default.

Files are over-ridden from the indicated folder in the following order:

- config.( json | yml | yaml | props | properties )
- default.( json | yml | yaml | props | properties )
- application.( json | yml | yaml | props | properties )
- appsettings.( json | yml | yaml | props | properties )
- {GO_ENV}.( json | yml | yaml | props | properties )
- appsettings.{GO_ENV}.( json | yml | yaml | props | properties )
- {GO_ENV}-{GO_APP_INSTANCE}.( json | yml | yaml | props | properties )
- appsettings.{GO_ENV}.{GO_APP_INSTANCE}.( json | yml | yaml | props | properties )
- application-{GO_PROFILES_ACTIVE[0]}.( json | yml | yaml | props | properties )
- application-{GO_PROFILES_ACTIVE[1]}.( json | yml | yaml | props | properties )
- environment variables (over-ridden into env)
- commandline arguments (over-ridden into args)


Environment variables and command line arguments, will over-ride values found in files, for example
`env.MY_VAR=someValue` in a `application.properties` file, or

`local-development.yaml`
```yaml
env:
  MY_VAR: someValue
```

will be over-ridden only if it exists on the host system, negating the need for 
setting local development environment variables or arguments.

### Placeholders, encrypted values, and remote config

Config values that include the common `${placeholder}` syntax, will resolve the inline
placeholders, so the `config.get('placeholder')'` path below will return `start.one.two.end`.

Config values that start with the prefix `enc.` will be decrypted with the
[gosypt](https://github.com/alt-golang/gosypt) package port, with the AES 16,32,64 byte passphrase being
sourced from the `GO_CONFIG_PASSPHRASE` environment variable.

Config values that start with the prefix `url.` are be fetched and resolved , and HTTP options can be specified as in 
the example config file.  

`config.yaml`
```yaml

  "key": "value"
  "one" : "one"
  "placeholder": "start.${one}.${nested.two}.end"
  "placeholderEncrypted": "start.${nested.encrypted}.end"
  "nested" : 
    "key" : "value"
    "two" : "two"
    "placeholder": "start.${one}.${nested.two}.end"
    "encrypted" : "enc.pxQ6z9s/LRpGB+4ddJ8bsq8RqELmhVU2"
    "encryptedWithSecret" : "enc./emLGkD3cbfqoSPijGZ0jh1p1SYIHQeJ"
  "jsonplaceholder": 
    "todos": "url.https://jsonplaceholder.typicode.com/todos/1"
  "fetchWithOpts" : 
    "url": "url.https://jsonplaceholder.typicode.com/todos/1"
    "method": "get",
    "body": "{}"
    "headers":
      "Authorization": "Basic dXNlcjpwYXNz"
      "Content-Type": "application/json"
```
<a name="license">License</a>
-----------------------------

May be freely distributed under the [MIT license](https://raw.githubusercontent.com/alt-javascript/config/main/LICENSE).

Copyright (c) 2022 Craig Parravicini    
