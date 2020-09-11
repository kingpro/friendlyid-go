# friendlyid - _go_ version of [friendlyId](https://github.com/Devskiller/friendly-id) 

`friendlyid` is a library to generate really unique and url friendly IDs
based on UUID and base62

This library is based on [url62](https://github.com/antonsekatskii/url62-go)  library with the traditional _base62_ alphabet (`0-9A-Za-z`) to achieve  compatibility with library _friendlyid_.

## Usage

```
import "github.com/mariuszs/friendlyid-go/friendlyid"

friendlyid.Encode("c3587ec5-0976-497f-8374-61e0c2ea3da5")
// 5wbwf6yUxVBcr48AMbz9cb

friendlyid.Decode("5wbwf6yUxVBcr48AMbz9cb")
// c3587ec5-0976-497f-8374-61e0c2ea3da5
```

## License

Copyright © 2013 8protons (developed by Anthony Sekatski)

Distributed under the Eclipse Public License either version 1.0 or (at
your option) any later version.
