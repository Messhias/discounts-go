# Discounts Goo

## How to Start the Project

### Prerequisites

- **Docker** and **Docker Compose** installed.
- **Git** to clone the repository.

### Steps

1. **Clone the Repository**

    ```bash
    git clone https://github.com/your-username/discounts-goo.git
    cd discounts-goo
    ```
    ```

2. **Build and Start Services with Docker Compose**

    ```bash
    docker-compose up --build -d
    ```

    This will:
    - Build the application image.
    - Run tests automatically during the build process.
    - Start the application and PostgreSQL database containers.

3. **Verify Test Results**

    After execution, test results will be available in the `test-results` folder in the project's root directory.
    
    - **Test Coverage:** `coverage.html`
    - **Detailed Results:** `test_results.txt`

4. **Access the Application**

    The API will be available at `http://localhost:8080/products`.

5. **Stop the Containers**

    To stop and remove containers and volumes:

    ```bash
    docker-compose down -v
    ```

## Decisions Made

### Project Structure

The project is organized with a clear separation between application logic (`internal/`), commands (`cmd/`), and tests (`tests/`). This organization facilitates maintenance and scalability of the code.

### Dockerization

- **Dockerfile:** Configured to run tests automatically during the build process. This ensures that only images with tested and approved code are created.
  
- **docker-compose.yml:** Defines the application and database services, and maps volumes to facilitate access to test results.

### Testing

- **Unit Tests:** Located in the `tests/` folder, using mocks to simulate external dependencies, ensuring that tests are fast and independent.
  
- **Test Coverage:** Generation of coverage reports in HTML and detailed test logs to facilitate analysis.

### Mocks

Using the `testify` library to create mocks for repositories, isolating the application logic and ensuring higher reliability of test results.

### Shared Volume

Mapping the `test-results` directory to the host, allowing easy access to test results outside the container. This facilitates viewing coverage reports and test logs directly in your IDE.

---

If you need further assistance or have additional questions, feel free to reach out.