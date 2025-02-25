package main

import (
	"api/database"
	"api/models"
	"api/routes"
	"fmt"
)

func main() {
	models.Player = []models.Players{
		{Name: "Weverton", History: "Jogador com mais títulos na história do Palmeiras (12, ao lado de Gustavo Gómez, Marcos Rocha, Mayke, Dudu, Ademir da Guia e Junqueira), Weverton é o 5º goleiro que mais atuou e o 5º que mais venceu pelo Palmeiras em todos os tempos, o 2º que mais jogou no Estádio Palestra Italia/Allianz Parque, além de ser o atleta que mais defendeu o Verdão em Libertadores e o que mais venceu pela competição continental. Além disso, Weverton igualou Emerson Leão como o goleiro com mais títulos brasileiros pelo Palmeiras (com três cada).", ShirtNumber: 21, Id: "1"},
		{Name: "Piquerez", History: "Com passagens pelas Seleções Uruguaias Sub-17 e Sub-23, Piquerez foi revelado pelo Defensor-URU e passou também pelo River Plate-URU antes de chegar ao Peñarol-URU, no começo de 2020. O jogador, que é o 21º uruguaio na história do Palmeiras, chegou ao clube em 2021 para substituir o compatriota Matías Viña, campeão paulista, da Copa do Brasil e da Libertadores pela equipe alviverde na temporada 2020.", ShirtNumber: 22, Id: "2"},
	}
	database.ConnectWithDB()

	fmt.Println("Server is listening on")
	routes.HandleRequest()
}
