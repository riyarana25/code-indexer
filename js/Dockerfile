# Use the official Node.js image
FROM node:24-slim

# Set working directory
WORKDIR /analyzer

# Copy package.json and install dependencies
COPY package.json package-lock.json ./
RUN npm ci

# Copy the rest of the application code
COPY . .

# Command to run the analyzer
CMD ["node", "analyzer.js"]
