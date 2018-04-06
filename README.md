# 📐 miter

## Instalation

```
go get github.com/rafamds/miter
```

## Documentation

To get the ratio and the percentage between two dimensions. The first one is the the width and the second argument is the height.

```
miter 1600 900
```

The result for this will be the following:

> The ratio is: 16:9

> The percentage is: 56.25%

To get the opposite dimension of a ratio use the following method. When you want to get a height, send the value after the flag `--width` and then the width and height. To achieve the opposite, send the height value after the `--height` flag. 

```
miter -w=16 1600 900
miter --width=16 1600 900

miter -ht=16 1600 900
miter --height=16 1600 900
```

The result for this will be something like this:

> The proportional height is:  9