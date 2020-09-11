# url62 - compatible with [friendlyId](https://github.com/Devskiller/friendly-id) 

url62 is a library to generate really unique and url friendly IDs
based on UUID and base62

## Usage

```
import "github.com/mariuszs/url62-go/url62"

url62.FromUUID("6ba7b814-9dad-11d1-80b4-00c04fd430c8")
// "U7PVVMkSSJGepn8RKilghA"

url62.ToUUID("5wbwf6yUxVBcr48AMbz9cb")
// c3587ec5-0976-497f-8374-61e0c2ea3da5
```

## License

Copyright © 2013 8protons (developed by Anthony Sekatski)

Distributed under the Eclipse Public License either version 1.0 or (at
your option) any later version.
