# Nova arquitetura de codigo

Foi criada uma nova arquitetura baseada nos design patterns Layered e Repository.

- **Status:**: accepted
- **Context**: A arquitetura inicial não estava muito clara e não seguia os principios do SOLID, tornando o codigo dificil de ser testado e de dar manutencao.
- **Decision**: Foi criado um novo repositorio no git e separamos uma sprint para adicionar as funcionalidades ja desenvolvidas na arquitetura anterior na atual.
- **Consequences**: O codigo possui testes unitarios e o code review passou a ficar mais simples. Agora os testes rodam no CircleCI a cada push.
