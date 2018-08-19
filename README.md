# currc
Simple CLI tool for fast and easy currency conversion.


Supports only USD/EUR to and from HRK (Croatian kuna) for now.

## Usage
```./currc <flag> <amount>```

Flags:
```
  -fd <dollars>
    	convert from <dollars> to hrk
  -fe <euros>
    	convert from <euros> hrk
  -td <hrk>
    	convert from <hrk> to dollars
  -te <hrk>
    	convert from <hrk> to euros
  -h	lists all commands
```

#### TODO
- export flags to file so that adding new flags is easier
- working with non-compact API
