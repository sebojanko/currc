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
## API key

To use the program an API key is needed. You can obtain one [here](https://free.currencyconverterapi.com/free-api-key).
The must be saved in an environment variable by the name CURRC_API_KEY.
```
export CURRC_API_KEY="insert_your_api_key_here"
```
Insert this command in your dotfiles if the variable doesn't persist between sessions.


#### TODO
- export flags to file so that adding new flags is easier
- working with non-compact API
