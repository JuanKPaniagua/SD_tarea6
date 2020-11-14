package main

import (
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "strconv"
)

type ContratoRef struct {
    Contrato_Id int    `json:"contrato_id"`
    Title  string      `json:"title"`
}

type Equipo struct {
    Id        int           `json:"_id"`
    Nombre    string        `json:"nombre"`
    Cuidad    string        `json:"cuidad"`
    Estadio   string        `json:"estadio"`
    Fundacion string        `json:"fundacion"`
    Contratos []ContratoRef `json:"contratos"`
}

var items []Equipo

var jsonData string = `[
     {
        "_id": 1,
		"nombre":"Club Sport Herediano",
        "cuidad": "Heredia",
        "estadio": "Rosabal Cordero",
        "fundacion": "1921",
		"contratos": [
            {
                "contrato_id": 3,
                "title": "Mariana Benavides - Activo"
            },
            {
                "contrato_id": 5,
                "title": "Mariana Benavides - Finalizado"
            },
            {
                "contrato_id": 8,
                "title": "Rosa Monge - Activo"
            }
        ]
    },
    {
        "_id": 2,
		"nombre":"Liga Deportiva Alajuelense",
        "cuidad": "Alajuela",
        "estadio": "Alejando Morera Soto",
        "fundacion": "1920",
		"contratos": [
            {
                "contrato_id": 1,
                "title": "Celia Ulloa - Activo"
            },
            {
                "contrato_id": 4,
                "title": "Mariana Benavides - Cancelado"
            }
        ]
    },
	{
        "_id": 3,
		"nombre":"Dimas Escazu",
        "cuidad": "Escazu",
        "estadio": "Nicolas Masis",
        "fundacion": "2005",
		"contratos": [
            {
                "contrato_id": 2,
                "title": "Celia Ulloa - Finalizado"
            },
            {
                "contrato_id": 7,
                "title": "Rosa Monge - Finalizado"
            }
        ]
    },
	{
        "_id": 4,
		"nombre":"Saprisa FC",
		"cuidad": "Tibas",
        "estadio": "Ricardo Saprisa",
        "fundacion": "1956",
		"contratos": [
            {
                "contrato_id": 6,
                "title": "Carolina Jimenez - Activo"
            }
        ]
    }
]`

func FindItem(id int) *Equipo {
    for _, item := range items {
        if item.Id == id {
            return &item
        }
    }
    return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    id := req.QueryStringParameters["id"]
    var data []byte
    if id == "" {
        data, _ = json.Marshal(items)
    } else {
        param, err := strconv.Atoi(id)
        if err == nil {
            item := FindItem(param)
            if item != nil {
                data, _ = json.Marshal(*item)
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
    _ = json.Unmarshal([]byte(jsonData), &items)
    lambda.Start(handler)
}