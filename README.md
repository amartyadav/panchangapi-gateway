## 🚧 Pre-Release Notice 🚧

**This project is currently in a pre-release stage and is not ready for production use.**

We are actively developing this API gateway in public, and the codebase is subject to significant changes. Public contributions are not accepted. If you wish to contribute, please contact me at my [email](mailto:contact@amartyadav.tech) please be aware that the project may not yet be stable or fully functional.

Stay tuned for updates as we continue to work towards a stable release!

---

# API Gateway for Panchang and Hindu Festival Data

Welcome to the API Gateway for accessing Panchang (Hindu calendar) and Hindu festival data. 
This project is under active development and aims to provide secure and efficient public access to the core Panchang API used by the QuickPanchang macOS app.

Check out the project development plan [here](https://aatmamartya.notion.site/8baed2e3ef3143a694025b0fd7ee1d3b?v=08bc8730f77d4627a2bad9378dcd0f73&pvs=74) to get updates on the project's development.

## Project Description

This project is an API gateway built using Go (Golang) following a microservices architecture. The main components include:

1. **User Registration and Authentication**: System with email verification to ensure secure access.
2. **API Key Management**: Enables users to generate and manage API keys for accessing Panchang and festival data.
3. **Rate Limiting and Request Tracking**: Monitors and controls API usage to prevent abuse.
4. **Caching Mechanism**: Optimizes data retrieval for better performance.
5. **Proxy Functionality**: Forwards authenticated requests to the existing core Panchang API.

### Key Features

- **Persistent Data Storage**: Uses PostgreSQL for storing user information, API keys, and usage analytics.
- **Temporary Data Storage**: Utilizes Redis for OTPs, session data, and caching.
- **Containerization**: Dockerized setup for easy deployment and orchestration with Docker Compose.
- **Secure and Scalable**: Incorporates password hashing, API key generation, and protection against fraudulent signups.

### Endpoints

- **User Registration**: Allows new users to register.
- **Email Verification**: Confirms user registration via email.
- **Data Retrieval**: Provides Panchang and Hindu festival information.
- **Usage Statistics**: Tracks API usage and provides analytics.

## QuickPanchang macOS App

To see the core Panchang API in action, check out the [QuickPanchang macOS app](https://apps.apple.com/in/app/quickpanchang-hindu-calendar/id6475807190?mt=12). This app utilizes the core API to provide comprehensive Panchang and festival data to users.

## Getting Started

Guide will be added after release.

## Join us on Matrix
Head over to our Matrix Space for discussions, suggestions and chat: [QuickPanchang on Matrix](https://matrix.to/#/#panchang:matrix.org)
## Contribution

Not currently accepting any contributions from the community. Sorry.
This will be enabled after v1 release. Right now this project is being developed in public by internal developers, namely:
- [Amartya Yadav](https://github.com/amartyadav)
- [Ishika Agarwal](https://github.com/ishikaubc)

## License

This project is licensed under the MIT License.

## Future Plans

We might open source the core Panchang API in the future. Stay tuned for updates!

## Contact

For any questions or suggestions, feel free to reach out at my [email](mailto:contact@amartyadav.tech) or open an issue.

## Activity

![](https://repobeats.axiom.co/api/embed/04e5f4e05f2daa51f29a1e052fd5ee7c1197c9c5.svg "Repobeats analytics image")

---

Thank you for checking out our project! We hope you find it useful and look forward to your contributions.

---

Happy Coding!
