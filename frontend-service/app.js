document.addEventListener("DOMContentLoaded", function () {
    fetch("/user")
        .then(response => response.json())
        .then(data => {
            document.getElementById("user-info").innerHTML = `
                <h2>User Information</h2>
                <p>Name: ${data.name}</p>
                <p>Email: ${data.email}</p>
            `;
        });

    fetch("/results")
        .then(response => response.json())
        .then(data => {
            document.getElementById("results-info").innerHTML = `
                <h2>Results</h2>
                <p>Subject: ${data.subject}</p>
                <p>Grade: ${data.grade}</p>
            `;
        });

    fetch("/notifications")
        .then(response => response.json())
        .then(data => {
            document.getElementById("notifications-info").innerHTML = `
                <h2>Notifications</h2>
                <p>${data.message}</p>
            `;
        });

    fetch("/info")
        .then(response => response.json())
        .then(data => {
            document.getElementById("info-info").innerHTML = `
                <h2>Student Info</h2>
                <p>${data.description}</p>
            `;
        });
});
