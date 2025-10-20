from typing import List
import os
import matplotlib.pyplot as plt
import pandas as pd
from conf import TYPE_LOGS, DIRECTORY_METRICS

class MetricParams:
    def __init__(self, graph_label: str, outputFilename: str):
        self.graph_label = graph_label
        self.outputFilename = outputFilename 


FILENAMES = {
    "container_cpu_usage_seconds_total.txt": MetricParams("Загрузка процессора в процентах", "cpu_usage_seconds_total.png"),
    "container_memory_usage_bytes.txt": MetricParams("Потребление памяти контейнером  (байты)", "memory_usage_bytes.png"),
    "goroutines_count.txt": MetricParams("Количество горутин", "goroutines_count.png"),
    "memory_allocations_bytes_heap.txt": MetricParams("Память выделенных объектов в куче (байты)", "memory_allocations_bytes_heap.png"),
    "memory_allocations_bytes_stack.txt": MetricParams("Память в стеках горутин (байты)", "memory_allocations_bytes_stack.png"),
    "total_http_request_counter.txt": MetricParams("Количество обработанных HTTP-запросов", "total_http_request_counter.png")
}
window = 10



def ParseFile(filename: str, res_list: List[List[float]]):
    ir = 0
    try:
        with open(filename, 'r') as f:
            line = f.readline().strip()
            assert(line)
            parts = line.split()
            assert(len(parts) == 2)
            start_timestamp, metric = int(parts[0]), float(parts[1])
            if ir >= len(res_list):
                res_list.append([])
            res_list[ir].append(metric)
            ir += 1

            for line in f:
                line = f.readline().strip()
                if line == '':
                    continue
                parts = line.split()
                assert(len(parts) == 2)

                timestamp, metric = int(parts[0]), float(parts[1])
                if ir >= len(res_list):
                    res_list.append([])
                res_list[ir].append(metric)
                ir += 1
    except (ValueError, IndexError) as e:
        print(f"Ошибка парсинга строки: {line} {e}")
        raise
    
def ParseFilesInDir(directory: str, shortFilename: str) -> List[float]:
    res_list = list()
    for subdir in os.listdir(directory):
        filepath = os.path.join(directory, subdir, shortFilename)
        # if int(subdir) <= 2:
        #     continue
        print(f"Обрабатывается файл: {filepath}")

        if os.path.isfile(filepath):
            ParseFile(filepath, res_list)
    
    # x = timestamp from 0, y = metric
    med_res = list()
    for lst in res_list:
        med_res.append(sum(lst) / len(lst))

    return med_res

def Plot(med_res: List[float], graph_label: str, outputFilename):
    timestamps = [i for i in range(len(med_res))]
    metrics = [med_res[i] for i in timestamps]

    plt.figure(figsize=(14, 8))
    # Строим линейный график
    plt.plot(timestamps, metrics, alpha=1, linewidth=1.5, color='blue', label=graph_label)
    # # Добавляем скользящее среднее для сглаживания
    # if len(timestamps) > 10:
    #     window_size = min(window, len(timestamps) // 10)
    #     df = pd.DataFrame({'x': timestamps, 'y': metrics})
    #     df['moving_avg'] = df['y'].rolling(window=window_size, center=True).mean()
    #     plt.plot(df['x'], df['moving_avg'], color='red', linewidth=2, 
    #             label=f'Скользящее среднее (окно={window_size})')
    
    plt.xlabel('Время (секунды)', fontsize=12)
    plt.ylabel(graph_label, fontsize=12)
    plt.title(graph_label, fontsize=14)
    plt.grid(True, alpha=0.3)
    plt.legend()

    plt.savefig(outputFilename, dpi=300, bbox_inches='tight')
    print(f"График сохранен как {outputFilename}")

def main_prom():
    for curTypeLogs in TYPE_LOGS:
        dir = f"{DIRECTORY_METRICS}/{curTypeLogs}"
        for shortFilename, params in FILENAMES.items():
            med_res = ParseFilesInDir(directory=dir, shortFilename=shortFilename)
            if len(med_res) != 0:
                Plot(med_res, params.graph_label, f"./img/{curTypeLogs}/{params.outputFilename}")

if __name__ == '__main__':
    main_prom()