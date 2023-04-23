# go-command-console
Aplicação que lê dados de collections no MongoDB e envia para exchange com routingkeys diferentes

[![LinkedIn][linkedin-shield]][linkedin-url]


<!-- SOBRE O PROJETO -->
## Sobre o Projeto

Aplicação que lê dados de duas collections no MongoDB e utilizando canais e goroutines faz o envio paralelizado para uma exchange com DLQ usando diferentes routingkeys 

### Utilizando

* [![Go][Go-badge]][Go-url]
* [![MongoDB][MongoDB-badge]][MongoDB-url]
* [![RabbitMQ][RabbitMQ-badge]][RabbitMQ-url]
* [![Docker][Docker-badge]][Docker-url]
* [![VS Code][VSCode-badge]][VSCode-url]


<!-- GETTING STARTED -->
## Getting Started

Instruções para execução da aplicação

### Prerequisites

Executar o comando para inicializar o MongoDB, RabbitMQ e a aplicação na porta selecionada
* docker
  ```sh
  docker-compose up -d
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/your_username_/Project-Name.git
   ```
2. exec
   ```sh
   go run main.go
   ```


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/felipe-fernandes-fca/
[Go-url]: https://golang.org/
[Go-badge]: https://img.shields.io/badge/go-%2300ADD8.svg?style=flat&logo=go&logoColor=white
[MongoDB-badge]: https://img.shields.io/badge/mongodb-%234ea94b.svg?style=flat&logo=mongodb&logoColor=white
[MongoDB-url]: https://www.mongodb.com/
[RabbitMQ-badge]: https://img.shields.io/badge/rabbitmq-%23ff6600.svg?style=flat&logo=rabbitmq&logoColor=white
[RabbitMQ-url]: https://www.rabbitmq.com/
[Docker-badge]: https://img.shields.io/badge/docker-%230db7ed.svg?style=flat&logo=docker&logoColor=white
[Docker-url]: https://www.docker.com/
[VSCode-badge]: https://img.shields.io/badge/VS_Code-007ACC?style=flat&logo=visual-studio-code&logoColor=white
[VSCode-url]: https://code.visualstudio.com/
