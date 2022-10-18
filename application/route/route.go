package route

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	Latitude  float64
	Longitude float64
}

type Route struct { //struct gera uma estrutura de dados
	ID        string
	ClientId  string
	Positions []Position
}

func (route *Route) LoadPositions() error {
	if route.ID == "" {
		return errors.New("route id not valid")
	}

	f, err := os.Open("destination" + route.ID + ".txt") //open file

	if err != nil {
		return err
	}

	defer f.Close() //defer espera tudo que estiver acima dele executar para assim executar o que for definido

	scanner := bufio.NewScanner(f) //scanner le o conteúdo do arquivo aberto anteriormente

	for scanner.Scan() { //vai ler linha a linha do arquivo enquanto tiver conteúdo
		data := strings.Split(scanner.Text(), ",")
		latitude, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}

		longitude, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}

		route.Positions = append(route.Positions, Position{
			Latitude:  latitude,
			Longitude: longitude,
		})
	}

	return nil

}
