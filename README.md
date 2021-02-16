# Traffic Editor API

Docker image available on dockerhub: docker pull weijuin/go-task:v1.9

Git clone[https://github.com/weijuinlee/go-task.git]

<br>

## Graph 

<br>

### 1. POST - /editor/graph

Description: Create new graph

<br>

### 2. PUT - /editor/graph/detailed/<graphID>

Description: Update one graph

<br>

```json
{
   "mapVerID": "b2a546f9-b7a1-4623-8c93-8a574b8db1f6",
   "scale": 0.05,
   "Location": "NCS A1",
   "Building": "NCS",
   "level": "1":,
   "lanes": {
       "0.0": {"1.0": 58.54130575371151},
       "0.1": {"1.0": 58.54130575371151},
       "1.0": {"0.0": 58.54130575371151, "2.0": 55.643275340786325},
       "1.0": {"0.1": 58.54130575371151, "2.0": 55.643275340786325},
       "2.0": {"1.0": 55.643275340786325, "3.0": 479.9329727520787},
       "3.0": {"2.0": 479.9329727520787, "4.0": 259.0517327067432, "6.0": 214.95190436162838},
       "4.0": {"3.0": 259.0517327067432, "5.0": 226.86823068055912},
       "5.0": {"4.0": 226.86823068055912, "6.0": 254.25660341657715},
       "6.0": {"3.0": 214.95190436162838, "5.0": 254.25660341657715}
   },
   "vertices": {
       "0": {"x": 35.57377049180328, "y": 727.8171428571429, "type": 0, "name": "Home", "intersectionDegree": 2},
       "0.0": {"x": 35.57377049180328, "y": 727.8171428571429, "type": 0, "name": "Home", "intersectionDegree": 2},
       "0.1": {"x": 35.57377049180328, "y": 727.8171428571429, "type": 0, "name": "Home", "intersectionDegree": 2},
       "1": {"x": 94.09931253626547, "y": 729.1756007149452, "type": 1, "name": "A", "intersectionDegree": 2},
       "2": {"x": 95.24531538675106, "y": 784.8070735348988, "type": 2, "name": "B", "intersectionDegree": 2},
       "3": {"x": 574.9178003414773, "y": 800.6173711919746, "type": 0, "name": "C", "intersectionDegree": 2},
       "4": {"x": 593.2786885245902, "y": 542.2171428571429, "type": 0, "name": "D", "intersectionDegree": 4},
       "5": {"x": 819.4102139602585, "y": 560.4853473738731, "type": 0, "name": "E", "intersectionDegree": 2},
       "6": {"x": 789.5140055094371, "y": 812.9781930174668, "type": 0, "name": "F", "intersectionDegree": 2}
   }
}
```
<br>

### Attributes:
### "type": 0 = null
### "type": 1 = door
### "type": 2 = elevator
### "type": 3 = waiting point
**Parent node "0" has a capacity of 2. Child nodes are "0.0" and "0.1"**


<br>

### 3. GET - /editor/graph/detailed

Description: Get all graphs(All details)

<br>

### 4. GET - /editor/graph/nondetailed

Description: Get all graphs(Names and location only)

<br>

### 5. GET - /editor/graph/detailed/<graphID>

Description: Get one graph(All details)

<br>

### 6. DELETE - /editor/graph/<graphID>

Description: Delete one graph

<br>

## Collection

<br>

### 1. POST - /editor/collection

Description: Create new collection

<br>

```json
{
   "name": "Sentosa Mall"
}
```

<br>

### 2. GET - /editor/collection

Description: Get all collection

<br>

### 3. DELETE - /editor/collection/<graphID>

Description: Delete one collection

<br>

## Patrol

<br>

### 1. POST - /editor/patrol

Description: Create new patrol

<br>

### 2. PUT - /editor/patrol/<patrolID>

Description: Update one patrol

<br>

```json
{
   "locationID": 1,
   "mapVerID": "b2a546f9-b7a1-4623-8c93-8a574b8db1f6",
   "name": “Patrol1”,
   "points": ["0", "1", "2", "3", "4", "5", "6"]
}
```

<br>

### 3. GET - /editor/patrol

Description: Get all patrol routes

<br>

### 4. GET - /editor/patrol/<graphID>

Description: Get all patrol routes linked to one graph

<br>

### 5. DELETE - /editor/patrol/<patrolID>

Description: Delete one patrol route

## Robots

<br>

### 1. GET - /editor/robot

Description: Get all robots available

<br>