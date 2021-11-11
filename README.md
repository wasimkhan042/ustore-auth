## U-STORE

## Scenario
Let’s suppose a small startup wants to build a U-STORE which would help the end user to rent common household and entertainment  items. It will be a subscription based system. The users can subscribe an item to rent for months or years. The system needs to be implemented in two microservices that are user-service and store-service.

User-service will be connected to a database to store and retrieve data of users and orders accordingly.

Store-service will have hard coded data about items that are present in the store and available for rent. The store-service should validate the user through internal communication with the user microservice. It is becuase to allow only registered users to search the present items and its details.

Let’s suppose you got the opportunity to join this startup as a backend Intern. You must provide clear and complete documentation about the application architecture and its working. You should be able to handle different routes in your app. A user should be able to register and then can login into the application. Provide a search based functionality. A user should be able to subscribe for an item. The subscription must be checked for expiry.
- Table of Content
[ToC]

## Requirements
### Functional Requirements
Following are the functions which the application will be able to perform.

* #### User signup /signin
User shall be able to sign up with his/her credentials for the first time and then can login to use the app features.

* #### Item operations
User shall be able to subscribe items (monthly or yearly), unsubscribe already rented items and check his/her record of all subscribed items.

* #### Search item
User shall be able to search for an items and can check its monthly and yearly prices.

* #### Profile read and modify details
User shall be able to view its record and modify his/her profile details except email and username for better security and performance of the application.

* #### Authentication calls to store-service
The store microservice should authenticate the call for registered users through internal communication with user microservice.

### Non-functional Requirements

Following are the non-functional requirements;

* #### Security
User password will be encrypted before storing in the database which cannot be understood if retrieved the data from database.

* #### Performance
The application should perform well and provide responsive messages on each operation.

## Detail Design and Architecture

* #### Entity Relationship Diagram
A NoSQL database that automatically handles sharding, replication, and fail-over with very little configuration is good but lack ACID and transaction. Considering the nature of U-STORE data and the single-database microservices architecture, RDBMS is selected for the application.

Folloing is the entity-relationship diagram of the given system.

![](https://i.imgur.com/rDBVudF.png)

In Figure 01, It is shown that user can subscribe multiple items and each subscription will belong to a single user.  The subscription will include single item and one item can be
subcribed by many people if available.

* #### Logical Schema
    A user will be assigned with an auto-id and will have unique email and username. An id helps in fast processing. The logical schema of user is given in Table-1.  ![](https://i.imgur.com/DL0ifCV.png)
    
    An Item table will have item_id as a primary key and will include the monthly and yearly subscription cost, along with an availability field for each item as shown in Table-2.
    ![](https://i.imgur.com/QPOEOJW.png)

    A Subscription table works as an order table and include time stemp fiels to track the subscription through cron jobs. Further it includes foreign keys that of item and user keys to relate the subscribed items to concerned users as shown in Table-03.
    ![](https://i.imgur.com/zGOh1hn.png)


* #### Component Diagram
    The following diagram show high-level interaction of user with the microservices and between microservices.
    ![](https://i.imgur.com/vvRVjSG.png)


* #### Use Case Diagram
    The following use case diagram shows all the possible interaction of the user with the microserverices and internal communication between micerservices.
    
    ![](https://i.imgur.com/mrGutB4.png)



