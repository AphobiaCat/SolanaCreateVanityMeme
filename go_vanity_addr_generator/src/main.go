package main

import (
	"encoding/binary"
	"strings"
	"time"

	"github.com/gagliardetto/solana-go"
)

// ----------------Global Parameter>

// any vanity addr you want, It is not recommended to be too long
var wanted_vanitys []string = []string{"Dunty", "dunty", "DUNTY", "haha"}

// fill your create token contract addr
var solana_contract_id string = "your smart contract"

// ----------------Function>

type vanity_addr struct {
	Vanity_addr string `json:"vanity_addr"`
	Random_str  string `json:"random_str"`
	Random_num1 uint64 `json:"random_num1"`
	Random_num2 uint64 `json:"random_num2"`
	Bump        uint32 `json:"bump"`
	Type        string `json:"yype"`
}

func createSeed(stringSeed string, uint64Seed1 uint64, uint64Seed2 uint64) [][]byte {
	var seeds [][]byte

	// convert string to slice
	seeds = append(seeds, []byte(stringSeed))

	// create & write frist uint64
	uint64Seed1Bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(uint64Seed1Bytes, uint64Seed1)
	seeds = append(seeds, uint64Seed1Bytes)

	// create & write second uint64
	uint64Seed2Bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(uint64Seed2Bytes, uint64Seed2)
	seeds = append(seeds, uint64Seed2Bytes)

	return seeds
}

func generator_vanity_addr(random_str string, prefixs ...string) vanity_addr {
	programID := solana.MustPublicKeyFromBase58(solana_contract_id)

	timestamp := time.Now().Unix()

	var randomNum uint64 = 1

	for {

		seeds := createSeed(random_str, uint64(timestamp), randomNum)

		pda, bump, err := solana.FindProgramAddress(seeds, programID)
		if err != nil {
			DBG_ERR("meet some error:", err)
		}

		pdaStr := pda.String()

		for _, val := range prefixs {
			containsPrefix := strings.Contains(pdaStr, val)
			startsWithPrefix := strings.HasPrefix(pdaStr, val)
			endsWithPrefix := strings.HasSuffix(pdaStr, val)

			if containsPrefix || startsWithPrefix || endsWithPrefix {

				type_ := ""

				if containsPrefix {
					type_ = "random"
				}

				if startsWithPrefix {
					type_ = "front"
				}

				if endsWithPrefix {
					type_ = "last"
				}

				return vanity_addr{
					Vanity_addr: pdaStr,
					Random_str:  random_str,
					Random_num1: uint64(timestamp),
					Random_num2: randomNum,
					Bump:        uint32(bump),
					Type:        type_,
				}
			}
		}

		randomNum++

		if randomNum%100000 == 0 {
			DBG_LOG("random:", randomNum)
		}
	}
}

func generator_thread(random_str string) {
	for {
		result := generator_vanity_addr(random_str, wanted_vanitys...)
		if result.Type == "last" || result.Type == "front" {
			DBG_LOG(Build_Json(&result))
		}
	}
}

func main() {

	go generator_thread("Dunty")
	go generator_thread("Is")
	go generator_thread("A")
	go generator_thread("Chain")
	go generator_thread("Dev")
	go generator_thread("Please")
	go generator_thread("Follow")
	go generator_thread("Me")
	go generator_thread("Github")
	go generator_thread("Dunty")
	go generator_thread("Hello")
	go generator_thread("World")
	generator_thread("Start")

	for {
		//DBG_LOG("hello wolrd")
		time.Sleep(1 * time.Second)
	}

}
