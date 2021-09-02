# GraphQL demo for Search 

## schema 

```graphql

interface SearchTamplate {
  id: ID! 
  type: String! ## search result type for example , 
  		## list data, ScoreCard data Poster Data
}

type ListData implements SearchTamplate {
  id: ID! 
  type: String!
  poster:[Poster!]
}
type Poster {
  id: ID!
  name: String!
  url: String!
}

type ScoreCardData implements SearchTamplate {
  id: ID! 
  type: String! 
  score_card:[ScoreCard!]
}
type ScoreCard {
  id: ID!
  teams:[Team!]
}

type Team { # teams data 
  id: ID!
  name: String!
}

type Query {
  search(search_word: String!): [SearchTamplate!]
}

```

## query using graphql


- query with score card data response

```graphql

query {
  search(search_word:"IPL"){
    id
    type
    ... on ScoreCardData {
      id 
      score_card{
        teams{
          id
          name
        }
      }
      
    }
  
    
    ... on ListData {
      poster {
        id 
        name
        url
      }
    }
  }
}


```
- score card response data 

```json
{
  "data": {
    "search": [
      {
        "id": "s1",
        "type": "ScoreCard",
        "score_card": [
          {
            "teams": [
              {
                "id": "player1",
                "name": "first player"
              },
              {
                "id": "player2",
                "name": "second player"
              }
            ]
          }
        ]
      }
    ]
  }
}
```

- query with list data response

```graphql
# return with list data 
query {
  search(search_word:"yeh"){
    id
    type
    ... on ScoreCardData {
      id 
      score_card{
        teams{
          id
          name
        }
      }
      
    }
    
    ... on ListData {
      poster {
        id 
        name
        url
      }
    }
  }
}

```

- list data response

```json
{
  "data": {
    "search": [
      {
        "id": "1",
        "type": "List",
        "poster": [
          {
            "id": "p1",
            "name": "luka",
            "url": "www.hotstar.com/movie/luka"
          },
          {
            "id": "p2",
            "name": "out-guy",
            "url": "www.hotstar.com/movie/out-guy"
          }
        ]
      }
    ]
  }
}



```
