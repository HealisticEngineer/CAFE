C.A.F.E. – Client-Aligned Framework for Environment
===================================================

C.A.F.E. is a lightweight configuration management system designed to provide 
client-specific configurations in a structured and scalable way. It ensures 
that the necessary database tables and indexes are created during startup and 
provides an HTTP API for retrieving configuration data.

---------------------------------------------------

Project Layout
```
├── main.go                // Entry point of the application
├── sql_tables.sql         // SQL script for creating database tables and indexes
├── readme.md              // Project documentation
│
├── db\
│   ├── db.go              // Database initialization and table setup
│
├── handlers\
│   ├── config_handler.go  // HTTP handler for configuration
│
├── utils\
│   ├── ip_utils.go        // Utility functions for IP checks
```
---------------------------------------------------

Features
--------

- Database Initialization: Automatically sets up the required tables and indexes 
  using the `sql_tables.sql` script.
- Configuration API: Provides an HTTP endpoint (`/config`) to retrieve 
  client-specific configuration data.
- IP Filtering: Ensures only requests from allowed IP ranges are processed.
- Modular Design: Organized into separate packages for better maintainability 
  and scalability.

---------------------------------------------------

## Getting Started

### 1. Set Up the Database
- Ensure you have a SQL Server instance running.
- Create environment-specific `.env` files in the root of the project:
  - `.env.development` for development:
    ```
    SQLSERVER_CONN=sqlserver://dev_user:dev_pass@localhost/DevDB
    ```
  - `.env.production` for production:
    ```
    SQLSERVER_CONN=sqlserver://prod_user:prod_pass@prod-server/ProdDB
    ```

### 2. Set the Environment
- Set the `APP_ENV` environment variable to specify the environment:
  - For development:
    ```bash
    export APP_ENV=development
    ```
  - For production:
    ```bash
    export APP_ENV=production
    ```

### 3. Run the Application:
   - Build and run the application using:
     go run main.go

### 4. Access the API:
   - Use the `/config` endpoint to retrieve configuration data for a client. 
     Example:
     curl -H "X-Client-Name: VM_Server1" http://localhost:3000/config

---------------------------------------------------

Contributing
------------

Contributions are welcome! Feel free to submit issues or pull requests to 
improve the project.

---------------------------------------------------

License
-------

This project is licensed under the MIT License.