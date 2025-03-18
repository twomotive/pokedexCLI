package main

import "fmt"

func commandMapB(cfg *PageConfig) error {
	if cfg.Previous == nil {
		fmt.Println("No previous location areas.")
		return nil
	}
	url := *cfg.Previous
	results, newCfg, err := getAreas(url)
	if err != nil {
		return fmt.Errorf("cannot mapb: %v", err)
	}
	for _, result := range results {
		fmt.Println(result.Name)
	}
	cfg.Next = newCfg.Next
	cfg.Previous = newCfg.Previous
	return nil
}
