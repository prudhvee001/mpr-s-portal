# Step 1: Base image
FROM node:16 AS frontend-service

# Step 2: Set working directory
WORKDIR /app

# Step 3: Copy package files
COPY package.json package-lock.json ./

# Step 4: Install dependencies
RUN npm install

# Step 5: Copy the rest of the code
COPY . .

# Step 6: Build
RUN npm run build

# Step 7: Expose port
EXPOSE 80

# Step 8: Start
CMD ["npm", "start"]
