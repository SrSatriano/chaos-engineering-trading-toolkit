# Chaos Engineering Toolkit para Trading Bots

Framework estilo Chaos Monkey para homologação: latência de rede, queda de DB e simulação de outage da corretora.

## Stack

- Go, Docker API, Bash

## Experimentos

| Experimento | Efeito |
|-------------|--------|
| `network_latency` | +50–500ms jitter |
| `db_kill` | Postgres pause 30s |
| `broker_outage` | Mock API 503 |

```bash
go run ./cmd/chaosd run --experiment broker_outage --duration 60s
```

## Disaster Recovery

Relatórios em `reports/` após cada run:

- Ordens órfãs detectadas
- Tempo até reconciliação
- MTTR (mean time to recovery)

Ver [docs/DISASTER_RECOVERY.md](docs/DISASTER_RECOVERY.md)

## Métricas de downtime

Prometheus: `chaos_experiment_active`, `orphan_orders_total`, `reconcile_duration_seconds`.

## Estrutura

| Pasta | Função |
|-------|--------|
| `cmd/chaosd/` | CLI |
| `pkg/experiments/` | Injetores |
| `reports/` | DR logs |
