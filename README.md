# WatchAndDo

This simple application provides an easy way to watch files changes and executa some commnad when it happens


## Instalation
1. clone this repo  
``` $ git clone github.com/lluz55/WatchAndDo```
2. Enter into clone folder  
``` $ cd WatchAndDo```
3. Build
``` $ go build -ldflags="-s -w"```  
4. Run  
``` $ .\wad.exe```

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

## TODO
- [ ] Watches for changes in some folder
- [ ] Allow watch multiple files
- [ ] Be able to run in background
