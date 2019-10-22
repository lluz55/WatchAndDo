# WatchAndDo

This simple application provides an easy way to watch files changes and executa some commnad when it happens

## Usage
```
  -c string
        Command that will be executed when file changes
  -f string
        Path to file that will be observed
  -i int
        Interval to watch file changes (default 3)
  -w    Wait until file be created when it doesn't exists
```

## Example
```$ wad.exe -f test.txt -c "xcopy test.txt c:\"```

This will copy the ```test.txt``` file from current folder into ```c:\``` drive