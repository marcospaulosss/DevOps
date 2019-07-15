# Envio de logs para o Rsyslog

Quando houver um servidor rsyslog disponível, os logs serão enviados para lá.
Inicialmente os logs eram enviados para um Elasticsearch na AWS e visualizados com o Kibana. Primeiro foi feito uma POC com o fluentd, porém ele consumia bastante memória. 
Foi trocado pelo fluentbit, funcionando como Sidercar no Kubernetes, colhendo os logs de todos os containers e enviando ao ES.
Como não fazia sentido salvar os logs de todos os containers, foi adicionado o fluentbit como um agente rodando em background no container de cada microserviço.

O próximo problema foi a visualização no Kibana. A versão fornecida pela AWS é antiga, lenta e tosca. Como o Cognito não está disponível na região sa-east-1, não dava para definir autenticação no Kibana e por isso era perigoso deixá-lo exposto na web. 
Configurar um Kibana bare metal requer uma máquina EC2 no mínimo uma small e é preciso fazer proxy para a API do ES da AWS. Configurar um ES bare metal também requer uma instância boa. Resumindo, as alternativas ELK e EFK para log exigem uma boa configuração de máquinas e demanda um esforço maior de gerenciamento.

A solução escolhida foi utilizar o Rsyslog. Basta configurar um server em qualquer server Linux ou mesmo rodando via container. O Rsyslog também possui um plugin para salvar os logs no Postgres. Golang já possui suporte ao syslog, então basta alterar a implementação de log no código do backend e configurar para enviar os logs ao servidor do Rsyslog. Assim, só vai ser preciso desenvolver uma interface gráfica para visualizar e filtrar os logs salvos no Postgres.

- **Status:**: accepted
- **Context**: Melhoria nos logs
- **Decision**: Alterei a implementação do log adicionando um hook de syslog no Logrus
- **Consequences**: Envia os logs para o servidor de log via UDP e caso o servidor esteja indisponível, envia apenas para o stdout.
