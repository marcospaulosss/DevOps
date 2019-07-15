# Documentacao utilizando API Blueprint

Estamos utilizando o Aglio para converter os arquivos markdown da documentacao da API do severino para HTML e jogando o arquivo no S3. 

- **Status:**: accepted
- **Context**: Estavamos utilizando o Postman SaaS mas o time encontrou problemas ao acessar com o servico indisponivel e nao era pratico de se atualizar.
- **Decision**: O formato API Blueprint permite utilizar markdown para documentar e armazenar os documentos no mesmo repositorio git do projeto. Durante o CI das branches master ou develop, é gerado o HTML e enviado ao S3. 
- **Consequences**: Cada desenvolvedor que alterar a assinatura da API deve modificar o arquivo da documentação. Foi necessário adicionar mais um passo no CI para renderizar a documentação e enviar ao S3.
