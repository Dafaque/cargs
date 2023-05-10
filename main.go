package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

type keys []string

func (k *keys) String() string {
	return strings.Join(*k, " ")
}

func (k *keys) Set(value string) error {
	*k = append(*k, value)
	return nil
}

func main() {
	var k keys
	var kv keys

	noDash := flag.Bool("no-dash", false, "prepend dashes into output flag keys")
	eq := flag.Bool("eq", false, "split flag and value with \"=\" mark")

	file := flag.String("f", "cargs.yml", "source file")
	encoding := flag.String("e", "", "source encoding; if not set then based on file ext")

	flag.Var(&k, "v", "key to extract; -v example.a => ${example.a}")
	flag.Var(&kv, "kv", "key to extract; return with flag name; -kv e=example.a => -e ${example.a}")
	flag.Parse()

	fmt.Fprintln(os.Stdout, strings.Join(run(*file, *encoding, *noDash, *eq, k, kv), " "))
}

func run(file, encoding string, noDash, eq bool, k, kv keys) []string {
	if len(encoding) == 0 {
		encoding = path.Ext(file)[1:]
	}
	var unmarshal func([]byte, interface{}) error
	switch encoding {
	case "yml", "yaml":
		unmarshal = yaml.Unmarshal
	case "json":
		unmarshal = json.Unmarshal
	default:
		log.Fatal("Unknown encoding " + encoding)
	}
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	var source map[string]interface{} = make(map[string]interface{})
	if err = unmarshal(data, &source); err != nil {
		log.Fatal(err)
	}
	var args []string = make([]string, 0)
	for _, key := range k {
		vals := resolveMapKey(source, key)
		args = append(args, vals...)
	}
	for _, keyvalue := range kv {
		comps := strings.Split(keyvalue, "=")
		if len(comps) != 2 {
			log.Fatal(fmt.Sprintf("invalid kv syntax at %s; Must be separate by \"=\"", keyvalue))
		}

		for _, v := range resolveMapKey(source, comps[1]) {
			separator := " "
			if eq {
				separator = "="
			}
			prefix := "-"
			if len(comps[0]) > 1 {
				prefix = "--"
			}
			if noDash {
				prefix = ""
			}

			args = append(args, fmt.Sprintf(
				"%s%s%s%s",
				prefix,
				comps[0],
				separator,
				v,
			))

		}
	}
	return args
}

func resolveMapKey(m map[string]interface{}, key string) []string {
	currentMap := m
	comonents := strings.Split(key, ".")
	currentKey := 0
	var result []string = make([]string, 0)
	for _, k := range comonents {
		val, exist := currentMap[comonents[currentKey]]
		if !exist {
			log.Fatal("map does not contain key " + k)
		}
		if currentKey != len(comonents)-1 {
			mv, casted := val.(map[string]interface{})
			if !casted {
				log.Fatal(k + " is not a map")
			}
			currentMap = mv
			currentKey++
			continue
		}
		if av, casted := val.([]interface{}); casted {
			for _, el := range av {
				result = append(result, fmt.Sprintf("%v", el))
			}
			break
		}
		result = append(result, fmt.Sprintf("%v", val))

	}
	return result
}
