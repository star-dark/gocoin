package main
import 

func main() {
	chain := blockchain{}
	chain.addBlock("Genesis Block")
	chain.addBlock("second Block")
	chain.addBlock("third Block")
	chain.listBlock()
}
