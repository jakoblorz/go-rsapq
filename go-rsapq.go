package main

import (
	"errors"
	"fmt"

	"github.com/caarlos0/env"
	eea "github.com/jakoblorz/go-eea/lib"
)

type RSAPQEnvironment struct {
	P int64 `env:"P"`
	Q int64 `env:"Q"`
}

func (env *RSAPQEnvironment) N() int64 {
	return env.P * env.Q
}

func (env *RSAPQEnvironment) Phi() int64 {
	return (env.P - 1) * (env.Q - 1)
}

func (env *RSAPQEnvironment) Exponents() (int64, int64, error) {

	phi := env.Phi()

	for e := int64(2); e < phi; e++ {
		table := (&eea.ExtendedEuclidianParameters{
			A: e,
			B: phi,
		}).Calculate()

		if table[len(table)-1].A == 1 {
			return e, table[0].S, nil
		}
	}

	return 0, 0, errors.New("Could not find any Exponent")
}

func main() {

	cfg := RSAPQEnvironment{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	fmt.Printf("generating pk(e,n), sk(e,n) for p=%d, q=%d with n=%d, phi=%d\n\n", cfg.P, cfg.Q, cfg.N(), cfg.Phi())

	e, d, err := cfg.Exponents()
	if err != nil {
		panic(err)
	}

	fmt.Printf("> pk(e,n) = (%d,%d)\n", e, cfg.N())
	fmt.Printf("> sk(d,n) = (%d,%d)\n", d, cfg.N())
}
