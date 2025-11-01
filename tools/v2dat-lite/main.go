package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	router "github.com/v2fly/v2ray-core/v5/app/router/routercommon"
	"google.golang.org/protobuf/proto"
)

func main() {
	outputPath := flag.String("o", "", "output file path")
	flag.Parse()

	inputs := flag.Args()
	if *outputPath == "" || len(inputs) == 0 {
		log.Fatalf("usage: v2dat-lite -o output.dat input1.dat [input2.dat ...]")
	}

	merged := make(map[string]*router.GeoSite)
	order := make([]string, 0, len(inputs))

	for _, path := range inputs {
		data, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("read %s: %v", path, err)
		}

		var list router.GeoSiteList
		if err := proto.Unmarshal(data, &list); err != nil {
			log.Fatalf("unmarshal %s: %v", path, err)
		}

		for _, site := range list.Entry {
			cc := site.CountryCode
			if existing, ok := merged[cc]; ok {
				existing.Domain = append(existing.Domain, site.Domain...)
			} else {
				cloned := proto.Clone(site).(*router.GeoSite)
				merged[cc] = cloned
				order = append(order, cc)
			}
		}
	}

	sort.Strings(order)

	result := &router.GeoSiteList{}
	for _, cc := range order {
		result.Entry = append(result.Entry, merged[cc])
	}

	out, err := proto.Marshal(result)
	if err != nil {
		log.Fatalf("marshal merged geosite: %v", err)
	}

	if err := os.WriteFile(*outputPath, out, 0o644); err != nil {
		log.Fatalf("write %s: %v", *outputPath, err)
	}

	fmt.Fprintf(os.Stderr, "merged %d lists into %s\n", len(order), *outputPath)
}
