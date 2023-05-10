Convert .YML and .JSON fields into cli arguments
Run `cargs -h` for help

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