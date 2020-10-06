# MATRIZ 90°

Este repo contiene un API que permite girar una matriz NxN en 90 sentido anti-horario°

Escrito en golang e impulsado con echo framework

Problema
---
Rotar matriz en sentido anti-horario

    Elaborar una API de tipo POST que reciba un array de números que conformen una matriz NXN
    y se devuelva la misma matriz pero rotada en sentido anti-horario (90°).
    Se debe validar que los datos de entrada estén en el formato correcto.
    
Ejecución
---

Uri:

    http://127.0.0.1:8080/api/v1/matriz
    
Method:

    POST

Body:

    {
        "Matriz": [
            [1,2,3],
            [4,5,6],
            [7,8,9]
        ]
    }

Result

    {
        "Matriz": [
            [3,6,9],
            [2,5,8],
            [1,4,7]
        ]
    }