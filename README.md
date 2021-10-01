# faux-browser

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/First-Foundation/faux-browser/agent-terminal-go-build?style=plastic)
![GitHub last commit](https://img.shields.io/github/last-commit/First-Foundation/faux-browser?style=plastic)
![GitHub](https://img.shields.io/github/license/First-Foundation/faux-browser?style=plastic)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/First-Foundation/faux-browser?style=plastic)

Fake Automated User eXperience for generating user browser activity.

## Rough Instructions For Use
- Run `go get golang.org/x/net/html`
- Add/remove Sites/Searches as you desire.
- Search for `USER TODO` and make changes as you desire.
- Run `./build.sh`
- Copy the appropriate binary to your machine, renaming it to the browser you want to emulate.
  - Check out the `NewBrowser()` function for implemented options.
