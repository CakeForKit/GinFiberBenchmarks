from pandora_analize import main_throughput
from prom_analize import main_prom
from logs_analize import main_logs

if __name__ == '__main__':
    main_logs()
    main_throughput()
    main_prom()