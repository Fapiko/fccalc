# FortressCraft Calculator
---

### Getting it
```go get github.com/Fapiko/fccalc```

### Building it
To rebuild the app after making changes and install the bash completion for recipe names:
```
cd $GOPATH/src/github.com/Fapiko/fccalc
make install
```

### Running it
First argument is the name of the item to look for, second argument (optional) is the quantity of items you want to
calculate for:
```
fccalc StorageHopper 10
StorageHopper
TinBar: 120 (12)
--TinOre: 1920 (16)
IronBar: 100 (10)
--IronOre: 1600 (16)
```
