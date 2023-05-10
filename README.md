Convert .YML and .JSON fields into cli arguments
Run `cargs -h` for help

# install
### with go
`go install github.com/Dafaque/cargs@v1.0.2`
### With Homebrew
`brew tap Dafaque/cargs && brew install cargs`

# example
urls.json
```
{
    "google": "https://google.com"
}
```
```
# curl --head $(cargs -f urls.json -kv location=google)
```