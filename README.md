# pathtreeprint

> CLI utility to print a `tree(1)` like graph from a list of file paths.

## Deprecation Notice

Around the same time I wrote this utility, [`tree` got a `--fromfile` option](http://mama.indstate.edu/users/ice/tree/changes.html).
This makes this project obsolete.

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)
- [License](#license)

## Install

```
go install github.com/fahrradflucht/pathtreeprint
```

## Usage

It is pretty simple:
```
➜  ~ pathtreeprint path/to/file.txt path/to/another/file.txt
.
└── path
    └── to
        ├── file.txt
        └── another
            └── file.txt
```

A more useful example would be printing a tree from `git ls-tree`:
```
➜  node-libs-browser git:(master) pathtreeprint $(git ls-tree -r --name-only master)
.
├── .gitattributes
├── .gitignore
├── LICENSE
├── README.md
├── index.js
├── mock
│   ├── buffer.js
│   ├── console.js
│   ├── dns.js
│   ├── empty.js
│   ├── net.js
│   ├── process.js
│   ├── punycode.js
│   ├── tls.js
│   └── tty.js
└── package.json
```

Combined with `find(1)` it can do most of the stuff `tree(1)` can do but a lot
more flexible.

## Maintainers

[@fahrradflucht](https://github.com/fahrradflucht)

## Contribute

PRs accepted.

## License

MIT © 2018 Mathis Wiehl
