# Stage 1: Build the frontend (if it's a React/Vue/Angular app)
FROM node:18-alpine AS build

# Set working directory
WORKDIR /app

# Copy source files
COPY . .

# Install dependencies and build the project
RUN npm install && npm run build

# Stage 2: Serve the build using nginx
FROM nginx:alpine

# Copy built frontend from stage 1
COPY --from=build /app/dist /usr/share/nginx/html  # For Angular
# COPY --from=build /app/build /usr/share/nginx/html  # For React

# Expose port
EXPOSE 80

# Start nginx
CMD ["nginx", "-g", "daemon off;"]
