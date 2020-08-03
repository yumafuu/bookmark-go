# bookmark-go

## install
```
$ go get bookmark-go
```

## Basic Usage
### show url list
```
$ bookmark-go show
```

### open url
```
$ bookmark-go open --id 1
```
or
```
$ bookmark-go open
id > {input}
```

### add url
```
$ bookmark-go add --url https://github.com/YumaFuu/bookmark-go --keyword bookmark-go
```
or
```
$ bookmark-go add
url > {input}
keyword > {input}
```

### delete url
```
$ bookmark-go delete --id 1
```
or
```
$ bookmark-go delete
id > {input}
```

## Advance Usage
```
$ bookmark-go show | fzf | bookmark-go open
$ bookmark-go show | peco | bookmark-go open
```
[![Image from Gyazo](https://i.gyazo.com/4fe94e15c423e6ffa6ec50f9ffa36b2b.gif)](https://gyazo.com/4fe94e15c423e6ffa6ec50f9ffa36b2b)
