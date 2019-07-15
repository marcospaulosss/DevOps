# Microserviço ecommerce

A idéia desse microserviço é gerenciar produtos, pedidos e fazer a integração com a multipagg.

- **Status:**: accepted
- **Context**: O app de áudio só vai poder ser lançado em produção quando houver integração com um gateway de pagamento para permitir que os usuários possam comprar uma assinatura. A Estratégia já possui um contrato com a Mundipagg e por esse motivo que vamos utilizar esse gateway de pagamento.
- **Decision**: Vamos colocar o gerenciamento de produtos, pedidos e integração com a Mundipagg no mesmo microserviço pois ainda há muita incerteza sobre o que precisamos criar e temos um prazo de 1 sprint para terminar isso.
- **Consequences**: O severino vai sofrer alterações.
