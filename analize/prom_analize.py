from typing import List
import os
import matplotlib.pyplot as plt
import pandas as pd
from conf import DIRECTORY_METRICS, GROUPS

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
                parts = line.strip().split()
                assert(len(parts) == 2)

                timestamp, metric = int(parts[0]), float(parts[1])
                if ir >= len(res_list):
                    res_list.append([])
                res_list[ir].append(metric)
                ir += 1
        # print(f"ir = {ir}")
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
        print("res_list: ", len(res_list))
    
    # x = timestamp from 0, y = metric
    med_res = list()
    for lst in res_list:
        med_res.append(sum(lst) / len(lst))

    return med_res

def AddGraph(time_starts, durations_ms, type, color1, color2):
    # Создаем линейный график
    # plt.plot(time_starts, durations_ms, alpha=0.2, linewidth=1.5, color=color1, label=f"{type}")
    # Добавляем скользящее среднее для тренда
    if len(time_starts) > 10:
        window_size = min(window, len(time_starts) // 10)
        df = pd.DataFrame({'time': time_starts, 'duration': durations_ms})
        df = df.sort_values('time')
        df['moving_avg'] = df['duration'].rolling(window=window_size, center=True).mean()
        
        plt.plot(df['time'], df['moving_avg'], color=color2, linewidth=2, 
                label=f'{type} скользящее среднее (окно={window_size})')


def Plot(fiber_med_res: List[float], gin_med_res: List[float], graph_label: str, outputFilename):
    fiber_timestamps = [i for i in range(len(fiber_med_res))]
    fiber_metrics = [fiber_med_res[i] for i in fiber_timestamps]
    gin_timestamps = [i for i in range(len(gin_med_res))]
    gin_metrics = [gin_med_res[i] for i in gin_timestamps]

    AddGraph(fiber_timestamps, fiber_metrics, "fiber", 'blue', 'red')
    AddGraph(gin_timestamps, gin_metrics, "gin", 'cyan', 'green')
    
    plt.xlabel('Время (секунды)', fontsize=12)
    plt.ylabel(graph_label, fontsize=12)
    plt.title(graph_label, fontsize=14)
    plt.grid(True, alpha=0.3)
    plt.legend()

    plt.savefig(outputFilename, dpi=300, bbox_inches='tight')
    plt.close()
    print(f"График сохранен как {outputFilename}")

def main_prom():
    for shortFilename, params in FILENAMES.items():
        for type, gr in GROUPS.items():
            fiber_med_res = ParseFilesInDir(directory=f"{DIRECTORY_METRICS}/{gr[0]}", shortFilename=shortFilename)
            gin_med_res = ParseFilesInDir(directory=f"{DIRECTORY_METRICS}/{gr[1]}", shortFilename=shortFilename)
            os.makedirs(f"./img/{type}", exist_ok=True)
            Plot(
                fiber_med_res, 
                gin_med_res,
                params.graph_label + " " + type, f"./img/{type}/{params.outputFilename}")
            

if __name__ == '__main__':
    main_prom()