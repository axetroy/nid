## Nid for go

[![Build Status](https://travis-ci.org/axetroy/Github.svg?branch=master)](https://travis-ci.org/axetroy/Github)
![License](https://img.shields.io/badge/license-Apache-green.svg)

Generate random number id in Golang

## Usage

```bash
go get -v github.com/axetroy/nid
```

```go
package main

import (
  "github.com/axetroy/nid"
  "fmt"
)

func main() {
  nid := nid.New(8)
  fmt.Print(nid)
}
```

## Contributing

[Contributing Guid](https://github.com/axetroy/nid/blob/master/CONTRIBUTING.md)

## Contributors

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
| [<img src="https://avatars1.githubusercontent.com/u/9758711?v=3" width="100px;"/><br /><sub>Axetroy</sub>](http://axetroy.github.io)<br />[💻](https://github.com/axetroy/nid/commits?author=axetroy) [🐛](https://github.com/axetroy/nid/issues?q=author%3Aaxetroy) 🎨 |
| :---: |
<!-- ALL-CONTRIBUTORS-LIST:END -->

## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Faxetroy%2Fnid.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Faxetroy%2Fnid?ref=badge_large)