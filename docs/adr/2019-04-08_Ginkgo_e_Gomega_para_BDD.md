# Ginkgo e Gomega para BDD

Adicionado Ginkgo como test suite e Gomega para asserts seguindo o formato BDD.

- **Status:**: accepted
- **Context**: A suite de testes da stdlib dificulta o entendimento e não fornece algo como `BeforeEach`.
- **Decision**: Adicionamos o Ginkgo e Gomega nos testes unitários escritos em Go.
- **Consequences**: Os testes estão mais legíveis e fácil de dar manutenção.
