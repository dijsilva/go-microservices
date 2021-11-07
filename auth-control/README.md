# auth-control

Cria, deleta e gerencia os tokens (JWT) gerados para cada usuário com base nas informações do usuário e do tipo de token gerado.

- [x] Gera um novo token do tipo `PROFILE`_LOGIN sempre que um usuário faz login na plataforma
- [x] Disponibiliza um endpoint para que outros microserviços possam checar se um token é valido e se um determinado endpoint pode ser acessado ou não com base no tipo de token enviado (`PROFILE`_LOGIN)
- [x] Disponibiliza um endpoint para que seja possível deletar um token (caso exista) quando um usuário é deslogado.