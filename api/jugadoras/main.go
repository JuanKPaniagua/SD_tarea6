package main

import (
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "strconv"
)

type ContratoRef struct {
    ContratoId int    `json:"contrato_id"`
    Title              string `json:"title"`
}

type Jugadora struct {
    Id          int       `json:"_id"`
    Nombre      string    `json:"nombre"`
    Genero      string    `json:"genero"`
    Nacimiento  int       `json:"nacimiento"`
    Residencia      string    `json:"residencia"`
    Contratos     []ContratoRef `json:"contratos"`
}

var items []Jugadora

var jsonData string = `[
    {
        "_id": 1,
		"nombre":"Celia Ulloa Vega",
        "genero": "Femenino",
        "nacimiento": "12/12/1992",
        "residencia": "Montes de Oca",
		"contratos": [
            {
                "contrato_id": 1,
                "title": "Liga - Activo"
            },
            {
                "contrato_id": 2,
                "title": "Dimas - Finalizado"
            }
        ]
    },
    {
        "_id": 2,
		"nombre":"Mariana Benavides Paniagua",
        "genero": "Femenino",
        "nacimiento": "20/12/1994",
        "residencia": "San Rafael",
		"contratos": [
            {
                "contrato_id": 3,
                "title": "Herediano - Activo"
            },
			{
                "contrato_id": 4,
                "title": "Liga - Cancelado"
            },
            {
                "contrato_id": 5,
                "title": "Herediano - Finalizado"
            }
        ]
    },
	{
        "_id": 3,
		"nombre":"Carolina Jimenez Martinez",
        "genero": "Femenino",
        "nacimiento": "19/07/2000",
        "residencia": "Heredia",
		"contratos": [
            {
                "contrato_id": 6,
                "title": "Saprisa - Activo"
            }
        ]
    },
	{
        "_id": 4,
		"nombre":"Rosa Monge Vargas",
		"genero": "Femenino",
        "nacimiento": "01/01/1998",
        "residencia": "Orosi",
		"contratos": [
            {
                "contrato_id": 7,
                "title": "Dimas - Finalizado"
            },
            {
                "contrato_id": 8,
                "title": "Herediano - Activo"
            }
        ]
    }
]`

func FindItem(id int) *Jugadora {
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