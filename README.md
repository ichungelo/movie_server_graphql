## HOW TO QUERY

### GET
- All movies
  ```json
  query {
    movies{
      id
      title
      year
      poster
      overview
    }
  }
  ```
- Detail movie
  ```json
  query{
    detailMovie(input:{
      id:"ID"
    }){
      id
      title
      year
      poster
      overview
      reviews{
        id
        movieId
        userId
        username
        review
        createdAt
        updatedAt
      }
    }
  }

  ```

### POST
- Register
  ```json
  mutation{
    register(input:{
      username:"string"
      email:"string"
      firstName:"string"
      lastName:"string"
      password:"string"
      confirmPassword:"string"
    })
  }
  ```
- Login
  ```json
  mutation{
    login(input:{
      username:"string"
      password:"string"
    })
  }
  ```
- New review
  ```json
  mutation{
    newReview(input:{
      movieId:"ID"
      review:"string"
    })
  }
  ```

### PUT
- Edit review
  ```json
  mutation{
    editReview(input:{
      movieId:"ID"
      review:"string"
    })
  }
  ```

### DELETE
- Delete review
  ```json
  mutation{
    deleteReview(input:{
      id:"ID"
    })
  }
  ```