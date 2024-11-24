Calculator - CLI Web Application
This project is a simple Calculator built using Go for the backend and HTML, CSS, JavaScript for the frontend. It allows users to perform basic arithmetic operations like addition, subtraction, multiplication, and division directly in their web browsers.

Features
Simple arithmetic calculations: Addition, Subtraction, Multiplication, and Division.
Responsive design, optimized for both desktop and mobile devices.
Input validation for numbers and operators.
User-friendly interface with interactive buttons for calculations.
Technology Stack
Backend: Go (Golang)
Frontend: HTML, CSS, JavaScript
Web Server: Go's net/http package for serving static files and handling requests.
Installation
Prerequisites
Make sure you have Go installed on your system. You can download it from the official Go website: https://golang.org/dl/.

Steps to Run the Project Locally
Clone the repository to your local machine:

git clone https://github.com/PawanKhanal/GO-Projects.git
Navigate to the project folder:

cd Calculator
Install any dependencies if needed (Go automatically handles dependencies).

Run the Go server:

go run main.go
Open a web browser and navigate to http://localhost:8080 to use the calculator.

Project Structure
The project follows a simple directory structure:

Calculator/
│
├── static/
│   ├── script.js       # JavaScript file for frontend logic
│   └── style.css       # CSS file for styling
│
├── main.go             # Go backend server file
├── index.html          # Main HTML file
└── README.md           # This readme file
main.go
This file contains the Go backend logic. It serves the HTML, CSS, and JavaScript files, and processes the arithmetic calculations when the user interacts with the web page.

index.html
The main HTML page that contains the structure of the calculator, including buttons for numbers and operations.

script.js
This JavaScript file handles user inputs, performs calculations, and updates the displayed result dynamically without refreshing the page.

style.css
CSS file for styling the calculator, making it mobile-friendly and visually appealing.

Usage
Once the server is running, open http://localhost:8080 in a web browser.
The calculator interface will appear with buttons for digits (0-9), operations (+, -, *, /), and an "=" button to display the result.
Use the calculator by clicking the buttons, and the result will appear immediately after pressing "=".
