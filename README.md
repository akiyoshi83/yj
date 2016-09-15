# yj

Convert YAML to JSON and JSON to YAML commadns.

## Usage

```
$ yj example.yml
{
  "a": 100,
  "b": "data",
}
```

```
$ jy example.json
a: 100
b: data
c:
- 1
- 2
- 3
d:
  d1: 1
  d2:
  - 1
  - 2
  - 3
  d3:
    foo: bar
```

## Install

```
go get github.com/akiyoshi83/yj/yj
go get github.com/akiyoshi83/yj/jy
```

