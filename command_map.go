package main

import "fmt"

func commandMap(cfg *PageConfig) error {
	url := "https://pokeapi.co/api/v2/location-area" // Varsayılan URL
	if cfg.Next != nil {                             // Eğer Next varsa, bir sonraki sayfaya git
		url = *cfg.Next
	}

	results, newCfg, err := getAreas(url)
	if err != nil {
		return fmt.Errorf("cannot map: %v", err)
	}

	// Sonuçları yazdır
	for i := range results {
		fmt.Printf("%s\n", results[i].Name)
	}

	// Yeni pagination bilgilerini cfg'ye ata
	cfg.Next = newCfg.Next
	cfg.Previous = newCfg.Previous

	return nil
}
