# Balky - Empowering Citizens, Connecting Communities

Welcome to Balky, your go-to app for accessing vital information and reporting public issues to your local government. With a robust technical foundation and a suite of user-friendly features, Balky is designed to enhance civic engagement and foster collaboration within your community.

## Technical Features:

### RabbitMQ Integration:
Balky leverages RabbitMQ as a reliable message broker and load balancer service, ensuring seamless communication and scalability as the platform grows. With RabbitMQ, Balky can efficiently handle high volumes of messages and distribute workload across multiple servers, guaranteeing optimal performance even during peak usage.

### High-Performance HTTP Server:
Powered by Go's standard library and the lightweight Chi router, Balky delivers lightning-fast response times and exceptional reliability. Our highly optimized HTTP server is engineered for maximum efficiency, enabling users to access information and report issues with minimal latency.

### Modular Monolith Architecture:
Inspired by NestJS, Balky adopts a modular monolith architecture that combines the benefits of encapsulation with the agility of a microservices approach. This architecture allows us to develop and maintain features independently while ensuring seamless integration and cohesive functionality across the platform.

### Dockerized Development and Deployment:
Balky embraces Docker for both development and deployment, streamlining the contribution process and simplifying maintenance tasks. With Docker, developers can spin up isolated environments effortlessly, while deployment becomes a breeze, ensuring a smooth and efficient workflow from development to production.

### Comprehensive Testing:
To uphold our commitment to quality and reliability, Balky undergoes rigorous unit and integration testing. Every function is meticulously tested to ensure that all critical features work as expected, guaranteeing a seamless user experience and robust performance under various scenarios.

## Business Features:

### Report Public Issues:
Empower yourself to make a difference by reporting public issues directly to your local government through Balky. Whether it's a pothole, a broken streetlight, or an environmental concern, Balky provides a convenient platform for citizens to voice their concerns and initiate action.

### Real-Time Monitoring and Contribution:
Stay informed and engaged with all current issues affecting your community. Balky enables users to monitor ongoing developments, contribute updates, and collaborate with fellow citizens and government officials to address issues effectively.

### Statistical Insights:
Gain valuable insights into the state of your community with Balky's comprehensive statistics and reports. Track trends, analyze data, and visualize key metrics to understand the impact of reported issues and identify areas for improvement.

### React and Comment:
Express your support, concern, or feedback by reacting to posts and adding comments to ongoing discussions. Balky fosters open dialogue and encourages constructive engagement, allowing users to participate actively in community initiatives and decision-making processes.

### Enhance Issue Resolution:
Provide additional context and information by adding comments to reported issues. Whether it's providing details, suggesting solutions, or sharing relevant resources, Balky empowers users to contribute valuable insights and facilitate the resolution process.

With Balky, you have the power to shape the future of your community. Join us in building a more connected, resilient, and vibrant society. Together, we can make a difference, one report at a time.

## Running the App:

To run Balky locally, follow these steps:

1. Clone the repository: `git clone https://github.com/mrspec7er/balky.git`
2. Navigate to the project directory: `cd balky`
3. Build the Docker image: `docker build -t balky .`
4. Run the Docker container with Docker Compose: `docker compose watch`
5. Access Balky in your browser at `http://localhost:8000/api`

