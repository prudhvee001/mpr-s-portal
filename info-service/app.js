const express = require('express');
const app = express();
const PORT = 3004;

app.get('/info', (req, res) => {
    res.json({
        title: "Student Info",
        description: "View detailed profile and general announcements."
    });
});

app.listen(PORT, () => {
    console.log(`🚀 Info service running on port ${PORT}`);
});
