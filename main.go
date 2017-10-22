package main

func main() {

	cfg, _ := leggiCFG("conf2.json")
	Login(cfg)
	GetEd2k(cfg.Serie[0])
}
