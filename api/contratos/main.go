package main

import (
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "strconv"
)

type Contrato struct {
    Id           int     `json:"_id"`
    Title        string  `json:"title"`
    Estado       string  `json:"estado"`
    P_Inicio     string  `json:"p_inicio"`
    P_Final      string  `json:"p_final"`
    Equipo       string  `json:"equipo"`
    Equipo_Id    int     `json:"equipo_id"`
    Jugador      string  `json:"jugador"`
    Jugador_Id   int     `json:"jugador_id"`
}

var contratos []Contrato

var jsonData string = `[
    {
        "_id": 1,
		"title":"Celia Ulloa - Liga - Activo",
        "estado": "Activo",
        "p_inicio": "Marzo 2016",
        "p_final": "Marzo 2022",
        "equipo": "Liga Deportiva Alajuelense",
        "equipo_id": 2,
        "jugador": "Celia Ulloa Vega",
        "jugador_id": 1
    },
    {
        "_id": 2,
		"title":"Celia Ulloa - Dimas - Finalizado",
        "estado": "Finalizado",
        "p_inicio": "Marzo 2015",
        "p_final": "Marzo 2016",
        "equipo": "Dimas Escazu",
        "equipo_id": 3,
        "jugador": "Celia Ulloa Vega",
        "jugador_id": 1
    },
	{
        "_id": 3,
		"title":"Mariana Benavides - Herediano - Activo",
        "estado": "Activo",
        "p_inicio": "Marzo 2020",
        "p_final": "Marzo 2022",
        "equipo": "Club Sport Herediano",
        "equipo_id": 1,
        "jugador": "Mariana Benavides Paniagua",
        "jugador_id": 2
    },
	{
        "_id": 4,
		"title":"Mariana Benavides - Liga - Cancelado",
        "estado": "Cancelado",
        "p_inicio": "Marzo 2018",
        "p_final": "Marzo 2021",
        "equipo": "Liga Deportiva Alajuelense",
        "equipo_id": 2,
        "jugador": "Mariana Benavides Paniagua",
        "jugador_id": 2
    },
	{
        "_id": 5,
		"title":"Mariana Benavides - Herediano - Finalizado",
        "estado": "Finalizado",
        "p_inicio": "Marzo 2016",
        "p_final": "Marzo 2018",
        "equipo": "Club Sport Herediano",
        "equipo_id": 1,
        "jugador": "Mariana Benavides Paniagua",
        "jugador_id": 2
    },
	{
        "_id": 6,
		"title":"Carolina Jimenez - Saprisa - Activo",
        "estado": "Activo",
        "p_inicio": "Marzo 2019",
        "p_final": "Marzo 2024",
        "equipo": "Saprisa FC",
        "equipo_id": 4,
        "jugador": "Carolina Jimenez Martinez",
        "jugador_id": 3
    },
	{
        "_id": 7,
		"title":"Rosa Monge - Dimas - Finalizado",
        "estado": "Finalizado",
        "p_inicio": "Marzo 2016",
        "p_final": "Marzo 2019",
        "equipo": "Dimas Escazu",
        "equipo_id": 3,
        "jugador": "Rosa Monge Vargas",
        "jugador_id": 4
    },
	{
        "_id": 8,
		"title":"Rosa Monge - Herediano - Activo",
        "estado": "Activo",
        "p_inicio": "Marzo 2019",
        "p_final": "Marzo 2024",
        "equipo": "Club Sport Herediano",
        "equipo_id": 1,
        "jugador": "Rosa Monge Vargas",
        "jugador_id": 4
    }
]`

func FindContrato(id int) *Contrato {
    for _, contrato := range contratos {
        if contrato.Id == id {
            return &contrato
        }
    }
    return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    id := req.QueryStringParameters["id"]
    var data []byte
    if id == "" {
        data, _ = json.Marshal(contratos)
    } else {
        param, err := strconv.Atoi(id)
        if err == nil {
            contrato := FindContrato(param)
            if contrato != nil {
                data, _ = json.Marshal(*contrato)
            } else {
                data = []byte("error\n")
            }
        }
    }
    return &events.APIGatewayProxyResponse{
        StatusCode:      200,
        Headers:         map[string]string{"Content-Type": "application/json"},
        Body:            string(data),
        IsBase64Encoded: false,
    }, nil
}

func main() {
    _ = json.Unmarshal([]byte(jsonData), &contratos)
    lambda.Start(handler)
}