# Base image
FROM python:3.10-slim

# Set working directory
WORKDIR /app

# Copy requirements and install
COPY requirements.txt .
RUN pip install -r requirements.txt

# Copy app code
COPY . .

# Expose port
EXPOSE 3001

# Run the app
CMD ["python", "app.py"]
