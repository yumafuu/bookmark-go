# bookmark-go
bookmark-go is a bookmark tool on cli.

## install
```
$ go get bookmark-go
```

This tool uses sqlite3. The default file path is `~/.bookmark.db`.
if you want to change file path, set env variable `BOOKMARK_GO_DB_PATH`

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

## Advanced Usage
```
$ bookmark-go show | fzf | bookmark-go open
$ bookmark-go show | peco | bookmark-go open
```
[![Image from Gyazo](https://i.gyazo.com/4fe94e15c423e6ffa6ec50f9ffa36b2b.gif)](https://gyazo.com/4fe94e15c423e6ffa6ec50f9ffa36b2b)
