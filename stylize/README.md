# ASCII Art Web Stylize

This project is a simple web server that serves static files, handles basic routes, and generates ASCII art based on input. The server is built using Go and provides several endpoints for viewing ASCII art and other static content.

## Features

- **Static File Serving:** Serves static files like CSS, images, and fonts.
- **Home Page:** Displays the home page with basic content.
- **ASCII Art:** Generates and displays ASCII art from given inputs.
- **Error Handling:** Serves custom error pages for common HTTP errors (400, 404, 500).
- **About Page:** Provides information about the application and its authors.
- **Authors Page:** Provides information about authors.

## Endpoints

- `/`: Serves the home page (`index.html`).
- `/static/`: Serves static files from the `static` directory.
- `/ascii-art`: Generates ASCII art based on input data.
- `/authors`: Displays a list of authors.
- `/about`: Provides information about the application.

## Directory Structure

- `templates/`: Contains HTML templates for different pages and error handling.
  - `index.html`: Home page template.
  - `about.html`: About page template.
  - `authors.html`: Authors page template.
  - `400.html`, `404.html`, `500.html`: Error page templates.
- `static/`: Contains static assets such as CSS files, images, and fonts.
- `main.go`: The main application file containing route handlers and server setup.
- `ascii.go`: Contains functions for generating and handling ASCII art.

## Setup

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/A-fethi/ascii-art-web.git
   cd ascii-art-web/stylize
    ```
2. **Install Dependencies:**

This project uses Go's standard library, so no additional dependencies are required beyond Go itself.

3. **Run the Application:**
```bash
    go run main.go
    Access the Server:
    Open your web browser and navigate to http://localhost:8080 to access the home page and explore the other endpoints.
```

4. **Error Handling**
The application provides custom error pages for the following status codes:

- 400 Bad Request: templates/400.html
- 404 Not Found: templates/404.html
- 500 Internal Server Error: templates/500.html

5. **Notes**
- Ensure that the templates and static directories contain the necessary files for the application to function correctly.
- The ascii.go file contains the logic for generating ASCII art from input text.