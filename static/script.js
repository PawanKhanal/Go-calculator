// Add number to the input field
function appendNumber(num) {
    document.getElementById("expression").value += num;
}

// Add operator to the input field
function appendOperator(operator) {
    document.getElementById("expression").value += operator;
}

// Clear the input field
function clearInput() {
    document.getElementById("expression").value = '';
}

// Calculate the result of the expression
function calculateResult() {
    var expression = document.getElementById("expression").value;

    // Sending the expression to the backend for evaluation
    fetch('/calculate', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ expression: expression })
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById("expression").value = data.result; // Show result in the input field
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Invalid input. Please enter a valid expression.');
    });
}
