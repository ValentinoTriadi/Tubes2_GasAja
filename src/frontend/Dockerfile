FROM node:20-alpine AS development

# Create app directory
WORKDIR /app/frontend

# Install app dependencies
COPY package*.json ./
RUN npm install

# Bundle app source
COPY . .

# Expose the port the app runs in
EXPOSE 3000

# Serve the app
CMD ["npm", "run", "dev"]