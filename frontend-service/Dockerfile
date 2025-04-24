# Step 1: Use an official Node.js image as the base
FROM node:16 AS frontend-service

# Step 2: Set the working directory in the container
WORKDIR /app

# Step 3: Copy the package.json and package-lock.json into the container
COPY ./frontend-service/package.json ./frontend-service/package-lock.json ./

# Step 4: Install the app dependencies
RUN npm install

# Step 5: Copy the entire frontend-service project into the container
COPY ./frontend-service ./

# Step 6: Build the React app
RUN npm run build

# Step 7: Expose the port for the frontend
EXPOSE 80

# Step 8: Start the React app
CMD ["npm", "start"]
