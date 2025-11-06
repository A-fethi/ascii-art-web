# 🎨 Ascii-Art Web

**Ascii-Art Web** is a web-based extension of the [ASCII-Art](https://github.com/A-fethi/ascii-art) project, allowing users to generate ASCII art directly from a browser using a graphical user interface (GUI).  
This project demonstrates how to integrate Go-based text processing with an HTTP server and dynamic HTML templates.

---

## 📜 Description

Ascii-Art Web provides a simple web interface to:  

- Enter text via a text input field.  
- Select a banner style (`standard`, `shadow`, or `thinkertoy`).  
- Generate ASCII art from the input and display the result on the webpage.  

The backend is written in Go and processes the text using the same logic as the ASCII-Art project.

---

## 👤 Authors

- **Abderrahmane Fethi** – Junior Full-Stack Developer  
  🌍 [LinkedIn](https://www.linkedin.com/in/abderrahmane-fethi)

---

## ⚙️ Usage

1. **Clone the repository**:

```bash
git clone https://github.com/A-fethi/ascii-art-web.git
cd ascii-art-web
```

2. **Run the Go server:**
```bash
go run main.go
```

3. **Open the browser:**

Navigate to http://localhost:8080

4. **Use the interface:**

- Enter text in the input field.

- Choose a banner style.

- Click the "Generate" button to see the ASCII art output.

## 🧰 Technologies Used

- Language: Go (Golang)

- Packages: Standard library only (net/http, html/template, os, fmt)

## 🎯 Learning Outcomes

- Building a **web server in Go**

- Handling **HTTP GET and POST requests**

- Integrating **HTML templates** with Go backend logic

- Using **ASCII-art algorithms** in a web environment

- Managing **HTTP status codes** and error handling