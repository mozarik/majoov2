# Majoo-V2 Test Case

This is an implementation of the test case from majoo implemented in Go. 

## TODO

### Documentation
---
- [ ] Entity Relationship Diagram
- [ ] Data Manipulation Language
- [ ] Activity Diagrams
- [ ] Use Case Diagrams

---
### Implementation
---

- [x] Make the Physical Model with GORM 
  - [x] Connect to Postgre
  - [x] Setup basic model and do automigrate with GORM db.Automigrate
    - [x] Products To Merchant 
    - [x] User To Merchant
    - [x] User To Outlet
    - [x] Merchant To Outlet
    - [x] MerchantProduct Table to Merchant and Product
    - [x] ProductOutlet to MerchantProduct
    - [x] Outlet to ProductOutlet   
-----
- [x] Make login and logout functionality 
  - [x] Implement the JWT token
  - [x] Implement the logout to clear the cache in the cookie
  - [x] Implement simple GET to dummy endpoint `/admin` that need JWT to access return not authorize http error code
- [x] Implement httponly to save the JWT in cookie 
  - [x] Make sure its automatically send the access token to Authorization header header when hitting our dummy `/admin` 
  - [x] Make sure after we hit `/logout` endpoint we cannot access the `/restricted` endpoint again

----  
- [x] Implement upload functionality
  - [x] Using direct binary and get the base64 encoded image binary ?
  - [ ] gRPC ? lol No 
  - [x] or just use echo upload file functionality

----
- [ ] Implement form-validation for all request body
  - [ ] Is there a way to better validate all the request body ? 
  - [ ] Use validator package ?
 
----
- [ ] CRUD of Master Data User
  - [x] Create (Register) User
  - [ ] Read User 
    - [ ] Read All User ? 
  - [ ] Update User
    - [ ] Make sure the user is log in (JWT)
  - [ ] Delete User
    - [ ] Make sure the user who delete is the user itself (JWT)
  ---
- [ ] CRUD of Master Data Produk
  - [ ] As a merchant i can add (CREATE) product (JWT)
  - [ ] As a merchant i can RETRIEVE all my products (JWT)
  - [ ] As a merchant i can UPDATE a product(s) (JWT)
  - [ ] As a merchant i can DELETE a product(s)  (JWT)