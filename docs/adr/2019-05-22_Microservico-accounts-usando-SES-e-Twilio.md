# Microserviço accounts usando SES e Twilio

Microserviço para autenticação, autorização e notificação via SMS e email.

- **Status:**: accepted
- **Context**: Os clientes não vão utilizar autenticação via password. Durante o cadastro eles deverão informar o email e celular e validar os códigos recebidos. Isso melhora a segurança pois evita que sejam utilizadas senhas fracas e compartilhadas.
- **Decision**: A princípio esse serviço também vai ser responsável por enviar as notificações/códigos via email e SMS. Para email será utilizo o AWS SES que possui um limite de 50 mil mensagens por mês. Para SMS vai ser o Twilio devido à ter uma melhor API e documentação.
- **Consequences**: Talvez o preço do Twilio se torne um problema.
