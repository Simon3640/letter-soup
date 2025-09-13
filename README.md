# Documentación de la prueba de programación - Sopa de letras

## 1. Descripción del problema
Se recibe como entrada:
1. Una lista de palabras a buscar.
2. Una matriz de caracteres representada como una lista de strings separados por comas.
El objetivo es identificar si cada palabra se encuentra dentro de la sopa, considerando todas las direcciones posibles:
- Horizontal: izquierda → derecha y derecha → izquierda.
- Vertical: arriba → abajo y abajo → arriba.
- Diagonal: ↘ ↙ ↖ ↗ (en los cuatro sentidos).
La salida es un listado de las palabras, clasificadas como encontradas (con su posición inicial, final y dirección) o no encontradas.


## 2. Tecnologías utilizadas
- Lenguaje de programación: Go (Golang) 1.25
- Framework base: [GoProjectSkeleton](https://github.com/simon3640/GoProjectSkeleton)
- Estándares: Clean Architecture, REST API / Domain-Driven Design (DDD) simplificado

## 3. Lógica de la solución
El principal componente de la logica se encuentra en el archivo `src/domain/utils/matrix_operations.go`

Se propone un algoritmo de fuerza bruta para extraer las palabras de la matriz, haciendo uso de álgebra básica se definen transformaciones ortogonales que permitan, "reflejar" "trasponer" "rotar" la matriz y de esta manera buscar de izquierda a derecha cada una de las palabras.

### 🔹 Representación de transformaciones

Se define un tipo Transform que contiene:
- Grid: matriz transformada.
- PosMap: mapeo inverso de posiciones para traducir coordenadas de vuelta al tablero original.
- Name: nombre de la transformación aplicada.

```
type Transform struct {
	Name   string
	Grid   [][]rune
	PosMap [][]Pos
}
```
### 🔹 Construcción de Transformaciones
La función BuildTransform permite crear una nueva matriz aplicando un mapeo de coordenadas (mapper), garantizando que siempre se pueda rastrear la posición original:
```
func BuildTransform(m [][]rune, Rnew, Cnew int, mapper func(r, c int) (int, int), Name string) Transform
```

Ejemplo de transformaciones aplicadas:
- Identidad (sin cambios).
- Reflexiones horizontales y verticales.
- Rotaciones y reflejos diagonales.
- Construcción de matrices diagonales (↘ y ↙) mediante BuildDiagonalMatrixSE y BuildDiagonalMatrixSW.

Esto permite recorrer todas las orientaciones posibles sin tener que programar un buscador distinto para cada dirección.

## 4 Algoritmo de búsqueda

ubicado en `src/application/service/solve_letter_soup.go`
Pasos principales:
- Validación de entrada (Rows, Columns, formato del grid).
- Conversión de la lista de strings en una matriz de runas ([][]rune) para trabajar carácter por carácter.
- Generación de transformaciones (horizontales, verticales y diagonales).
- Búsqueda secuencial:
    - Para cada transformación, se recorre la matriz.
    - Se compara cada secuencia de caracteres con la palabra buscada.
    - Al encontrar coincidencias, se registra la palabra con:
        - Posición inicial y final en la matriz original.
        - Dirección (nombre de la transformación aplicada).
- Construcción de la respuesta:
    - FoundWords: lista de coincidencias.
    - NotFoundWords: palabras no localizadas.

Ejemplo de respuesta:
```
{
    "data": {
        "Rows": 14,
        "Columns": 14,
        "user_id": 52,
        "grid": [
            "NDEKICANGUROGE",
            "SXRYKVIIQGWQOD",
            "JAGUARZWBNKOUA",
            "MLELEFANTEHOGW",
            "LOBONUTRIAOUSU",
            "WWOSOGATOVRTMO",
            "HLZNCTYZEOXAUR",
            "CECYTIBURONSRO",
            "CONEJOYUSMRSHT",
            "YNIFEFPTEZOOSF",
            "OSSERPIENTEFLG",
            "PPVDDXUFALCONY",
            "MONOCUQWMANATI",
            "NNXHEBPMUPERRO"
        ],
        "words": [
            "MANATI",
            "LEON",
            "PERRO",
            "LORO",
            "GATO",
            "TORO",
            "CONEJO",
            "ORUGA",
            "TIBURON",
            "ELEFANTE",
            "ALCON",
            "SERPIENTE",
            "JAGUAR",
            "CANGURO",
            "LOBO",
            "MONO",
            "NUTRIA",
            "NOENCONTRADO"
        ],
        "id": 6,
        "created_at": "2025-09-13T13:24:30.190135392Z",
        "updated_at": "2025-09-13T13:24:30.190135392Z",
        "deleted_at": "0001-01-01T00:00:00Z",
        "founded_words": [
            "MANATI",
            "PERRO",
            "GATO",
            "CONEJO",
            "TIBURON",
            "ELEFANTE",
            "ALCON",
            "SERPIENTE",
            "JAGUAR",
            "CANGURO",
            "LOBO",
            "MONO",
            "NUTRIA",
            "LEON",
            "TORO",
            "LORO",
            "ORUGA"
        ],
        "not_founded_words": [
            "NOENCONTRADO"
        ]
    },
    "details": "Password created successfully."
}
```

## 5. Guía de ejecución

### 🔹 Requisitos
- Docker y Docker compose
- Go 1.25 (para desarrollo local)

### 🔹 Configuración

1. Clonar el repositorio:
   ```bash
   git clone https://github.com/Simon3640/letter-soup.git
   cd letter-soup
   ```
2. Crear un archivo `.env` basado en el `.env.example` y ajustar las variables según sea necesario.

3. Construir la imagen de Docker:
   ```bash
   docker-compose -f docker/docker-compose.yml build
   ```

4. Iniciar los contenedores:
   ```bash
   docker-compose -f docker/docker-compose.yml up
   ```

5. Acceder a la aplicación:
   - La aplicación estará disponible en `http://localhost:8080`

6. Acceder a la documentación de la API:
   - Swagger UI: `http://localhost:8080/docs/index.html`

7. Cargue collecciones de postman en `tests/e2e/collection.json` para probar la API.


### 🔹 Uso
- Endpoint para resolver sopa de letras: `POST /api/letter-soup`
- Enviar un JSON con la estructura requerida (rows, columns, grid, words).
```
    {
    "columns": 14,
    "rows": 14,
    "grid": [
        "N,D,E,K,I,C,A,N,G,U,R,O,G,E,S,X,R,Y,K,V,I,I,Q,G,W,Q,O,D,J,A,G,U,A,R,Z,W,B,N,K,O,U,A,M,L,E,L,E,F,A,N,T,E,H,O,G,W,L,O,B,O,N,U,T,R,I,A,O,U,S,U,W,W,O,S,O,G,A,T,O,V,R,T,M,O,H,L,Z,N,C,T,Y,Z,E,O,X,A,U,R,C,E,C,Y,T,I,B,U,R,O,N,S,R,O,C,O,N,E,J,O,Y,U,S,M,R,S,H,T,Y,N,I,F,E,F,P,T,E,Z,O,O,S,F,O,S,S,E,R,P,I,E,N,T,E,F,L,G,P,P,V,D,D,X,U,F,A,L,C,O,N,Y,M,O,N,O,C,U,Q,W,M,A,N,A,T,I,N,N,X,H,E,B,P,M,U,P,E,R,R,O"
    ],
    "user_id": {{userId}},
    "words": [
        "MANATI",
        "LEON",
        "PERRO",
        "LORO",
        "GATO",
        "TORO",
        "CONEJO",
        "ORUGA",
        "TIBURON",
        "ELEFANTE",
        "ALCON",
        "SERPIENTE",
        "JAGUAR",
        "CANGURO",
        "LOBO",
        "MONO",
        "NUTRIA",
        "NOENCONTRADO"
    ]
}
```

## 6. Recursos Especiales
- Se reutilizó la estructura hexagonal de GoProjectSkeleton para organizar la solución.
- Se desarrolló un sistema de transformaciones de matrices con mapeo inverso, lo que evita duplicación de lógica y asegura posiciones correctas en el grid original.
- Se implementó una estrategia de diagonales combinadas para representar de forma compacta las diagonales ↘ y ↙ dentro de una matriz cuadrada.