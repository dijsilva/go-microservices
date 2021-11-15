# Go - Microsserviços

Microsserviços em Go para estudo da linguagem

## Descrição
Trata-se de uma aplicação para análises de dados espectrais, oriundos da tecnologia NIRS, utilizando aprendizado de máquina para classificaçao dos espectros.

Pelo frontend, construído em React, é possível se cadastrar e enviar espectros (por meio de uma arquivo .csv)

O espectro é então enviando para um microserviço (`spectra`) que armazena este espectro em um banco de dados (MongoDB) e publica uma mensagem no RabbitMQ com o id do espectro.

Um microserviço construído para realizar as predições (`spectra-prediction`) consome a mensagem publicada na fila do RabbitMQ. Nesta mensagem há o id do espectro enviado e com ele este microserviço solicita o espectro para o microserviço que armazena os espectros (`spectra`). Depois de buscar as informações do espectro, este microserviço realiza alguns preprocessamentos nos dados para que então seja de fato utilizado como input para que o modelo realize a predição. Neste projeto, foi utilizado RandomForest como método de aprendizado.

Depois de realizar a predição, o modelo faz uma requisição para o microserviço `spectra` enviando os dados da predição. O microserviço recebe estes dados e armazena no banco de dados, para que possa ser consumido para frontend.
## Arquitetura

![Arquitetura](./docs/images/spectra.png)


## Instalação

### Pré requisitos
- Docker (versão >= 20.19.9) 
- docker-compose (versão >= 1.29.2)
- git (versão >= 2.33.0)

Para executar o projeto, faça o download do projeto executando `git clone https://github.com/dijsilva/spectra-microservices.git`

Entre na pasta do projeto que contém os scripts `cd spectra-microservices/scripts`

Faça com que o script possa ser executado utilizando `chmod +x init_app.sh`

Execute o script que irá executar o projeto `./init_app.sh`