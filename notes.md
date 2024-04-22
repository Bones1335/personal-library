
# Personal Library Project

- database inventory of all books
    - unique ID
    - title
    - author
    - summary
    - publication year
    - publisher
    - ISBN
    - genre
    - loan status
    - have/want
    - user reviews

- user profiles
    - admin
    - regular user

## TODOs
    - create database schema
        - book instance
        - author instance
        - genre instance
        - publisher instance
        - user instance
        - loan instance
    - create go server with database calls
        - GET user info
        - GET book info
        - GET loan info
        - UPDATE user info
        - UPDATE book info
        - UPDATE loan info
        - UPDATE loan history
        - ADD new user
        - ADD new book
        - ADD wish list item 
        - ADD loan history
        - DELETE user
        - DELETE book
        - DELETE author
        - DELETE genre
    - create html templates
        - connect templates with database data for each action
    - create search functionality
        - search books
        - search genres
        - search authors
        - search ISBN
    - create book citation
