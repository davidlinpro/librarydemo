# Library Management System
    * API for CRUD of a book ( title, author, isbn, description )
    * Ability to manage books through a web interface 
    * Ability to check in and check out a book  ( guest checks out a book; admin returns a book )
    * Ability to track state changes for a book 
    * Report that contains the current state of all books ( show inventory on admin screen )


# Install:
    * run `make docker`
        * docker-compose.yml: setup docker for app and db
    * run `make seed`
        * db/library.sql: setup database with schema and initial data


# Database (MariaDB):
    * Books: title, author, isbn, description
    * Inventory: all books in the library management system  (id, book_id, status, guest_id, last_updated)
    * Users: name, email, phone, role[deleted=0,admin=1/guest=2, last_login
    * ( activity log of state changes are managed through EventStore )


# Website (Vue.js):
    * http://localhost:8080/


# API (Golang):
    * http://localhost:8081/


# Eventstore:
    * http://localhost:2113/web/index.html#/dashboard  (admin / changeit)





# Web Pages:


## Login View:
    * Login to use system  ( for this sample app, we will not have real password login authentication, we will just use their 4-digit library card id)


## Admin View:
    * Homepage: shows all books in the library ( display book info; TODO: enable pagination and filters )
    * Inventory: 
        * shows all books in inventory ( display book title and sublist of each instance with status and which guest has it checked out )
            * can return a checked out book on behalf of the guest
            * can delete a book from the inventory
        * search by guest_id ( what books are checked out by user)
            * return a book thats checked out
        * search by book title
            * return a book thats checked out
    * Books: 
        * Create: new book ( title, author, isbn, description )
            * Note: can only specify the number of books to add to the library inventory during book creation time. To reduce the number of books in the inventory, admin needs to "delete" books from inventory section.
        * Edit: attributes 
        * Remove: all instances of book
    * Users:
        * Create: new user ( name, email, phone, role[admin/guest] )
        * Edit: attributes
        * Remove: delete user (set user role=0 for deleted)
    * Profile(TODO):
        * See and edit own profile
    * Logout


## Guest View:
    * Homepage: shows all books in the library ( display book info; TODO: enable pagination and filters )
        * Allow guest to click a "check out" button if there are available instances of that book ( or if not available then "add to wish list" / wait list )
    * Wish List (TODO): shows all books we added to wish list ( wait list )
        * Allows guest to view and remove books from their wish list ( send guest an email when book becomes available for checkout via EventStore )
    * Profile (TODO):
        * See and edit own profile
    * Logout




# API (Golang):
    * Notes:
        * for this sample app, we will not use swagger for api documentation; TODO: https://github.com/swaggo/swag 
        * also no ACL.  TODO: implement API permissions based on role https://github.com/casbin/casbin
        * todo: better error logging
    * Auth:
        * Login [ POST: /login ]
        * Logout [ POST: /logout ]
    * User:
        * Get all users  [ GET: /users ]
        * Get user by id [ GET: /users/{id} ]
        * Create new user [ POST: /users ]
        * Edit user by id [ PUT: /users/{id} ]
        * Delete user by id [ DELETE: /users/{id} ]
    * Books:
        * Get all books [ GET: /books ]
        * Get book by id [ GET: /books/{id} ]
        * Create new book [ POST: /books ]
        * Edit book by id [ PUT: /books/{id} ]
        * Delete book by id [ DELETE: /books/{id} ]
    * Inventory:
        * Get all inventory books [ GET: /inventory ]   ( TODO: filtering/sorting/pagination )
        * Get inventory book by id [ GET: /inventory/{id} ]
        * Create new inventory book entry [ POST: /inventory ]
        * Edit: User checkouts a book [ PUT: /inventory/{id} ]
        * Edit: Admin returns a book [ PUT: /inventory/{id} ]
        * Delete: Admin removes a book from inventory [ DELETE: /inventory/{id} ]




# Future Enhancements:
- kubernetes instead of docker-compose
- caddy instead of nginx for frontend  https://caddyserver.com/
- ES for searching 
- Redis for caching
- some CDN for edge delivery
- monitoring:  https://docs.datadoghq.com/integrations/eventstore/
- maybe use kafka instead of eventstore https://blog.softwaremill.com/event-sourcing-using-kafka-53dfd72ad45d   https://stackoverflow.com/questions/17708489/using-kafka-as-a-cqrs-eventstore-good-idea

