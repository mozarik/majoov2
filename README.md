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

- [x] Make the Physical Model with  GORM 
  - [x] Connect to Postgre
  - [x] Setup basic model and do automigrate with GORM db.Automigrate
- [ ] Make login functionality 
  - [ ] Implement the JWT token
  - [ ] Implement the logout to clear the cache in the cookie
  - [ ] Implement simple GET to dummy endpoint `/restricted` that need JWT to access return not authorize http error code
- [ ] Implement httponly to save the JWT in cookie 
  - [ ] Make sure its automatically send the `Authorization` header when hitting our dummy `/restricted` 
  - [ ] Make sure after we hit `/logout` endpoint we cannot access the `/restricted` endpoint again
  
- [ ] Implement upload functionality
  - [ ] Using direct binary and get the base64 encoded image binary ?
  - [ ] gRPC ? 
  - [ ] or just use echo upload file functionality
- [ ] Implement form-validation

- [ ] CRUD of Master Data User
  - [ ] 
- [ ] CRUD of Master Data Produk